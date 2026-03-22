package main

import (
	"image"
	"os"
	"path/filepath"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"github.com/ddkwork/HyperDbg/debugger/driver"
	"github.com/ddkwork/ux/app"
	"github.com/ddkwork/ux/wdk"
	"github.com/ddkwork/ux/widget/button"
	"github.com/ddkwork/ux/widget/card"
	"github.com/ddkwork/ux/widget/chip"
	"github.com/ddkwork/ux/widget/input"
	"github.com/ddkwork/ux/widget/logview"
	"golang.org/x/exp/shiny/materialdesign/icons"
)

type DriverTool struct {
	provider driver.Api

	driverPathInput  *input.Input
	serviceNameInput *input.Input
	deviceNameInput  *input.Input

	installButton    *button.Button
	uninstallButton  *button.Button
	startButton      *button.Button
	stopButton       *button.Button
	connectButton    *button.Button
	disconnectButton *button.Button

	statusChip *chip.Chip
	logView    *logview.LogView

	driverPath  string
	serviceName string
	deviceName  string
	connected   bool
}

func NewDriverTool() *DriverTool {
	d := &DriverTool{}

	d.driverPathInput = input.FilledTextInput()
	d.driverPathInput.LabelText = "Driver Path"

	d.serviceNameInput = input.FilledTextInput()
	d.serviceNameInput.LabelText = "Service Name"

	d.deviceNameInput = input.FilledTextInput()
	d.deviceNameInput.LabelText = "Device Name"

	d.installButton = button.Filled()
	d.uninstallButton = button.Elevated()
	d.startButton = button.Filled()
	d.stopButton = button.Elevated()
	d.connectButton = button.Filled()
	d.disconnectButton = button.Elevated()

	d.statusChip = chip.Assist()
	d.logView = logview.New()

	exePath, _ := os.Executable()
	exeDir := filepath.Dir(exePath)
	defaultDriverPath := filepath.Join(exeDir, "driver.sys")

	d.driverPathInput.Editor.SetText(defaultDriverPath)
	d.serviceNameInput.Editor.SetText("MyDriver")
	d.deviceNameInput.Editor.SetText(`\\.\MyDevice`)

	return d
}

func (d *DriverTool) Update(gtx layout.Context) {
	d.driverPath = d.driverPathInput.Editor.GetText()
	d.serviceName = d.serviceNameInput.Editor.GetText()
	d.deviceName = d.deviceNameInput.Editor.GetText()

	if d.installButton.Clicked(gtx) {
		d.logView.Info("Installing driver...")
		d.provider = driver.New(d.driverPath, d.serviceName, d.deviceName)
		d.provider.Install()
		d.logView.Info("Driver installed successfully")
	}

	if d.uninstallButton.Clicked(gtx) {
		d.logView.Info("Uninstalling driver...")
		if d.provider != nil {
			d.provider.Uninstall()
			d.logView.Info("Driver uninstalled successfully")
		} else {
			d.logView.Warn("No driver loaded")
		}
	}

	if d.startButton.Clicked(gtx) {
		d.logView.Info("Starting driver...")
		if d.provider != nil {
			d.provider.Start()
			d.logView.Info("Driver started successfully")
		} else {
			d.logView.Error("No driver loaded, install first")
		}
	}

	if d.stopButton.Clicked(gtx) {
		d.logView.Info("Stopping driver...")
		if d.provider != nil {
			d.provider.Stop()
			d.logView.Info("Driver stopped successfully")
		} else {
			d.logView.Warn("No driver loaded")
		}
	}

	if d.connectButton.Clicked(gtx) {
		d.logView.Info("Connecting to device...")
		if d.provider != nil {
			d.connected = d.provider.IsConnected()
			if d.connected {
				d.logView.Info("Connected to device successfully")
			} else {
				d.logView.Error("Failed to connect")
			}
		}
	}

	if d.disconnectButton.Clicked(gtx) {
		d.logView.Info("Disconnecting...")
		if d.provider != nil {
			d.provider.Close()
			d.connected = false
			d.logView.Info("Disconnected successfully")
		}
	}

	if d.provider != nil {
		d.connected = d.provider.IsConnected()
	}
}

func (d *DriverTool) Layout(gtx layout.Context) layout.Dimensions {
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
						return d.layoutConfigCard(gtx)
					}),
					layout.Rigid(layout.Spacer{Height: unit.Dp(16)}.Layout),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return d.layoutControlCard(gtx)
					}),
					layout.Rigid(layout.Spacer{Height: unit.Dp(16)}.Layout),
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						return d.layoutLogCard(gtx)
					}),
				)
			})
		}),
	)
}

func (d *DriverTool) layoutHeader(gtx layout.Context) layout.Dimensions {
	materialTheme := wdk.GetMaterialTheme(gtx)
	return layout.UniformInset(unit.Dp(16)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				icon := wdk.RequireIconWidget(icons.HardwareMemory)
				return icon(gtx, materialTheme.Scheme.Background.OnColor)
			}),
			layout.Rigid(layout.Spacer{Width: unit.Dp(12)}.Layout),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return wdk.TitleL(gtx, "Driver Tool")
			}),
			layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				return layout.Dimensions{}
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				if d.connected {
					return d.statusChip.Layout(gtx, "Connected")
				}
				return d.statusChip.Layout(gtx, "Disconnected")
			}),
		)
	})
}

func (d *DriverTool) layoutConfigCard(gtx layout.Context) layout.Dimensions {
	c := &card.Card{}
	c.Kind = card.Elevated
	return c.Layout(gtx, card.Content(func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(unit.Dp(16)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return wdk.LabelM(gtx, "Configuration")
				}),
				layout.Rigid(layout.Spacer{Height: unit.Dp(12)}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return d.driverPathInput.Layout(gtx)
				}),
				layout.Rigid(layout.Spacer{Height: unit.Dp(8)}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return d.serviceNameInput.Layout(gtx)
				}),
				layout.Rigid(layout.Spacer{Height: unit.Dp(8)}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return d.deviceNameInput.Layout(gtx)
				}),
			)
		})
	}))
}

func (d *DriverTool) layoutControlCard(gtx layout.Context) layout.Dimensions {
	c := &card.Card{}
	c.Kind = card.Elevated
	return c.Layout(gtx, card.Content(func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(unit.Dp(16)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return wdk.LabelM(gtx, "Control")
				}),
				layout.Rigid(layout.Spacer{Height: unit.Dp(12)}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return d.installButton.Layout(gtx, "Install")
						}),
						layout.Rigid(layout.Spacer{Width: unit.Dp(8)}.Layout),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return d.uninstallButton.Layout(gtx, "Uninstall")
						}),
						layout.Rigid(layout.Spacer{Width: unit.Dp(16)}.Layout),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return d.startButton.Layout(gtx, "Start")
						}),
						layout.Rigid(layout.Spacer{Width: unit.Dp(8)}.Layout),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return d.stopButton.Layout(gtx, "Stop")
						}),
						layout.Rigid(layout.Spacer{Width: unit.Dp(16)}.Layout),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return d.connectButton.Layout(gtx, "Connect")
						}),
						layout.Rigid(layout.Spacer{Width: unit.Dp(8)}.Layout),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return d.disconnectButton.Layout(gtx, "Disconnect")
						}),
					)
				}),
			)
		})
	}))
}

func (d *DriverTool) layoutLogCard(gtx layout.Context) layout.Dimensions {
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
	d := NewDriverTool()
	app.Run("Driver Tool", func(gtx layout.Context) {
		d.Update(gtx)
		d.Layout(gtx)
	})
}
