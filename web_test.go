package web_test

import (
  "github.com/cypriss/mars_web"
  . "launchpad.net/gocheck"
  "testing"
  "net/http"
  "net/http/httptest"
  "fmt"
  "strings"
)

//
// This file is the "driver" for the test suite. We're using gocheck.
// This file will contain helpers and general things the rest of the suite needs
//

//
// gocheck: hook into "go test"
//
func Test(t *testing.T) { TestingT(t) }

// Make a testing request
func newTestRequest(method, path string) (*httptest.ResponseRecorder, *http.Request) {
  request, _ := http.NewRequest(method, path, nil)
  recorder := httptest.NewRecorder()

  return recorder, request
}

func assertResponse(c *C, rr *httptest.ResponseRecorder, body string, code int) {
  c.Assert(strings.TrimSpace(string(rr.Body.Bytes())), Equals, body)
  c.Assert(rr.Code, Equals, code)
}

//
// Some default contexts and possible error handlers / actions
//
type Context struct {}

type AdminContext struct {
  *Context
}

type ApiContext struct {
  *Context
}

type SiteContext struct {
  *Context
}

type TicketsContext struct {
  *AdminContext
}

func (c *Context) ErrorHandler(w *web.ResponseWriter, r *web.Request, err interface{}) {
  fmt.Fprintf(w, "My Error")
}

func (c *Context) ErrorAction(w *web.ResponseWriter, r *web.Request) {
  var x, y int
  fmt.Fprintln(w, x / y)
}

func (c *AdminContext) ErrorHandler(w *web.ResponseWriter, r *web.Request, err interface{}) {
  fmt.Fprintf(w, "Admin Error")
}

func (c *AdminContext) ErrorAction(w *web.ResponseWriter, r *web.Request) {
  var x, y int
  fmt.Fprintln(w, x / y)
}

func (c *ApiContext) ErrorHandler(w *web.ResponseWriter, r *web.Request, err interface{}) {
  fmt.Fprintf(w, "Api Error")
}

func (c *ApiContext) ErrorAction(w *web.ResponseWriter, r *web.Request) {
  var x, y int
  fmt.Fprintln(w, x / y)
}

func (c *TicketsContext) ErrorAction(w *web.ResponseWriter, r *web.Request) {
  var x, y int
  fmt.Fprintln(w, x / y)
}