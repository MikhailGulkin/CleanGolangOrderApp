generate-customer:
	mkdir -p ./servicespb
	protoc --proto_path=servicespb \
 		  --go_out=servicespb --go_opt=paths=source_relative \
          --go-grpc_out=servicespb --go-grpc_opt=paths=source_relative \
			customer.proto