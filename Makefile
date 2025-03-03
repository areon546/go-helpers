

run: 
	go run .

coverage:
	cd helpers && go test -coverprofile cover.out && go tool cover -html=cover.out

test:
	cd helpers && go test