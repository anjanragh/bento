package utils

import (
	"errors"
	"net"
	"os"
	"time"
)

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}

func SocketActive(socketPath string) bool {
	conn, err := net.DialTimeout("unix", socketPath, 100*time.Millisecond)

	if err != nil {
		return false
	}
	conn.Close()
	return true
}
