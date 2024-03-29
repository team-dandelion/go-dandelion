proto:
	protoc -I. -I${GOPATH}/src -I ../../ -I ./ \
	  --gofast_out=. --gofast_opt=paths=source_relative \
	  --rpcx_out=. --rpcx_opt=paths=source_relative *.proto
