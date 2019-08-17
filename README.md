## Test Query param and Write
``` s
   Gin_loadsingle and writehandler: 728 Bytes
   Echo_loadsingle and writehandler: 2600 Bytes
   Beego_loadsingle and writehandler: 1256 Bytes
goos: linux
goarch: amd64
pkg: github.com/dtmkeng/benchmark_queryparam
BenchmarkGin_Query_Handler        	30000000	        55.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkEcho_Query_Handler       	 1000000	      1525 ns/op	    1408 B/op	      14 allocs/op
BenchmarkBeego_Query_Handler      	 2000000	       659 ns/op	     352 B/op	       3 allocs/op
BenchmarkGin_Query_WriteHandler   	 3000000	       497 ns/op	     464 B/op	       4 allocs/op
BenchmarkEcho_Query_WriteHandler  	 1000000	      1609 ns/op	    1408 B/op	      14 allocs/op
BenchmarkBeego_Query_WriteHandler 	 2000000	       622 ns/op	     360 B/op	       4 allocs/op
PASS
ok  	github.com/dtmkeng/benchmark_queryparam	10.729s

```