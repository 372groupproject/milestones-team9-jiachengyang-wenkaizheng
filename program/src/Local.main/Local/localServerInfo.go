package Local

import "strconv"

type ServerInfo struct {
	Server     string `json:"server"`
	ServerPort int    `json:"server_port"`
	LocalPort  int    `json:"local_port"`
	Password   string `json:"password"`
	Timeout    int    `json:"timeout"`
	UserName   string `json:"username"`
}

func (s ServerInfo) GetServerAddr() string {
	return s.Server + ":" + strconv.Itoa(s.ServerPort)
}

func (s ServerInfo) GetLocalAddr() string {
	return "127.0.0.1:" + strconv.Itoa(s.LocalPort)
}

func (s ServerInfo) GetUserName() string {
	return s.UserName
}

func (s ServerInfo) GetPassword() string {
	return s.Password
}
