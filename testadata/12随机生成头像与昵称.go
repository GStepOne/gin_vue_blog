package main

import (
	"fmt"
	"github.com/DanPlayer/randomname"
	"github.com/disintegration/letteravatar"
	"github.com/golang/freetype"
	"image/png"
	"log"
	"os"
	"path"
	"unicode/utf8"
)

func main() {
	//name := randomname.GenerateName()
	//fmt.Println(name)
	//
	//img, err := letteravatar.Draw(100, 'A', &letteravatar.Options{
	//	Palette: []color.Color{
	//		color.RGBA{255, 0, 0, 255},
	//		color.RGBA{0, 255, 0, 255},
	//		color.RGBA{0, 0, 255, 255},
	//	},
	//})
	GenerateNameAvatar()
	//DrawAvatar("李日天", "uploads/avatar")

	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(img)
}

func GenerateNameAvatar() {
	dir := "uploads/avatar"
	for _, s := range randomname.AdjectiveSlice {
		DrawAvatar(s, dir)
	}

	for _, s := range randomname.PersonSlice {
		DrawAvatar(s, dir)
	}
}

func DrawAvatar(name string, dir string) {
	fontFile, err := os.ReadFile("uploads/font/hanyifengbomili.ttf")
	fmt.Println(err)
	font, err := freetype.ParseFont(fontFile)
	options := &letteravatar.Options{
		Font: font,
		//Palette: []color.Color{
		//	color.RGBA{255, 0, 0, 255},
		//	color.RGBA{0, 255, 0, 255},
		//	color.RGBA{0, 0, 255, 255},
		//},
	}
	//绘制文字
	firstLetter, _ := utf8.DecodeRuneInString(name)
	img, err := letteravatar.Draw(140, firstLetter, options)
	if err != nil {
		log.Fatal(err)
	}

	filePath := path.Join(dir, name+".png")
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = png.Encode(file, img)
	if err != nil {
		fmt.Println("对不起啊生成失败了")
	}
}

func GenerateName() string {
	name := randomname.GenerateName()
	return name
}
