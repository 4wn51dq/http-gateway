.PHONY: protos

protos: 
	protoc -I=protos/currency \
  --go_out=protos/ \
  --go-grpc_out=protos/ \
  protos/currency/currency.proto