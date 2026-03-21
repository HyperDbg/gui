#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <windows.h>
#include <windef.h>

// 定义数据类型
typedef unsigned char uint8;
typedef unsigned short uint16;
typedef unsigned int uint32;

#pragma pack(1) // 设置1字节对齐
typedef struct _IDSector {
    uint16 wGenConfig;
    uint16 wNumCyls;
    uint16 wReserved;
    uint16 wNumHeads;
    uint16 wBytesPerTrack;
    uint16 wBytesPerSector;
    uint16 wSectorsPerTrack;
    uint16 wVendorUnique[3];
    char   sSerialNumber[20];
    uint16 wBufferType;
    uint16 wBufferSize;
    uint16 wECCSize;
    char   sFirmwareRev[8];
    char   sModelNumber[40];
    uint16 wMoreVendorUnique;
    uint16 wDoubleWordIO;
    uint16 wCapabilities;
    uint16 wReserved1;
    uint16 wPIOTiming;
    uint16 wDMATiming;
    uint16 wBS;
    uint16 wNumCurrentCyls;
    uint16 wNumCurrentHeads;
    uint16 wNumCurrentSectorsPerTrack;
    uint32 ulCurrentSectorCapacity;
    uint16 wMultSectorStuff;
    uint32 ulTotalAddressableSectors;
    uint16 wSingleWordDMA;
    uint16 wMultiWordDMA;
    uint8  bReserved[128];
} IDSECTOR, *PIDSECTOR;

typedef struct HDiskInfo_ {
    uint8  module[128];
    uint8  firmware[128];
    uint8  serialno[32];
    uint32 capacity;
} HDiskInfo;
#pragma pack()  // 恢复默认对齐

#define DFP_GET_VERSION        0x00074080
#define DFP_SEND_DRIVE_COMMAND 0x0007c084
#define DFP_RECEIVE_DRIVE_DATA 0x0007c088

#pragma pack(1) // 设置1字节对齐
typedef struct _GETVERSIONOUTPARAMS {
    BYTE  bVersion;
    BYTE  bRevision;
    BYTE  bReserved;
    BYTE  bIDEDeviceMap;
    DWORD fCapabilities;
    DWORD dwReserved[4];
} GETVERSIONOUTPARAMS, *PGETVERSIONOUTPARAMS, *LPGETVERSIONOUTPARAMS;
#pragma pack()

void ChangeByteOrder(uint8* szString, int uscStrSize) {
    int i;
    uint8 temp;

    for (i = 0; i < uscStrSize; i += 2) {
        temp = szString[i];
        szString[i] = szString[i + 1];
        szString[i + 1] = temp;
    }
}

void hdid9x(HDiskInfo* pinfo) {
    PIDSECTOR phdinfo;
    GETVERSIONOUTPARAMS vers;
    SENDCMDINPARAMS in;
    SENDCMDOUTPARAMS out;
    HANDLE h;
    DWORD i;
    BYTE j;

    ZeroMemory(&vers, sizeof(vers));
    h = CreateFile("\\\\.\\Smartvsd", 0, 0, 0, CREATE_NEW, 0, 0);
    if (h == INVALID_HANDLE_VALUE) {
        return;
    }

    if (!DeviceIoControl(h, DFP_GET_VERSION, 0, 0, &vers, sizeof(vers), &i, 0)) {
        CloseHandle(h);
        return;
    }

    if (!(vers.fCapabilities & 1)) {
        CloseHandle(h);
        return;
    }

    for (j = 0; j < 4; j++) {
        ZeroMemory(&in, sizeof(in));
        ZeroMemory(&out, sizeof(out));
        if (j & 1) {
            in.irDriveRegs.bDriveHeadReg = 0xb0;
        } else {
            in.irDriveRegs.bDriveHeadReg = 0xa0;
        }
        if (vers.fCapabilities & (16 >> j)) {
            continue;
        } else {
            in.irDriveRegs.bCommandReg = 0xec;
        }
        in.bDriveNumber = j;
        in.irDriveRegs.bSectorCountReg = 1;
        in.irDriveRegs.bSectorNumberReg = 1;
        in.cBufferSize = 512;
        if (!DeviceIoControl(h, DFP_RECEIVE_DRIVE_DATA, &in, sizeof(in), &out, sizeof(out), &i, 0)) {
            CloseHandle(h);
            return;
        }
        phdinfo = (PIDSECTOR)out.bBuffer;
        memcpy(pinfo->module, phdinfo->sModelNumber, 40);
        memcpy(pinfo->firmware, phdinfo->sFirmwareRev, 8);
        memcpy(pinfo->serialno, phdinfo->sSerialNumber, 20);
        pinfo->capacity = phdinfo->ulTotalAddressableSectors / 2 / 1024;
    }

    CloseHandle(h);
}

