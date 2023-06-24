package qiniu

import (
	"fmt"
	"github.com/ryze2048/oss/common"
	"path/filepath"
	"time"
)

func (q *Qiniu) GetPath(info string) string {
	return filepath.Join(q.BasePath, time.Now().Format("2006-01-02"), fmt.Sprintf("%s%s", common.GenerateRandomFilename(16), common.GetFileExtension(info)))
}
func (q *Qiniu) GetResponseContentTypeSuffix(contentType string) string {
	return filepath.Join(q.BasePath, time.Now().Format("2006-01-02"), fmt.Sprintf("%s%s", common.GenerateRandomFilename(16), common.GetFileExtensionToContentType(contentType)))
}
