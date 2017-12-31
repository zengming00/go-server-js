var redis = require('redis')
var utils = require('utils')

var r = redis.dial('tcp', ':6379')
if (r.err) {
  throw r.err;
}
var conn = r.value;
r = conn.do('get', 'str')
if (r.err) {
  throw r.err;
}
r = redis.string(r.value)
if (r.err) {
  conn.do('set', 'str', 'helloworld', 'ex', '3')
  console.log(">>>>>>>   set value")
  throw r.err;
} else {
  log(r.value)
}

r = conn.do('keys', '*')
if (r.err) {
  throw r.err;
}
if (Array.isArray(r.value)) {
  r.value.forEach(function (item) {
    println(utils.toString(item))
  });
}
r = conn.close()
if (r) {
  throw r;
}

log(r)

function log(data) {
  console.log("log:", data)
  console.log("log:", JSON.stringify(data, null, 2))
}

function println(data) {
  utils.print(data);
  utils.print("\r\n");
  response.write(data + "\r\n");
}