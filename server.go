package main

import (
	"github.com/yyquick/page-pipe/gzip"
	"github.com/yyquick/page-pipe/dispatcher"
	"github.com/yyquick/page-pipe/parser"
	"io"
	"net/http"
	"strings"
	"time"
)

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
	warmup int
}

func (w *gzipResponseWriter) Write(b []byte) (int, error) {
	if w.warmup == 0 {
		flen := 1024 - len(b)
		if flen > 0 {
			fb := make([]byte, 1024)
			for i := range fb {
				if i < flen {
					fb[i] = ' '
				} else {
					fb[i] = b[i - flen]
				}
			}
			b = fb
		}
		w.warmup = 1
	}
	ret, err := w.Writer.Write(b)
	w.Flush()
	return ret, err
}

func (w *gzipResponseWriter) Flush() {
	w.Writer.(*gzip.Writer).Flush()
	w.ResponseWriter.(http.Flusher).Flush()
}

func makeGzipHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			fn(w, r)
			return
		}
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Set("Content-Type", "text/html")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		gzr := gzipResponseWriter{Writer: gz, ResponseWriter: w, warmup: 0}
		fn(&gzr, r)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("This is a test.<br>"))
	// time.Sleep(1 * time.Second)
	// w.Write([]byte(r.URL.Host))
	// time.Sleep(1 * time.Second)
	// w.Write([]byte(r.URL.Path))
	// time.Sleep(1 * time.Second)
	// w.Write([]byte(r.URL.RawQuery))

	results := make(chan string)

	for i := range results {
		w.Write([]byte(i))
	}
}

func main() {
	http.ListenAndServe(":8080", makeGzipHandler(handler))
}