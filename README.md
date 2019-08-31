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

   Gin_loadsingle and writehandler: 792 Bytes
   Echo_loadsingle and writehandler: 2600 Bytes
   Beego_loadsingle and writehandler: 1256 Bytes
   Bmux_loadsingle and writehandler: 18928 Bytes
goos: linux
goarch: amd64
pkg: github.com/dtmkeng/benchmark_queryparam
BenchmarkGin_Query_Handler        	20000000	       100 ns/op	       0 B/op	       0 allocs/op
BenchmarkEcho_Query_Handler       	  500000	      2808 ns/op	    1408 B/op	      14 allocs/op
BenchmarkBeego_Query_Handler      	 1000000	      1052 ns/op	     352 B/op	       3 allocs/op
BenchmarkBmux_Query_Handler       	20000000	        73.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_Query_WriteHandler   	 2000000	       767 ns/op	     464 B/op	       4 allocs/op
BenchmarkEcho_Query_WriteHandler  	 1000000	      2463 ns/op	    1408 B/op	      14 allocs/op
BenchmarkBeego_Query_WriteHandler 	 1000000	      1127 ns/op	     360 B/op	       4 allocs/op
BenchmarkBmux_Query_WriteHandler  	20000000	        77.2 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/dtmkeng/benchmark_queryparam	14.820s

# update 
   Gin_loadsingle and writehandler: 792 Bytes
   Echo_loadsingle and writehandler: 2600 Bytes
   Beego_loadsingle and writehandler: 1256 Bytes
   Bmux_loadsingle and writehandler: 18928 Bytes
goos: linux
goarch: amd64
pkg: github.com/dtmkeng/benchmark_queryparam
BenchmarkGin_Query_Handler        	30000000	        53.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkEcho_Query_Handler       	 1000000	      1425 ns/op	    1408 B/op	      14 allocs/op
BenchmarkBeego_Query_Handler      	 3000000	       568 ns/op	     352 B/op	       3 allocs/op
BenchmarkBmux_Query_Handler       	30000000	        41.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_Query_WriteHandler   	 3000000	       425 ns/op	     464 B/op	       4 allocs/op
BenchmarkEcho_Query_WriteHandler  	 1000000	      1429 ns/op	    1408 B/op	      14 allocs/op
BenchmarkBeego_Query_WriteHandler 	 2000000	       640 ns/op	     360 B/op	       4 allocs/op
BenchmarkBmux_Query_WriteHandler  	30000000	        41.3 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/dtmkeng/benchmark_queryparam	13.062s

```