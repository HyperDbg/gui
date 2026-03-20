package driver

import (
	"bytes"
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"syscall"
	"unsafe"

	"github.com/ddkwork/golibrary/std/mylog"
	"golang.org/x/sys/windows"
)

const (
	BufferSize = 0x1000
)

// good  C:\Users\Admin\go\pkg\mod\golang.org\x\sys@v0.42.0\windows\zerrors_windows.go
// C:\Program Files\Go\src\syscall\syscall_windows.go
// procFormatMessageW                     = modkernel32.NewProc("FormatMessageW")
type api interface {
	Install()
	Uninstall()
	Start()
	Stop()
	IsConnected() bool
	Close()
	Send(buffer *bytes.Buffer, ioctlCode uint32)
	Receive(ioctlCode uint32) *bytes.Buffer
}

type Provider struct {
	scManager    windows.Handle // 服务控制管理器句柄，用于管理驱动的安装、卸载、启动、停止
	deviceHandle syscall.Handle // 设备文件句柄，用于通过 DeviceIoControl 与驱动程序通信
	DriverPath   string         // 驱动程序文件的完整路径
	DeviceName   string         // 设备名称，用于打开设备文件进行通信
	ServiceName  string         // 服务名称，用于在 Windows 服务管理器中标识驱动服务
}

