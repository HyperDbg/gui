package main

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"github.com/ddkwork/HyperDbg/walker/hardwareinfo"
	"github.com/ddkwork/ux/app"
	"github.com/ddkwork/ux/wdk"
	"github.com/ddkwork/ux/widget/button"
	"github.com/ddkwork/ux/widget/card"
	"github.com/ddkwork/ux/widget/chip"
	"github.com/ddkwork/ux/widget/logview"
	"golang.org/x/exp/shiny/materialdesign/icons"
)

type HardwareInfoUI struct {
	hw *hardwareinfo.HardwareInfo

	refreshButton *button.Button
	statusChip    *chip.Chip
	logView       *logview.LogView

	cpuInfo *cpuInfoPanel
	ssdInfo *ssdInfoPanel
	macInfo *macInfoPanel
}

type cpuInfoPanel struct {
	vendorChip   *chip.Chip
	modelChip    *chip.Chip
	coresChip    *chip.Chip
	featuresChip *chip.Chip
}

type ssdInfoPanel struct {
	modelChip   *chip.Chip
	serialChip  *chip.Chip
	versionChip *chip.Chip
}

type macInfoPanel struct {
	addrChip *chip.Chip
}

func NewHardwareInfoUI() *HardwareInfoUI {
	d := &HardwareInfoUI{}
	d.hw = hardwareinfo.New()

	d.refreshButton = button.Filled()
	d.statusChip = chip.Assist()
	d.logView = logview.New()

	d.cpuInfo = &cpuInfoPanel{
		vendorChip:   chip.Assist(),
		modelChip:    chip.Assist(),
		coresChip:    chip.Assist(),
		featuresChip: chip.Assist(),
	}

	d.ssdInfo = &ssdInfoPanel{
		modelChip:   chip.Assist(),
		serialChip:  chip.Assist(),
		versionChip: chip.Assist(),
	}

	d.macInfo = &macInfoPanel{
		addrChip: chip.Assist(),
	}

	return d
}

func (d *HardwareInfoUI) Update(gtx layout.Context) {
	if d.refreshButton.Clicked(gtx) {
		d.logView.Info("Refreshing hardware information...")

		if d.hw.CpuInfo.Get() {
			d.logView.Info("CPU info retrieved")
		}

		if d.hw.SsdInfo.Get() {
			d.logView.Info("SSD info retrieved")
		}

		if d.hw.MacInfo.Get() {
			d.logView.Info("MAC info retrieved")
		}

		d.logView.Info("Refresh complete")
	}
}

func (d *HardwareInfoUI) Layout(gtx layout.Context) layout.Dimensions {
	materialTheme := wdk.GetMaterialTheme(gtx)
	paint.FillShape(gtx.Ops, materialTheme.Scheme.Background.Color.AsNRGBA(), clip.Rect(image.Rect(0, 0, gtx.Constraints.Max.X, gtx.Constraints.Max.Y)).Op())

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return d.layoutHeader(gtx)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(unit.Dp(16)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Vertical, Spacing: layout.SpaceBetween}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return d.layoutCPUCard(gtx)
					}),
					layout.Rigid(layout.Spacer{Height: unit.Dp(12)}.Layout),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return d.layoutSSDCard(gtx)
					}),
					layout.Rigid(layout.Spacer{Height: unit.Dp(12)}.Layout),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return d.layoutMACCard(gtx)
					}),
					layout.Rigid(layout.Spacer{Height: unit.Dp(12)}.Layout),
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						return d.layoutLogCard(gtx)
					}),
				)
			})
		}),
	)
}

func (d *HardwareInfoUI) layoutHeader(gtx layout.Context) layout.Dimensions {
	materialTheme := wdk.GetMaterialTheme(gtx)
	return layout.UniformInset(unit.Dp(16)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				icon := wdk.RequireIconWidget(icons.HardwareDesktopWindows)
				return icon(gtx, materialTheme.Scheme.Background.OnColor)
			}),
			layout.Rigid(layout.Spacer{Width: unit.Dp(12)}.Layout),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return wdk.TitleL(gtx, "Hardware Info")
			}),
			layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				return layout.Dimensions{}
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return d.refreshButton.Layout(gtx, "Refresh")
			}),
		)
	})
}

func (d *HardwareInfoUI) layoutCPUCard(gtx layout.Context) layout.Dimensions {
	c := &card.Card{}
	c.Kind = card.Elevated
	return c.Layout(gtx, card.Content(func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(unit.Dp(16)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return wdk.LabelM(gtx, "CPU Information")
				}),
				layout.Rigid(layout.Spacer{Height: unit.Dp(12)}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return d.layoutInfoRow(gtx, "Vendor", d.cpuInfo.vendorChip)
				}),
				layout.Rigid(layout.Spacer{Height: unit.Dp(8)}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return d.layoutInfoRow(gtx, "Model", d.cpuInfo.modelChip)
				}),
				layout.Rigid(layout.Spacer{Height: unit.Dp(8)}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return d.layoutInfoRow(gtx, "Cores", d.cpuInfo.coresChip)
				}),
			)
		})
	}))
}

func (d *HardwareInfoUI) layoutSSDCard(gtx layout.Context) layout.Dimensions {
	c := &card.Card{}
	c.Kind = card.Elevated
	return c.Layout(gtx, card.Content(func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(unit.Dp(16)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return wdk.LabelM(gtx, "SSD Information")
				}),
				layout.Rigid(layout.Spacer{Height: unit.Dp(12)}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return d.layoutInfoRow(gtx, "Model", d.ssdInfo.modelChip)
				}),
				layout.Rigid(layout.Spacer{Height: unit.Dp(8)}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return d.layoutInfoRow(gtx, "Serial", d.ssdInfo.serialChip)
				}),
				layout.Rigid(layout.Spacer{Height: unit.Dp(8)}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return d.layoutInfoRow(gtx, "Version", d.ssdInfo.versionChip)
				}),
			)
		})
	}))
}

func (d *HardwareInfoUI) layoutMACCard(gtx layout.Context) layout.Dimensions {
	c := &card.Card{}
	c.Kind = card.Elevated
	return c.Layout(gtx, card.Content(func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(unit.Dp(16)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return wdk.LabelM(gtx, "Network Information")
				}),
				layout.Rigid(layout.Spacer{Height: unit.Dp(12)}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return d.layoutInfoRow(gtx, "MAC Address", d.macInfo.addrChip)
				}),
			)
		})
	}))
}

func (d *HardwareInfoUI) layoutInfoRow(gtx layout.Context, label string, valueChip *chip.Chip) layout.Dimensions {
	return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return wdk.LabelS(gtx, label)
		}),
		layout.Rigid(layout.Spacer{Width: unit.Dp(12)}.Layout),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return valueChip.Layout(gtx, "-")
		}),
	)
}

func (d *HardwareInfoUI) layoutLogCard(gtx layout.Context) layout.Dimensions {
	c := &card.Card{}
	c.Kind = card.Elevated
	return c.Layout(gtx, card.Content(func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(unit.Dp(16)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return wdk.LabelM(gtx, "Log")
				}),
				layout.Rigid(layout.Spacer{Height: unit.Dp(12)}.Layout),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return d.logView.Layout(gtx)
				}),
			)
		})
	}))
}

func main() {
	d := NewHardwareInfoUI()
	app.Run("Hardware Info", func(gtx layout.Context) {
		d.Update(gtx)
		d.Layout(gtx)
	})
}
