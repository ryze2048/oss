package obs

import "github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"

func (o *Obs) Delete(key string) (err error) {
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
