package main

import (
	"context"
	teacher "github.com/Qi118/cloudwego/kitex_gen/teacher"
)

// TeacherServiceImpl implements the last service interface defined in the IDL.
type TeacherServiceImpl struct{}

// Query implements the TeacherServiceImpl interface.
func (s *TeacherServiceImpl) Query(ctx context.Context, req *teacher.QueryReq) (resp *teacher.Teacher, err error) {
	// TODO: Your code here...
	resp = &teacher.Teacher{
		Name: "Teacher Name: " + req.Name,
	}
	return
}
