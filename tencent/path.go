package tencent

import (
	"fmt"
	"github.com/ryze2048/oss/common"
	"path/filepath"
	"time"
)

func (t *TencentOSS) GetPath(info string) string {
	return filepath.Join(t.BaseURL, time.Now().Format("2006-01-02"), fmt.Sprintf("%s%s", common.GenerateRandomFilename(16), common.GetFileExtension(info)))
}
func (t *TencentOSS) GetResponseContentTypeSuffix(contentType string) string {
	return filepath.Join(t.BaseURL, time.Now().Format("2006-01-02"), fmt.Sprintf("%s%s", common.GenerateRandomFilename(16), common.GetFileExtensionToContentType(contentType)))
}
