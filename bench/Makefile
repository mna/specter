FILE=fib
LOOPS=1000
PROFFILE=cpu.prof
PROF=$(GOPATH)/src/github.com/PuerkitoBio/specter/bench/$(PROFFILE)
SPECTER=$(GOPATH)/bin/cmd
TVMI=$(SUBWRKSPC)/tinyvm/bin/tvmi
RESULTS=$(GOPATH)/src/github.com/PuerkitoBio/specter/bench/results
EXDIR=$(GOPATH)/src/github.com/PuerkitoBio/specter/cmd/examples

.SILENT:
	# Do not print Make commands

.NOTPARALLEL: 
	# Make sure there is no parallelism

bench:
	# Times the execution of the specified file for both implementations, printing only 
	# the timing information.
	# Not needed anymore, C time will not change...
	"./xtime.sh" "./loop_tvm.sh" $(FILE) $(LOOPS) | tail -n 0	
	"./xtime.sh" "./loop_specter.sh" $(FILE) $(LOOPS) | tail -n 0	

all:
	# Run the benchmark (timing the execution) for all .vm files
	echo This will take a few minutes...
	echo 
	$(MAKE) bench FILE=fact
	$(MAKE) bench FILE=fib
	$(MAKE) bench FILE=jsr
	$(MAKE) bench FILE=loop
	$(MAKE) bench FILE=loop2
	$(MAKE) bench FILE=loop3 LOOPS=10 # takes a lot of time
	$(MAKE) bench FILE=nop
	$(MAKE) bench FILE=primes
	$(MAKE) bench FILE=stack_bench
	$(MAKE) bench FILE=euler1
	$(MAKE) bench FILE=euler1_nodiv
	$(MAKE) bench FILE=euler2
	$(MAKE) bench FILE=euler7 LOOPS=4 # euler7 takes over 30secs, limit to 4 loops only!

install:
	# Build and install the packages (VM and Main)
	go install github.com/PuerkitoBio/specter/...

prof:
	# Run in profiling mode
	# $(SPECTER) -cpu=$(PROF) $(EXDIR)/$(FILE).vm
	go tool pprof $(SPECTER) $(PROF)

run:
	# Run a single iteration, not timed, of the specified file for each implementation
	$(SPECTER) $(GOPATH)/src/github.com/PuerkitoBio/specter/cmd/examples/$(FILE).vm
	$(TVMI) $(GOPATH)/src/github.com/PuerkitoBio/specter/cmd/examples/$(FILE).vm

gdb:
	gdb $(SPECTER) -d $(GOROOT)

clear:
	rm $(SPECTER)
	rm $(GOPATH)/pkg/darwin_amd64/github.com/PuerkitoBio/specter/vm.a

cmp:
	# Run and save the output for the specified file for each implementation, and compare the results.
	# This is to make sure both have the same output.
	echo No output means all is ok.
	$(SPECTER) $(GOPATH)/src/github.com/PuerkitoBio/specter/cmd/examples/$(FILE).vm > $(RESULTS)/tmp_specter
	$(TVMI) $(GOPATH)/src/github.com/PuerkitoBio/specter/cmd/examples/$(FILE).vm	> $(RESULTS)/tmp_tvmi
	diff $(RESULTS)/tmp_specter $(RESULTS)/tmp_tvmi

.PHONY: bench all prof run install cmp gdb
