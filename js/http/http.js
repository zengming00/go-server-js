var http = require("http")
var file = require("file")
var utils = require("utils")
var url = require('url')

var data = file.read('./js/b.js')
console.log(utils.toString(data))
console.log(url.queryEscape("asdf/asdf"))

var obj = {
  ticket: '99c936b0-9c35-11e7-be48-43b6cddd43aa'
};

var header = {
  "Content-Type": "application/json",
  "x-api-version": "1",
};

var method = "POST"
var url = "http://esee.rabbitpre.com/api/rabuser/login"
var body = JSON.stringify(obj)
var data = http.request(method, url, header, body, 3000)

data.body = utils.toString(data.body);
console.log("data:", JSON.stringify(data, null, 2))
console.log("Set-Cookie:", data.header['Set-Cookie'][0])

obj = JSON.parse(data.body)
var avatar = obj.result.avatar
console.log("avatar:", avatar)

data = http.request("GET", avatar, {}, '', 3000)

file.write("./a.jpg", data.body)

console.log('done.');

