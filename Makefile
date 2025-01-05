obu:
	@go build -o bin/obu obu/main.go
	@./bin/obu

reciever:
	@go build -o bin/receiver ./data_reciever
	@./bin/receiver


.PHONY: obu invoicer