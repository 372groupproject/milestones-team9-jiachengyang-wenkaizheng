package Core

import (
	"Logging"
	"errors"
	"net"
)

type ConnectionHandler struct {
	proxy             *Proxy
	localTcpConn      *net.TCPConn
	serverTcpConn     *net.TCPConn
	localTcpComplete  chan int
	serverTcpComplete chan int
	isLocalRunning    bool
	isServerRunning   bool
}

func NewConnectionHandler(proxy *Proxy, local, server *net.TCPConn) *ConnectionHandler {
	return &ConnectionHandler{
		proxy:             proxy,
		localTcpConn:      local,
		serverTcpConn:     server,
		localTcpComplete:  make(chan int),
		serverTcpComplete: make(chan int),
		isLocalRunning:    false,
		isServerRunning:   false,
	}
}

func (h *ConnectionHandler) dealAsServer() error {
	Logging.NormalLogger.Print("Going to deal as server ")
	Logging.NormalLogger.Println(*h)
	if h.isServerRunning {
		return errors.New("server is already running")
	}
	h.isServerRunning = true
	h.serverTcpComplete <- 0
	var device, t = h.proxy.device, type0
	Transfer(h.proxy, h.localTcpConn, h.serverTcpConn, device, t)
	h.serverTcpComplete <- 0
	return nil
}

func (h *ConnectionHandler) dealAsClient() error {
	Logging.NormalLogger.Print("Going to deal as client ")
	Logging.NormalLogger.Println(*h)
	if h.isLocalRunning {
		return errors.New("client is already running")
	}
	h.isLocalRunning = true
	h.localTcpComplete <- 0
	var device, t int = h.proxy.device, type1
	Transfer(h.proxy, h.serverTcpConn, h.localTcpConn, device, t)
	h.localTcpComplete <- 0
	return nil
}

func (h *ConnectionHandler) DealAsBoth() {
	go func() {
		if err := h.dealAsServer(); err != nil {
			Logging.ErrorLogger.Fatal(err)
		}
	}()
	go func() {
		if err := h.dealAsClient(); err != nil {
			Logging.ErrorLogger.Fatal(err)
		}
	}()
	<-h.localTcpComplete
	<-h.serverTcpComplete
	if err := h.Wait(); err != nil {
		Logging.ErrorLogger.Fatal(err)
	}
	if err := h.Close(); err != nil {
		Logging.ErrorLogger.Fatal(err)
	}
}

func (h *ConnectionHandler) Wait() error {
	Logging.NormalLogger.Print("Going to wait client and server ")
	Logging.NormalLogger.Println(*h)
	if !h.isLocalRunning {
		return errors.New("client is not running")
	}
	if !h.isServerRunning {
		return errors.New("server is not running")
	}
	<-h.localTcpComplete
	<-h.serverTcpComplete
	h.isLocalRunning = false
	h.isServerRunning = false
	return nil
}

func (h *ConnectionHandler) GetProxy() *Proxy {
	return h.proxy
}

func (h *ConnectionHandler) Close() error {
	Logging.NormalLogger.Print("Going to close client and server ")
	Logging.NormalLogger.Println(*h)
	if h.isServerRunning {
		return errors.New("need to wait server first")
	}
	if h.isLocalRunning {
		return errors.New("need to wait local first")
	}
	if err := h.serverTcpConn.Close(); err != nil {
		return err
	}
	if err := h.localTcpConn.Close(); err != nil {
		return err
	}
	return nil
}
