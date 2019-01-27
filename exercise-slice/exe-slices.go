package main


//实现 `Pic`。它返回一个 slice 的长度 `dy`，和 slice 中每个元素的长度的 8 位无符号整数 `dx`。当执行这个程序，它会将整数转换为灰度（好吧，蓝度）图片进行展示。
//
//图片的实现已经完成。可能用到的函数包括 (x+y)/2 、 x*y 和 `x^y`（使用 math.Pow 计算最后的函数）。
//
//（需要使用循环来分配 [][]uint8 中的每个 `[]uint8`。）
//
//（使用 uint8(intValue) 在类型之间进行转换。）

import (
	"bytes"
	"image"
	"image/png"
	"io/ioutil"
)

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dx, dy)

	for x:=range a {
		b :=make([]uint8,dx)
		for y:=range b{
			b[y] = uint8(x*y -1)
		}

		a[x]=b
	}

	return a
}

func localShow(f func(int, int) [][]uint8) {
	const (
		dx = 256
		dy = 256
	)
	data := f(dx, dy)
	m := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			m.Pix[i] = v
			m.Pix[i+1] = v
			m.Pix[i+2] = 255
			m.Pix[i+3] = 255
		}
	}

	localmakeImage(m)
}

func localmakeImage(m image.Image) {
	var buf bytes.Buffer
	err := png.Encode(&buf, m)
	if err != nil {
		panic(err)
	}

	_ = ioutil.WriteFile("./output1.jpg", buf.Bytes(), 0666) //buffer输出到jpg文件中

	/*
	enc := base64.StdEncoding.EncodeToString(buf.Bytes()) //buf转成 string
	dist, _ := base64.StdEncoding.DecodeString(enc) //string 转成 byte
	_ = ioutil.WriteFile("./output.jpg", dist, 0666) //buffer输出到jpg文件中
	*/
}

func main() {
	localShow(Pic)
}