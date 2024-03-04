env: 

build: env
	@export
	@go fmt ./...
	@echo "Traval Diary build process started..."
	@go build -o build/diary-srv -a cmd/diary-srv/diary-srv.go
	@echo "Traval Diary build process Completed..."