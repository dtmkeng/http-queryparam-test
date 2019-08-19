package main

import (
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strings"
	"testing"
)

var (
	ginQuery   http.Handler
	echoQuery  http.Handler
	beegoQuery http.Handler
	bmuxQuery  http.Handler
)
var benchRe *regexp.Regexp

func isTested(name string) bool {
	if benchRe == nil {
		// Get -test.bench flag value (not accessible via flag package)
		bench := ""
		for _, arg := range os.Args {
			if strings.HasPrefix(arg, "-test.bench=") {
				// ignore the benchmark name after an underscore
				bench = strings.SplitN(arg[12:], "_", 2)[0]
				break
			}
		}

		// Compile RegExp to match Benchmark names
		var err error
		benchRe, err = regexp.Compile(bench)
		if err != nil {
			panic(err.Error())
		}
	}
	return benchRe.MatchString(name)
}

func calcMem(name string, load func()) {
	if !isTested(name) {
		return
	}

	m := new(runtime.MemStats)

	// before
	runtime.GC()
	runtime.ReadMemStats(m)
	before := m.HeapAlloc

	load()

	// after
	runtime.GC()
	runtime.ReadMemStats(m)
	after := m.HeapAlloc
	println("   "+name+":", after-before, "Bytes")
}
func benchRequest(b *testing.B, router http.Handler, r *http.Request) {
	w := new(mockResponseWriter)
	u := r.URL
	rq := u.RawQuery
	r.RequestURI = u.RequestURI()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		u.RawQuery = rq
		router.ServeHTTP(w, r)
	}
}
func init() {
	calcMem("Gin_loadsingle and writehandler", func() {
		ginQuery = loadGinSingle("GET", "/user/?name=test", ginHandleWrite)
	})
	calcMem("Echo_loadsingle and writehandler", func() {
		echoQuery = loadEchoSingle("GET", "/user/?name=test", echoHandlerWrite)
	})
	calcMem("Beego_loadsingle and writehandler", func() {
		beegoQuery = loadBeegoSingle("GET", "/user/?name=test", beegoHandlerWrite)
	})
	calcMem("Bmux_loadsingle and writehandler", func() {
		bmuxQuery = loadBmuxSingle("GET", "/user/?name=test", bmuxHandlerWrite)
	})
}
func BenchmarkGin_Query_Handler(b *testing.B) {
	router := loadGinSingle("GET", "/user", ginHandle)

	r, _ := http.NewRequest("GET", "/user/?name=test", nil)
	benchRequest(b, router, r)
}
func BenchmarkEcho_Query_Handler(b *testing.B) {
	router := loadEchoSingle("GET", "/user", echoHandler)

	r, _ := http.NewRequest("GET", "/user/?name=test", nil)
	benchRequest(b, router, r)
}
func BenchmarkBeego_Query_Handler(b *testing.B) {
	router := loadBeegoSingle("GET", "/user", beegoHandler)

	r, _ := http.NewRequest("GET", "/user/?name=test", nil)
	benchRequest(b, router, r)
}
func BenchmarkBmux_Query_Handler(b *testing.B) {
	router := loadBmuxSingle("GET", "/user", bmuxHandler)

	r, _ := http.NewRequest("GET", "/user/?name=test", nil)
	benchRequest(b, router, r)
}
func BenchmarkGin_Query_WriteHandler(b *testing.B) {
	router := loadGinSingle("GET", "/user", ginHandleWrite)

	r, _ := http.NewRequest("GET", "/user/?name=test", nil)
	benchRequest(b, router, r)
}
func BenchmarkEcho_Query_WriteHandler(b *testing.B) {
	router := loadEchoSingle("GET", "/user", echoHandlerWrite)

	r, _ := http.NewRequest("GET", "/user/?name=test", nil)
	benchRequest(b, router, r)
}
func BenchmarkBeego_Query_WriteHandler(b *testing.B) {
	router := loadBeegoSingle("GET", "/user", beegoHandlerWrite)

	r, _ := http.NewRequest("GET", "/user/?name=test", nil)
	benchRequest(b, router, r)
}

func BenchmarkBmux_Query_WriteHandler(b *testing.B) {
	router := loadBmuxSingle("GET", "/user", bmuxHandlerWrite)

	r, _ := http.NewRequest("GET", "/user/?name=test", nil)
	benchRequest(b, router, r)
}
