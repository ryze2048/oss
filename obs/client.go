package obs

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
)

type Obs struct {
	BasePath  string `json:"base_path"`
	Bucket    string `json:"bucket"`
	Endpoint  string `json:"endpoint"`
	AccessKey string `json:"access-key"`
	SecretKey string `json:"secret-key"`
}

func (o *Obs) NewClient() (client *obs.ObsClient, err error) {
	return obs.New(o.AccessKey, o.SecretKey, o.Endpoint)
}
