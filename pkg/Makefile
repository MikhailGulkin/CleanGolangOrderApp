generate-customer:
	mkdir -p ./customer/servicespb
	protoc --proto_path=customer \
 		  --go_out=customer/servicespb --go_opt=paths=source_relative \
          --go-grpc_out=customer/servicespb --go-grpc_opt=paths=source_relative \
			customer.proto