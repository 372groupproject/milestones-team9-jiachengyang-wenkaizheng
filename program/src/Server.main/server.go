package main

import (
	"Authentication"
	"Core"
	"Logging"
	"errors"
	"net"
)

var DataPath = "./data.csv"

func signInUser(localTcpConn *net.TCPConn) (bool, error) {
	name := make([]byte, 128)
	key := make([]byte, 128)
	// read password and name first and then read decode , encode
	check1, check2 := Core.ReadAll(name, localTcpConn, 128)
	if (check1 == -1 && check2 != nil) || (check1 == 0 && check2 == nil) {
		return false, nil // todo make a new error
	}
	username := Core.ConvertByteTOString(name)
	check1, check2 = Core.ReadAll(key, localTcpConn, 128)
	if (check1 == -1 && check2 != nil) || (check1 == 0 && check2 == nil) {
		return false, nil // todo make a new error
	}
	password := Core.ConvertByteTOString(key)
	return Authentication.Verify(username, password)
}

func readEncryptionTable(localTcpConn *net.TCPConn, proxy *Core.Proxy) error {
	encode := make([]byte, 256)
	decode := make([]byte, 256)
	check1, check2 := Core.ReadAll(encode, localTcpConn, 256)
	if (check1 == -1 && check2 != nil) || (check1 == 0 && check2 == nil) {
		return nil // todo make a new error
	}
	proxy.EncryptionTable.SetEncodeArr(encode)
	check1, check2 = Core.ReadAll(decode, localTcpConn, 256)
	if (check1 == -1 && check2 != nil) || (check1 == 0 && check2 == nil) {
		return nil // todo make a new error
	}
	proxy.EncryptionTable.SetDecodeArr(decode)
	return nil // todo make a new error
}

func shakeHand(localTcpConn *net.TCPConn, proxy *Core.Proxy) error {
	var serverTcpConn *net.TCPConn
	request := make([]byte, 256)
	// now we expect socks5 protocol, first step is confirm socks5
	//readLength, err := localTcpConn.Read(request)
	readLength, err := localTcpConn.Read(request)
	// we need to use proxy to
	if err != nil {
		Logging.NormalLogger.Println("encounter error when reading")
		return err
	}

	decodedRequest := proxy.EncryptionTable.Decode(request[0:readLength])
	// we need to use proxy to
	if decodedRequest[0] != 0x5 {
		return nil //todo error
	}

	Logging.NormalLogger.Println("1st request", decodedRequest, "length'", readLength)

	// give response to local and do not need key and password
	encodeResponse := proxy.EncryptionTable.Encode([]byte{0x5, 0x0})
	Logging.NormalLogger.Println("going to write response")
	_, err = localTcpConn.Write(encodeResponse) // todo int return value
	//todo problem when connect 2 different websites
	/**
		+----+-----+-------+------+----------+----------+
	        |VER | CMD |  RSV  | ATYP | DST.ADDR | DST.PORT |
	        +----+-----+-------+------+----------+----------+
	        | 1  |  1  | X'00' |  1   | Variable |    2     |
	        +----+-----+-------+------+----------+----------+
		**/
	Logging.NormalLogger.Println("waiting for new package")
	// clean the buffer
	request = make([]byte, 256)
	readLength, err = localTcpConn.Read(request)
	if err != nil {
		return err
	}
	if readLength < 7 {
		Logging.NormalLogger.Println("It should be at least seven bytes")
	}
	decodedRequest = proxy.EncryptionTable.Decode(request[0:readLength])
	Logging.NormalLogger.Println("request", decodedRequest, "length'", readLength)

	// only support connect as method
	if decodedRequest[1] != 0x1 {
		return errors.New("100th connect is only support method")
	}
	//realRequest, length, err := parseRequestAddress(localTcpConn, proxy, decodedRequest, readLength, nextReadByte)

	tcpAddress := proxy.ConnectToRealServer(decodedRequest, readLength)
	serverTcpConn, err = net.DialTCP("tcp", nil, tcpAddress)
	if err != nil {
		return err
	}
	// we need to set this to server proxy struct
	proxy.SetServerHost(tcpAddress)

	response2 := []byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	encodeResponse = proxy.EncryptionTable.Encode(response2)
	//localTcpConn.Core.WriteAll(encodeResponse)
	_, err = localTcpConn.Write(encodeResponse)
	if err != nil {
		return err
	}

	connection := Core.NewConnectionHandler(proxy, localTcpConn, serverTcpConn)
	go connection.DealAsBoth()
	return err
}

func main() {
	Logging.NormalLogger.Println("server is running")
	Authentication.LoadCSV(DataPath)
	proxy, err := Core.NewServerProxy("127.0.0.1:6204") // todo may need to read from file
	if err != nil {
		Logging.NormalLogger.Println("encounter a error when starting server proxy")
		return
	}
	// as a server for localhost
	tcpListener, err := net.ListenTCP("tcp", proxy.LocalHost)
	if err != nil {
		Logging.NormalLogger.Println("encounter error when opening TCP")
		Logging.ErrorLogger.Println(err)
		return
	}
	// add defer close
	defer func() {
		if err := tcpListener.Close(); err != nil {
			Logging.NormalLogger.Println("cannot close tcp listener")
			Logging.ErrorLogger.Println(err)
		}
	}()

	// todo comment
	//var count int = 0
	ipMap := make(map[string]int)
	var ip string = ""
	for {
		localTcpConn, err := tcpListener.AcceptTCP()
		ip =  localTcpConn.RemoteAddr().String()
		if err != nil {
			Logging.NormalLogger.Println("encounter error when accepting TCP")
			Logging.ErrorLogger.Println(err)
			return
		}
		if _, exist := ipMap[ip]; !exist {
			//sign in user
			if rc, err := signInUser(localTcpConn); rc == false || err != nil {
				Logging.NormalLogger.Println("could not sign in user")
				Logging.ErrorLogger.Fatal(err)
			}

			//get the encryption table
			if err := readEncryptionTable(localTcpConn, proxy); err != nil {
				Logging.NormalLogger.Println("could not get encryption table")
				Logging.ErrorLogger.Fatal(err)
			}
			//count += 1
			ipMap[ip] = 1
		}
		if err := shakeHand(localTcpConn, proxy); err != nil {
			Logging.NormalLogger.Println("could not shake hands")
			Logging.ErrorLogger.Fatal(err)
		}
	}
}
