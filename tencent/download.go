package tencent

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"net/http"
	"net/url"
	"time"
)

// GetTemporaryLink 获取临时的访问链接
// expired 单位秒
func (t *TencentOSS) GetTemporaryLink(key string, expired int64) (path string, err error) {
	var client *cos.Client
	if client, err = t.NewClient(); err != nil {
		return "", err
	}
	var presignedURL *url.URL
	if presignedURL, err = client.Object.GetPresignedURL(t.Ctx, http.MethodGet, key, t.SecretID, t.SecretKey, time.Second*time.Duration(expired), nil); err != nil {
		return "", err
	}
	return presignedURL.String(), err
}

// DownloadByte 流式下载
func (t *TencentOSS) DownloadByte(key string) (data []byte, err error) {
	var client *cos.Client
	if client, err = t.NewClient(); err != nil {
		return nil, err
	}
	var response *cos.Response
	if response, err = client.Object.Get(context.Background(), key, nil); err != nil {
		panic(err)
	}

	defer func() {
		_ = response.Body.Close()
	}()
	return io.ReadAll(response.Body)
}
