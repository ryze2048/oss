package baidu

import (
	"github.com/baidubce/bce-sdk-go/services/bos"
	"github.com/baidubce/bce-sdk-go/services/bos/api"
	"io"
)

func (b *BaiduOSS) DownloadByte(key string) (data []byte, err error) {
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
	var client *bos.Client
	if client, err = b.NewClient(); err != nil {
		return "", err
	}
	return client.BasicGeneratePresignedUrl(b.BucketName, key, int(expired)), err
}
