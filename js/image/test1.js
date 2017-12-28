var image = require('image')
var os = require('os')
var png = require('png')

function makeImage() {
  var width = 256
  var height = 256
  var rect = image.rect(0, 0, width, height)
  var img = image.newRGBA(rect)

  for (var y = 0; y < height; y++) {
    for (var x = 0; x < width; x++) {
      var r = (x + y) & 255
      var g = (x + y) << 1 & 255
      var b = (x + y) << 2 & 255
      var a = 255
      img.setRGBA(x, y, r, g, b, a)
    }
  }
  return img
}

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

writeToResponse(makeImage())










