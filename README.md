# a js server
一种新的javascript写服务端程序的方案，没有回调，完全不同于node.js

这可能是最简单的服务器方案了，因为它不需要任何配置，只有一个文件，运行它就有了完整的服务器、数据库和缓存系统，并且可以运行在go语言支持的所有平台上，例如:路由器上，相对于其它语言开发的服务器软件而言是非常有优势的，我曾经在路由器上配置过php服务器，花了好几个小时的时间，并且无法及时更新到新版本。这不像pc机那么容易的。

优点：
1. 免安装，无需配置任何环境，自带sqlite数据库和一个简易的缓存系统
2. go语言开发，无限扩展功能，可以自由定制
3. 跨平台，支持linux、windows、mac，支持x86/arm/mips等指令集的cpu（运行在安卓手机、树莓派、路由器、国产龙芯。。。）
4. 完全不同于node.js，没有回调，程序更易维护和编写，推荐使用typescript

缺点：
1. 性能不高，和node.js完全不是一个级别的，目前能完全满足小应用的需求，在这里能找到一份测试报告：https://github.com/zengming00/go-server-js-testShop
2. api目前不够完善，我只是需要什么就往上面加什么，你也可以
3. 目前没有文档支持，没空写，哈哈
4. 没有debug功能，调试不方便，这是个很严重的问题，目前没有办法

写node一年了，感觉node的异步优势很少用到，坑爹的回调让人非常痛苦，虽然说有async/await但仍然会时不时接触到，用async/await其实还是在写同步的代码，所谓的异步并发优势只能很少数场合能用到，有的时候甚至是得不偿失的。

在不了解node之前，在我的想象中以为node是像php那样写的，学过之后发现完全不是这样的，我曾经去找过类似这种东西，但是没找到，也许是我的方法不对，网上有个fibjs，那不是我想要的。并且，我是在发布go-server-js之后才听说的

所以我决定自己做一个，为此我曾经注册了serverjs.cn域名，我尝试过用c语言来写，C语言门槛太高了只做了一个能够连上redis的东西，还得靠cgi模式来提供服务，最终放弃了，后来接触了go语言，于是有了实现它的可能，最早是用的otto，后来才用的goja。

从开始到完成0.0.2版本应该花了一个多月吧，基本上都是在学go语言和尝试一些功能细节，最后还花时间将我一年前写的一个nodejs写的商城用我自创的go-server-js技术重写了一遍（ 项目地址：https://github.com/zengming00/go-server-js-testShop ），并且和go-server-js 0.0.2捆绑发布。

再后来，为了验证go-server-js写的项目能不能够方便的移植到原生go语言项目，我又把这个商城用go语言写了一遍（ 项目地址：https://github.com/zengming00/go-testShop ），得益于当时选择照抄go语言的编程风格，只需要将go-server-js封装好的一些功能实现，js代码到go语言代码的转换是很方便的

# 下载试用 (download)，下载后运行go-server-js，打开 http://127.0.0.1:8080/ 就是一个商城
https://github.com/zengming00/go-server-js/releases

## 入门之helloworld
整个服务器就是go-server-js这个文件，不需要任何其它东西，运行它会在当前目录下生成一个config.json，这个是用来修改一些服务器配置的，比如端口号之类的。

在当前目录下创建test.js内容为

```js
response.write('helloworld');
```
运行go-server-js然后在浏览器打开 http://127.0.0.1:8080/test.js 就可以看到helloworld了，修改文件后不需要重启服务器
更多的教程不如直接看附带的商城源码，一些api与go语言是完全一样的，因此为后期移植到原生go语言提供了极大便利

# 用go-server-js写的项目
一个商城：https://github.com/zengming00/go-server-js-testShop

wooyun本地镜像：https://github.com/zengming00/go-server-js-wooyun

# 用go-server-js写的项目移植为go语言代码，两种等效代码对比

https://github.com/zengming00/go-testShop

![两种等效代码对比](https://github.com/zengming00/go-testShop/raw/master/public/uploads/1.png)


# 如果你想尝试自己编译，需要准备点东西

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


## 一些实现上的细节
```go
    // 在go语言中如果是返回一个error，要经过转换才能给js使用
       err := db.Close()
       if err != nil {
           return runtime.ToValue(lib.NewError(runtime, err))
       }
       return nil
	
    // 如果类型无法处理，应该用这种方式抛出
       panic(runtime.NewTypeError("p0 is not a string type:%T", args[0]))
	
    // 如果有多个返回值，应该返回一个对象供js使用
       tx, err := db.Begin()
       if err != nil {
           return lib.MakeErrorValue(runtime, err)
       }
       return lib.MakeReturnValue(runtime, NewTx(runtime, tx))
		
    // 动态参数
       args := lib.GetAllArgs(&call)
       err := rows.Scan(args...)
       
    // go语言原生类型的传递
        p0 := GetNativeType(runtime, &call, 0)
        if err, ok := p0.(error); ok {
            return runtime.ToValue(os.IsNotExist(err))
        }
        panic(runtime.NewTypeError("p0 is not error type:%T", p0))
    // 注意函数签名的不同
        func(call goja.FunctionCall) goja.Value {}
        func(call goja.ConstructorCall) *Object {}
```

# docker
```Dockerfile
FROM scratch
COPY . /
EXPOSE 8080
CMD [ "/go-server-js" ]
```
```sh
$ docker pull zengming00/go-server-js
$ docker run -d --rm -p 80:8080 zengming00/go-server-js
```

