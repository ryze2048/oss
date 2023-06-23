package baidu

import "github.com/baidubce/bce-sdk-go/services/bos"

func (b *BaiduOSS) Delete(key string) (err error) {
	var client *bos.Client
	if client, err = b.NewClient(); err != nil {
		return err
	}
	return client.DeleteObject(b.BucketName, key)
}
