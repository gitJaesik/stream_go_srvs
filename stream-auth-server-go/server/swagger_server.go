package server

import (
	"mime"
	"net/http"

	"github.com/rakyll/statik/fs"

	// Statik files
	_ "github.com/gitJaesik/stream_go_srvs/stream-auth-server-go/statik"
)

func GetOpenAPIHandler() http.Handler {
	mime.AddExtensionType(".svg", "image/svg+xml")

	statikFS, err := fs.New()
	if err != nil {
		panic("creating OpenAPI filesystem: " + err.Error())
	}

	return http.FileServer(statikFS)
}
