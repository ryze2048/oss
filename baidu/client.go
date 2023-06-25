package baidu

import "github.com/baidubce/bce-sdk-go/services/bos"

type BaiduOSS struct {
	AccessKey        string `json:"access-key"`
	SecretKey        string `json:"secret_key"`
	Endpoint         string `json:"endpoint"`
	BucketName       string `json:"bucket-name"`
	RedirectDisabled bool   `json:"redirect_disabled"` // 是否关闭重定向，true关闭。
	BasePath         string `json:"base_path"`         // 一级目录
}

func (b *BaiduOSS) NewClient() (*bos.Client, error) {
	clientConfig := bos.BosClientConfiguration{
		Ak:               b.AccessKey,
		Sk:               b.SecretKey,
		Endpoint:         b.Endpoint,
		RedirectDisabled: b.RedirectDisabled,
	}
	return bos.NewClientWithConfig(&clientConfig)
}
