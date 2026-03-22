package hardwareinfo

var name = `
typedef struct _IDINFO {
    USHORT GenConfig;
    USHORT NumCyls;
    USHORT Reserved2;
    USHORT NumHeads;
    USHORT Reserved4;
    USHORT Reserved5;
    USHORT NumSectorsPerTrack;
    USHORT VendorUnique[3];
    CHAR SerialNumber[20];
    USHORT BufferType;
    USHORT BufferSize;
    USHORT ECCSize;
    CHAR FirmwareRev[8];
    CHAR ModelNumber[40];
    USHORT MoreVendorUnique;
    USHORT Reserved48;
    struct {
        USHORT reserved1: 8;
        USHORT DMA: 1;
        USHORT LBA: 1;
        USHORT DisIORDY: 1;
        USHORT IORDY: 1;
        USHORT SoftReset: 1;
        USHORT Overlap: 1;
        USHORT Queue: 1;
        USHORT InlDMA: 1;
    } Capabilities;
    USHORT Reserved1;
    USHORT PIOTiming;
    USHORT DMATiming;
    struct {
        USHORT CHSNumber: 1;
        USHORT CycleNumber: 1;
        USHORT UnltraDMA: 1;
        USHORT reserved: 13;
    } FieldValidity;
    USHORT NumCurCyls;
    USHORT NumCurHeads;
    USHORT NumCurSectorsPerTrack;
    USHORT CurSectorsLow;
    USHORT CurSectorsHigh;
    struct {
        USHORT CurNumber: 8;
        USHORT Multi: 1;
        USHORT reserved1: 7;
    } MultSectorStuff;
    ULONG dwTotalSectors;
    USHORT SingleWordDMA;
    struct {
        USHORT Mode0: 1;
        USHORT Mode1: 1;
        USHORT Mode2: 1;
        USHORT Reserved1: 5;
        USHORT Mode0Sel: 1;
        USHORT Mode1Sel: 1;
        USHORT Mode2Sel: 1;
        USHORT Reserved2: 5;
    } MultiWordDMA;
    struct {
        USHORT AdvPOIModes: 8;
        USHORT reserved: 8;
    } PIOCapacity;
    USHORT MinMultiWordDMACycle;
    USHORT RecMultiWordDMACycle;
    USHORT MinPIONoFlowCycle;
    USHORT MinPOIFlowCycle;
    USHORT Reserved69[11];
    struct {
        USHORT Reserved1: 1;
        USHORT ATA1: 1;
        USHORT ATA2: 1;
        USHORT ATA3: 1;
        USHORT ATA4: 1;
        USHORT ATA5: 1;
        USHORT ATA6: 1;
        USHORT ATA7: 1;
        USHORT ATA8: 1;
        USHORT ATA9: 1;
        USHORT ATA10: 1;
        USHORT ATA11: 1;
        USHORT ATA12: 1;
        USHORT ATA13: 1;
        USHORT ATA14: 1;
        USHORT Reserved2: 1;
    } MajorVersion;
    USHORT MinorVersion;
    USHORT Reserved82[6];
    struct {
        USHORT Mode0: 1;
        USHORT Mode1: 1;
        USHORT Mode2: 1;
        USHORT Mode3: 1;
        USHORT Mode4: 1;
        USHORT Mode5: 1;
        USHORT Mode6: 1;
        USHORT Mode7: 1;
        USHORT Mode0Sel: 1;
        USHORT Mode1Sel: 1;
        USHORT Mode2Sel: 1;
        USHORT Mode3Sel: 1;
        USHORT Mode4Sel: 1;
        USHORT Mode5Sel: 1;
        USHORT Mode6Sel: 1;
        USHORT Mode7Sel: 1;
    } UltraDMA;
    USHORT Reserved89[167];
} IDINFO, *PIDINFO;
typedef struct _DRIVERSTATUS {
    BYTE DriverError;
    BYTE IDEError;
    BYTE Reserved[2];
    DWORD dwReserved[2];
} DRIVERSTATUS, *PDRIVERSTATUS, *LPDRIVERSTATUS;
typedef struct _SENDCMDOUTPARAMS {
    DWORD BufferSize;
    DRIVERSTATUS DriverStatus;
    BYTE Buffer[1];
} SENDCMDOUTPARAMS, *PSENDCMDOUTPARAMS, *LPSENDCMDOUTPARAMS;
typedef struct _GETVERSIONINPARAMS {
    BYTE Version;
    BYTE Revision;
    BYTE Reserved1;
    BYTE IDEDeviceMap;
    DWORD fCapabilities;
    DWORD Reserved2[4];
} GETVERSIONINPARAMS, *PGETVERSIONINPARAMS, *LPGETVERSIONINPARAMS;
typedef struct _IDEREGS {
    BYTE FeaturesReg;
    BYTE SectorCountReg;
    BYTE SectorNumberReg;
    BYTE CylLowReg;
    BYTE CylHighReg;
    BYTE DriveHeadReg;
    BYTE CommandReg;
    BYTE Reserved;
} IDEREGS, *PIDEREGS, *LPIDEREGS;
typedef struct _SENDCMDINPARAMS {
    DWORD BufferSize;
    IDEREGS irDriveRegs;
    BYTE DriveNumber;
    BYTE Reserved1[3];
    DWORD Reserved2[4];
    BYTE Buffer[1];
} SENDCMDINPARAMS, *PSENDCMDINPARAMS, *LPSENDCMDINPARAMS;
`
