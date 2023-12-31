// Code generated by Kitex v0.5.2. DO NOT EDIT.

package viewerservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	viewersapi "testviewerservice/kitex_gen/viewersapi"
)

func serviceInfo() *kitex.ServiceInfo {
	return viewerServiceServiceInfo
}

var viewerServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "ViewerService"
	handlerType := (*viewersapi.ViewerService)(nil)
	methods := map[string]kitex.MethodInfo{
		"getuniqueviewernames8020":    kitex.NewMethodInfo(getuniqueviewernames8020Handler, newViewerServiceGetuniqueviewernames8020Args, newViewerServiceGetuniqueviewernames8020Result, false),
		"getuniqueviewernames2080":    kitex.NewMethodInfo(getuniqueviewernames2080Handler, newViewerServiceGetuniqueviewernames2080Args, newViewerServiceGetuniqueviewernames2080Result, false),
		"getuniqueviewernames8099920": kitex.NewMethodInfo(getuniqueviewernames8099920Handler, newViewerServiceGetuniqueviewernames8099920Args, newViewerServiceGetuniqueviewernames8099920Result, false),
		"getuniqueviewernames2099980": kitex.NewMethodInfo(getuniqueviewernames2099980Handler, newViewerServiceGetuniqueviewernames2099980Args, newViewerServiceGetuniqueviewernames2099980Result, false),
		"getuniqueviewernames80k20k":  kitex.NewMethodInfo(getuniqueviewernames80k20kHandler, newViewerServiceGetuniqueviewernames80k20kArgs, newViewerServiceGetuniqueviewernames80k20kResult, false),
		"getuniqueviewernames20k80k":  kitex.NewMethodInfo(getuniqueviewernames20k80kHandler, newViewerServiceGetuniqueviewernames20k80kArgs, newViewerServiceGetuniqueviewernames20k80kResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "viewersapi",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.5.2",
		Extra:           extra,
	}
	return svcInfo
}

