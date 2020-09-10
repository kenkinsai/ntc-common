package middleware

import (
	"io"
	"net/http"

	"github.com/kenkinsai/ntc-common/config"
)

var cfg = config.GetConfig()

type bodyDumpResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w *bodyDumpResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}
