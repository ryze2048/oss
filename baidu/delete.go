package baidu

import (
	"github.com/baidubce/bce-sdk-go/services/bos"
	"github.com/ryze2048/oss/common"
)

func (b *BaiduOSS) Delete(key string) (err error) {
	if key == common.NULL {
		return common.KeyError
	}
	var client *bos.Client
	if client, err = b.NewClient(); err != nil {
		return err
	}
	return client.DeleteObject(b.BucketName, key)
}
