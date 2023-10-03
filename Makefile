install :
#if macos - Darwin
	brew install buf

generate:
	buf generate proto
	buf generate --template buf.gen.tag.yaml proto
