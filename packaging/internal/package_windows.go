package internal

import (
	"archive/zip"
	"fmt"
	"image"
	"io"
	"os"
	"strings"
	"time"

	"github.com/ddkwork/golibrary/mylog"
	"github.com/tc-hib/winres"
	"github.com/tc-hib/winres/version"
	"golang.org/x/image/draw"
)

const (
	AppVersion = "v0.0.5"
	AppCmdName = "hyperdbg"
)

func platformPackage() {
	if len(os.Args) > 1 && os.Args[1] == "-z" {
		compress()
		return
	}
	rs := &winres.ResourceSet{}
	rs.SetManifest(winres.AppManifest{
		Description:    "HyperDbg is a debugger ",
		Compatibility:  winres.Win7AndAbove,
		ExecutionLevel: winres.AsInvoker,
		DPIAwareness:   winres.DPIAware,
	})
	addWindowsIcon(rs)
	addWindowsVersion(rs)
	f := mylog.Check2(os.Create("rsrc_windows_amd64.syso"))
	defer func() { mylog.Check(f.Close()) }()
	mylog.Check(rs.WriteObject(f, winres.ArchAMD64))
}

func compress() {
	name := fmt.Sprintf("HyperDbg-%s-windows.zip", "v0.0.5")
	mylog.Check(os.Remove(name))
	f := mylog.Check2(os.Create(name))
	defer func() { mylog.Check(f.Close()) }()
	zw := zip.NewWriter(f)
	var fw io.Writer
	hdr := &zip.FileHeader{
		Name:     "HyperDbg.exe",
		Method:   zip.Deflate,
		Modified: time.Now(),
	}
	hdr.SetMode(0o755)
	mylog.Check2(zw.CreateHeader(hdr))
	in := mylog.Check2(os.Open("HyperDbg.exe"))
	defer func() { mylog.Check(in.Close()) }()
	mylog.Check2(io.Copy(fw, in))
	mylog.Check(zw.Close())
}

func addWindowsIcon(rs *winres.ResourceSet) {
	winIcon := mylog.Check2(winres.NewIconFromImages([]image.Image{Scale(appImg, 256, 256)}))
	mylog.Check(rs.SetIconTranslation(winres.Name("APP"), 0, winIcon))
}

func Scale(img image.Image, width, height int) image.Image {
	bounds := img.Bounds()
	w := bounds.Dx()
	h := bounds.Dy()
	if w == width && h == height {
		return img
	}
	scaled := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.CatmullRom.Scale(scaled, scaled.Bounds(), img, bounds, draw.Over, nil)
	return scaled
}

func addWindowsVersion(rs *winres.ResourceSet) {
	var vi version.Info
	vi.SetFileVersion(AppVersion)
	vi.SetProductVersion(AppVersion)
	CopyrightHolder := "xxx"
	mylog.Check(vi.Set(version.LangDefault, version.CompanyName, CopyrightHolder))
	mylog.Check(vi.Set(version.LangDefault, version.FileDescription, "HyperDbg is a debugger"))
	mylog.Check(vi.Set(version.LangDefault, version.FileVersion, shortAppVersion()))
	mylog.Check(vi.Set(version.LangDefault, version.InternalName, AppCmdName))
	mylog.Check(vi.Set(version.LangDefault, version.LegalCopyright, Copyright()))
	mylog.Check(vi.Set(version.LangDefault, version.LegalTrademarks, "HyperDbg is a debugger"))
	mylog.Check(vi.Set(version.LangDefault, version.OriginalFilename, AppCmdName+".exe"))
	mylog.Check(vi.Set(version.LangDefault, version.ProductName, "HyperDbg"))
	mylog.Check(vi.Set(version.LangDefault, version.ProductVersion, shortAppVersion()))
	rs.SetVersionInfo(vi)
}

func Copyright() string {
	var dot string
	CopyrightHolder := "xxx"
	if !strings.HasSuffix(CopyrightHolder, ".") {
		dot = "."
	}
	return fmt.Sprintf("Copyright © %[1]s by %[2]s%[3]s All rights reserved.", ResolveCopyrightYears(),
		CopyrightHolder, dot)
}

func ResolveCopyrightYears() string {
	return "2021"
	//if CopyrightYears != "" {
	//	return CopyrightYears
	//}
	//years := CopyrightStartYear
	//if CopyrightEndYear != "" && CopyrightEndYear != CopyrightStartYear {
	//	if years == "" {
	//		years = CopyrightEndYear
	//	} else {
	//		years += "-" + CopyrightEndYear
	//	}
	//}
	//return years
}
