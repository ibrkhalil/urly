run:
	go run server.go
watch:
	nodemon -x go run server.go --signal SIGTERM -e go --verbose

