package web

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed build
var embeddedFS embed.FS

func AttachWeb(r *gin.Engine) {
	subFS, _ := fs.Sub(embeddedFS, "build")
	fileHandler := http.FileServer(http.FS(subFS))

	// Kind of a hack
	r.NoRoute(func(c *gin.Context) {
		fileHandler.ServeHTTP(c.Writer, c.Request)
	})
}
