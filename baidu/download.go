package baidu

import (
	"github.com/baidubce/bce-sdk-go/services/bos"
	"github.com/baidubce/bce-sdk-go/services/bos/api"
	"github.com/ryze2048/oss/common"
	"io"
)

func (b *BaiduOSS) DownloadByte(key string) (data []byte, err error) {
	if key == common.NULL {
		return nil, common.KeyError
	}
	var client *bos.Client
	if client, err = b.NewClient(); err != nil {
		return nil, err
	}
	var result *api.GetObjectResult
	if result, err = client.BasicGetObject(b.BucketName, key); err != nil {
		return nil, err
	}
	return io.ReadAll(result.Body)
}

// GetTemporaryLink 获取临时的访问链接
// expired 单位秒
func (b *BaiduOSS) GetTemporaryLink(key string, expired int64) (path string, err error) {
	if key == common.NULL {
		return path, common.KeyError
	}
	if expired <= 0 {
		return "", common.ExpiredError
	}
	var client *bos.Client
	if client, err = b.NewClient(); err != nil {
		return "", err
	}
	return client.BasicGeneratePresignedUrl(b.BucketName, key, int(expired)), err
}
