genpb:
	 @protoc -I=. -I=$(GOPATH)/src -I=$(GOPATH)/src/github.com/gogo/protobuf --gogoslick_out=. workpb/*.proto
