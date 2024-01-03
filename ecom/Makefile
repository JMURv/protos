pb-all:
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative ecom/user/user.proto &&\
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative ecom/category/category.proto &&\
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative ecom/item/item.proto &&\
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative ecom/favorite/favorite.proto &&\
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative ecom/review/review.proto