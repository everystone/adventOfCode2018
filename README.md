## Performance testing
go get -u github.com/google/pprof  
go test -cpuprofile cpu.prof -memprofile mem.prof -bench .  
go tool pprof cpu.prof  
pprof -http=localhost:5522 cpu.prof  