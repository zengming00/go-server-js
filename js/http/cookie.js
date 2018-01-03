
var Name = 'test'
var Value = 'testval'
var Path = '/'
var MaxAge = 10
var HttpOnly = true


// response.setCookie(Name, Value, Path, MaxAge, HttpOnly)
response.setCookie(Name, Value, Path, MaxAge)
// response.setCookie(Name, Value, Path)


var cks = request.cookies()
var str = JSON.stringify(cks, null, 2)
log(str)

var r = request.cookie('test')
if (r.err) {
  throw r.err;
}
var cookie = r.value
log(cookie.name)
log(cookie.value)


function log(d) {
  console.log("%j", d);
  response.write(d);
  response.write('\r\n');
}

