all: go python

go:
	protoc -Iproto --go_out=./pb --go_opt=paths=source_relative --go-grpc_out=./pb --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false proto/urlmap.proto

python:
	pip install -r requirements.txt
	python -m grpc_tools.protoc -Iproto --python_out=pb --grpc_python_out=pb proto/*.proto

doc:
	# Before 'make doc', install protoc-gen-doc to your $PATH
	# Repository: https://github.com/pseudomuto/protoc-gen-doc
	protoc -I. -Iproto --doc_out=. --doc_opt=markdown,urlmap.md proto/*.proto
