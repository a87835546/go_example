package go_rpc_exmaple

import (
	"errors"
	"io"
	"log"
	"net"
	"reflect"
	"time"
)

type Server struct {
	addr  string
	funcs map[string]reflect.Value
}

func NewServer(addr string) *Server {
	return &Server{
		addr:  addr,
		funcs: make(map[string]reflect.Value),
	}
}

func (s *Server) Register(rpcName string, f any) {
	v := reflect.ValueOf(f)
	if v.Kind() != reflect.Func {
		panic(errors.New("invalid Typeof Func"))
	}
	if _, ok := s.funcs[rpcName]; ok == false {
		panic(errors.New("repeated rpc name"))
	}
	s.funcs[rpcName] = v
}

func (s *Server) Run() {
	if listen, err := net.Listen("tcp", s.addr); err != nil {
		panic(errors.New("run server err  -->>> " + err.Error()))
	} else {
		for {
			conn, err := listen.Accept()
			if err != nil {
				continue
			}
			go s.process(conn)
		}
	}

}

func (s *Server) process(conn net.Conn) {
	session := NewSession(conn)
	session.conn.SetDeadline(time.Time{})
	for {
		b, err := session.Read()
		if err == io.EOF {
			log.Print("EOF received")
			return
		}
		if err != nil {
			log.Printf("Error reading session err :" + err.Error())
			return
		}
		data := new(UserData)
		if err := data.Decode(b); err != nil {
			log.Printf("Error decoding session err :" + err.Error())
			return
		}

		f, ok := s.funcs[data.Name]
		if ok == false {
			log.Printf("Error call func not find, func name: " + data.Name)
			return
		}

		t := f.Type()
		if t.NumIn() != len(data.Args) {
			log.Printf("Error NumIn Lens Not Equal")
		}

		inArgs := make([]reflect.Value, 0, len(data.Args))
		for i, arg := range data.Args {
			v := reflect.ValueOf(arg)
			if t.In(i).Kind() != v.Type().Kind() {
				log.Printf("Err parameter type mistach need: %s but : %s", t.In(i).Kind(), v.Type().Kind())
			}
			inArgs = append(inArgs, v)
		}
		out := f.Call(inArgs)
		outArgs := make([]any, 0, len(out))
		for _, value := range out {
			outArgs = append(outArgs, value.Interface())
		}

		response := NewUserData(data.Name, outArgs)
		bytes, err := response.Encode()
		if err != nil {
			log.Printf("Encode error :%s", err.Error())
			return
		}
		if err := session.Write(bytes); err != nil {
			log.Printf("Send error %s", err.Error())
		}

	}
}
