var utils = require('utils')

function log(d) {
  console.log("%j", d);
  response.write(d);
  response.write('\r\n');
}

log("getHeader('content-type'): " + request.header.get('content-type')) // application/x-www-form-urlencoded
log("headers:" + JSON.stringify(request.headers, null, 2))
log("headers:" + JSON.stringify(request.headers['Haha'], null, 2))
var r = request.header.getRaw()
log(utils.toString(r))



