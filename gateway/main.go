// Code generated by hertz generator.

package main

import (
	"gateway/biz/handler"
	"github.com/cloudwego/hertz/pkg/app/server"
	"tests/constants"
)

func main() {
	h := server.Default(server.WithHostPorts(constants.APIGATEWAY_URL))
	handler.InitPSISvcInfo() //initialise the map of PSI svc info
	register(h)
	h.Spin()
}