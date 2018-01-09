var io = require('io')
var os = require('os')

if (request.method === 'POST') {
  var err = request.parseMultipartForm(1024 * 1024)
  if (err) {
    throw err
  }
  var haha = request.formValue('haha')
  console.log('haha:', haha)
  var o = request.formFile('file')
  if (o.err) {
    throw o.err
  }
  console.log(o.name)
  var r = os.openFile(o.name, os.O_CREATE|os.O_WRONLY, 0666)
  var file = r.value
  io.copy(file, o.file)
  o.file.close()
  file.close()
  response.write(JSON.stringify(o.header, null, 2))

} else {
  response.write("hello /")
}