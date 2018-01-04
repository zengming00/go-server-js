package lib

import (
	"image"
	"time"

	"image/color"
	"math"
	mrand "math/rand"
	"unicode/utf8"
)

func rand(min, max int) int {
	return mrand.Intn(max + 1)
}

func randColor() color.Color {
	return &color.RGBA{
		R: uint8(rand(0, 255)),
		G: uint8(rand(0, 255)),
		B: uint8(rand(0, 255)),
		A: 255,
	}
}

func MakeCapcha() *image.RGBA {
	mrand.Seed(time.Now().UnixNano())
	const imgw = 100
	const imgh = 40
	img := image.NewRGBA(image.Rect(0, 0, imgw, imgh))
	DrawCircle(img, rand(0, 100), rand(0, 40), rand(10, 40), randColor())
	// 边框
	DrawRect(img, 0, 0, imgw-1, imgh-1, randColor())
	FillRect(img, rand(0, 100), rand(0, 40), rand(10, 35), rand(10, 35), randColor())
	DrawLine(img, rand(0, 100), rand(0, 40), rand(0, 100), rand(0, 40), randColor())
	// return img;
	// 画曲线
	var w = float64(imgw / 2)
	var h = float64(imgh)
	var color = randColor()
	var y1 = float64(rand(-5, 5))
	var w2 = float64(rand(10, 15))
	var h3 = float64(rand(4, 6))
	var bl = rand(1, 5)
	for i1 := -w; i1 < w; i1 += 0.1 {
		var y1 = math.Floor(h/h3*math.Sin(i1/w2) + h/2 + y1)
		var x1 = math.Floor(i1 + w)
		for j := 0; j < bl; j++ {
			img.Set(int(x1), int(y1+float64(j)), color)
		}
	}
	var p = "ABCDEFGHKMNPQRSTUVWXYZ3456789"
	var plegn = utf8.RuneCountInString(p)
	var str = ""
	for i2 := 0; i2 < 5; i2++ {
		str += string(charAt(p, mrand.Intn(plegn)))
	}
	fonts := []*Font{Font8x16, Font12x24, Font16x32}
	x := 15
	y := 8
	arr := []rune(str)
	for i := 0; i < len(arr); i++ {
		var ch = arr[i]
		var f = fonts[mrand.Intn(len(fonts))]
		y = 8 + rand(-10, 10)
		DrawChar(img, ch, x, y, f, randColor())
		x += f.w + rand(2, 8)
	}
	return img
}

type DrawPoint interface {
	Set(x, y int, c color.Color)
}

func charAt(str string, i int) rune {
	arr := []rune(str)
	return arr[i]
}

func indexRune(str string, r rune) int {
	finded := false
	n := -1
	for _, v := range str {
		n++
		if v == r {
			finded = true
			break
		}
	}
	if finded {
		return n
	}
	return -1
}

func DrawLineH(img DrawPoint, x1, x2, y int, color color.Color) {
	if x1 > x2 {
		var tmp = x2
		x2 = x1
		x1 = tmp
	}
	for ; x1 <= x2; x1++ {
		img.Set(x1, y, color)
	}
}

func DrawLineV(img DrawPoint, y1, y2, x int, color color.Color) {
	if y1 > y2 {
		var tmp = y2
		y2 = y1
		y1 = tmp
	}
	for ; y1 <= y2; y1++ {
		img.Set(x, y1, color)
	}
}

func DrawLine(img DrawPoint, x1, y1, x2, y2 int, color color.Color) {
	var x, y, dx, dy, s1, s2, p, temp, i int
	var interchange bool
	x = x1
	y = y1

	if x2 > x1 {
		dx = x2 - x1
	} else {
		dx = x1 - x2
	}
	if y2 > y1 {
		dy = y2 - y1
	} else {
		dy = y1 - y2
	}
	if x2 > x1 {
		s1 = 1
	} else {
		s1 = -1
	}
	if y2 > y1 {
		s2 = 1
	} else {
		s2 = -1
	}
	if dy > dx {
		temp = dx
		dx = dy
		dy = temp
		interchange = true
	} else {
		interchange = false
	}
	p = (dy << 1) - dx
	for i = 0; i <= dx; i++ {
		img.Set(x, y, color)
		if p >= 0 {
			if interchange {
				x = x + s1
			} else {
				y = y + s2
			}
			p = p - (dx << 1)
		}
		if interchange {
			y = y + s2
		} else {
			x = x + s1
		}
		p = p + (dy << 1)
	}
}

func DrawRect(img DrawPoint, x1, y1, x2, y2 int, color color.Color) {
	DrawLineH(img, x1, x2, y1, color)
	DrawLineH(img, x1, x2, y2, color)
	DrawLineV(img, y1, y2, x1, color)
	DrawLineV(img, y1, y2, x2, color)
}

func FillRect(img DrawPoint, x1, y1, x2, y2 int, color color.Color) {
	var x int
	if x1 > x2 {
		var tmp = x2
		x2 = x1
		x1 = tmp
	}
	if y1 > y2 {
		var tmp = y2
		y2 = y1
		y1 = tmp
	}
	for ; y1 <= y2; y1++ {
		for x = x1; x <= x2; x++ {
			img.Set(x, y1, color)
		}
	}
}

func DrawCircle(img DrawPoint, x, y, r int, color color.Color) {
	var a, b, c int
	a = 0
	b = r
	//   c = 1.25 - r;
	c = 3 - 2*r
	for a < b {
		img.Set(x+a, y+b, color)
		img.Set(x-a, y+b, color)
		img.Set(x+a, y-b, color)
		img.Set(x-a, y-b, color)
		img.Set(x+b, y+a, color)
		img.Set(x-b, y+a, color)
		img.Set(x+b, y-a, color)
		img.Set(x-b, y-a, color)
		if c < 0 {
			c = c + 4*a + 6
		} else {
			c = c + 4*(a-b) + 10
			b--
		}
		a = a + 1
	}
	if a == b {
		img.Set(x+a, y+b, color)
		img.Set(x-a, y+b, color)
		img.Set(x+a, y-b, color)
		img.Set(x-a, y+b, color)
		img.Set(x+b, y+a, color)
		img.Set(x-b, y+a, color)
		img.Set(x+b, y-a, color)
		img.Set(x-b, y-a, color)
	}
}

func DrawChar(img DrawPoint, ch rune, x int, y int, font *Font, color color.Color) {
	var index = indexRune(font.fonts, ch)
	if index < 0 {
		return
	}
	var fontData = font.data[index]
	var y0 = y
	var x0 = x
	for i := 0; i < len(fontData); i++ {
		var data = fontData[i]
		x0 = x
		for b := data; b > 0; b <<= 1 {
			if b&0x80 != 0 {
				img.Set(x0, y0, color)
			}
			x0++
		}
		y0++
		if (y0 - y) >= font.h {
			y0 = y
			x += 8
		}
	}
}

func DrawString(img DrawPoint, str string, x int, y int, font *Font, color color.Color) {
	for _, v := range str {
		DrawChar(img, rune(v), x, y, font, color)
		x += font.w
	}
}
