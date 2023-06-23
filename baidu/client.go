package baidu

import "github.com/baidubce/bce-sdk-go/services/bos"

type BaiduOSS struct {
	AccessKey        string `json:"access-key"`
	SecretKey        string `json:"secret_key"`
	Endpoint         string `json:"endpoint"`
	BucketName       string `json:"bucket-name"`
	RedirectDisabled bool   `json:"redirect_disabled"`
	BasePath         string `json:"base_path"`
}

// Secret Key	 94958f43103148aab2b4f9a66436d1bf
// Access Key    ALTAK7EsV3XJpFYzG0phF1es0g

func (b *BaiduOSS) NewClient() (*bos.Client, error) {
	clientConfig := bos.BosClientConfiguration{
		Ak:               b.AccessKey,
		Sk:               b.SecretKey,
		Endpoint:         b.Endpoint,
		RedirectDisabled: b.RedirectDisabled,
	}
	return bos.NewClientWithConfig(&clientConfig)
}
