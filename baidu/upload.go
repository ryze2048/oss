package baidu

import (
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/bos"
	"io/fs"
	"mime/multipart"
	"net/http"
	"os"
)

// UploadFile 文件上传
// object 表单文件
func (b *BaiduOSS) UploadFile(object *multipart.FileHeader) (path string, err error) {
	var client *bos.Client
	if client, err = b.NewClient(); err != nil {
		return "", err
	}
	var file multipart.File
	if file, err = object.Open(); err != nil {
		return "", err
	}
	defer func() {
		_ = file.Close()
	}()
	path = b.GetPath(object.Filename)
	var body *bce.Body
	if body, err = bce.NewBodyFromSizedReader(file, object.Size); err != nil {
		return "", err
	}
	_, err = client.PutObject(b.BucketName, path, body, nil)
	return path, err
}

// UploadLocal 本地文件上传
// localPath 本地文件路径
func (b *BaiduOSS) UploadLocal(localPath string) (path string, err error) {
	var file *os.File
	if file, err = os.Open(localPath); err != nil {
		return "", err
	}
	defer func() {
		_ = file.Close()
	}()
	var client *bos.Client
	if client, err = b.NewClient(); err != nil {
		return "", err
	}
	var fileInfo fs.FileInfo
	if fileInfo, err = file.Stat(); err != nil {
		return "", err
	}
	path = b.GetPath(localPath)
	var body *bce.Body
	if body, err = bce.NewBodyFromSizedReader(file, fileInfo.Size()); err != nil {
		return "", err
	}
	_, err = client.PutObject(b.BucketName, path, body, nil)
	return path, err
}

// UploadLink 根据图片链接进行上传
// link 链接地址 ｜ 需要一个包含后缀名的链接
func (b *BaiduOSS) UploadLink(link string) (path string, err error) {
	var response *http.Response
	if response, err = http.Get(link); err != nil {
		return "", err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	var client *bos.Client
	if client, err = b.NewClient(); err != nil {
		return "", err
	}
	path = b.GetResponseContentTypeSuffix(response.Header.Get("Content-Type"))
	var body *bce.Body
	if body, err = bce.NewBodyFromSizedReader(response.Body, response.ContentLength); err != nil {
		return "", err
	}
	_, err = client.PutObject(b.BucketName, path, body, nil)
	return path, err
}
