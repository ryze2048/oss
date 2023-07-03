package aliyun

import (
	"fmt"
	"github.com/ryze2048/oss/common"
	"path/filepath"
	"time"
)

func (a *AliyunOSS) GetPath(info string) (path string) {
	return filepath.Join(a.BasePath, time.Now().Format("2006-01-02"), fmt.Sprintf("%s%s", common.GenerateRandomFilename(16), common.GetFileExtension(info)))
}

func (a *AliyunOSS) GetResponseContentTypeSuffix(contentType string) string {
	return filepath.Join(a.BasePath, time.Now().Format("2006-01-02"), fmt.Sprintf("%s%s", common.GenerateRandomFilename(16), common.GetFileExtensionToContentType(contentType)))
}

// GetBasePath
// suffix .png|jpg....
func (a *AliyunOSS) GetBasePath(suffix string) (path string) {
	return filepath.Join(a.BasePath, time.Now().Format("2006-01-02"), fmt.Sprintf("%s%s", common.GenerateRandomFilename(16), suffix))
}
