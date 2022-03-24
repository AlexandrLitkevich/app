start:
	cd ./client; npm run build
	cd ./server; go run main.go


back: 
	cd ./server; go run main.go


front: 
	cd ./client; npm run start
