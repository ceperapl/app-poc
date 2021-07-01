package dist

import (
	"embed"
)

//go:embed index.html assets
var staticFiles embed.FS

func GetEmbedFS() embed.FS {
	return staticFiles
}
