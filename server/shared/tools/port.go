package tools

import (
	"net"
	"tank_war/server/shared/consts"
)

func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", consts.FreePortAddress)
	if err != nil {
		return 0, err
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}
