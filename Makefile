all: 
	go run race-condition.go	

race:
	go run -race race-condition.go
