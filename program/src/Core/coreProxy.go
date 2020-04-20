package Core

import (
	"Encryption"
	"Logging"
	"encoding/binary"
	"math/rand"
	"net"
	"time"
)

type Proxy struct {
	LocalHost       *net.TCPAddr
	ServerHost      *net.TCPAddr
	EncryptionTable *Encryption.Table
	device          int // 1 is server 0 is local
}

func NewServerProxy(local string) (*Proxy, error) {
	// as a server we don't need ip address just port
	addr0, err := net.ResolveTCPAddr("tcp4", local)
	if err != nil {
		return nil, err
	}
	return &Proxy{LocalHost: addr0, ServerHost: nil, EncryptionTable: Encryption.NewEmptyEncryptionTable(), device: Server}, nil
}

func NewLocalProxy(local, server string) (*Proxy, error) {
	rand.Seed(time.Now().Unix())
	//rand.Seed(0)
	// as a server we don't need ip address just port
	addr0, err := net.ResolveTCPAddr("tcp", local)
	if err != nil {
		return nil, err
	}
	// as a client we need both ip address and port
	addr1, err := net.ResolveTCPAddr("tcp", server)
	if err != nil {
		return nil, err
	}
	return &Proxy{LocalHost: addr0, ServerHost: addr1, EncryptionTable: Encryption.NewEncryptionTable(), device: Local}, nil
}

func (p *Proxy) ConnectToRealServer(request []byte, length int) *net.TCPAddr {
	Logging.NormalLogger.Println("going to connect to real server")
	Logging.NormalLogger.Println("request ", request, "length ", length)
	port := int(binary.BigEndian.Uint16(request[length-2:]))
	Logging.NormalLogger.Println("get port", port)
	var ip []byte
	if request[3] == 0x1 {
		Logging.NormalLogger.Println("ipv4 request")
		ip = request[4 : 4+net.IPv4len]
	} else if request[3] == 0x3 {
		Logging.NormalLogger.Println("0x3 request")
		ip1, err := net.ResolveIPAddr("ip", string(request[5:length-2]))
		if err != nil {
			return nil
		}
		ip = ip1.IP
	} else if request[3] == 0x4 {
		Logging.NormalLogger.Println("ipv6 request")
		ip = request[4 : 4+net.IPv6len]
	}
	return &net.TCPAddr{
		IP:   ip,
		Port: port}
}
func (p *Proxy) SetServerHost(copy *net.TCPAddr) {
	p.ServerHost = copy
}
func (p *Proxy) SetLocalHost(copy *net.TCPAddr) {
	p.LocalHost = copy
}
