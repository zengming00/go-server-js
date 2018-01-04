var os = require('os')
var png = require('png')
var font = require('./js/image/font.js')
var image = require('./js/image/image.js')



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








var cnfonts = {
  w: 16,
  h: 16,
  fonts: '中国',
  data: [
    [0x01, 0x01, 0x01, 0x01, 0x3F, 0x21, 0x21, 0x21, 0x21, 0x21, 0x3F, 0x21, 0x01, 0x01, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00, 0xF8, 0x08, 0x08, 0x08, 0x08, 0x08, 0xF8, 0x08, 0x00, 0x00, 0x00, 0x00],
    [0x00, 0x7F, 0x40, 0x40, 0x5F, 0x41, 0x41, 0x4F, 0x41, 0x41, 0x41, 0x5F, 0x40, 0x40, 0x7F, 0x40, 0x00, 0xFC, 0x04, 0x04, 0xF4, 0x04, 0x04, 0xE4, 0x04, 0x44, 0x24, 0xF4, 0x04, 0x04, 0xFC, 0x04],
  ],
};
// 测试字库
function makeImg2() {
  var img = new image.Image(300, 140);
  img.drawString('helloworld', 20, 10, font.font8x16, { r: 255, g: 0, b: 0, a: 255 });
  img.drawString('helloworld', 20, 25, font.font12x24, { r: 0, g: 255, b: 0, a: 255 });
  img.drawString('helloworld', 20, 50, font.font16x32, { r: 0, g: 0, b: 255, a: 255 });
  img.drawString('中国', 20, 85, cnfonts, { r: 255, g: 255, b: 0, a: 255 });
  return img;
}
// 仿PHP的rand函数
function rand(min, max) {
  return Math.random() * (max - min + 1) + min | 0;
}
// 制造验证码图片
function makeCapcha() {
  var img = new image.Image(100, 40);
  img.drawCircle(rand(0, 100), rand(0, 40), rand(10, 40), rand(0, 0xffffff));
  // 边框
  img.drawRect(0, 0, img.w - 1, img.h - 1, rand(0, 0xffffff));
  img.fillRect(rand(0, 100), rand(0, 40), rand(10, 35), rand(10, 35), rand(0, 0xffffff));
  img.drawLine(rand(0, 100), rand(0, 40), rand(0, 100), rand(0, 40), rand(0, 0xffffff));
  // return img;
  // 画曲线
  var w = img.w / 2;
  var h = img.h;
  var color = rand(0, 0xffffff);
  var y1 = rand(-5, 5); // Y轴位置调整
  var w2 = rand(10, 15); // 数值越小频率越高
  var h3 = rand(4, 6); // 数值越小幅度越大
  var bl = rand(1, 5);
  for (var i_1 = -w; i_1 < w; i_1 += 0.1) {
    var y_1 = Math.floor(h / h3 * Math.sin(i_1 / w2) + h / 2 + y1);
    var x_1 = Math.floor(i_1 + w);
    for (var j = 0; j < bl; j++) {
      img.drawPoint(x_1, y_1 + j, color);
    }
  }
  var p = 'ABCDEFGHKMNPQRSTUVWXYZ3456789';
  var str = '';
  for (var i_2 = 0; i_2 < 5; i_2++) {
    str += p.charAt(Math.random() * p.length | 0);
  }
  var fonts = [font.font8x16, font.font12x24, font.font16x32];
  var x = 15, y = 8; // tslint:disable-line
  for (var _i = 0, str_1 = str; _i < str_1.length; _i++) {
    var ch = str_1[_i];
    var f = fonts[Math.random() * fonts.length | 0];
    y = 8 + rand(-10, 10);
    img.drawChar(ch, x, y, f, rand(0, 0xffffff));
    x += f.w + rand(2, 8);
  }
  return img;
}

// 测试生成验证码的效率
// var start = Date.now();
// var i = 0;
// while ((Date.now() - start) < 1000) {
//   makeCapcha();
//   // makeImg2();
//   i++;
// }
// console.log('1秒钟生成：' + i);


// writeToResponse(makeImg2().img)
writeToResponse(makeCapcha().img)