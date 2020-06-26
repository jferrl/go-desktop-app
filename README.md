# go-desktop-app

Golang desktop app example based on https://github.com/webview/webview

## Build ui

esc -o ../../api/static.go -pkg api ./

## Build go

### MacOS

go build -o desktopapp.app/Contents/MacOS/desktopapp cmd/go-desktopapp/main.go
