build:
	mkdir -p bin
	go build -o bin/voip voip/*
	go build -o bin/app  app/*.go

run-voip:
	./bin/voip

run-app:
	./bin/app