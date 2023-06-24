package qiniu

import (
	"github.com/ryze2048/oss/common"
	"mime/multipart"
	"net/http"
)

// UploadFile 文件上传
// object 表单文件
func (q *Qiniu) UploadFile(object *multipart.FileHeader) (path string, err error) {
	if path == common.NULL {
		return path, common.KeyError
	}
	var data = q.init()
	var file multipart.File
	if file, err = object.Open(); err != nil {
		return "", err
	}
	defer func() {
		_ = file.Close()
	}()
	path = q.GetPath(object.Filename)
	if err = data.FormUploader.Put(q.Ctx, &data.PutRet, data.Token, path, file, object.Size, &data.PutExtra); err != nil {
		return "", err
	}
	return data.PutRet.Key, nil
}

// UploadLocal 本地文件上传
// localPath 本地文件路径
func (q *Qiniu) UploadLocal(localPath string) (path string, err error) {
	if localPath == common.NULL {
		return path, common.LocalPathError
	}
	var data = q.init()
	path = q.GetPath(localPath)
	if err = data.FormUploader.PutFile(q.Ctx, &data.PutRet, data.Token, path, localPath, &data.PutExtra); err != nil {
		return "", err
	}
	return data.PutRet.Key, nil
}

// UploadLink 根据图片链接进行上传
// link 链接地址 ｜ 需要一个包含后缀名的链接
func (q *Qiniu) UploadLink(link string) (path string, err error) {
	if link == common.NULL {
		return path, common.LinkError
	}
	var data = q.init()
	var response *http.Response
	if response, err = http.Get(link); err != nil {
		return "", err
	}
	defer func() {
		_ = response.Body.Close()
	}()
	path = q.GetResponseContentTypeSuffix(response.Header.Get("Content-Type"))
	if err = data.FormUploader.Put(q.Ctx, &data.PutRet, data.Token, path, response.Body, response.ContentLength, &data.PutExtra); err != nil {
		return "", err
	}
	return data.PutRet.Key, nil
}
