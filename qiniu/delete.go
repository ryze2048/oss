package qiniu

import (
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/ryze2048/oss/common"
)

func (q *Qiniu) Delete(key string, days int) (err error) {
	if key == common.NULL {
		return common.KeyError
	}
	var data = q.init()
	bucketManager := storage.NewBucketManager(data.Mac, data.Config)
	return bucketManager.DeleteAfterDays(q.Bucket, key, days)
}
