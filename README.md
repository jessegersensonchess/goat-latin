# Profiling
 - go test -cpuprofile cpu.prof -memprofile mem.prof -bench .
 - go tool pprof cpu.prof # top20 -cum
 - go tool pprof mem.prof # top20
