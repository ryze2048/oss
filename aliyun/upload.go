package aliyun

import (
	"bytes"
	"encoding/base64"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/ryze2048/oss/common"
	"mime/multipart"
	"net/http"
	"os"
)

// UploadFile 文件上传
// object 表单文件
func (a *AliyunOSS) UploadFile(object *multipart.FileHeader) (path string, err error) {
	if object == nil {
		return "", common.ObjectError
	}
	var bucket *oss.Bucket
	if bucket, err = a.NewBucket(); err != nil {
		return "", err
	}
	var file multipart.File
	if file, err = object.Open(); err != nil {
		return "", err
	}
	defer func() {
		_ = file.Close()
	}()
	path = a.GetPath(object.Filename)
	if err = bucket.PutObject(path, file); err != nil {
		return "", err
	}
	return path, nil
}

// UploadLocal 本地文件上传
// localPath 本地文件路径
func (a *AliyunOSS) UploadLocal(localPath string) (path string, err error) {
	if localPath == common.NULL {
		return "", common.KeyError
	}
	var bucket *oss.Bucket
	if bucket, err = a.NewBucket(); err != nil {
		return "", err
	}
	var file *os.File
	if file, err = os.Open(localPath); err != nil {
		return "", err
	}
	defer func() {
		_ = file.Close()
	}()
	path = a.GetPath(file.Name())
	if err = bucket.PutObject(path, file); err != nil {
		return "", err
	}
	return path, nil
}

// UploadLink 根据图片链接进行上传
// link 链接地址 ｜ 需要一个包含后缀名的链接
func (a *AliyunOSS) UploadLink(link string) (path string, err error) {
	if link == common.NULL {
		return "", common.LinkError
	}
	var response *http.Response
	if response, err = http.Get(link); err != nil {
		return "", err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	var bucket *oss.Bucket
	if bucket, err = a.NewBucket(); err != nil {
		return "", err
	}
	path = a.GetResponseContentTypeSuffix(response.Header.Get("Content-Type"))
	if err = bucket.PutObject(path, response.Body); err != nil {
		return "", err
	}
	return path, err
}

func (a *AliyunOSS) UploadBase64(base *string, suffix string) (path string, err error) {
	if base == nil || suffix == common.NULL {
		return "", common.Base64Error
	}

	var byteInfo = make([]byte, 0)
	if byteInfo, err = base64.StdEncoding.DecodeString(*base); err != nil {
		return "", err
	}
	var reader = bytes.NewReader(byteInfo)
	path = a.GetBasePath(suffix)
	var bucket *oss.Bucket
	if bucket, err = a.NewBucket(); err != nil {
		return "", err
	}
	return path, bucket.PutObject(path, reader)
}
