package redis

import (
	"log"
	"testing"
)

func TestInitClient(t *testing.T) {
	addr := "81.69.255.188:6379"
	err := InitClient(addr)
	log.Println(err)
}
