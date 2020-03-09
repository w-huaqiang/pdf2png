### pdf2png

##### 注意

本程序的使用环境为windows,需要mutool.exe

1. windows编译

```shell
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go -o pdf2p.exe
```

2. 使用

将需要转换的pdf文件放在当前pdf目录中，双击可执行文件 pdf2p.exe即可。若需调整转换出的png大小，参考修改`pdf2png.ini`参数文件。

也可直接下载编译好的工具使用: [release](http://githu.com)