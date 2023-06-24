package baidu

import (
	"fmt"
	"github.com/ryze2048/oss/common"
	"path/filepath"
	"time"
)

func (b *BaiduOSS) GetPath(info string) string {
	return filepath.Join(b.BasePath, time.Now().Format("2006-01-02"), fmt.Sprintf("%s%s", common.GenerateRandomFilename(16), common.GetFileExtension(info)))
}

func (b *BaiduOSS) GetResponseContentTypeSuffix(contentType string) string {
	return filepath.Join(b.BasePath, time.Now().Format("2006-01-02"), fmt.Sprintf("%s%s", common.GenerateRandomFilename(16), common.GetFileExtensionToContentType(contentType)))
}
