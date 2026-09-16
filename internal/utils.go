package utils

import (
	"encoding/json"
	"errors"
	"fmt"
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

func SetPaused(paused bool, socketPath string) error {
	conn, err := net.DialTimeout("unix", socketPath, 100*time.Millisecond)
	if err != nil {
		return fmt.Errorf("Bento is currently not playing!")
	}
	defer conn.Close()

	request := struct {
		Command []any `json:"command"`
	}{
		Command: []any{"set_property", "pause", paused},
	}

	err = json.NewEncoder(conn).Encode(request)
	if err != nil {
		return fmt.Errorf("send playback command: %w", err)
	}

	var response struct {
		Error string `json:"error"`
	}

	if err := json.NewDecoder(conn).Decode(&response); err != nil {
		return fmt.Errorf("read mpv response: %w", err)
	}

	if response.Error != "success" {
		return fmt.Errorf("mpv rejected command: %s", response.Error)
	}
	
	return nil
}
