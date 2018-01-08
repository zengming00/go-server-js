
function log(d) {
  console.log("%j", d);
  response.write(d);
  response.write('\r\n');
}

// http://localhost:8080/js/http/request.js?a=b&c=d&c=22#aa

log("method: " + request.method) // POST
log("uri: " + request.uri) // /js/http/request.js?a=b&c=d
log("host: " + request.host) // localhost:8080
log("userAgent: " + request.userAgent()) // PostmanRuntime/7.1.1
log("remoteAddr: " + request.remoteAddr) // [::1]:63716

// 若调用 parseForm() 则body参数必需以'application/x-www-form-urlencoded'方式传递
// 不调用，则可以由'application/x-www-form-urlencoded'和'multipart/form-data'传递
// request.parseForm()
log("test=" + request.formValue('test')) // from body  test=1111 : 1111

log("a=" + request.formValue('a')) // b
log("c=" + request.formValue('c')) // d

var values = request.form.gets('c')
log(JSON.stringify(values)) // ["d","22"]