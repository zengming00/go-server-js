var os = require('os')

console.log('args', os.args)

function getwd() {
  var wd = os.getwd()
  if (wd.err) {
    throw err;
  }
  console.log('wd:', wd.value)
}

getwd()
os.chdir("d:/")
getwd()

console.log('tmpdir:', os.tempDir())
console.log('path:', os.getEnv('path'))

var r = os.hostname()
if (r.err) {
  throw err;
}
console.log('hostname:', r.value)

var r = os.mkdir("est", 0666)
if (r) {
  throw r;
}
r = os.remove('est')
if (r) {
  throw r;
}