func New(driverPath, serviceName, deviceName string) api {

	if serviceName == "" {
		panic("serviceName 不能为空")
	}

	if deviceName == "" {
		panic("deviceName 不能为空")
	}

	if !strings.HasPrefix(deviceName, `\\.\`) {
		panic(fmt.Sprintf("deviceName 格式错误，必须以 \\\\.\\ 开头，当前值: %s", deviceName))
	}

	absPath := filepath.ToSlash(mylog.Check2(filepath.Abs(driverPath)))
	scManager := mylog.Check2(windows.OpenSCManager(nil, nil, windows.SC_MANAGER_ALL_ACCESS))

	return &Provider{
		scManager:    scManager,
		deviceHandle: syscall.Handle(windows.InvalidHandle),
		DeviceName:   deviceName,
		ServiceName:  serviceName,
		DriverPath:   absPath,
	}
}

func (p *Provider) Install() {
	var schService windows.Handle
	var lastError error

	schService, lastError = windows.CreateService(
		p.scManager,
		windows.StringToUTF16Ptr(p.ServiceName),
		windows.StringToUTF16Ptr(p.ServiceName),
		windows.SERVICE_ALL_ACCESS,
		windows.SERVICE_KERNEL_DRIVER,
		windows.SERVICE_DEMAND_START,
		windows.SERVICE_ERROR_NORMAL,
		windows.StringToUTF16Ptr(p.DriverPath),
		nil,
		nil,
		nil,
		nil,
		nil,
	)

	if lastError != nil {
		switch {
		case errors.Is(lastError, windows.ERROR_SERVICE_EXISTS):
			mylog.Info("", "服务（驱动）已存在")
			mylog.Info("", "尝试先移除旧的驱动实例")

			p.Stop()
			p.Uninstall()

			mylog.Info("", "重新安装驱动")
			p.Install()
			return
		case errors.Is(lastError, windows.ERROR_SERVICE_MARKED_FOR_DELETE):
			mylog.Info("", "服务标记为删除，尝试强制删除")
			p.Uninstall()
			mylog.Info("", "重新安装驱动")
			p.Install()
			return
		default:
			mylog.Check(lastError)
			return
		}
	}

	if schService != windows.Handle(0) {
		mylog.Check(windows.CloseServiceHandle(schService))
	}

	mylog.Success("", "驱动安装成功")
}

func (p *Provider) Uninstall() {
	schService := mylog.Check2(windows.OpenService(p.scManager, windows.StringToUTF16Ptr(p.ServiceName), windows.SERVICE_ALL_ACCESS))

	err := windows.DeleteService(schService)
	if err != nil && !errors.Is(err, windows.ERROR_SERVICE_MARKED_FOR_DELETE) {
		mylog.Check(err)
	}

	if schService != windows.Handle(0) {
		mylog.Check(windows.CloseServiceHandle(schService))
	}
}

func (p *Provider) Start() {
	schService := mylog.Check2(windows.OpenService(p.scManager, windows.StringToUTF16Ptr(p.ServiceName), windows.SERVICE_ALL_ACCESS))

	err := windows.StartService(schService, 0, nil)
	if err != nil {
		switch {
		case errors.Is(err, windows.ERROR_SERVICE_ALREADY_RUNNING):
		case errors.Is(err, windows.ERROR_PATH_NOT_FOUND):
			mylog.Info("", "错误，找不到驱动路径，或对驱动文件的访问受限")
			mylog.Info("", "大多数情况下，这是因为杀毒软件尚未完成对驱动的扫描，")
			mylog.Info("", "因此，如果您尝试再次加载驱动（重新输入上一个命令），问题将得到解决")
			mylog.Check(err)
		case errors.Is(err, windows.ERROR_INVALID_IMAGE_HASH):
			mylog.Info("", "错误，加载驱动失败")
			mylog.Info("", "这是因为启用了驱动程序签名强制执行或 HVCI 阻止了驱动程序加载")
			mylog.Info("", "您应该通过附加 WinDbg 或从启动菜单禁用驱动程序签名强制执行")
			mylog.Info("", "如果禁用了驱动程序签名强制执行，HVCI 可能会阻止驱动程序加载")
			mylog.Info("", "HyperDbg 不兼容基于虚拟化的安全性 (VBS)")
			mylog.Info("", "请遵循以下说明：https://docs.hyperdbg.org/getting-started/build-and-install")
			mylog.Check(err)
		case err.Error() == "A certificate was explicitly revoked by its issuer.":
			mylog.Info("", "错误，驱动证书已被吊销")
			mylog.Info("", "这是因为驱动程序的数字证书已被吊销，无法加载")
			mylog.Info("", "请使用有效的签名证书重新编译驱动程序")
			mylog.Info("", "或者在测试环境中禁用驱动程序签名强制执行")
			mylog.Check(err)
		default:
			mylog.Check(err)
		}
	}

	if schService != windows.Handle(0) {
		mylog.Check(windows.CloseServiceHandle(schService))
	}

	if p.deviceHandle == syscall.Handle(windows.InvalidHandle) {
		deviceNamePtr := mylog.Check2(syscall.UTF16PtrFromString(p.DeviceName))
		handle := mylog.Check2(windows.CreateFile(
			deviceNamePtr,
			windows.GENERIC_READ|windows.GENERIC_WRITE,
			windows.FILE_SHARE_READ|windows.FILE_SHARE_WRITE,
			nil,
			windows.OPEN_EXISTING,
			windows.FILE_ATTRIBUTE_NORMAL,
			0,
		))
		p.deviceHandle = syscall.Handle(handle)
	}
}

func (p *Provider) Stop() {
	schService := mylog.Check2(windows.OpenService(p.scManager, windows.StringToUTF16Ptr(p.ServiceName), windows.SERVICE_ALL_ACCESS))

	var serviceStatus windows.SERVICE_STATUS
	err := windows.ControlService(schService, windows.SERVICE_CONTROL_STOP, &serviceStatus)
	if err != nil && !errors.Is(err, windows.ERROR_SERVICE_NOT_ACTIVE) {
		mylog.Check(err)
	}

	if schService != windows.Handle(0) {
		mylog.Check(windows.CloseServiceHandle(schService))
	}
}

func (p *Provider) IsConnected() bool { return p.deviceHandle != syscall.InvalidHandle }

func (p *Provider) Close() {
	if p.deviceHandle != syscall.InvalidHandle {
		mylog.Check(windows.CloseHandle(windows.Handle(p.deviceHandle)))
		p.deviceHandle = syscall.InvalidHandle
	}
	if p.scManager != windows.Handle(0) {
		mylog.Check(windows.CloseServiceHandle(p.scManager))
		p.scManager = windows.Handle(0)
	}
}

func (p *Provider) Send(buffer *bytes.Buffer, ioctlCode uint32) {
	var bytesReturned uint32
	mylog.Check(windows.DeviceIoControl(
		windows.Handle(p.deviceHandle),
		ioctlCode,
		unsafe.SliceData(buffer.Bytes()),
		uint32(buffer.Len()),
		unsafe.SliceData(buffer.Bytes()),
		uint32(buffer.Len()),
		&bytesReturned,
		nil,
	))
}

func (p *Provider) Receive(ioctlCode uint32) *bytes.Buffer {
	buffer := make([]byte, BufferSize)
	var bytesReturned uint32
	mylog.Check(windows.DeviceIoControl(
		windows.Handle(p.deviceHandle),
		ioctlCode,
		nil,
		0,
		unsafe.SliceData(buffer),
		uint32(len(buffer)),
		&bytesReturned,
		nil,
	))
	return bytes.NewBuffer(buffer[:bytesReturned])
}
