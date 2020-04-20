package main

import (
	"Authentication"
	"Core"
	"FileParser"
	"Local.main/Local"
	"Logging"
	"net"
)

var ConfigPath = "./config.json"

func readJson(serverInfo *Local.ServerInfo) {
	Logging.NormalLogger.Println("Going to read Json")
	if err := FileParser.GetJasonConfig(ConfigPath, serverInfo); err != nil {
		Logging.ErrorLogger.Fatal("encounter a error when reading Json file")
	}
}

func signIn(serverInfo Local.ServerInfo, serverTcpConn *net.TCPConn) {
	// we need to send user name and key to verify
	// we need to make sure decode and encode table will be sent
	// also we need to send key and password as verify
	Logging.NormalLogger.Println("going to send username")
	EncodeName := Core.ConvertStringTOByte(Authentication.EncodeUsername(serverInfo.GetUserName()))
	check1, check2 := Core.WriteAll(EncodeName, serverTcpConn, 128)
	if check1 == -1 && check2 != nil {
		Logging.ErrorLogger.Fatal("encounter a error when sending username")
	}

	Logging.NormalLogger.Println("going to send password")
	EncodeKey := Core.ConvertStringTOByte(Authentication.EncodePassword(serverInfo.GetPassword()))
	check1, check2 = Core.WriteAll(EncodeKey, serverTcpConn, 128)
	if check1 == -1 && check2 != nil {
		Logging.ErrorLogger.Fatal("encounter a error when sending password")
	}
}

func sendEncryptionTable(proxy *Core.Proxy, serverTcpConn *net.TCPConn) {
	Logging.NormalLogger.Println("going to send encode arr")
	encode := proxy.EncryptionTable.GetEncodeArr()
	check1, check2 := Core.WriteAll(encode, serverTcpConn, 256)
	if check1 == -1 && check2 != nil {
		Logging.ErrorLogger.Fatal("encounter a error when sending encode arr")
	}

	Logging.NormalLogger.Println("going to send decode arr")
	decode := proxy.EncryptionTable.GetDecodeArr()
	check1, check2 = Core.WriteAll(decode, serverTcpConn, 256)
	if check1 == -1 && check2 != nil {
		Logging.ErrorLogger.Fatal("encounter a error when sending decode arr")
	}
	Logging.NormalLogger.Println(encode)
	Logging.NormalLogger.Println(decode)
}

func listenConnection(proxy *Core.Proxy, tcpListener *net.TCPListener, serverTcpConn *net.TCPConn) {
	var err error
	var i = 0
	Logging.NormalLogger.Println("local is waiting for connection")
	for {
		localTcpConn, _ := tcpListener.AcceptTCP()
		if i != 0 {
			serverTcpConn, err = net.DialTCP("tcp", nil, proxy.ServerHost)
			if err != nil {
				Logging.ErrorLogger.Fatal("terminating: server is not running")
			}
		} else {
			i = 1
		}

		Logging.NormalLogger.Println("accepted a connection")
		connection := Core.NewConnectionHandler(proxy, localTcpConn, serverTcpConn)
		go connection.DealAsBoth()
	}
}

func main() {
	// should be get in configuration
	var err error
	var serverInfo Local.ServerInfo

	readJson(&serverInfo)

	Logging.NormalLogger.Println("starting local proxy")
	// we need to encode this key and password
	proxy, err := Core.NewLocalProxy(serverInfo.GetLocalAddr(), serverInfo.GetServerAddr())
	if err != nil {
		Logging.ErrorLogger.Fatal("encounter a error when starting local proxy")
	}

	serverTcpConn, err := net.DialTCP("tcp", nil, proxy.ServerHost)
	if err != nil {
		Logging.ErrorLogger.Fatal("terminating: server is not running")
	}

	signIn(serverInfo, serverTcpConn)

	sendEncryptionTable(proxy, serverTcpConn)

	Logging.NormalLogger.Printf("going to listen: %s:%d\n", proxy.LocalHost.IP.String(), proxy.LocalHost.Port)
	// as a server for localhost
	tcpListener, _ := net.ListenTCP("tcp", proxy.LocalHost) // todo why ignore error?
	// add defer close
	defer func() {
		if err := tcpListener.Close(); err != nil {
			Logging.NormalLogger.Println("cannot close tcp listener")
			Logging.ErrorLogger.Println(err)
		}
	}()

	listenConnection(proxy, tcpListener, serverTcpConn)
}
