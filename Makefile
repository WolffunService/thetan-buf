install :
#if macos - Darwin
	brew install buf
	brew install protoc-gen-go

generate:
	buf generate proto
	buf generate --template buf.gen.tag.yaml proto
