all: run

run: tidy
	go build && ./one-time-processor

tidy:
	clear && rm -f ./one-time-processor
