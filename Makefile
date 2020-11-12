.PHONY: cli
all: front bindata back cli
	mkdir -p build
	cp server/build/treasury build/
	cp -r cli/build/* build/

docker:
	docker build . -t treasury/treasury

run_debug: fake_bindata
	cd webapp/treasury && npm start &
	cd server && go run .

# Converts the backend build results into a go asset file
bindata:
	cd webapp/treasury/build && go-bindata -o ../../../server/assets.go ./...

fake_bindata:
	cp server/assets.go_ server/assets.go

front:
	cd webapp/treasury && npm run-script build

back:
	cd server && mkdir -p build && CGO_ENABLED=0 go build -o build/treasury .

cli:
	cd cli && mkdir -p build && CGO_ENABLED=0 go build -o build/treasury-linux-x64 .
	cd cli && mkdir -p build && GOOS=windows CGO_ENABLED=0 go build -o build/treasury-win-x64.exe .
	
.PHONY: docker
