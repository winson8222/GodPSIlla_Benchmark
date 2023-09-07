// Code generated by hertz generator.

	package main
	
	import (
		handler "gateway/biz/handler"
	
		"github.com/cloudwego/hertz/pkg/app/server"
	)
	
	// customizeRegister registers customize routers.
	func customizedRegister(r *server.Hertz) {
		r.GET("/ping", handler.Ping)
		r.POST("/PSI", handler.PSI)
	
		// your code ...
	}
	