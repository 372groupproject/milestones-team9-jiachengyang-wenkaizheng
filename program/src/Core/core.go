package Core

import (
	"Logging"
	"net"
)

const (
	Server = iota
	Local  = iota
)

const (
	type0 = iota
	type1 = iota
)

func ConvertStringTOByte(str string) []byte {
	return []byte(str)

}
func ConvertByteTOString(bt []byte) string {
	return string(bt)
}

func ReadAll(buffer []byte, socket *net.TCPConn, size int) (int, error) {
	// 256 is one packet size
	readLength, err := socket.Read(buffer)
	if err != nil {
		return -1, err
	}
	if readLength == 0 {
		return 0, nil
	}
	for {
		if readLength == size {
			break
		}
		twice, err := socket.Read(buffer[readLength:size])
		if err != nil {
			return -1, err
		}
		if twice == 0 {
			return 0, nil
		}
		readLength += twice
	}
	return 1, nil
}
func WriteAll(buffer []byte, socket *net.TCPConn, size int) (int, error) {
	// 256 is one packet size
	writeLength, err := socket.Write(buffer)
	if err != nil {
		return -1, err
	}
	for {
		if writeLength == size {
			break
		}
		twice, err := socket.Write(buffer[writeLength:size])
		if err != nil {
			return -1, err
		}
		writeLength += twice
	}
	return 1, nil

}

func Transfer(proxy *Proxy, conn1, conn2 *net.TCPConn, device, types int) {
	for {
		Logging.NormalLogger.Println("encode arr", proxy.EncryptionTable.GetEncodeArr())
		Logging.NormalLogger.Println("decode arr", proxy.EncryptionTable.GetDecodeArr())
		request := make([]byte, 256)
		readLen, err := conn1.Read(request)
		// we need to use proxy to
		if err != nil {
			Logging.NormalLogger.Println(*proxy, "device and types", device, types, "got an error when reading request", "get length", readLen)
			Logging.ErrorLogger.Println(err)
			break
		}
		// connection close by user
		if readLen == 0 {
			Logging.NormalLogger.Println(*proxy, "device and types", device, types, "connection closed by user")
			break
		}

		Logging.NormalLogger.Println(request)
		// if it is local and works as a server
		// or if it is server and works as a client
		if (device == Local && types == type0) || (device == Server && types == type1) {
			request = proxy.EncryptionTable.Encode(request)
			Logging.NormalLogger.Print(*proxy, "device and types", device, types, "got encoded request")
			Logging.NormalLogger.Println(request)
		}
		// if it is a local and works as a client
		// or it is a server and works as a server

		if (device == Local && types == type1) || (device == Server && types == type0) {
			request = proxy.EncryptionTable.Decode(request)
			Logging.NormalLogger.Print(*proxy, "device and types", device, types, "got decoded request ")
			Logging.NormalLogger.Println(request)
		}

		// we send this byte to sp
		writeLength, err := conn2.Write(request[0:readLen])
		if err != nil {
			Logging.NormalLogger.Println(*proxy, "device and types", device, types, "got an error when writing request")
			Logging.ErrorLogger.Println(err)
			break
		}
		for writeLength != readLen {
			Logging.NormalLogger.Println(*proxy, "device and types", device, types, "going to write request")
			if writeLength != readLen {
				Logging.NormalLogger.Println(*proxy, "device and types", device, types, "write length != readLen")
				twiceLength, err := conn2.Write(request[writeLength:readLen])
				if err != nil {
					Logging.NormalLogger.Println(*proxy, "device and types", device, types, "got an error when writing request")
					Logging.ErrorLogger.Println(err)
					break
				}
				writeLength += twiceLength
			}
		}
	}
}
