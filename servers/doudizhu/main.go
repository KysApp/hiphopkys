package main

import (
	"github.com/astaxie/beego"
	_ "hiphopkys/servers/horse/routers"
	"runtime"
)

func init() {
	cpuNum := runtime.NumCPU()
	runtime.GOMAXPROCS(cpuNum)
}
func main() {
	beego.Run()
}
