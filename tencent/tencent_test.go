package tencent

import (
	"context"
	"fmt"
	"testing"
)

func TestTencentOSS_NewClient(t *testing.T) {
	tencentOss := TencentOSS{}
	client, err := tencentOss.NewClient()
	if err != nil {
		t.Error("err --> ", err)
	}
	fmt.Println(client)
}

func TestTencentOSS_UploadLocal(t *testing.T) {
	var ctx = context.Background()
	tencentOss := TencentOSS{
		Bucket:    "",
		Region:    "ap-guangzhou",
		SecretID:  "",
		SecretKey: "",
		BaseURL:   "test",
		Ctx:       ctx,
	}
	path, err := tencentOss.UploadLocal("dAmZNXqEdXTvUvSlyrPTLtvCmjCkTp.png")
	if err != nil {
		t.Error("err --> ", err)
	}
	fmt.Println(path)
}

func TestTencentOSS_GetTemporaryLink(t *testing.T) {
	var ctx = context.Background()
	tencentOss := TencentOSS{
		Bucket:    "",
		Region:    "ap-guangzhou",
		SecretID:  "",
		SecretKey: "",
		BaseURL:   "test",
		Ctx:       ctx,
	}
	var path string
	var err error
	if path, err = tencentOss.GetTemporaryLink("test/2023-06-19/db4778ba063313cfcc45a1e96378881c.png", 10); err != nil {
		t.Error("err --> ", err)
	}
	fmt.Println(path)
}

func TestTencentOSS_UploadFile(t *testing.T) {
	var ctx = context.Background()
	tencentOss := TencentOSS{
		Bucket:    "",
		Region:    "ap-guangzhou",
		SecretID:  "",
		SecretKey: "",
		BaseURL:   "test",
		Ctx:       ctx,
	}
	var path string
	var err error
	if path, err = tencentOss.UploadLink(`https://fanyiapp.cdn.bcebos.com/cms/image/7b9655bb7a32a6c1a472c3591428855d.jpg`); err != nil {
		t.Error("err --> ", err)
	}
	fmt.Println(path)

}
