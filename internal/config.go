package internal

import "github.com/gzdzh-cn/dzhcore"

var (
	Version = "v1.2.0"
)

func init() {
	dzhcore.SetVersions("dzhgo", Version)
}
