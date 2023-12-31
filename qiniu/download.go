package qiniu

import (
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/ryze2048/oss/common"
	"time"
)

// GetTemporaryLink 获取临时的访问链接
// expired 单位秒
func (q *Qiniu) GetTemporaryLink(key string, expired int64) (path string, err error) {
	if key == common.NULL {
		return path, common.KeyError
	}
	var mac = q.NewClient()
	deadline := time.Now().Add(time.Second * time.Duration(expired)).Unix()
	return storage.MakePrivateURL(mac, q.Domain, key, deadline), nil
}
