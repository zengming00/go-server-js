var utils = require('utils')

function log(d) {
  console.log("%j", d);
  var r = response.write(d);
  if (r.err) {
    throw r.err;
  }
  response.write('\r\n');
}

var r = request.getRawBody()
if (r.err) {
  throw r.err;
}

var value = utils.toString(r.value)
log(value)
var obj = JSON.parse(value)
log(obj.a)



