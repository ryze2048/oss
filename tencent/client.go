package tencent

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
)

type TencentOSS struct {
	Bucket    string          `json:"bucket"`
	Region    string          `json:"region"`
	SecretID  string          `json:"secret-id"`
	SecretKey string          `json:"secret-key"`
	BaseURL   string          `json:"base-url"`
	BasePath  string          `json:"base-path"` // 一级目录
	Ctx       context.Context `json:"ctx"`
}

func (t *TencentOSS) NewClient() (client *cos.Client, err error) {
	var netUrl *url.URL
	if netUrl, err = url.Parse("https://" + t.Bucket + ".cos." + t.Region + ".myqcloud.com"); err != nil {
		return nil, err
	}
	var baseURL = &cos.BaseURL{BucketURL: netUrl}
	return cos.NewClient(baseURL, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  t.SecretID,
			SecretKey: t.SecretKey,
		},
	}), nil
}
