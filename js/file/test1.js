var os = require('os')

var r = os.create('te.txt')
if(r.err){
  throw r.err
}

var file = r.file
file.writeString("helloworld go-serverjs")
file.close()