func getuniqueviewernames8020Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*viewersapi.ViewerServiceGetuniqueviewernames8020Args)
	realResult := result.(*viewersapi.ViewerServiceGetuniqueviewernames8020Result)
	success, err := handler.(viewersapi.ViewerService).Getuniqueviewernames8020(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newViewerServiceGetuniqueviewernames8020Args() interface{} {
	return viewersapi.NewViewerServiceGetuniqueviewernames8020Args()
}

func newViewerServiceGetuniqueviewernames8020Result() interface{} {
	return viewersapi.NewViewerServiceGetuniqueviewernames8020Result()
}

func getuniqueviewernames2080Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*viewersapi.ViewerServiceGetuniqueviewernames2080Args)
	realResult := result.(*viewersapi.ViewerServiceGetuniqueviewernames2080Result)
	success, err := handler.(viewersapi.ViewerService).Getuniqueviewernames2080(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newViewerServiceGetuniqueviewernames2080Args() interface{} {
	return viewersapi.NewViewerServiceGetuniqueviewernames2080Args()
}

func newViewerServiceGetuniqueviewernames2080Result() interface{} {
	return viewersapi.NewViewerServiceGetuniqueviewernames2080Result()
}

func getuniqueviewernames8099920Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*viewersapi.ViewerServiceGetuniqueviewernames8099920Args)
	realResult := result.(*viewersapi.ViewerServiceGetuniqueviewernames8099920Result)
	success, err := handler.(viewersapi.ViewerService).Getuniqueviewernames8099920(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newViewerServiceGetuniqueviewernames8099920Args() interface{} {
	return viewersapi.NewViewerServiceGetuniqueviewernames8099920Args()
}

func newViewerServiceGetuniqueviewernames8099920Result() interface{} {
	return viewersapi.NewViewerServiceGetuniqueviewernames8099920Result()
}

func getuniqueviewernames2099980Handler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*viewersapi.ViewerServiceGetuniqueviewernames2099980Args)
	realResult := result.(*viewersapi.ViewerServiceGetuniqueviewernames2099980Result)
	success, err := handler.(viewersapi.ViewerService).Getuniqueviewernames2099980(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newViewerServiceGetuniqueviewernames2099980Args() interface{} {
	return viewersapi.NewViewerServiceGetuniqueviewernames2099980Args()
}

func newViewerServiceGetuniqueviewernames2099980Result() interface{} {
	return viewersapi.NewViewerServiceGetuniqueviewernames2099980Result()
}

func getuniqueviewernames80k20kHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*viewersapi.ViewerServiceGetuniqueviewernames80k20kArgs)
	realResult := result.(*viewersapi.ViewerServiceGetuniqueviewernames80k20kResult)
	success, err := handler.(viewersapi.ViewerService).Getuniqueviewernames80k20k(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newViewerServiceGetuniqueviewernames80k20kArgs() interface{} {
	return viewersapi.NewViewerServiceGetuniqueviewernames80k20kArgs()
}

func newViewerServiceGetuniqueviewernames80k20kResult() interface{} {
	return viewersapi.NewViewerServiceGetuniqueviewernames80k20kResult()
}

func getuniqueviewernames20k80kHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*viewersapi.ViewerServiceGetuniqueviewernames20k80kArgs)
	realResult := result.(*viewersapi.ViewerServiceGetuniqueviewernames20k80kResult)
	success, err := handler.(viewersapi.ViewerService).Getuniqueviewernames20k80k(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newViewerServiceGetuniqueviewernames20k80kArgs() interface{} {
	return viewersapi.NewViewerServiceGetuniqueviewernames20k80kArgs()
}

func newViewerServiceGetuniqueviewernames20k80kResult() interface{} {
	return viewersapi.NewViewerServiceGetuniqueviewernames20k80kResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Getuniqueviewernames8020(ctx context.Context, req *viewersapi.Request) (r *viewersapi.Response, err error) {
	var _args viewersapi.ViewerServiceGetuniqueviewernames8020Args
	_args.Req = req
	var _result viewersapi.ViewerServiceGetuniqueviewernames8020Result
	if err = p.c.Call(ctx, "getuniqueviewernames8020", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Getuniqueviewernames2080(ctx context.Context, req *viewersapi.Request) (r *viewersapi.Response, err error) {
	var _args viewersapi.ViewerServiceGetuniqueviewernames2080Args
	_args.Req = req
	var _result viewersapi.ViewerServiceGetuniqueviewernames2080Result
	if err = p.c.Call(ctx, "getuniqueviewernames2080", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Getuniqueviewernames8099920(ctx context.Context, req *viewersapi.Request) (r *viewersapi.Response, err error) {
	var _args viewersapi.ViewerServiceGetuniqueviewernames8099920Args
	_args.Req = req
	var _result viewersapi.ViewerServiceGetuniqueviewernames8099920Result
	if err = p.c.Call(ctx, "getuniqueviewernames8099920", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Getuniqueviewernames2099980(ctx context.Context, req *viewersapi.Request) (r *viewersapi.Response, err error) {
	var _args viewersapi.ViewerServiceGetuniqueviewernames2099980Args
	_args.Req = req
	var _result viewersapi.ViewerServiceGetuniqueviewernames2099980Result
	if err = p.c.Call(ctx, "getuniqueviewernames2099980", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Getuniqueviewernames80k20k(ctx context.Context, req *viewersapi.Request) (r *viewersapi.Response, err error) {
	var _args viewersapi.ViewerServiceGetuniqueviewernames80k20kArgs
	_args.Req = req
	var _result viewersapi.ViewerServiceGetuniqueviewernames80k20kResult
	if err = p.c.Call(ctx, "getuniqueviewernames80k20k", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Getuniqueviewernames20k80k(ctx context.Context, req *viewersapi.Request) (r *viewersapi.Response, err error) {
	var _args viewersapi.ViewerServiceGetuniqueviewernames20k80kArgs
	_args.Req = req
	var _result viewersapi.ViewerServiceGetuniqueviewernames20k80kResult
	if err = p.c.Call(ctx, "getuniqueviewernames20k80k", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
