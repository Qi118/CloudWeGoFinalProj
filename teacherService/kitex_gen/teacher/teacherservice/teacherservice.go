// Code generated by Kitex v0.6.1. DO NOT EDIT.

package teacherservice

import (
	"context"
	teacher "github.com/Qi118/cloudwego/kitex_gen/teacher"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return teacherServiceServiceInfo
}

var teacherServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "TeacherService"
	handlerType := (*teacher.TeacherService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Query": kitex.NewMethodInfo(queryHandler, newTeacherServiceQueryArgs, newTeacherServiceQueryResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "teacher",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.6.1",
		Extra:           extra,
	}
	return svcInfo
}

func queryHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*teacher.TeacherServiceQueryArgs)
	realResult := result.(*teacher.TeacherServiceQueryResult)
	success, err := handler.(teacher.TeacherService).Query(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newTeacherServiceQueryArgs() interface{} {
	return teacher.NewTeacherServiceQueryArgs()
}

func newTeacherServiceQueryResult() interface{} {
	return teacher.NewTeacherServiceQueryResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Query(ctx context.Context, req *teacher.QueryReq) (r *teacher.Teacher, err error) {
	var _args teacher.TeacherServiceQueryArgs
	_args.Req = req
	var _result teacher.TeacherServiceQueryResult
	if err = p.c.Call(ctx, "Query", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
