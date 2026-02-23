# yolo

yolo - Golang project scaffolding

## install

```
go install github.com/gd365/yolo@latest
```
Explanation: By default, `yolo` is installed in the `$GOPATH/bin` directory.

## usage

```
$GOBIN/yolo init -name <project_name> -port <project_port>
```

## run

```
cd <project_name>/cmd
go run server.go
```

## test

```
curl http://127.0.0.1:<project_port>/v1/liveness
```
