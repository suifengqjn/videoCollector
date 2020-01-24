#!/usr/bin/env bash

rm -rf ./build/*

mkdir build/vc_win32
mkdir build/vc_win32/conf

mkdir build/vc_win64
mkdir build/vc_win64/conf

mkdir build/vc_mac
mkdir build/vc_mac/conf

cp ./conf/config1.toml ./build/vc_win32/conf/config.toml
cp ./conf/config1.toml ./build/vc_win64/conf/config.toml
cp ./conf/config1.toml ./build/vc_mac/conf/config.toml

cp ./conf/ssr1.txt ./build/vc_win32/conf/ssr.txt
cp ./conf/ssr1.txt ./build/vc_win64/conf/ssr.txt
cp ./conf/ssr1.txt ./build/vc_mac/conf/ssr.txt


CGO_ENABLED=0 GOOS=windows GOARCH=386 go build  -o ./build/vc_win32/vc.exe main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build  -o ./build/vc_win64/vc.exe main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./build/vc_mac/vc main.go


zip -q -r win32.zip ./build/vc_win32
zip -q -r win64.zip ./build/vc_win64
zip -q -r mac.zip ./build/vc_mac