package qiniu

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/ryze2048/oss/common"
)

type Qiniu struct {
	Zone          string          `json:"zone"`            // 存储区域
	Bucket        string          `json:"bucket"`          // 空间名称
	ImgPath       string          `json:"img-path"`        // CDN加速域名
	UseHTTPS      bool            `json:"use-https"`       // 是否使用https
	AccessKey     string          `json:"access-key"`      // 秘钥AK
	SecretKey     string          `json:"secret-key"`      // 秘钥SK
	UseCdnDomains bool            `json:"use-cdn-domains"` // 上传是否使用CDN上传加速
	BasePath      string          `json:"base_path"`       // 第一层目录
	Domain        string          `json:"domain"`          // 域名
	Ctx           context.Context `json:"ctx"`
}

type Data struct {
	Mac          *qbox.Mac
	PutPolicy    storage.PutPolicy
	Token        string
	Config       *storage.Config
	FormUploader *storage.FormUploader
	PutRet       storage.PutRet
	PutExtra     storage.PutExtra
}

func (q *Qiniu) NewClient() *qbox.Mac {
	return qbox.NewMac(q.AccessKey, q.SecretKey)
}

func (q *Qiniu) config() *storage.Config {
	cfg := storage.Config{
		UseHTTPS:      q.UseHTTPS,
		UseCdnDomains: q.UseCdnDomains,
	}
	switch q.Zone {
	case common.Huadong:
		cfg.Zone = &storage.ZoneHuadong
	case common.Huabei:
		cfg.Zone = &storage.ZoneHuabei
	case common.Huanan:
		cfg.Zone = &storage.ZoneHuanan
	case common.Beimei:
		cfg.Zone = &storage.ZoneBeimei
	case common.Xinjiapo:
		cfg.Zone = &storage.ZoneXinjiapo
	}
	return &cfg
}

func (q *Qiniu) init() *Data {
	data := &Data{
		Mac:       q.NewClient(),
		PutPolicy: storage.PutPolicy{Scope: q.Bucket},
		Config:    q.config(),
		PutRet:    storage.PutRet{},
		PutExtra: storage.PutExtra{
			Params: map[string]string{
				"x:name": "github logo",
			},
		},
	}
	data.Token = data.PutPolicy.UploadToken(data.Mac)
	data.FormUploader = storage.NewFormUploader(data.Config)
	return data
}
