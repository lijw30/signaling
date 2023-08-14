package main

import (
	"flag"
	"github.com/lijw30/signaling/src/framework"
)

func main() {
	flag.Parse()
	err := framework.Init("./conf/framework.toml")
	if err != nil {
		panic(err)
	}

	// 启动http
	go startHttp()

	// 启动https
	startHttps()
}

func startHttp() {
	err := framework.StartHttp()
	if err != nil {
		panic(err)
	}
}

func startHttps() {
	err := framework.StartHttps()
	if err != nil {
		panic(err)
	}
}
