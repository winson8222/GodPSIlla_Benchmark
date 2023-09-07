// Code generated by hertz generator. DO NOT EDIT.

package viewersapi

import (
	viewersapi "gateway/biz/handler/viewersapi"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_viewerservice := root.Group("/ViewerService", _viewerserviceMw()...)
		_viewerservice.POST("/getuniqueviewernames2080", append(_getuniqueviewernames2080Mw(), viewersapi.Getuniqueviewernames2080)...)
		_viewerservice.POST("/getuniqueviewernames2099980", append(_getuniqueviewernames2099980Mw(), viewersapi.Getuniqueviewernames2099980)...)
		_viewerservice.POST("/getuniqueviewernames20k80k", append(_getuniqueviewernames20k80kMw(), viewersapi.Getuniqueviewernames20k80k)...)
		_viewerservice.POST("/getuniqueviewernames8020", append(_getuniqueviewernames8020Mw(), viewersapi.Getuniqueviewernames8020)...)
		_viewerservice.POST("/getuniqueviewernames8099920", append(_getuniqueviewernames8099920Mw(), viewersapi.Getuniqueviewernames8099920)...)
		_viewerservice.POST("/getuniqueviewernames80k20k", append(_getuniqueviewernames80k20kMw(), viewersapi.Getuniqueviewernames80k20k)...)
	}
}