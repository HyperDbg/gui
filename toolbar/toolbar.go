package toolbar

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/ddkwork/golibrary/src/fynelib/canvasobjectapi"
)

type (
	Interface interface {
		canvasobjectapi.Interface
		//Fn() (ok bool)
	}
	res interface {
		i1() fyne.Resource
		i2() fyne.Resource
		i3() fyne.Resource
		i4() fyne.Resource
		i5() fyne.Resource
		i6() fyne.Resource
		i7() fyne.Resource
		i8() fyne.Resource
		i9() fyne.Resource
		i10() fyne.Resource
		i11() fyne.Resource
		i12() fyne.Resource
		i13() fyne.Resource
		i14() fyne.Resource
		i15() fyne.Resource
		i16() fyne.Resource
		i17() fyne.Resource
		i18() fyne.Resource
		i19() fyne.Resource
		i20() fyne.Resource
		i21() fyne.Resource
		i22() fyne.Resource
		i23() fyne.Resource
		i24() fyne.Resource
		i25() fyne.Resource
		i26() fyne.Resource
		i27() fyne.Resource
	} //todo delete
	resObj struct {
	}
	object struct {
	}
)

func (o *object) CanvasObject(window fyne.Window) fyne.CanvasObject {
	r := newResObj()
	return widget.NewToolbar(
		widget.NewToolbarAction(r.i1(), func() {}),
		widget.NewToolbarAction(r.i2(), func() {}),
		widget.NewToolbarAction(r.i3(), func() {}),
		widget.NewToolbarAction(r.i4(), func() {}),
		widget.NewToolbarAction(r.i5(), func() {}),
		widget.NewToolbarAction(r.i6(), func() {}),
		widget.NewToolbarAction(r.i7(), func() {}),
		widget.NewToolbarAction(r.i8(), func() {}),
		widget.NewToolbarAction(r.i9(), func() {}),
		widget.NewToolbarAction(r.i10(), func() {}),
		widget.NewToolbarAction(r.i11(), func() {}),
		widget.NewToolbarAction(r.i12(), func() {}),
		widget.NewToolbarAction(r.i13(), func() {}),
		widget.NewToolbarAction(r.i14(), func() {}),
		widget.NewToolbarAction(r.i15(), func() {}),
		widget.NewToolbarAction(r.i16(), func() {}),
		widget.NewToolbarAction(r.i17(), func() {}),
		widget.NewToolbarAction(r.i18(), func() {}),
		widget.NewToolbarAction(r.i19(), func() {}),
		widget.NewToolbarAction(r.i20(), func() {}),
		widget.NewToolbarAction(r.i21(), func() {}),
		widget.NewToolbarAction(r.i22(), func() {}),
		widget.NewToolbarAction(r.i23(), func() {}),
		widget.NewToolbarAction(r.i24(), func() {}),
		widget.NewToolbarAction(r.i25(), func() {}),
		widget.NewToolbarAction(r.i26(), func() {}),
		widget.NewToolbarAction(r.i27(), func() {}),
	)
}

func New() Interface { return &object{} }

func newResObj() res {
	return &resObj{}
}

