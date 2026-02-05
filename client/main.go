package main

import (
	"fmt"
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
	return nil
}

func main() {
	req := Req{Num1: 10, Num2: 201}
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		fmt.Println(err)
		return
	}
	var resp Res
	err = client.Call("Server.Add1", req, &resp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.Num)
}
