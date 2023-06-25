package aliyun

import "github.com/aliyun/aliyun-oss-go-sdk/oss"

type AliyunOSS struct {
	Endpoint        string `json:"endpoint"`
	AccessKeyId     string `json:"access-key-id"`
	AccessKeySecret string `json:"access-key-secret"`
	BucketName      string `json:"bucket-name"`
	BasePath        string `json:"base_path"` // 一级目录 可以不传
}

func (a *AliyunOSS) NewClient() (client *oss.Client, err error) {
	return oss.New(a.Endpoint, a.AccessKeyId, a.AccessKeySecret)
}

func (a *AliyunOSS) NewBucket() (bucket *oss.Bucket, err error) {
	var client *oss.Client
	if client, err = a.NewClient(); err != nil {
		return nil, err
	}
	return client.Bucket(a.BucketName)
}
