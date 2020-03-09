### pdf2png

##### 注意

本程序的使用环境为windows,需要mutool.exe

1. windows编译

```shell
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

2. 使用

将需要转换的pdf文件放在当前pdf目录中，双击可执行文件 pdf2p.exe即可。若需调整转换出的png大小，参考修改`pdf2png.ini`参数文件。

也可直接下载编译好的工具使用: [release](https://github.com/w-huaqiang/pdf2png/releases)

* pdf2png.ini

> \#dir 为转换文件的目录名称
>
> \#imgResolution 为转换成png的像素，像素越小，png文件越小
>
> \#extname 转换文件后缀名称，一般后缀为pdf