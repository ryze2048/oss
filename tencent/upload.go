package tencent

import (
	"github.com/ryze2048/oss/common"
	"github.com/tencentyun/cos-go-sdk-v5"
	"mime/multipart"
	"net/http"
	"os"
)

// UploadFile 文件上传
// object 表单文件
func (t *TencentOSS) UploadFile(object *multipart.FileHeader) (path string, err error) {
	if object == nil {
		return path, common.ObjectError
	}
	var client *cos.Client
	if client, err = t.NewClient(); err != nil {
		return "", err
	}
	var file multipart.File
	if file, err = object.Open(); err != nil {
		return "", err
	}
	defer func() {
		_ = file.Close()
	}()
	path = t.GetPath(object.Filename)
	_, err = client.Object.Put(t.Ctx, path, file, nil)
	return path, err
}

// UploadLocal 本地文件上传
// localPath 本地文件路径
func (t *TencentOSS) UploadLocal(localPath string) (path string, err error) {
	if localPath == common.NULL {
		return path, common.LocalPathError
	}
	var client *cos.Client
	if client, err = t.NewClient(); err != nil {
		return "", err
	}
	var file *os.File
	if file, err = os.Open(localPath); err != nil {
		return "", err
	}
	defer func() {
		_ = file.Close()
	}()
	path = t.GetPath(localPath)
	_, err = client.Object.Put(t.Ctx, path, file, nil)
	return path, err
}

// UploadLink 根据图片链接进行上传
// link 链接地址 ｜ 需要一个包含后缀名的链接
func (t *TencentOSS) UploadLink(link string) (path string, err error) {
	if link == common.NULL {
		return path, common.LinkError
	}
	var response *http.Response
	if response, err = http.Get(link); err != nil {
		return "", err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	var client *cos.Client
	if client, err = t.NewClient(); err != nil {
		return "", err
	}
	path = t.GetResponseContentTypeSuffix(response.Header.Get("Content-Type"))
	_, err = client.Object.Put(t.Ctx, path, response.Body, nil)
	return path, err
}
