FILE=fib
LOOPS=1000
CPUPROF=$(GOPATH)/src/github.com/PuerkitoBio/specter/bench/cpu.prof
SPECTER=$(GOPATH)/bin/cmd
TVMI=$(HOME)/subwrkspc/tinyvm/bin/tvmi
RESULTS=$(GOPATH)/src/github.com/PuerkitoBio/specter/bench/results

.SILENT:
	
bench:
	echo Running benchmark for $(FILE).vm $(LOOPS) times
	echo
	echo TinyVM
	"./xtime.sh" "./loop_tvm.sh" $(FILE) $(LOOPS)
	echo 
	echo Specter
	"./xtime.sh" "./loop_specter.sh" $(FILE) $(LOOPS)
	echo

all:
	echo This will take a few minutes...
	echo 
	$(MAKE) bench FILE=fact
	$(MAKE) bench FILE=fib
	$(MAKE) bench FILE=jsr
	$(MAKE) bench FILE=loop
	$(MAKE) bench FILE=nop
	$(MAKE) bench FILE=primes
	$(MAKE) bench FILE=stack_bench
	$(MAKE) bench FILE=euler1
	$(MAKE) bench FILE=euler1_nodiv
	$(MAKE) bench FILE=euler2
	$(MAKE) bench FILE=euler7 LOOPS=4 # euler7 takes over 30secs, limit to 4 loops only!

install:
	go install github.com/PuerkitoBio/specter/...

run-prof:
	go test github.com/PuerkitoBio/specter/vm --bench "BenchmarkForProfiling" --cpuprofile $(CPUPROF)

prof:
	go tool pprof $(GOPATH)/bin/cmd $(CPUPROF)

run:
	$(SPECTER) $(GOPATH)/src/github.com/PuerkitoBio/specter/cmd/examples/$(FILE).vm

run-tvmi:
	$(TVMI) $(GOPATH)/src/github.com/PuerkitoBio/specter/cmd/examples/$(FILE).vm

PHONY: bench all prof run install run-prof run-tvmi
