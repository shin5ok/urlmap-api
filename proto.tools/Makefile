all: go python

go:
	protoc -Iproto --go_out=plugins=grpc:../pb -I ../proto ../proto/*.proto

python:
	pip install -r requirements.txt
	python -m grpc_tools.protoc -Iproto --python_out=../pb --grpc_python_out=../pb proto/*.proto
