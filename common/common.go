package common

import (
	"crypto/rand"
	"encoding/hex"
	"path"
)

// GetFileExtension 获取文件后缀名
func GetFileExtension(fileName string) string {
	return path.Ext(fileName)
}

// GenerateRandomFilename 生层随机的文件名
func GenerateRandomFilename(len int) string {
	randomBytes := make([]byte, len)
	if _, err := rand.Read(randomBytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(randomBytes)
}

// GetFileExtensionToContentType 根据相应头获取文件后缀名
func GetFileExtensionToContentType(contentType string) string {
	switch contentType {
	case "text/plain":
		return ".txt"
	case "text/html":
		return ".html"
	case "text/css":
		return ".css"
	case "text/javascript":
		return ".js"
	case "application/json":
		return ".json"
	case "image/jpeg":
		return ".jpg"
	case "image/png":
		return ".png"
	case "image/gif":
		return ".gif"
	case "image/svg+xml":
		return ".svg"
	case "audio/mpeg":
		return ".mp3"
	case "audio/wav":
		return ".wav"
	case "audio/midi":
		return ".midi"
	case "video/mp4":
		return ".mp4"
	case "video/quicktime":
		return ".mov"
	case "video/webm":
		return ".webm"
	case "application/pdf":
		return ".pdf"
	case "application/zip":
		return ".zip"
	case "application/octet-stream":
		return ".bin"
	case "application/x-www-form-urlencoded":
		return ".html"
	default:
		return ".bin"
	}
}
