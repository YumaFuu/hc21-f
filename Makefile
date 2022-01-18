.PHONY: run

run:
	go run cmd/job/main.go friends; \
		go run cmd/job/main.go result
