#!/usr/bin/env bash

rm -rf ./build/*

mkdir build/vc_win
mkdir build/vc_win/conf

mkdir build/vc_mac
mkdir build/vc_mac/conf

cp ./conf/config1.toml ./build/vc_win/conf/config.toml
cp ./conf/config1.toml ./build/vc_mac/conf/config.toml

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build  -o ./build/vc_win/vc.exe main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./build/vc_mac/vc main.go


zip -q -r win.zip ./build/vc_win
zip -q -r mac.zip ./build/vc_mac