void hdidnt(HDiskInfo* pinfo) {
    char hd[80];
    PIDSECTOR phdinfo;
    GETVERSIONOUTPARAMS vers;
    SENDCMDINPARAMS in;
    SENDCMDOUTPARAMS out;
    HANDLE h;
    DWORD i;
    BYTE j;

    ZeroMemory(&vers, sizeof(vers));
    for (j = 0; j < 4; j++) {
        sprintf(hd, "\\\\.\\PhysicalDrive%d", j);
        h = CreateFile(hd, GENERIC_READ | GENERIC_WRITE, FILE_SHARE_READ | FILE_SHARE_WRITE, 0, OPEN_EXISTING, 0, 0);
        if (h == INVALID_HANDLE_VALUE) {
            continue;
        }
        if (!DeviceIoControl(h, DFP_GET_VERSION, 0, 0, &vers, sizeof(vers), &i, 0)) {
            CloseHandle(h);
            continue;
        }

        if (!(vers.fCapabilities & 1)) {
            CloseHandle(h);
            return;
        }

        ZeroMemory(&in, sizeof(in));
        ZeroMemory(&out, sizeof(out));
        if (j & 1) {
            in.irDriveRegs.bDriveHeadReg = 0xb0;
        } else {
            in.irDriveRegs.bDriveHeadReg = 0xa0;
        }
        if (vers.fCapabilities & (16 >> j)) {
            continue;
        } else {
            in.irDriveRegs.bCommandReg = 0xec;
        }
        in.bDriveNumber = j;
        in.irDriveRegs.bSectorCountReg = 1;
        in.irDriveRegs.bSectorNumberReg = 1;
        in.cBufferSize = 512;
        if (!DeviceIoControl(h, DFP_RECEIVE_DRIVE_DATA, &in, sizeof(in), &out, sizeof(out), &i, 0)) {
            CloseHandle(h);
            return;
        }

        phdinfo = (PIDSECTOR)out.bBuffer;
        memcpy(pinfo->module, phdinfo->sModelNumber, 40);
        memcpy(pinfo->firmware, phdinfo->sFirmwareRev, 8);
        memcpy(pinfo->serialno, phdinfo->sSerialNumber, 20);
        pinfo->capacity = phdinfo->ulTotalAddressableSectors / 2 / 1024;

        CloseHandle(h);
    }
    return;
}

int read_harddisk_info(HDiskInfo* pinfo) {
    OSVERSIONINFO VersionInfo;

    if (!pinfo) return -1;
    memset(pinfo, 0, sizeof(HDiskInfo));

    ZeroMemory(&VersionInfo, sizeof(VersionInfo));
    VersionInfo.dwOSVersionInfoSize = sizeof(VersionInfo);
    GetVersionEx(&VersionInfo);

    switch (VersionInfo.dwPlatformId) {
    case VER_PLATFORM_WIN32s:
        return -100;
    case VER_PLATFORM_WIN32_WINDOWS:
        hdid9x(pinfo);
        break;
    case VER_PLATFORM_WIN32_NT:
        hdidnt(pinfo);
        break;
    }
    return 0;
}

int main() {
    HDiskInfo hddInfo;
    int status = read_harddisk_info(&hddInfo);

    if (status == 0) {
        printf("硬盘信息获取成功:\n");
        printf("型号: %s\n", hddInfo.module);
        printf("固件版本: %s\n", hddInfo.firmware);
        printf("序列号: %s\n", hddInfo.serialno);
        printf("容量: %u MB\n", hddInfo.capacity);
    } else {
        printf("硬盘信息获取失败，错误码: %d\n", status);
    }

    return 0;
}
