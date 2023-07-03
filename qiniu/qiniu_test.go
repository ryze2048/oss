package qiniu

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
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
	path, err = q.GetTemporaryLink("test/2023-06-25/f54f24d70d3340e0c1868729e8c15d05.png", 360)
	if err != nil {
		t.Error("err --> ", err)
	}
	fmt.Println(path)
}

func TestQiniu_UploadBase64(t *testing.T) {
	filePath := "dAmZNXqEdXTvUvSlyrPTLtvCmjCkTp.png" // 替换为你的本地文件路径

	// 读取本地文件
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 将文件内容转换为 Base64 字符串
	base64String := base64.StdEncoding.EncodeToString(fileBytes)
	fmt.Println("Base64 encoded string:", base64String)
	var ctx = context.Background()
	client := Qiniu{
		Ctx: ctx,
	}
	path, err := client.UploadBase64(&base64String, ".png")
	if err != nil {
		t.Error("upload base err --> ", err)
	}

	fmt.Println(path)
}
