curl -sK -v http://localhost:6060/debug/pprof/allocs > allocs
curl -sK -v http://localhost:6060/debug/pprof/goroutine > goroutine
curl -sK -v http://localhost:6060/debug/pprof/heap > heap
curl -sK -v http://localhost:6060/debug/pprof/threadcreate > threadcreate

go tool pprof -png allocs > allocs.png
go tool pprof -png goroutine > goroutine.png
go tool pprof -png heap > heap.png
go tool pprof -png threadcreate > threadcreate.png
