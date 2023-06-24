package obs

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"github.com/ryze2048/oss/common"
)

func (o *Obs) Delete(key string) (err error) {
	if key == common.NULL {
		return common.KeyError
	}
	var client *obs.ObsClient
	if client, err = o.NewClient(); err != nil {
		return err
	}
	input := &obs.DeleteObjectInput{}
	input.Bucket = o.Bucket
	input.Key = key
	_, err = client.DeleteObject(input)
	return err
}
