package knet_test

import (
	"bytes"
	"log"
	"testing"

	"github.com/iakud/knoll/knet"
)

func TestCodec(t *testing.T) {
	buffer := bytes.NewBuffer(nil)
	var c knet.Codec = &knet.StdCodec{}
	message := "hello"
	if err := c.Write(buffer, []byte(message)); err != nil {
		log.Fatalln(err)
	}
	log.Println("codec write", message)
	b, err := c.Read(buffer)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("codec read", string(b))
}
