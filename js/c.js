var file = require('file')
var time = require('time')
var utils = require('./js/utils.js')



// 汉字
var data = file.read('./js/c.js')

response.setCookie("testCookie", "cookieValueTest:" + new Date())
var header = response.header()
// response.writeHeader(403)
header.set('haha', 'hehe')

// response.write(data)
data = [97, 98, 100, 0x0d, 0x0a, 97]
// data = [97, 98, 1024]
// console.log(Array.isArray(data))

write(data)
write("method:" + request.method)
write("host:" + request.host)
write("requestURI:" + request.requestURI)

write("userAgent:" + request.userAgent())

write("header[x]:" + request.header.get('x'))
write("header[Content-Type]:" + request.header.get('Content-Type'))

// 可以获取query上的参数, 目前不支持json
// request.parseForm()
write("form[a]:" + request.formValue('a'))
write("form[cc]:" + request.formValue('cc'))

var d = parseInt(request.formValue('sleep'))
if (!isNaN(d)) {
  time.sleep(1000 * d)
}


try {
  var cookie = request.cookie("testCookie")
  utils.log(cookie)
  write("cookie[testCookie]:" + JSON.stringify(cookie, null, 2))
} catch (e) {

}


function write(v) {
  response.write(v)
  response.write("\r\n")
}

new Date()