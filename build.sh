#!/bin/bash
# garble混淆安装: go install mvdan.cc/garble@latest
GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 garble -literals -tiny build -ldflags "-w -s" -o build/cursor-vip_darwin_amd64;
GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 garble -literals -tiny build -ldflags "-w -s" -o build/cursor-vip_darwin_arm64;
# rsrc 应用程序图标安装: go install github.com/akavel/rsrc@latest
rsrc -arch amd64 -ico rsrc.ico -o rsrc.syso;
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -o build/cursor-vip_windows_amd64.exe;
rm rsrc.syso;
rsrc -arch arm -ico rsrc.ico -o rsrc.syso;
GOOS=windows GOARCH=arm64 CGO_ENABLED=0 go build -ldflags "-w -s" -o build/cursor-vip_windows_arm64.exe;
rm rsrc.syso;

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 garble -literals -tiny build -ldflags "-w -s" -o build/cursor-vip_linux_amd64;
GOOS=linux GOARCH=arm64 CGO_ENABLED=0 garble -literals -tiny build -ldflags "-w -s" -o build/cursor-vip_linux_arm64;

chmod +x build/install.sh;
chmod +x build/cursor-vip_darwin_amd64;
chmod +x build/cursor-vip_darwin_arm64;
chmod +x build/cursor-vip_windows_amd64.exe;
chmod +x build/cursor-vip_windows_arm64.exe;
chmod +x build/cursor-vip_linux_amd64;
chmod +x build/cursor-vip_linux_arm64;
