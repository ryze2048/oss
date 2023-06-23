package qiniu

import (
	"context"
	"fmt"
	"testing"
)

func TestQiniu_UploadLocal(t *testing.T) {
	var err error
	var ctx = context.Background()
	q := Qiniu{
		Ctx: ctx,
	}
	var path string
	path, err = q.UploadLocal("dAmZNXqEdXTvUvSlyrPTLtvCmjCkTp.png")
	if err != nil {
		t.Error("err --> ", err)
	}
	fmt.Println(path)
}

func TestQiniu_UploadLink(t *testing.T) {
	var err error
	var ctx = context.Background()
	q := Qiniu{
		Ctx: ctx,
	}
	var path string
	path, err = q.UploadLink("https://bce.bdstatic.com/login/banner_20230518.png")
	if err != nil {
		t.Error("err --> ", err)
	}
	fmt.Println(path)
}

func TestQiniu_GetTemporaryLink(t *testing.T) {
	var err error
	var ctx = context.Background()
	q := Qiniu{
		Ctx: ctx,
	}
	var path string
	path, err = q.GetTemporaryLink("test/2023-06-20/f7d3f5a8e8008608d5430c331fb0be92.png", 360)
	if err != nil {
		t.Error("err --> ", err)
	}
	fmt.Println(path)
}
