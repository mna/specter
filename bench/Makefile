FILE=fib
LOOPS=100

.SILENT:
	
bench:
	echo Running benchmark for $(FILE).vm $(LOOPS) times
	echo
	echo TinyVM
	time "./loop_tvm.sh" $(FILE) $(LOOPS)
	echo 
	echo Specter
	time "./loop_specter.sh" $(FILE) $(LOOPS)
	echo

all:
	$(MAKE) bench FILE=fact
	$(MAKE) bench FILE=fib
	$(MAKE) bench FILE=jsr
	$(MAKE) bench FILE=loop
	$(MAKE) bench FILE=nop
	$(MAKE) bench FILE=primes
	$(MAKE) bench FILE=stack_bench

PHONY: bench all
