/*
	Created by XiaoInk at 2021/06/11 16:11:00
	GitHub: https://github.com/XiaoInk
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	Port int
}

func main() {
	server := &Server{}

	flag.IntVar(&server.Port, "port", 8080, "服务端口")

	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("%d", server.Port)))

		log.Println(r.RemoteAddr, r.Method, r.RequestURI)
	})

	http.ListenAndServe(fmt.Sprintf(":%d", server.Port), nil)
}
