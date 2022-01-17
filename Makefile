.PHONY: run sort

run:
	go run cmd/job/main.go friends; \
		go run cmd/job/main.go result && \
		make sort

sort:
	cat data/row_result.csv | \
		grep -v ",,," | \
		awk 'NR==1;NR>1{print $0|"sort -k1nr -t,"}' > data/result.csv
