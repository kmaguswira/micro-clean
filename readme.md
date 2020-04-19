require:

HOMEBREW_NO_AUTO_UPDATE=1 brew install protobuf
HOMEBREW_NO_AUTO_UPDATE=1 brew install consul
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u github.com/micro/protoc-gen-micro