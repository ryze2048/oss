package aliyun

import "github.com/aliyun/aliyun-oss-go-sdk/oss"

func (a *AliyunOSS) Delete(key string) (err error) {
	var bucket *oss.Bucket
	if bucket, err = a.NewBucket(); err != nil {
		return err
	}
	return bucket.DeleteObject(key)
}
