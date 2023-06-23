package obs

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"mime/multipart"
	"net/http"
)

// UploadFile 文件上传
// object 表单文件
func (o *Obs) UploadFile(object *multipart.FileHeader) (path string, err error) {
	var client *obs.ObsClient
	if client, err = o.NewClient(); err != nil {
		return "", err
	}

	var file multipart.File
	if file, err = object.Open(); err != nil {
		return "", err
	}
	defer func() {
		_ = file.Close()
		_ = client.Close
	}()
	path = o.GetPath(object.Filename)
	input := &obs.PutObjectInput{
		PutObjectBasicInput: obs.PutObjectBasicInput{
			ObjectOperationInput: obs.ObjectOperationInput{
				Bucket: o.Bucket,
				Key:    path,
			},
			ContentType: object.Header.Get("Content-Type"),
		},
		Body: file,
	}
	if _, err = client.PutObject(input); err != nil {
		return "", err
	}
	return path, err
}

// UploadLocal 本地文件上传
// localPath 本地文件路径
func (o *Obs) UploadLocal(localPath string) (path string, err error) {
	var client *obs.ObsClient
	if client, err = o.NewClient(); err != nil {
		return "", err
	}
	defer func() {
		_ = client.Close
	}()
	path = o.GetPath(localPath)
	input := &obs.PutFileInput{
		PutObjectBasicInput: obs.PutObjectBasicInput{
			ObjectOperationInput: obs.ObjectOperationInput{
				Bucket: o.Bucket,
				Key:    path,
			},
		},
		SourceFile: localPath,
	}
	if _, err = client.PutFile(input); err != nil {
		return "", err
	}
	return path, nil
}

// UploadLink 根据图片链接进行上传
// link 链接地址 ｜ 需要一个包含后缀名的链接
func (o *Obs) UploadLink(link string) (path string, err error) {
	var client *obs.ObsClient
	if client, err = o.NewClient(); err != nil {
		return "", err
	}
	var response *http.Response
	if response, err = http.Get(link); err != nil {
		return "", err
	}
	defer func() {
		_ = response.Body.Close()
		_ = client.Close
	}()

	path = o.GetResponseContentTypeSuffix(response.Header.Get("Content-Type"))
	input := &obs.PutObjectInput{
		PutObjectBasicInput: obs.PutObjectBasicInput{
			ObjectOperationInput: obs.ObjectOperationInput{
				Bucket: o.Bucket,
				Key:    path,
			},
			ContentType: response.Header.Get("Content-Type"),
		},
		Body: response.Body,
	}

	if _, err = client.PutObject(input); err != nil {
		return "", err
	}
	return path, err
}
