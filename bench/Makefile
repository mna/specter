FILE=fib
LOOPS=100

bench:
	@echo Running benchmark for $(FILE).vm $(LOOPS) times
	@echo
	@echo TinyVM
	@time "./loop_tvm.sh" $(FILE) $(LOOPS)

	@echo 
	@echo Specter
	@time "./loop_specter.sh" $(FILE) $(LOOPS)

PHONY: bench
