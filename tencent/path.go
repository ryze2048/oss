package tencent

import (
	"fmt"
	"github.com/ryze2048/oss/common"
	"path/filepath"
	"time"
)

func (t *TencentOSS) GetPath(info string) string {
	return filepath.Join(t.BasePath, time.Now().Format("2006-01-02"), fmt.Sprintf("%s%s", common.GenerateRandomFilename(16), common.GetFileExtension(info)))
}
func (t *TencentOSS) GetResponseContentTypeSuffix(contentType string) string {
	return filepath.Join(t.BasePath, time.Now().Format("2006-01-02"), fmt.Sprintf("%s%s", common.GenerateRandomFilename(16), common.GetFileExtensionToContentType(contentType)))
}

// GetBasePath
// suffix .png|jpg....
func (t *TencentOSS) GetBasePath(suffix string) (path string) {
	return filepath.Join(t.BasePath, time.Now().Format("2006-01-02"), fmt.Sprintf("%s%s", common.GenerateRandomFilename(16), suffix))
}
