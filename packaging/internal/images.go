package internal

import (
	"bytes"
	_ "embed"
	"image"

	"github.com/ddkwork/golibrary/mylog"
)

//go:embed embedded/app-1024.png
var appImgBytes []byte
var appImg image.Image

//go:embed embedded/doc-1024.png
var docImgBytes []byte
var docImg image.Image

func loadBaseImages() {
	appImg, _ = mylog.Check3(image.Decode(bytes.NewBuffer(appImgBytes)))
	docImg, _ = mylog.Check3(image.Decode(bytes.NewBuffer(docImgBytes)))
}
