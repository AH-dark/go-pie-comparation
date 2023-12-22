MAIN_FILE=main.go
NON_PIE_EXECUTABLE=bin/math_test_non_pie
PIE_EXECUTABLE=bin/math_test_pie

.PHONY: build bench_time bench_mem clean

build: $(MAIN_FILE)
	@echo "Building non-PIE executable"
	@go build -o $(NON_PIE_EXECUTABLE) $(MAIN_FILE)

	@echo "Building PIE executable"
	@go build -buildmode=pie -o $(PIE_EXECUTABLE) $(MAIN_FILE)

bench_time: build
	@echo "Benchmarking time for non-PIE executable"
	@time $(PIE_EXECUTABLE)

	@echo "Benchmarking time for PIE executable"
	@time $(NON_PIE_EXECUTABLE)

bench_mem: build
	@echo "Benchmarking memory for non-PIE executable"
	@valgrind --tool=massif --stacks=yes $(NON_PIE_EXECUTABLE)

	@echo "Benchmarking memory for PIE executable"
	@valgrind --tool=massif --stacks=yes $(PIE_EXECUTABLE)

clean:
	@rm -f $(NON_PIE_EXECUTABLE) $(PIE_EXECUTABLE)
