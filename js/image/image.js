

var image = require('image');

var Image = (function () {
    function Image(w, h) {
        this.w = w;
        this.h = h;

        var rect = image.rect(0, 0, this.w, this.h)
        this.img = image.newRGBA(rect)
    }
    Image.prototype.drawPoint = function (x, y, rgba) {
        if (x >= this.w || y >= this.h || x < 0 || y < 0) {
            return;
        }
        if (!(rgba instanceof Object)) {
            rgba = {
                b: rgba & 0xFF, // 蓝
                g: (rgba >> 8) & 0xFF, // 绿
                r: (rgba >> 16) & 0xFF, // 红
                a: 255,
            }
        }
        this.img.setRGBA(x, y, rgba.r, rgba.g, rgba.b, rgba.a);
    };
    Image.prototype.drawLineH = function (x1, x2, y, rgba) {
        if (x1 > x2) {
            var tmp = x2;
            x2 = x1;
            x1 = tmp;
        }
        for (; x1 <= x2; x1++) {
            this.drawPoint(x1, y, rgba);
        }
    };
    Image.prototype.drawLineV = function (y1, y2, x, rgba) {
        if (y1 > y2) {
            var tmp = y2;
            y2 = y1;
            y1 = tmp;
        }
        for (; y1 <= y2; y1++) {
            this.drawPoint(x, y1, rgba);
        }
    };
    Image.prototype.drawLine = function (x1, y1, x2, y2, rgba) {
        var x, y, dx, dy, s1, s2, p, temp, interchange, i;
        x = x1;
        y = y1;
        dx = x2 > x1 ? (x2 - x1) : (x1 - x2);
        dy = y2 > y1 ? (y2 - y1) : (y1 - y2);
        s1 = x2 > x1 ? 1 : -1;
        s2 = y2 > y1 ? 1 : -1;
        if (dy > dx) {
            temp = dx;
            dx = dy;
            dy = temp;
            interchange = true;
        }
        else {
            interchange = false;
        }
        p = (dy << 1) - dx;
        for (i = 0; i <= dx; i++) {
            this.drawPoint(x, y, rgba);
            if (p >= 0) {
                if (interchange) {
                    x = x + s1;
                }
                else {
                    y = y + s2;
                }
                p = p - (dx << 1);
            }
            if (interchange) {
                y = y + s2;
            }
            else {
                x = x + s1;
            }
            p = p + (dy << 1);
        }
    };
    Image.prototype.drawRect = function (x1, y1, x2, y2, rgba) {
        this.drawLineH(x1, x2, y1, rgba);
        this.drawLineH(x1, x2, y2, rgba);
        this.drawLineV(y1, y2, x1, rgba);
        this.drawLineV(y1, y2, x2, rgba);
    };
    Image.prototype.fillRect = function (x1, y1, x2, y2, rgba) {
        var x;
        if (x1 > x2) {
            var tmp = x2;
            x2 = x1;
            x1 = tmp;
        }
        if (y1 > y2) {
            var tmp = y2;
            y2 = y1;
            y1 = tmp;
        }
        for (; y1 <= y2; y1++) {
            for (x = x1; x <= x2; x++) {
                this.drawPoint(x, y1, rgba);
            }
        }
    };
    Image.prototype.drawCircle = function (x, y, r, rgba) {
        var a, b, c;
        a = 0;
        b = r;
        //   c = 1.25 - r;
        c = 3 - 2 * r;
        while (a < b) {
            this.drawPoint(x + a, y + b, rgba);
            this.drawPoint(x - a, y + b, rgba);
            this.drawPoint(x + a, y - b, rgba);
            this.drawPoint(x - a, y - b, rgba);
            this.drawPoint(x + b, y + a, rgba);
            this.drawPoint(x - b, y + a, rgba);
            this.drawPoint(x + b, y - a, rgba);
            this.drawPoint(x - b, y - a, rgba);
            if (c < 0) {
                c = c + 4 * a + 6;
            }
            else {
                c = c + 4 * (a - b) + 10;
                b -= 1;
            }
            a = a + 1;
        }
        if (a === b) {
            this.drawPoint(x + a, y + b, rgba);
            this.drawPoint(x - a, y + b, rgba);
            this.drawPoint(x + a, y - b, rgba);
            this.drawPoint(x - a, y + b, rgba);
            this.drawPoint(x + b, y + a, rgba);
            this.drawPoint(x - b, y + a, rgba);
            this.drawPoint(x + b, y - a, rgba);
            this.drawPoint(x - b, y - a, rgba);
        }
    };
    Image.prototype.drawChar = function (ch, x, y, font, rgba) {
        var index = font.fonts.indexOf(ch);
        if (index < 0) {
            return;
        }
        var fontData = font.data[index];
        var y0 = y;
        var x0 = x;
        for (var _i = 0, fontData_1 = fontData; _i < fontData_1.length; _i++) {
            var data = fontData_1[_i];
            x0 = x;
            for (var b = data; b > 0; b <<= 1) {
                if (b & 0x80) {
                    this.drawPoint(x0, y0, rgba);
                }
                x0++;
            }
            y0++;
            if ((y0 - y) >= font.h) {
                y0 = y;
                x += 8;
            }
        }
    };
    Image.prototype.drawString = function (str, x, y, font, rgba) {
        for (var _i = 0, str_1 = str; _i < str_1.length; _i++) {
            var c = str_1[_i];
            this.drawChar(c, x, y, font, rgba);
            x += font.w;
        }
    };
    return Image;
}());
exports.Image = Image;
