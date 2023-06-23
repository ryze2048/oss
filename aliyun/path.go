package aliyun

import (
	"fmt"
	"github.com/ryze2048/oss/common"
	"time"
)

func (a *AliyunOSS) GetPath(info string) string {
	return fmt.Sprintf("%s/%s/%s%s", a.BasePath, time.Now().Format("2006-01-02"), common.GenerateRandomFilename(16), common.GetFileExtension(info))
}

func (a *AliyunOSS) GetResponseContentTypeSuffix(contentType string) string {
	return fmt.Sprintf("%s/%s/%s%s", a.BasePath, time.Now().Format("2006-01-02"), common.GenerateRandomFilename(16), common.GetFileExtensionToContentType(contentType))
}
