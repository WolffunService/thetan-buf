install:
#if macos - Darwin
	brew install buf
	brew install protoc-gen-go
	go get github.com/srikrsna/protoc-gen-gotag

generate:
	buf generate proto
	buf generate --template buf.gen.tag.yaml proto
