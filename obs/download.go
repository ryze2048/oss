package obs

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"github.com/ryze2048/oss/common"
	"io"
)

// DownloadByte 流式下载
func (o *Obs) DownloadByte(key string) (byteInfo []byte, err error) {
	if key == common.NULL {
		return nil, common.KeyError
	}
	var client *obs.ObsClient
	if client, err = o.NewClient(); err != nil {
		return nil, err
	}

	input := &obs.GetObjectInput{}
	input.Bucket = o.Bucket
	input.Key = key

	var output *obs.GetObjectOutput
	if output, err = client.GetObject(input); err != nil {
		return nil, err
	}
	defer func() {
		_ = output.Body.Close()
	}()
	if byteInfo, err = io.ReadAll(output.Body); err != nil {
		return nil, err
	}
	return byteInfo, nil
}

// GetTemporaryLink 获取临时的访问链接
// expired 单位秒
func (o *Obs) GetTemporaryLink(key string, expires int) (path string, err error) {
	if key == "" {
		return path, common.KeyError
	}
	if expires <= 0 {
		return path, common.ExpiredError
	}
	var client *obs.ObsClient
	if client, err = o.NewClient(); err != nil {
		return "", err
	}

	putObjectInput := &obs.CreateSignedUrlInput{
		Method:  obs.HttpMethodGet,
		Bucket:  o.Bucket,
		Key:     key,
		Expires: expires,
	}
	var output *obs.CreateSignedUrlOutput
	if output, err = client.CreateSignedUrl(putObjectInput); err != nil {
		return "", err
	}
	return output.SignedUrl, err
}
