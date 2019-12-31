#!/usr/bin/env bash

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build  -o ./build/win/vc.exe main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./build/mac/vc main.go

rm -f ./win.zip
rm -f ./mac.zip

cp ./conf/config1.toml ./build/win/
cp ./conf/config1.toml ./build/mac/

zip -q -r win.zip ./build/win
zip -q -r mac.zip ./build/mac