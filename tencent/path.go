package tencent

import (
	"fmt"
	"github.com/ryze2048/oss/common"
	"time"
)

func (t *TencentOSS) GetPath(info string) string {
	return fmt.Sprintf("%s/%s/%s%s", t.BaseURL, time.Now().Format("2006-01-02"), common.GenerateRandomFilename(16), common.GetFileExtension(info))
}
func (t *TencentOSS) GetResponseContentTypeSuffix(contentType string) string {
	return fmt.Sprintf("%s/%s/%s%s", t.BaseURL, time.Now().Format("2006-01-02"), common.GenerateRandomFilename(16), common.GetFileExtensionToContentType(contentType))
}
