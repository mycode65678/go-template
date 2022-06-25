package main

import (
	"go-hash/a"
	"go-hash/orm"
	"go-hash/web"
	"os"
	"os/signal"
)

// 验证规则
//https://github.com/asaskevich/govalidator

func main() {
	a.Start()
	orm.Start()
	go web.StartHttp()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
}
