package main

import "testing"

// 单元测试以验证 IDSECTOR 结构体的字段设置和获取
func TestIDSECTOR(t *testing.T) {
	t.Skip()
	var bf BitField
	var idsektor IDSECTOR

	// 设置 wVendorUnique 字段
	vendorUniqueValues := [3]USHORT{0x0001, 0x0002, 0x0003}
	for i, v := range vendorUniqueValues {
		idsektor.wVendorUnique[i] = v
		bf.SetBytesFromValue(i*2, 2, uint64(v)) // 将每个USHORT写入
	}

	// 获取并验证 wVendorUnique 字段
	gotVendorUnique := [3]USHORT{
		USHORT(bf.GetBytes(0, 2)),
		USHORT(bf.GetBytes(2, 2)),
		USHORT(bf.GetBytes(4, 2)),
	}

	for i := range vendorUniqueValues {
		if gotVendorUnique[i] != vendorUniqueValues[i] {
			t.Errorf("expected wVendorUnique[%d] to be %d, got %d", i, vendorUniqueValues[i], gotVendorUnique[i])
		}
	}

	// 设置 sSerialNumber 字段
	serialNumber := [20]CHAR{'S', 'E', 'R', 'I', 'A', 'L', '0', '0', '0', '0'}
	for i, c := range serialNumber {
		idsektor.sSerialNumber[i] = c
		bf.SetBytesFromValue(6+i, 1, uint64(c))
	}

	// 获取并验证 sSerialNumber 字段
	for i, c := range idsektor.sSerialNumber {
		gotChar := bf.GetBytes(6+i, 1)
		if CHAR(gotChar) != c {
			t.Errorf("expected sSerialNumber[%d] to be %c, got %c", i, c, CHAR(gotChar))
		}
	}

	// 设置 sFirmwareRev 字段
	firmwareRev := [8]CHAR{'F', 'W', '1', '0', ' ', ' ', ' ', ' '}
	for i, c := range firmwareRev {
		idsektor.sFirmwareRev[i] = c
		bf.SetBytesFromValue(26+i, 1, uint64(c))
	}

	// 获取并验证 sFirmwareRev 字段
	for i, c := range idsektor.sFirmwareRev {
		gotChar := bf.GetBytes(26+i, 1)
		if CHAR(gotChar) != c {
			t.Errorf("expected sFirmwareRev[%d] to be %c, got %c", i, c, CHAR(gotChar))
		}
	}

	// 设置 sModelNumber 字段
	modelNumber := [40]CHAR{'M', 'O', 'D', 'E', 'L', '0', '0', '0', '1'}
	for i, c := range modelNumber {
		idsektor.sModelNumber[i] = c
		bf.SetBytesFromValue(34+i, 1, uint64(c))
	}

	// 获取并验证 sModelNumber 字段
	for i, c := range idsektor.sModelNumber {
		gotChar := bf.GetBytes(34+i, 1)
		if CHAR(gotChar) != c {
			t.Errorf("expected sModelNumber[%d] to be %c, got %c", i, c, CHAR(gotChar))
		}
	}
}
