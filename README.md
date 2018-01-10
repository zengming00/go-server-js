# a js server

# china user , git proxy (optional)
## set
```
git config --global http.proxy http://127.0.0.1:8087

git config --global https.proxy https://127.0.0.1:8087
```
## unset
```
git config --global --unset http.proxy

git config --global --unset https.proxy
```
## disable ssl 
git config --global http.sslVerify false

## dep
http_proxy=http://127.0.0.1:8087  dep init -v

## 下载 (download)
https://github.com/zengming00/go-server-js/releases
