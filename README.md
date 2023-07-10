## Thetan Buf
`Azzzzzzzzzzzzzzzzzz`

### Formatting Protobuf files
- https://buf.build/docs/format/style/

### Install tools
```sh
#brew install buf
make install
```

### Generate code
```sh
#buf generate proto
make generate
```

### Unity Release Package

#### Flow release Unity

```shell
buf generate proto
```

```shell
cp ./unity/package.json ./gen/csharp/package.json | cp ./unity/Wolffun.Protobuf.asmdef ./gen/csharp/Wolffun.Protobuf.asmdef
```

```shell
git add . | git commit -m "????"
```

```shell
./tools/unity-meta-check-darwin-amd64 ./gen/csharp | ./tools/unity-meta-autofix-darwin-amd64 -root-dir ./gen/csharp .
```

```sh
./unity/deploy.sh --semver 0.0.1-preview.12
```
