package main

type (
	USHORT uint16
	ULONG  uint32
	BYTE   uint8
	CHAR   byte
)

// IDSECTOR 结构体定义
type IDSECTOR struct {
	wGenConfig                 USHORT
	wNumCyls                   USHORT
	wReserved                  USHORT
	wNumHeads                  USHORT
	wBytesPerTrack             USHORT
	wBytesPerSector            USHORT
	wSectorsPerTrack           USHORT
	wVendorUnique              [3]USHORT
	sSerialNumber              [20]CHAR
	wBufferType                USHORT
	wBufferSize                USHORT
	wECCSize                   USHORT
	sFirmwareRev               [8]CHAR
	sModelNumber               [40]CHAR
	wMoreVendorUnique          USHORT
	wDoubleWordIO              USHORT
	wCapabilities              USHORT
	wReserved1                 USHORT
	wPIOTiming                 USHORT
	wDMATiming                 USHORT
	wBS                        USHORT
	wNumCurrentCyls            USHORT
	wNumCurrentHeads           USHORT
	wNumCurrentSectorsPerTrack USHORT
	ulCurrentSectorCapacity    ULONG
	wMultSectorStuff           USHORT
	ulTotalAddressableSectors  ULONG
	wSingleWordDMA             USHORT
	wMultiWordDMA              USHORT
	bReserved                  [128]BYTE
}

type BitField struct {
	data [16]byte // 支持到16字节
}

// 从偏移位置设置指定字节数的值
func (b *BitField) SetBytesFromValue(offset int, byteCount int, value uint64) {
	if offset+byteCount > 128 {
		panic("Field exceeds the maximum size of 128 bits")
	}

	// 逐字节设置
	for i := range byteCount {
		b.data[offset+i] = uint8(value >> (8 * (byteCount - 1 - i))) // 从高字节到低字节
	}
}

// 从偏移位置获取指定字节数的值
func (b *BitField) GetBytes(offset int, byteCount int) uint64 {
	if offset+byteCount > 128 {
		panic("Field exceeds the maximum size of 128 bits")
	}
	var result uint64 = 0

	// 逐字节获取
	for i := range byteCount {
		result |= uint64(b.data[offset+i]) << (8 * (byteCount - 1 - i)) // 从高字节到低字节
	}
	return result
}
