package web

import (
	"net/http"
	"io"
)
type statusCoder interface {
	StatusCode() int
}
type ResponseWriter interface {
	http.ResponseWriter
	statusCoder
}
type Response interface {
	response
	statusCoder
}
type response interface {
	http.ResponseWriter
	http.Flusher
	http.Hijacker
	http.CloseNotifier
	ReadFrom(src io.Reader) (n int64, err error)
//	Header() http.Header // required by http.ResponseWriter
//	WriteHeader(code int) // required by http.ResponseWriter
//	Write(data []byte) (n int, err error) // required by http.ResponseWriter
	WriteString(data string) (n int, err error)
//	Flush() // required by http.Flusher
//	Hijack() (rwc net.Conn, buf *bufio.ReadWriter, err error) // required by http.Hijacker
//	CloseNotify() <-chan bool // required by http.Hijacker
}
type AppResponseWriter struct {
	response
	statusCode int
	written    bool
}

// Don't need this yet because we get it for free:
func (w *AppResponseWriter) Write(data []byte) (n int, err error) {
	if !w.written {
		w.statusCode = http.StatusOK
		w.written = true
	}
	return w.response.Write(data)
}

func (w *AppResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.written = true
	w.response.WriteHeader(statusCode)
}

func (w *AppResponseWriter) StatusCode() int {
	return w.statusCode
}
