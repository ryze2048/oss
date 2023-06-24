package aliyun

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/ryze2048/oss/common"
	"io"
)

// DownloadByte 流式下载
func (a *AliyunOSS) DownloadByte(key string) (data []byte, err error) {
	if key == common.NULL {
		return nil, common.KeyError
	}
	var bucket *oss.Bucket
	if bucket, err = a.NewBucket(); err != nil {
		return data, err
	}
	var body io.ReadCloser
	if body, err = bucket.GetObject(key); err != nil {
		return nil, err
	}
	defer func() {
		_ = body.Close()
	}()
	return io.ReadAll(body)
}

// GetTemporaryLink 获取临时的访问链接
// expired 单位秒
func (a *AliyunOSS) GetTemporaryLink(key string, expired int64) (path string, err error) {
	if key == common.NULL {
		return "", common.KeyError
	}
	if expired <= 0 {
		return "", common.ExpiredError
	}
	var bucket *oss.Bucket
	if bucket, err = a.NewBucket(); err != nil {
		return path, err
	}
	return bucket.SignURL(key, oss.HTTPGet, expired)
}
