package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type Server struct {
}

type Req struct {
	Num1 int
	Num2 int
}

type Res struct {
	Num int
}

func (s *Server) Add(r Req, resp *Res) error {
	resp.Num = r.Num1 + r.Num2
	fmt.Println("request is coming", r)
	return nil
}

func main() {
	rpc.Register(new(Server))
	rpc.HandleHTTP()
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println(err)
	}
	http.Serve(listen, nil)
}
