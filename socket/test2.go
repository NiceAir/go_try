package main

import (
	"fmt"
	"net"
)

func main()  {
	// 查ip地址
	r, _ := net.ResolveTCPAddr("tcp", "www.gooogle.com:80")

	fmt.Println(r)
}
