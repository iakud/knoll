package klog

import (
	"fmt"
	"path/filepath"
	"testing"
	"time"
)

func TestFileWriter(t *testing.T) {
	fw := NewFileWriter(filepath.Join("tests", t.Name()+".log"), 1024, 1)
	defer fw.Close()
	s := fmt.Sprintf("open file: %s", time.Now())
	fw.Write([]byte(s))
}

func TestFileWriterFlush(t *testing.T) {
	fw := NewFileWriter(filepath.Join("tests", t.Name()+".log"), 1024, 1)
	s := fmt.Sprintf("open file: %s\n", time.Now())
	fw.Write([]byte(s))
	fw.Write([]byte("flush\n"))
	fw.Flush()
	fw.Write([]byte("after flush\n"))
}

func TestFileWriterRolls(t *testing.T) {
	maxRolls := 2

	fw := NewFileWriter(filepath.Join("tests", t.Name()+".log"), 4, maxRolls)
	defer fw.Close()
	time.Sleep(time.Second * 1)
	fw.Write([]byte("test 1"))
	time.Sleep(time.Second * 1)
	fw.Write([]byte("test 2"))
	time.Sleep(time.Second * 1)
	fw.Write([]byte("test 3"))
	time.Sleep(time.Second * 1)
	fw.Write([]byte("test 4"))
}
