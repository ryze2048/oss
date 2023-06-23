package qiniu

import "github.com/qiniu/go-sdk/v7/storage"

func (q *Qiniu) Delete(key string, days int) (err error) {
	var data = q.init()
	bucketManager := storage.NewBucketManager(data.Mac, data.Config)
	return bucketManager.DeleteAfterDays(q.Bucket, key, days)
}
