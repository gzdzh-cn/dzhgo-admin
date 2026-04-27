package internal

import "github.com/gzdzh-cn/dzhcore"

var (
	Version = "v1.2.3"
)

func init() {
	dzhcore.SetVersions("dzhgo", Version)
}
