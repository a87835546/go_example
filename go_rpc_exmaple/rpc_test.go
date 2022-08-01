package go_rpc_exmaple

import (
	"encoding/gob"
	"fmt"
	"log"
	"testing"
)

type User struct {
	Name string
	Age  int
}

func SayHello(user User) string {
	return fmt.Sprintf("hello %s:%d", user.Name, user.Age)
}
func FuzzName(f *testing.F) {
	log.Println("test starting ---------")
	gob.Register(User{})
	server := NewServer(":8080")
	server.Register("SayHello", SayHello)
	server.Run()
}
