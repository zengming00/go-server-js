var file = require('file')

// 汉字
var data = file.read('./js/c.js')

var header = response.header()
// response.writeHeader(403)
header.set('haha', 'hehe')
// response.write(data)
data = [97, 98, 100, 0x0d, 0x0a]
// data = [97, 98, 1024]
console.log(Array.isArray(data))

var len = response.write(data)
new Date()
