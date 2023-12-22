MAIN_FILE=main.go
NON_PIE_EXECUTABLE=bin/math_test_non_pie
PIE_EXECUTABLE=bin/math_test_pie

GO=go
TIME=time
VALGRIND=valgrind
DIFF=diff

.PHONY: build bench_time bench_mem clean objdump

build: $(MAIN_FILE)
	@echo "Building non-PIE executable"
	@$(GO) build -o $(NON_PIE_EXECUTABLE) $(MAIN_FILE)

	@echo "Building PIE executable"
	@$(GO) build -buildmode=pie -o $(PIE_EXECUTABLE) $(MAIN_FILE)

bench_time: build
	@echo "Benchmarking time for non-PIE executable"
	@$(TIME) $(PIE_EXECUTABLE)

	@echo "Benchmarking time for PIE executable"
	@$(TIME) $(NON_PIE_EXECUTABLE)

bench_mem: build
	@echo "Benchmarking memory for non-PIE executable"
	@-$(VALGRIND) --tool=massif --stacks=yes --max-stackframe=687211821072 $(NON_PIE_EXECUTABLE) > $(NON_PIE_EXECUTABLE).valgrind.txt

	@echo "Benchmarking memory for PIE executable"
	@-$(VALGRIND) --tool=massif --stacks=yes --max-stackframe=687211821072 $(PIE_EXECUTABLE) > $(PIE_EXECUTABLE).valgrind.txt

objdump: build
	@$(GO) tool objdump $(NON_PIE_EXECUTABLE) > $(NON_PIE_EXECUTABLE).objdump
	@$(GO) tool objdump $(PIE_EXECUTABLE) > $(PIE_EXECUTABLE).objdump
	@-$(DIFF) $(NON_PIE_EXECUTABLE).objdump $(PIE_EXECUTABLE).objdump > objdump.diff

clean:
	@rm -f $(NON_PIE_EXECUTABLE) $(PIE_EXECUTABLE) $(NON_PIE_EXECUTABLE).objdump $(PIE_EXECUTABLE).objdump
