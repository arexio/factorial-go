## Run server example. Usage 'make run-server-example'
run-server-example: ; $(info Starting server example...)
	go run examples/server/*.go

## Run persistence example. Usage 'make run-persistence-example'
run-persistence-example: ; $(info Starting persistence example...)
	go run examples/persistence/*.go
