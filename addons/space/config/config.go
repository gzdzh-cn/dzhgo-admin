package config

import (
	"dzhgo/addons/file_upload"
	"github.com/gzdzh-cn/dzhcore"
)

func init() {
	dzhcore.SetVersions("space", file_upload.Version)
}
