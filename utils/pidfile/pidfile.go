package pidfile

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
)

type PidFile struct {
	path string
}

func Create(path string) (*PidFile, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return nil, err
	}

	if pid, err := getpid(path); err == nil {
		if process, err := os.FindProcess(pid); err == nil {
			if err := process.Signal(syscall.Signal(0)); err == nil {
				return nil, fmt.Errorf("process %d exists", pid)
			}
		}
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	_, err = file.WriteString(strconv.Itoa(os.Getpid()))

	return &PidFile{path: path}, nil
}

func Remove(p *PidFile) error {
	return os.Remove(p.path)
}

func getpid(path string) (int, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}
	pid, err := strconv.Atoi(string(data))
	if err != nil {
		return 0, err
	}
	return pid, nil
}
