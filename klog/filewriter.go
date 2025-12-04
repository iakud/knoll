package klog

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

var ErrClosed = errors.New("klog: file writer already closed")

const bufferSize = 256 * 1024
const flushInterval = 3 * time.Second
const rollInterval = 24 * time.Hour
const fileRotateLayout = "20060102-150405"
const DefaultRollSize = 10240000

type FileWriter struct {
	dir    string
	name   string
	cancel context.CancelFunc

	mutex  sync.Mutex
	file   *os.File
	buffer *bufio.Writer
	closed bool

	rollTime     time.Time
	writtenBytes int
	rollSize     int
	maxRolls     int
	history      []string
	filePeriod   time.Time
}

// maxRools: if <= 0, unlimited
func NewFileWriter(path string, rollSize int, maxRolls int) *FileWriter {
	ctx, cancel := context.WithCancel(context.Background())
	dir, name := filepath.Split(path)
	fw := &FileWriter{
		dir:  filepath.Dir(dir),
		name: name,

		cancel: cancel,

		rollTime: time.Unix(0, 0),
		rollSize: rollSize,
		maxRolls: maxRolls,
	}
	if maxRolls > 0 {
		if history, err := fw.historyRolls(); err == nil {
			fw.history = history
		}
		fw.removeOldRolls()
	}
	fw.rollFile(time.Now())
	go fw.flushPeriodically(ctx)
	return fw
}

func (fw *FileWriter) Write(p []byte) (int, error) {
	fw.mutex.Lock()
	defer fw.mutex.Unlock()
	if fw.closed {
		return 0, ErrClosed
	}

	n, err := fw.buffer.Write(p)
	fw.writtenBytes += n

	now := time.Now()
	if fw.writtenBytes > fw.rollSize {
		fw.rollFile(now)
	} else if period := fw.period(now); period != fw.filePeriod {
		fw.rollFile(now)
	}
	return n, err
}

func (fw *FileWriter) Sync() error {
	fw.mutex.Lock()
	defer fw.mutex.Unlock()
	if fw.closed {
		return ErrClosed
	}
	if fw.buffer == nil {
		return nil
	}
	fw.buffer.Flush()
	return fw.file.Sync()
}

func (fw *FileWriter) Flush() error {
	fw.mutex.Lock()
	defer fw.mutex.Unlock()
	if fw.closed {
		return ErrClosed
	}
	if fw.buffer == nil {
		return nil
	}
	return fw.buffer.Flush()
}

func (fw *FileWriter) Close() error {
	fw.mutex.Lock()
	defer fw.mutex.Unlock()
	if fw.closed {
		return ErrClosed
	}
	fw.closed = true
	fw.cancel()
	if fw.buffer == nil {
		return nil
	}
	fw.buffer.Flush()
	return fw.file.Close()
}

func (fw *FileWriter) flushPeriodically(ctx context.Context) {
	ticker := time.NewTicker(flushInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			fw.Flush()
		case <-ctx.Done():
			return
		}
	}
}

func (fw *FileWriter) rollFile(now time.Time) {
	if !now.After(fw.rollTime) {
		return
	}

	if fw.buffer != nil {
		fw.buffer.Flush()
		fw.buffer = nil
		filename := fw.file.Name()
		fw.file.Close()
		fw.file = nil
		if fw.maxRolls > 0 {
			fw.history = append(fw.history, filepath.Base(filename))
			fw.removeOldRolls()
		}
	}
	file, err := fw.createFile(now)
	if err != nil {
		panic(err)
	}
	fw.file = file
	fw.buffer = bufio.NewWriterSize(file, bufferSize)
	fw.writtenBytes = 0

	period := fw.period(now)
	fw.rollTime = now
	fw.filePeriod = period
}

func (fw *FileWriter) createFile(t time.Time) (*os.File, error) {
	ext := filepath.Ext(fw.name)
	prefix := strings.TrimSuffix(fw.name, ext)
	name := fmt.Sprintf("%s.%s%s", prefix, t.Format(fileRotateLayout), ext)
	filename := filepath.Join(fw.dir, name)

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return nil, fmt.Errorf("klog: open file error: %v", err)
	}

	symlink := filepath.Join(fw.dir, fw.name)
	os.Remove(symlink) // ignore err
	if err := os.Symlink(name, symlink); err != nil {
		os.Link(filename, symlink)
	}
	return file, nil
}

func (fw *FileWriter) historyRolls() ([]string, error) {
	ext := filepath.Ext(fw.name)
	prefix := strings.TrimSuffix(fw.name, ext)
	f, err := os.Open(fw.dir)
	if err != nil {
		return nil, err
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	var history []string
	for _, file := range files {
		// regular files
		if file.Mode()&os.ModeType != 0 {
			continue
		}
		// filter
		if strings.HasPrefix(file.Name(), prefix) {
			history = append(history, file.Name())
		}
	}
	sort.Strings(history)
	return history, nil
}

func (fw *FileWriter) removeOldRolls() {
	if nRolls := len(fw.history); nRolls > fw.maxRolls {
		removeRools := nRolls - fw.maxRolls
		for _, filename := range fw.history[:removeRools] {
			name := filepath.Join(fw.dir, filename)
			os.Remove(name) // ignore err
		}
		fw.history = fw.history[removeRools:]
	}
}

func (fw *FileWriter) period(now time.Time) time.Time {
	if rollInterval == 24*time.Hour {
		year, month, day := now.Date()
		return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	}
	return now.Truncate(rollInterval)
}
