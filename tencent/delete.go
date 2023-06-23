package tencent

import "github.com/tencentyun/cos-go-sdk-v5"

func (t *TencentOSS) Delete(key string) (err error) {
	var client *cos.Client
	if client, err = t.NewClient(); err != nil {
		return err
	}
	_, err = client.Object.Delete(t.Ctx, key)
	return err
}
