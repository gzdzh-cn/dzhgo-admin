package internal

import "github.com/gzdzh-cn/dzhcore"

var (
	Version = "v1.2.2"
)

func init() {
	dzhcore.SetVersions("dzhgo", Version)
}
