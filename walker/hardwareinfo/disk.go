package hardwareinfo

// func disk() {
//	// 打开硬盘设备
//	h, err := windows.CreateFile(windows.StringToUTF16Ptr("\\\\.\\PhysicalDrive0"), windows.GENERIC_READ|windows.GENERIC_WRITE, windows.FILE_SHARE_READ|windows.FILE_SHARE_WRITE, nil, windows.OPEN_EXISTING, 0, 0)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer windows.CloseHandle(h)
//
//	// 定义输入和输出缓冲区
//	var inBuf [8]byte
//	var outBuf [512]byte
//
//	// 设置输入缓冲区为 SMART 的命令
//	inBuf[0] = 0xEC // SMART READ DATA
//	inBuf[2] = 1    // Sector count
//	inBuf[3] = 1    // LBA low
//	inBuf[4] = 0x4F // LBA mid
//	inBuf[5] = 0xC2 // LBA high
//	inBuf[6] = 0xB0 // Device register
//
//	// 调用 DeviceIoControl 函数
//	var bytesReturned uint32
//	err = windows.DeviceIoControl(h, windows.IOCTL_ATA_PASS_THROUGH, &inBuf[0], uint32(len(inBuf)), &outBuf[0], uint32(len(outBuf)), &bytesReturned, nil)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	// 解析输出缓冲区中的数据，获取序列号和型号
//	var serialNumber [20]byte
//	var modelNumber [40]byte
//
//	copy(serialNumber[:], outBuf[20:40])
//	copy(modelNumber[:], outBuf[54:94])
//
//	fmt.Println("Serial Number:", string(serialNumber[:]))
//	fmt.Println("Model Number:", string(modelNumber[:]))
// }
