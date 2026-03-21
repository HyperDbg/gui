#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <windows.h>
#include <ntddstor.h>
#include <winioctl.h>

typedef unsigned char uint8;
typedef unsigned short uint16;
typedef unsigned int uint32;

#pragma pack(1)
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

int read_harddisk_info(HDiskInfo* pinfo) {
    HANDLE hDevice;
    DWORD bytesReturned;
    BOOL result;

    if (!pinfo) return -1;
    memset(pinfo, 0, sizeof(HDiskInfo));

    hDevice = CreateFile("\\\\.\\PhysicalDrive0",
        GENERIC_READ | GENERIC_WRITE,
        FILE_SHARE_READ | FILE_SHARE_WRITE,
        NULL, OPEN_EXISTING, 0, NULL);

    if (hDevice == INVALID_HANDLE_VALUE) {
        printf("无法打开物理驱动器。错误码: %d\n", GetLastError());
        return -100;
    }

    STORAGE_PROPERTY_QUERY query;
    query.PropertyId = StorageDeviceProperty;
    query.QueryType = PropertyStandardQuery;

    STORAGE_DEVICE_DESCRIPTOR* deviceDescriptor;
    BYTE buffer[512];

    result = DeviceIoControl(
        hDevice,
        IOCTL_STORAGE_QUERY_PROPERTY,
        &query,
        sizeof(STORAGE_PROPERTY_QUERY),
        buffer,
        sizeof(buffer),
        &bytesReturned,
        NULL
    );

    if (!result) {
        printf("无法查询硬盘属性。错误码: %d\n", GetLastError());
        CloseHandle(hDevice);
        return -100;
    }

    deviceDescriptor = (STORAGE_DEVICE_DESCRIPTOR*)buffer;
    if (deviceDescriptor->SerialNumberOffset != 0) {
        strcpy(pinfo->serialno, (char*)(buffer + deviceDescriptor->SerialNumberOffset));
    }
    if (deviceDescriptor->ProductIdOffset != 0) {
        strcpy(pinfo->module, (char*)(buffer + deviceDescriptor->ProductIdOffset));
    }
    if (deviceDescriptor->ProductRevisionOffset != 0) {
        strcpy(pinfo->firmware, (char*)(buffer + deviceDescriptor->ProductRevisionOffset));
    }

    // Querying capacity
    result = DeviceIoControl(
        hDevice,
        IOCTL_DISK_GET_DRIVE_GEOMETRY,
        NULL,
        0,
        buffer,
        sizeof(buffer),
        &bytesReturned,
        NULL
    );

    if (result) {
        DISK_GEOMETRY* geom = (DISK_GEOMETRY*)buffer;
        pinfo->capacity = (geom->Cylinders.QuadPart * geom->TracksPerCylinder *
                          geom->SectorsPerTrack * geom->BytesPerSector) / (1024 * 1024);
    }

    CloseHandle(hDevice);
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
