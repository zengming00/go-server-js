# a js server
一种新的javascript写服务端程序的方案，没有回调，完全不同于node.js

优点：
1. 免安装，无需配置任何环境，自带sqlite数据库和一个简易的缓存系统
2. go语言开发，无限扩展功能，可以自由定制
3. 跨平台，支持linux、windows、mac，支持x86/arm/mips等指令集的cpu
4. 完全不同于node.js，没有回调，程序更易维护和编写

缺点(还有更多)：
1. 性能不高，和node.js完全不是一个级别的
2. api目前不够完善
3. 目前没有文档支持
4. 没有debug功能，调试不方便

写node一年了，感觉node的异步很少用到，坑爹的回调让人非常痛苦，在不了解node之前，在我的想象中node就是像php那样写的，但实际上不是这样的，我曾经去找过类似这种东西，但没找到，可能是我的方法不对，所以我决定自己做一个，尝试过用c语言来写，但是那个门槛太高了最终放弃，后来接触了go语言，发现了goja这个开源项目，于是做这个东西变为可能。

如果你觉得这东西没卵用，请闭嘴，且不说有多少实际价值，起码我把我的想法变成了现实，曾经js写服务器只能选node（也许有其它），现在，有了新的选择。

# 下载试用 (download)
https://github.com/zengming00/go-server-js/releases

# 用go-server-js写的一个项目
https://github.com/zengming00/go-server-js-testShop

# 用go-server-js写的项目移植为go语言代码，两种等效代码对比

https://github.com/zengming00/go-testShop

![两种等效代码对比](https://github.com/zengming00/go-testShop/raw/master/public/uploads/1.png)


# 编译前的准备

因为使用的一些包在国内环境无法直接go get到，所以要做一些特殊的操作

**方法一，在GOROOT目录下执行下列命令(可以通过go env命令获取此目录的路径)**

```
mkdir -p src/golang.org/x/
cd src/golang.org/x/
git clone https://github.com/golang/text
git clone https://github.com/golang/net
```


**方法二，设置git代理**
* set
```
git config --global http.proxy http://127.0.0.1:8087

git config --global https.proxy https://127.0.0.1:8087
```
* unset
```
git config --global --unset http.proxy

git config --global --unset https.proxy
```
* disable ssl 
```
git config --global http.sslVerify false
```

* 为dep设置代理
```
http_proxy=http://127.0.0.1:8087  dep ...
```
## 获取源码
```
go get -v github.com/zengming00/go-server-js
```
也可以手动克隆此项目，建议使用dep工具安装依赖的包

默认不会编译sqlite，在windows下编译sqlite需要安装 TDM-GCC 并 set CGO_ENABLED = 1

如果要编译sqlite，在windows下运行build.bat，在linux下运行build_linux.sh
