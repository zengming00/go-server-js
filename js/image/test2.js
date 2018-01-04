var os = require('os')
var png = require('png')
var image = require('image')


function writeToFile(img) {
  var r = os.create('image.png')
  if (r.err) {
    console.log(r.err)
    return
  }
  var file = r.file

  var err = png.encode(file, img)
  if (err) {
    console.log(err)
    return
  }
  var err = file.close()
  if (err) {
    console.log(err)
    return
  }
  console.log("done.")
}

function writeToResponse(img) {
  var err = png.encode(response, img)
  if (err) {
    console.log(err)
    return
  }
}

// var start = Date.now();
// var i = 0;
// while ((Date.now() - start) < 1000) {
//   image.makeCapcha()
//   i++;
// }
// console.log('1秒钟生成：' + i);

writeToResponse(image.makeCapcha())
