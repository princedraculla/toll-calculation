obu:
	@go build -o bin/obu obu/main.go
	@./bin/obu

reciever:
	@go build -o bin/receiver ./data_reciever
	@./bin/receiver

calculator: 
	@go build -o bin/calculator ./distance_calculator
	@./bin/calculator
agg:
	@go build -o bin/agg ./aggregator
	@./bin/agg

.PHONY: obu invoicer