func (r *resObj) i1() fyne.Resource  { return fyne.NewStaticResource("i1", i1) }
func (r *resObj) i2() fyne.Resource  { return fyne.NewStaticResource("i2", i2) }
func (r *resObj) i3() fyne.Resource  { return fyne.NewStaticResource("i3", i3) }
func (r *resObj) i4() fyne.Resource  { return fyne.NewStaticResource("i4", i4) }
func (r *resObj) i5() fyne.Resource  { return fyne.NewStaticResource("i5", i5) }
func (r *resObj) i6() fyne.Resource  { return fyne.NewStaticResource("i6", i6) }
func (r *resObj) i7() fyne.Resource  { return fyne.NewStaticResource("i7", i7) }
func (r *resObj) i8() fyne.Resource  { return fyne.NewStaticResource("i8", i8) }
func (r *resObj) i9() fyne.Resource  { return fyne.NewStaticResource("i9", i9) }
func (r *resObj) i10() fyne.Resource { return fyne.NewStaticResource("i0", i10) }
func (r *resObj) i11() fyne.Resource { return fyne.NewStaticResource("i11", i11) }
func (r *resObj) i12() fyne.Resource { return fyne.NewStaticResource("i12", i12) }
func (r *resObj) i13() fyne.Resource { return fyne.NewStaticResource("i13", i13) }
func (r *resObj) i14() fyne.Resource { return fyne.NewStaticResource("i14", i14) }
func (r *resObj) i15() fyne.Resource { return fyne.NewStaticResource("i15", i15) }
func (r *resObj) i16() fyne.Resource { return fyne.NewStaticResource("i16", i16) }
func (r *resObj) i17() fyne.Resource { return fyne.NewStaticResource("i17", i17) }
func (r *resObj) i18() fyne.Resource { return fyne.NewStaticResource("i18", i18) }
func (r *resObj) i19() fyne.Resource { return fyne.NewStaticResource("i19", i19) }
func (r *resObj) i20() fyne.Resource { return fyne.NewStaticResource("i20", i20) }
func (r *resObj) i21() fyne.Resource { return fyne.NewStaticResource("i21", i21) }
func (r *resObj) i22() fyne.Resource { return fyne.NewStaticResource("i22", i22) }
func (r *resObj) i23() fyne.Resource { return fyne.NewStaticResource("i23", i23) }
func (r *resObj) i24() fyne.Resource { return fyne.NewStaticResource("i24", i24) }
func (r *resObj) i25() fyne.Resource { return fyne.NewStaticResource("i25", i25) }
func (r *resObj) i26() fyne.Resource { return fyne.NewStaticResource("i26", i26) }
func (r *resObj) i27() fyne.Resource { return fyne.NewStaticResource("i27", i27) }

var (
	//go:embed "OllyICE 1.10/(1).jpg"
	i1 []byte

	//go:embed "OllyICE 1.10/(2).jpg"
	i2 []byte

	//go:embed "OllyICE 1.10/(3).jpg"
	i3 []byte

	//go:embed "OllyICE 1.10/(4).jpg"
	i4 []byte

	//go:embed "OllyICE 1.10/(5).jpg"
	i5 []byte

	//go:embed "OllyICE 1.10/(6).jpg"
	i6 []byte

	//go:embed "OllyICE 1.10/(7).jpg"
	i7 []byte

	//go:embed "OllyICE 1.10/(8).jpg"
	i8 []byte

	//go:embed "OllyICE 1.10/(9).jpg"
	i9 []byte

	//go:embed "OllyICE 1.10/(10).jpg"
	i10 []byte

	//go:embed "OllyICE 1.10/(11).jpg"
	i11 []byte

	//go:embed "OllyICE 1.10/(12).jpg"
	i12 []byte

	//go:embed "OllyICE 1.10/(13).jpg"
	i13 []byte

	//go:embed "OllyICE 1.10/(14).jpg"
	i14 []byte

	//go:embed "OllyICE 1.10/(15).jpg"
	i15 []byte

	//go:embed "OllyICE 1.10/(16).jpg"
	i16 []byte

	//go:embed "OllyICE 1.10/(17).jpg"
	i17 []byte

	//go:embed "OllyICE 1.10/(18).jpg"
	i18 []byte

	//go:embed "OllyICE 1.10/(19).jpg"
	i19 []byte

	//go:embed "OllyICE 1.10/(20).jpg"
	i20 []byte

	//go:embed "OllyICE 1.10/(21).jpg"
	i21 []byte

	//go:embed "OllyICE 1.10/(22).jpg"
	i22 []byte

	//go:embed "OllyICE 1.10/(23).jpg"
	i23 []byte

	//go:embed "OllyICE 1.10/(24).jpg"
	i24 []byte

	//go:embed "OllyICE 1.10/(25).jpg"
	i25 []byte

	//go:embed "OllyICE 1.10/(26).jpg"
	i26 []byte

	//go:embed "OllyICE 1.10/(27).jpg"
	i27 []byte
)
