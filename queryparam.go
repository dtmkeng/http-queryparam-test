package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/astaxie/beego"
	bc "github.com/astaxie/beego/context"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo"
)

type route struct {
	method string
	path   string
}

var nullLogger *log.Logger
var loadTestHandler = false

type mockResponseWriter struct{}

func (m *mockResponseWriter) Header() (h http.Header) {
	return http.Header{}
}

func (m *mockResponseWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *mockResponseWriter) WriteString(s string) (n int, err error) {
	return len(s), nil
}
func (m *mockResponseWriter) WriteHeader(int) {}
func init() {
	runtime.GOMAXPROCS(1)

	// makes logging 'webscale' (ignores them)
	log.SetOutput(new(mockResponseWriter))
	nullLogger = log.New(new(mockResponseWriter), "", 0)

	initBeego()
	initGin()
}

func ginHandle(_ *gin.Context) {}

func ginHandleWrite(c *gin.Context) {
	io.WriteString(c.Writer, c.Query("name"))
}
func initGin() {
	gin.SetMode(gin.ReleaseMode)
}

func loadGinSingle(method, path string, handle gin.HandlerFunc) http.Handler {
	router := gin.New()
	router.Handle(method, path, handle)
	return router
}

// Beego
func beegoHandler(ctx *bc.Context) {}
func beegoHandlerWrite(ctx *bc.Context) {
	ctx.WriteString(ctx.Input.Query("name"))
}

func beegoHandlerTest(ctx *bc.Context) {
	ctx.WriteString(ctx.Request.RequestURI)
}

func initBeego() {
	beego.BConfig.RunMode = beego.PROD
	// beego
}

func loadBeegoSingle(method, path string, handler beego.FilterFunc) http.Handler {
	app := beego.NewControllerRegister()
	switch method {
	case "GET":
		app.Get(path, handler)
	default:
		// panic("Unknow HTTP method: " + method)
	}
	return app
}

// Echo
func echoHandler(c echo.Context) error {
	return nil
}
func echoHandlerWrite(c echo.Context) error {
	io.WriteString(c.Response(), c.QueryParam("name"))
	return nil
}

func loadEchoSingle(method, path string, h echo.HandlerFunc) http.Handler {
	e := echo.New()
	switch method {
	case "GET":
		e.GET(path, h)
	default:
		// panic("Unknow HTTP method: " + method)
	}
	return e
}
func main() {
	fmt.Println("Usage: go test -bench=. -timeout=20m")
	os.Exit(1)
}
