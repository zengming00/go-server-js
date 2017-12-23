var redis = require('redis')
var utils = require('utils')

var conn = redis.dial('tcp', ':6379')

var v = conn.do('get', 'str')
log(redis.string(v))

v = conn.do('keys', '*')
if(Array.isArray(v)){
  v.forEach(function(item){
    println(utils.toString(item))
  });
}
var ret = conn.close()
log(conn)
log(ret)

function log(data) {
  console.log("log:", data)
  console.log("log:", JSON.stringify(data, null, 2))
}

function println(data){
  utils.print(data);
  utils.print("\r\n");
}