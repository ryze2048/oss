package obs

import (
	"fmt"
	"github.com/ryze2048/oss/common"
	"path/filepath"
	"time"
)

func (o *Obs) GetPath(info string) string {
	return filepath.Join(o.BasePath, time.Now().Format("2006-01-02"), fmt.Sprintf("%s%s", common.GenerateRandomFilename(16), common.GetFileExtension(info)))
}

func (o *Obs) GetResponseContentTypeSuffix(contentType string) string {
	return filepath.Join(o.BasePath, time.Now().Format("2006-01-02"), fmt.Sprintf("%s%s", common.GenerateRandomFilename(16), common.GetFileExtensionToContentType(contentType)))
}

// GetBasePath
// suffix .png|jpg....
func (o *Obs) GetBasePath(suffix string) (path string) {
	return filepath.Join(o.BasePath, time.Now().Format("2006-01-02"), fmt.Sprintf("%s%s", common.GenerateRandomFilename(16), suffix))
}
