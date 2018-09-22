all: run

run: tidy
	#go run definitions.go utils.go main.go > out.txt
	go run definitions.go utils.go main.go

tidy:
	clear
