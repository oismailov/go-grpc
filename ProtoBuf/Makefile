release:
		protoc -I . landingcrew.proto3 *.proto3 --go_out=plugins=grpc:../landingcrew/workflow/
		protoc -I . landingcrew.proto3 *.proto3 --go_out=plugins=grpc:../api.landingcrew.com/workflow/
		-protoc --ruby_out=build fields.proto3

deps:
		brew install prototool