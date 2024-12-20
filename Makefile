generate_grpc_code:
	protoc \
	--go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	proto/$(PKG)/$(PKG).proto

rebuild_service:
	docker compose up --build $(PKG)_service -d