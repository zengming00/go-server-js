console.log("haha")
var types = require('types')

var v = types.newInt()
var b = types.newBool()

log("v",v)
log("types.intValue",types.intValue(v))
log("b",b)
log("types.boolValue", types.boolValue(b))

types.scan(v, b)

log("v",v)
log("types.intValue",types.intValue(v))
log("b",b)
log("types.boolValue", types.boolValue(b))


function log(label, data) {
  console.log(label + ":", data)
  console.log(label + " json:", JSON.stringify(data, null, 2))
}