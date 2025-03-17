go1.23.0.txt:
	go1.23.0 test -bench=. -benchmem -count=10 -timeout 20m > go1.23.0.txt

go1.24.0.txt:
	go1.24.0 test -bench=. -benchmem -count=10 -timeout 20m > go1.24.0.txt

benchmark_comparison.txt: results_noswissmap.txt results.txt
	benchstat go1.24.0.txt go1.23.0.txt > diff.txt
