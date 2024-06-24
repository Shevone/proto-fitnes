.PHONY: all generate git-add git-commit git-push

all: generate git-add git-commit git-push



generate:
	protoc -I proto proto/accounts/accounts.proto --go_out=gen/go --go-grpc_out=gen/go --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative
	protoc -I proto proto/training/training --go_out=gen/go --go-grpc_out=gen/go --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative
	protoc -I proto proto/auth/auth.proto --go_out=gen/go --go-grpc_out=gen/go --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative
git-add:
	git add gen/go

git-commit:
	git commit -m "Generate protobuf files"

git-push:
	git push -u origin main
