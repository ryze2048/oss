package aliyun

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/ryze2048/oss/common"
)

func (a *AliyunOSS) Delete(key string) (err error) {
	if key == common.NULL {
		return common.KeyError
	}
	var bucket *oss.Bucket
	if bucket, err = a.NewBucket(); err != nil {
		return err
	}
	return bucket.DeleteObject(key)
}
