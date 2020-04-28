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
		request := make([]byte, 256)
		readLen, err := conn1.Read(request)
		if err != nil {
			break
		}
		// connection close by user
		if readLen == 0 {
			break
		}
		if (device == Local && types == type0) || (device == Server && types == type1) {
			request = proxy.EncryptionTable.Encode(request)
		}
		if (device == Local && types == type1) || (device == Server && types == type0) {
			request = proxy.EncryptionTable.Decode(request)
		}
		writeLength, err := conn2.Write(request[0:readLen])
		if err != nil {
			break
		}
		for writeLength != readLen {
			if writeLength != readLen {
				twiceLength, err := conn2.Write(request[writeLength:readLen])
				if err != nil {
					break
				}
				writeLength += twiceLength
			}
		}
	}
}
