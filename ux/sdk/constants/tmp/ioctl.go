package constants

type IoctlKind int

const (
	MONITOR_IOCTL_ENABLE_MONITOR                     IoctlKind = 0x00120004
	MONITOR_IOCTL_DISABLE_MONITOR                              = 0x00120008
	IOCTL_INTERNAL_KEYBOARD_CONNECT                            = 0x000B0203
	IOCTL_INTERNAL_KEYBOARD_DISCONNECT                         = 0x000B0403
	IOCTL_INTERNAL_KEYBOARD_ENABLE                             = 0x000B0803
	IOCTL_INTERNAL_KEYBOARD_DISABLE                            = 0x000B1003
	IOCTL_INTERNAL_MOUSE_CONNECT                               = 0x000F0203
	IOCTL_INTERNAL_MOUSE_DISCONNECT                            = 0x000F0403
	IOCTL_INTERNAL_MOUSE_ENABLE                                = 0x000F0803
	IOCTL_INTERNAL_MOUSE_DISABLE                               = 0x000F1003
	IOCTL_DO_KERNELMODE_SAMPLES                                = 0x00222000
	IOCTL_REGISTER_CALLBACK                                    = 0x00222004
	IOCTL_UNREGISTER_CALLBACK                                  = 0x00222008
	IOCTL_GET_CALLBACK_VERSION                                 = 0x0022200C
	IOCTL_GET_VERSION                                          = 0x00222000
	IOCTL_RESET                                                = 0x00222007
	IOCTL_READWRITE_MAILBOX                                    = 0x00222008
	IOCTL_MAILBOX_WAIT                                         = 0x0022200C
	IOCTL_READ_DMA                                             = 0x00226012
	IOCTL_WRITE_DMA                                            = 0x0022A015
	IOCTL_ACPI_ASYNC_EVAL_METHOD                               = 0x0032C000
	IOCTL_ACPI_EVAL_METHOD                                     = 0x0032C004
	IOCTL_ACPI_ACQUIRE_GLOBAL_LOCK                             = 0x0032C010
	IOCTL_ACPI_RELEASE_GLOBAL_LOCK                             = 0x0032C014
	IOCTL_ACPI_EVAL_METHOD_EX                                  = 0x0032C018
	IOCTL_ACPI_ASYNC_EVAL_METHOD_EX                            = 0x0032C01C
	IOCTL_ACPI_ENUM_CHILDREN                                   = 0x0032C020
	IOCTL_AVC_UPDATE_VIRTUAL_SUBUNIT_INFO                      = 0x002A0000
	IOCTL_AVC_REMOVE_VIRTUAL_SUBUNIT_INFO                      = 0x002A0004
	IOCTL_AVC_BUS_RESET                                        = 0x002A0008
	IOCTL_DOT4_CREATE_SOCKET                                   = 0x003A2022
	IOCTL_DOT4_DESTROY_SOCKET                                  = 0x003A202A
	IOCTL_DOT4_WAIT_FOR_CHANNEL                                = 0x003A2026
	IOCTL_DOT4_OPEN_CHANNEL                                    = 0x003A2006
	IOCTL_DOT4_CLOSE_CHANNEL                                   = 0x003A2008
	IOCTL_DOT4_READ                                            = 0x003A200E
	IOCTL_DOT4_WRITE                                           = 0x003A2011
	IOCTL_DOT4_ADD_ACTIVITY_BROADCAST                          = 0x003A2014
	IOCTL_DOT4_REMOVE_ACTIVITY_BROADCAST                       = 0x003A2018
	IOCTL_DOT4_WAIT_ACTIVITY_BROADCAST                         = 0x003A201E
	IOCTL_MPDSM_REGISTER                                       = 0x736DC004
	IOCTL_MPDSM_DEREGISTER                                     = 0x736DC008
	IOCTL_MOUNTDEV_QUERY_UNIQUE_ID                             = 0x004D0000
	IOCTL_MOUNTDEV_QUERY_SUGGESTED_LINK_NAME                   = 0x004D000C
	IOCTL_MOUNTDEV_LINK_CREATED                                = 0x004DC010
	IOCTL_MOUNTDEV_LINK_DELETED                                = 0x004DC014
	IOCTL_MOUNTDEV_QUERY_STABLE_GUID                           = 0x004D0018
	IOCTL_MOUNTMGR_CREATE_POINT                                = 0x006DC000
	IOCTL_MOUNTMGR_DELETE_POINTS                               = 0x006DC004
	IOCTL_MOUNTMGR_QUERY_POINTS                                = 0x006D0008
	IOCTL_MOUNTMGR_DELETE_POINTS_DBONLY                        = 0x006DC00C
	IOCTL_MOUNTMGR_NEXT_DRIVE_LETTER                           = 0x006DC010
	IOCTL_MOUNTMGR_AUTO_DL_ASSIGNMENTS                         = 0x006DC014
	IOCTL_MOUNTMGR_VOLUME_MOUNT_POINT_CREATED                  = 0x006DC018
	IOCTL_MOUNTMGR_VOLUME_MOUNT_POINT_DELETED                  = 0x006DC01C
	IOCTL_MOUNTMGR_CHANGE_NOTIFY                               = 0x006D4020
	IOCTL_MOUNTMGR_KEEP_LINKS_WHEN_OFFLINE                     = 0x006DC024
	IOCTL_MOUNTMGR_CHECK_UNPROCESSED_VOLUMES                   = 0x006D4028
	IOCTL_MOUNTMGR_VOLUME_ARRIVAL_NOTIFICATION                 = 0x006D402C
	IOCTL_MOUNTDEV_QUERY_DEVICE_NAME                           = 0x004D0008
	IOCTL_MOUNTMGR_QUERY_DOS_VOLUME_PATH                       = 0x006D0030
	IOCTL_MOUNTMGR_QUERY_DOS_VOLUME_PATHS                      = 0x006D0034
	IOCTL_MOUNTMGR_SCRUB_REGISTRY                              = 0x006DC038
	IOCTL_MOUNTMGR_QUERY_AUTO_MOUNT                            = 0x006D003C
	IOCTL_MOUNTMGR_SET_AUTO_MOUNT                              = 0x006DC040
	IOCTL_MOUNTMGR_BOOT_DL_ASSIGNMENT                          = 0x006DC044
	IOCTL_MOUNTMGR_TRACELOG_CACHE                              = 0x006D4048
	IOCTL_AVIO_ALLOCATE_STREAM                                 = 0x00000000
	IOCTL_AVIO_FREE_STREAM                                     = 0x00000000
	IOCTL_AVIO_MODIFY_STREAM                                   = 0x00000000
	NLB_IOCTL_REGISTER_HOOK                                    = 0xC0C08048
	NLB_PUBLIC_IOCTL_CLIENT_STICKINESS_NOTIFY                  = 0xC0C08054
	IOCTL_GET_TUPLE_DATA                                       = 0x00042EE0
	IOCTL_SOCKET_INFORMATION                                   = 0x00042EF0
	IOCTL_PCMCIA_HIDE_DEVICE                                   = 0x0004AF08
	IOCTL_PCMCIA_REVEAL_DEVICE                                 = 0x0004AF0C
	IOCTL_SD_SUBMIT_REQUEST                                    = 0x00043073
	FSCTL_REQUEST_OPLOCK_LEVEL_1                               = 0x00090000
	FSCTL_REQUEST_OPLOCK_LEVEL_2                               = 0x00090004
	FSCTL_REQUEST_BATCH_OPLOCK                                 = 0x00090008
	FSCTL_OPLOCK_BREAK_ACKNOWLEDGE                             = 0x0009000C
	FSCTL_OPBATCH_ACK_CLOSE_PENDING                            = 0x00090010
	FSCTL_OPLOCK_BREAK_NOTIFY                                  = 0x00090014
	FSCTL_LOCK_VOLUME                                          = 0x00090018
	FSCTL_UNLOCK_VOLUME                                        = 0x0009001C
	FSCTL_DISMOUNT_VOLUME                                      = 0x00090020
	FSCTL_IS_VOLUME_MOUNTED                                    = 0x00090028
	FSCTL_IS_PATHNAME_VALID                                    = 0x0009002C
	FSCTL_MARK_VOLUME_DIRTY                                    = 0x00090030
	FSCTL_QUERY_RETRIEVAL_POINTERS                             = 0x0009003B
	FSCTL_GET_COMPRESSION                                      = 0x0009003C
	FSCTL_SET_COMPRESSION                                      = 0x0009C040
	FSCTL_SET_BOOTLOADER_ACCESSED                              = 0x0009004F
	FSCTL_OPLOCK_BREAK_ACK_NO_2                                = 0x00090050
	FSCTL_INVALIDATE_VOLUMES                                   = 0x00090054
	FSCTL_QUERY_FAT_BPB                                        = 0x00090058
	FSCTL_REQUEST_FILTER_OPLOCK                                = 0x0009005C
	FSCTL_FILESYSTEM_GET_STATISTICS                            = 0x00090060
	FSCTL_GET_NTFS_VOLUME_DATA                                 = 0x00090064
	FSCTL_GET_NTFS_FILE_RECORD                                 = 0x00090068
	FSCTL_GET_VOLUME_BITMAP                                    = 0x0009006F
	FSCTL_GET_RETRIEVAL_POINTERS                               = 0x00090073
	FSCTL_MOVE_FILE                                            = 0x00090074
	FSCTL_IS_VOLUME_DIRTY                                      = 0x00090078
	FSCTL_ALLOW_EXTENDED_DASD_IO                               = 0x00090083
	FSCTL_FIND_FILES_BY_SID                                    = 0x0009008F
	FSCTL_SET_OBJECT_ID                                        = 0x00090098
	FSCTL_GET_OBJECT_ID                                        = 0x0009009C
	FSCTL_DELETE_OBJECT_ID                                     = 0x000900A0
	FSCTL_SET_REPARSE_POINT                                    = 0x000900A4
	FSCTL_GET_REPARSE_POINT                                    = 0x000900A8
	FSCTL_DELETE_REPARSE_POINT                                 = 0x000900AC
	FSCTL_ENUM_USN_DATA                                        = 0x000900B3
	FSCTL_SECURITY_ID_CHECK                                    = 0x000940B7
	FSCTL_READ_USN_JOURNAL                                     = 0x000900BB
	FSCTL_SET_OBJECT_ID_EXTENDED                               = 0x000900BC
	FSCTL_CREATE_OR_GET_OBJECT_ID                              = 0x000900C0
	FSCTL_SET_SPARSE                                           = 0x000900C4
	FSCTL_SET_ZERO_DATA                                        = 0x000980C8
	FSCTL_QUERY_ALLOCATED_RANGES                               = 0x000940CF
	FSCTL_ENABLE_UPGRADE                                       = 0x000980D0
	FSCTL_SET_ENCRYPTION                                       = 0x000900D7
	FSCTL_ENCRYPTION_FSCTL_IO                                  = 0x000900DB
	FSCTL_WRITE_RAW_ENCRYPTED                                  = 0x000900DF
	FSCTL_READ_RAW_ENCRYPTED                                   = 0x000900E3
	FSCTL_CREATE_USN_JOURNAL                                   = 0x000900E7
	FSCTL_READ_FILE_USN_DATA                                   = 0x000900EB
	FSCTL_WRITE_USN_CLOSE_RECORD                               = 0x000900EF
	FSCTL_EXTEND_VOLUME                                        = 0x000900F0
	FSCTL_QUERY_USN_JOURNAL                                    = 0x000900F4
	FSCTL_DELETE_USN_JOURNAL                                   = 0x000900F8
	FSCTL_MARK_HANDLE                                          = 0x000900FC
	FSCTL_SIS_COPYFILE                                         = 0x00090100
	FSCTL_SIS_LINK_FILES                                       = 0x0009C104
	FSCTL_RECALL_FILE                                          = 0x00090117
	FSCTL_READ_FROM_PLEX                                       = 0x0009411E
	FSCTL_FILE_PREFETCH                                        = 0x00090120
	FSCTL_MAKE_MEDIA_COMPATIBLE                                = 0x00098130
	FSCTL_SET_DEFECT_MANAGEMENT                                = 0x00098134
	FSCTL_QUERY_SPARING_INFO                                   = 0x00090138
	FSCTL_QUERY_ON_DISK_VOLUME_INFO                            = 0x0009013C
	FSCTL_SET_VOLUME_COMPRESSION_STATE                         = 0x00090140
	FSCTL_TXFS_MODIFY_RM                                       = 0x00098144
	FSCTL_TXFS_QUERY_RM_INFORMATION                            = 0x00094148
	FSCTL_TXFS_ROLLFORWARD_REDO                                = 0x00098150
	FSCTL_TXFS_ROLLFORWARD_UNDO                                = 0x00098154
	FSCTL_TXFS_START_RM                                        = 0x00098158
	FSCTL_TXFS_SHUTDOWN_RM                                     = 0x0009815C
	FSCTL_TXFS_READ_BACKUP_INFORMATION                         = 0x00094160
	FSCTL_TXFS_WRITE_BACKUP_INFORMATION                        = 0x00098164
	FSCTL_TXFS_CREATE_SECONDARY_RM                             = 0x00098168
	FSCTL_TXFS_GET_METADATA_INFO                               = 0x0009416C
	FSCTL_TXFS_GET_TRANSACTED_VERSION                          = 0x00094170
	FSCTL_TXFS_SAVEPOINT_INFORMATION                           = 0x00098178
	FSCTL_TXFS_CREATE_MINIVERSION                              = 0x0009817C
	FSCTL_TXFS_TRANSACTION_ACTIVE                              = 0x0009418C
	FSCTL_SET_ZERO_ON_DEALLOCATION                             = 0x00090194
	FSCTL_SET_REPAIR                                           = 0x00090198
	FSCTL_GET_REPAIR                                           = 0x0009019C
	FSCTL_WAIT_FOR_REPAIR                                      = 0x000901A0
	FSCTL_INITIATE_REPAIR                                      = 0x000901A8
	FSCTL_CSC_INTERNAL                                         = 0x000901AF
	FSCTL_SHRINK_VOLUME                                        = 0x000901B0
	FSCTL_SET_SHORT_NAME_BEHAVIOR                              = 0x000901B4
	FSCTL_DFSR_SET_GHOST_HANDLE_STATE                          = 0x000901B8
	FSCTL_TXFS_LIST_TRANSACTIONS                               = 0x000941E4
	FSCTL_QUERY_PAGEFILE_ENCRYPTION                            = 0x000901E8
	FSCTL_RESET_VOLUME_ALLOCATION_HINTS                        = 0x000901EC
	FSCTL_QUERY_DEPENDENT_VOLUME                               = 0x000901F0
	FSCTL_SD_GLOBAL_CHANGE                                     = 0x000901F4
	FSCTL_TXFS_READ_BACKUP_INFORMATION2                        = 0x000901F8
	FSCTL_LOOKUP_STREAM_FROM_CLUSTER                           = 0x000901FC
	FSCTL_TXFS_WRITE_BACKUP_INFORMATION2                       = 0x00090200
	FSCTL_FILE_TYPE_NOTIFICATION                               = 0x00090204
	FSCTL_GET_BOOT_AREA_INFO                                   = 0x00090230
	FSCTL_GET_RETRIEVAL_POINTER_BASE                           = 0x00090234
	FSCTL_SET_PERSISTENT_VOLUME_STATE                          = 0x00090238
	FSCTL_QUERY_PERSISTENT_VOLUME_STATE                        = 0x0009023C
	FSCTL_REQUEST_OPLOCK                                       = 0x00090240
	FSCTL_CSV_TUNNEL_REQUEST                                   = 0x00090244
	FSCTL_IS_CSV_FILE                                          = 0x00090248
	FSCTL_QUERY_FILE_SYSTEM_RECOGNITION                        = 0x0009024C
	FSCTL_CSV_GET_VOLUME_PATH_NAME                             = 0x00090250
	FSCTL_CSV_GET_VOLUME_NAME_FOR_VOLUME_MOUNT_POINT           = 0x00090254
	FSCTL_CSV_GET_VOLUME_PATH_NAMES_FOR_VOLUME_NAME            = 0x00090258
	FSCTL_IS_FILE_ON_CSV_VOLUME                                = 0x0009025C
	FSCTL_LMR_GET_LINK_TRACKING_INFORMATION                    = 0x001400E8
	FSCTL_LMR_SET_LINK_TRACKING_INFORMATION                    = 0x001400EC
	IOCTL_LMR_ARE_FILE_OBJECTS_ON_SAME_SERVER                  = 0x001400F0
	FSCTL_PIPE_ASSIGN_EVENT                                    = 0x00110000
	FSCTL_PIPE_DISCONNECT                                      = 0x00110004
	FSCTL_PIPE_LISTEN                                          = 0x00110008
	FSCTL_PIPE_PEEK                                            = 0x0011400C
	FSCTL_PIPE_QUERY_EVENT                                     = 0x00110010
	FSCTL_PIPE_TRANSCEIVE                                      = 0x0011C017
	FSCTL_PIPE_WAIT                                            = 0x00110018
	FSCTL_PIPE_IMPERSONATE                                     = 0x0011001C
	FSCTL_PIPE_SET_CLIENT_PROCESS                              = 0x00110020
	FSCTL_PIPE_QUERY_CLIENT_PROCESS                            = 0x00110024
	FSCTL_PIPE_GET_PIPE_ATTRIBUTE                              = 0x00110028
	FSCTL_PIPE_SET_PIPE_ATTRIBUTE                              = 0x0011002C
	FSCTL_PIPE_GET_CONNECTION_ATTRIBUTE                        = 0x00110030
	FSCTL_PIPE_SET_CONNECTION_ATTRIBUTE                        = 0x00110034
	FSCTL_PIPE_GET_HANDLE_ATTRIBUTE                            = 0x00110038
	FSCTL_PIPE_SET_HANDLE_ATTRIBUTE                            = 0x0011003C
	FSCTL_PIPE_FLUSH                                           = 0x00118040
	FSCTL_PIPE_INTERNAL_READ                                   = 0x00115FF4
	FSCTL_PIPE_INTERNAL_WRITE                                  = 0x00119FF8
	FSCTL_PIPE_INTERNAL_TRANSCEIVE                             = 0x0011DFFF
	FSCTL_PIPE_INTERNAL_READ_OVFLOW                            = 0x00116000
	FSCTL_MAILSLOT_PEEK                                        = 0x000C4003
	IOCTL_REDIR_QUERY_PATH                                     = 0x0014018F
	IOCTL_REDIR_QUERY_PATH_EX                                  = 0x00140193
	IOCTL_VOLSNAP_FLUSH_AND_HOLD_WRITES                        = 0x0053C000
	IOCTL_INTERNAL_PARALLEL_PORT_ALLOCATE                      = 0x0016002C
	IOCTL_INTERNAL_GET_PARALLEL_PORT_INFO                      = 0x00160030
	IOCTL_INTERNAL_PARALLEL_CONNECT_INTERRUPT                  = 0x00160034
	IOCTL_INTERNAL_PARALLEL_DISCONNECT_INTERRUPT               = 0x00160038
	IOCTL_INTERNAL_RELEASE_PARALLEL_PORT_INFO                  = 0x0016003C
	IOCTL_INTERNAL_GET_MORE_PARALLEL_PORT_INFO                 = 0x00160044
	IOCTL_INTERNAL_PARCHIP_CONNECT                             = 0x00160048
	IOCTL_INTERNAL_PARALLEL_SET_CHIP_MODE                      = 0x0016004C
	IOCTL_INTERNAL_PARALLEL_CLEAR_CHIP_MODE                    = 0x00160050
	IOCTL_INTERNAL_GET_PARALLEL_PNP_INFO                       = 0x00160054
	IOCTL_INTERNAL_INIT_1284_3_BUS                             = 0x00160058
	IOCTL_INTERNAL_SELECT_DEVICE                               = 0x0016005C
	IOCTL_INTERNAL_DESELECT_DEVICE                             = 0x00160060
	IOCTL_INTERNAL_GET_PARPORT_FDO                             = 0x00160074
	IOCTL_INTERNAL_PARCLASS_CONNECT                            = 0x00160078
	IOCTL_INTERNAL_PARCLASS_DISCONNECT                         = 0x0016007C
	IOCTL_INTERNAL_DISCONNECT_IDLE                             = 0x00160080
	IOCTL_INTERNAL_LOCK_PORT                                   = 0x00160094
	IOCTL_INTERNAL_UNLOCK_PORT                                 = 0x00160098
	IOCTL_INTERNAL_PARALLEL_PORT_FREE                          = 0x001600A0
	IOCTL_INTERNAL_PARDOT3_CONNECT                             = 0x001600A4
	IOCTL_INTERNAL_PARDOT3_DISCONNECT                          = 0x001600A8
	IOCTL_INTERNAL_PARDOT3_RESET                               = 0x001600AC
	IOCTL_INTERNAL_PARDOT3_SIGNAL                              = 0x001600B0
	IOCTL_INTERNAL_REGISTER_FOR_REMOVAL_RELATIONS              = 0x001600C8
	IOCTL_INTERNAL_UNREGISTER_FOR_REMOVAL_RELATIONS            = 0x001600CC
	IOCTL_INTERNAL_LOCK_PORT_NO_SELECT                         = 0x001600D0
	IOCTL_INTERNAL_UNLOCK_PORT_NO_DESELECT                     = 0x001600D4
	IOCTL_INTERNAL_DISABLE_END_OF_CHAIN_BUS_RESCAN             = 0x001600D8
	IOCTL_INTERNAL_ENABLE_END_OF_CHAIN_BUS_RESCAN              = 0x001600DC
	IOCTL_WPD_MESSAGE_READWRITE_ACCESS                         = 0x0040C108
	IOCTL_WPD_MESSAGE_READ_ACCESS                              = 0x00404108
	IOCTL_SCSISCAN_CMD                                         = 0x00190012
	IOCTL_SCSISCAN_LOCKDEVICE                                  = 0x00190016
	IOCTL_SCSISCAN_UNLOCKDEVICE                                = 0x0019001A
	IOCTL_SCSISCAN_SET_TIMEOUT                                 = 0x0019001C
	IOCTL_SCSISCAN_GET_INFO                                    = 0x00190022
	IOCTL_SWENUM_INSTALL_INTERFACE                             = 0x002A0000
	IOCTL_SWENUM_REMOVE_INTERFACE                              = 0x002A0004
	IOCTL_SWENUM_GET_BUS_ID                                    = 0x002A400B
	IOCTL_CANCEL_IO                                            = 0x80002004
	IOCTL_WAIT_ON_DEVICE_EVENT                                 = 0x80002008
	IOCTL_READ_REGISTERS                                       = 0x8000200C
	IOCTL_WRITE_REGISTERS                                      = 0x80002010
	IOCTL_GET_CHANNEL_ALIGN_RQST                               = 0x80002014
	IOCTL_GET_DEVICE_DESCRIPTOR                                = 0x80002018
	IOCTL_RESET_PIPE                                           = 0x8000201C
	IOCTL_GET_USB_DESCRIPTOR                                   = 0x80002020
	IOCTL_SEND_USB_REQUEST                                     = 0x80002024
	IOCTL_GET_PIPE_CONFIGURATION                               = 0x80002028
	IOCTL_SET_TIMEOUT                                          = 0x8000202C
	IOCTL_EHSTOR_DEVICE_SET_AUTHZ_STATE                        = 0x002D1404
	IOCTL_EHSTOR_DEVICE_GET_AUTHZ_STATE                        = 0x002D1408
	IOCTL_EHSTOR_DEVICE_SILO_COMMAND                           = 0x002D140C
	IOCTL_EHSTOR_DEVICE_ENUMERATE_PDOS                         = 0x002D1410
	IOCTL_KS_PROPERTY                                          = 0x002F0003
	IOCTL_KS_ENABLE_EVENT                                      = 0x002F0007
	IOCTL_KS_DISABLE_EVENT                                     = 0x002F000B
	IOCTL_KS_METHOD                                            = 0x002F000F
	IOCTL_KS_WRITE_STREAM                                      = 0x002F8013
	IOCTL_KS_READ_STREAM                                       = 0x002F4017
	IOCTL_KS_RESET_STATE                                       = 0x002F001B
	IOCTL_KS_HANDSHAKE                                         = 0x002F001F
	IOCTL_INTERNAL_I8042_HOOK_KEYBOARD                         = 0x000B3FC3
	IOCTL_INTERNAL_I8042_HOOK_MOUSE                            = 0x000F3FC3
	IOCTL_INTERNAL_I8042_KEYBOARD_WRITE_BUFFER                 = 0x000B3FC7
	IOCTL_INTERNAL_I8042_MOUSE_WRITE_BUFFER                    = 0x000F3FC7
	IOCTL_INTERNAL_I8042_CONTROLLER_WRITE_BUFFER               = 0x000B3FCB
	IOCTL_INTERNAL_I8042_KEYBOARD_START_INFORMATION            = 0x000B3FCF
	IOCTL_INTERNAL_I8042_MOUSE_START_INFORMATION               = 0x000F3FCF
	IOCTL_BEEP_SET                                             = 0x00010000
	IOCTL_CDROM_UNLOAD_DRIVER                                  = 0x00025008
	IOCTL_CDROM_READ_TOC                                       = 0x00024000
	IOCTL_CDROM_SEEK_AUDIO_MSF                                 = 0x00024004
	IOCTL_CDROM_STOP_AUDIO                                     = 0x00024008
	IOCTL_CDROM_PAUSE_AUDIO                                    = 0x0002400C
	IOCTL_CDROM_RESUME_AUDIO                                   = 0x00024010
	IOCTL_CDROM_GET_VOLUME                                     = 0x00024014
	IOCTL_CDROM_PLAY_AUDIO_MSF                                 = 0x00024018
	IOCTL_CDROM_SET_VOLUME                                     = 0x00024028
	IOCTL_CDROM_READ_Q_CHANNEL                                 = 0x0002402C
	IOCTL_CDROM_GET_CONTROL                                    = 0x00024034
	OBSOLETE_IOCTL_CDROM_GET_CONTROL                           = 0x00024034
	IOCTL_CDROM_GET_LAST_SESSION                               = 0x00024038
	IOCTL_CDROM_RAW_READ                                       = 0x0002403E
	IOCTL_CDROM_DISK_TYPE                                      = 0x00020040
	IOCTL_CDROM_GET_DRIVE_GEOMETRY                             = 0x0002404C
	IOCTL_CDROM_GET_DRIVE_GEOMETRY_EX                          = 0x00024050
	IOCTL_CDROM_READ_TOC_EX                                    = 0x00024054
	IOCTL_CDROM_GET_CONFIGURATION                              = 0x00024058
	IOCTL_CDROM_EXCLUSIVE_ACCESS                               = 0x0002C05C
	IOCTL_CDROM_SET_SPEED                                      = 0x00024060
	IOCTL_CDROM_GET_INQUIRY_DATA                               = 0x00024064
	IOCTL_CDROM_ENABLE_STREAMING                               = 0x00024068
	IOCTL_CDROM_SEND_OPC_INFORMATION                           = 0x0002C06C
	IOCTL_CDROM_GET_PERFORMANCE                                = 0x00024070
	IOCTL_CDROM_CHECK_VERIFY                                   = 0x00024800
	IOCTL_CDROM_MEDIA_REMOVAL                                  = 0x00024804
	IOCTL_CDROM_EJECT_MEDIA                                    = 0x00024808
	IOCTL_CDROM_LOAD_MEDIA                                     = 0x0002480C
	IOCTL_CDROM_RESERVE                                        = 0x00024810
	IOCTL_CDROM_RELEASE                                        = 0x00024814
	IOCTL_CDROM_FIND_NEW_DEVICES                               = 0x00024818
	IOCTL_CDROM_SIMBAD                                         = 0x0002400C
	IOCTL_DVD_START_SESSION                                    = 0x00335000
	IOCTL_DVD_READ_KEY                                         = 0x00335004
	IOCTL_DVD_SEND_KEY                                         = 0x00335008
	IOCTL_DVD_END_SESSION                                      = 0x0033500C
	IOCTL_DVD_SET_READ_AHEAD                                   = 0x00335010
	IOCTL_DVD_GET_REGION                                       = 0x00335014
	IOCTL_DVD_SEND_KEY2                                        = 0x0033D018
	IOCTL_AACS_READ_MEDIA_KEY_BLOCK_SIZE                       = 0x003350C0
	IOCTL_AACS_READ_MEDIA_KEY_BLOCK                            = 0x003350C4
	IOCTL_AACS_START_SESSION                                   = 0x003350C8
	IOCTL_AACS_END_SESSION                                     = 0x003350CC
	IOCTL_AACS_SEND_CERTIFICATE                                = 0x003350D0
	IOCTL_AACS_GET_CERTIFICATE                                 = 0x003350D4
	IOCTL_AACS_GET_CHALLENGE_KEY                               = 0x003350D8
	IOCTL_AACS_SEND_CHALLENGE_KEY                              = 0x003350DC
	IOCTL_AACS_READ_VOLUME_ID                                  = 0x003350E0
	IOCTL_AACS_READ_SERIAL_NUMBER                              = 0x003350E4
	IOCTL_AACS_READ_MEDIA_ID                                   = 0x003350E8
	IOCTL_AACS_READ_BINDING_NONCE                              = 0x003350EC
	IOCTL_AACS_GENERATE_BINDING_NONCE                          = 0x0033D0F0
	IOCTL_DVD_READ_STRUCTURE                                   = 0x00335140
	IOCTL_STORAGE_SET_READ_AHEAD                               = 0x002D4400
	IOCTL_CHANGER_GET_PARAMETERS                               = 0x00304000
	IOCTL_CHANGER_GET_STATUS                                   = 0x00304004
	IOCTL_CHANGER_GET_PRODUCT_DATA                             = 0x00304008
	IOCTL_CHANGER_SET_ACCESS                                   = 0x0030C010
	IOCTL_CHANGER_GET_ELEMENT_STATUS                           = 0x0030C014
	IOCTL_CHANGER_INITIALIZE_ELEMENT_STATUS                    = 0x00304018
	IOCTL_CHANGER_SET_POSITION                                 = 0x0030401C
	IOCTL_CHANGER_EXCHANGE_MEDIUM                              = 0x00304020
	IOCTL_CHANGER_MOVE_MEDIUM                                  = 0x00304024
	IOCTL_CHANGER_REINITIALIZE_TRANSPORT                       = 0x00304028
	IOCTL_CHANGER_QUERY_VOLUME_TAGS                            = 0x0030C02C
	IOCTL_DISK_GET_DRIVE_GEOMETRY                              = 0x00070000
	IOCTL_DISK_GET_PARTITION_INFO                              = 0x00074004
	IOCTL_DISK_SET_PARTITION_INFO                              = 0x0007C008
	IOCTL_DISK_GET_DRIVE_LAYOUT                                = 0x0007400C
	IOCTL_DISK_SET_DRIVE_LAYOUT                                = 0x0007C010
	IOCTL_DISK_VERIFY                                          = 0x00070014
	IOCTL_DISK_FORMAT_TRACKS                                   = 0x0007C018
	IOCTL_DISK_REASSIGN_BLOCKS                                 = 0x0007C01C
	IOCTL_DISK_PERFORMANCE                                     = 0x00070020
	IOCTL_DISK_IS_WRITABLE                                     = 0x00070024
	IOCTL_DISK_LOGGING                                         = 0x00070028
	IOCTL_DISK_FORMAT_TRACKS_EX                                = 0x0007C02C
	IOCTL_DISK_HISTOGRAM_STRUCTURE                             = 0x00070030
	IOCTL_DISK_HISTOGRAM_DATA                                  = 0x00070034
	IOCTL_DISK_HISTOGRAM_RESET                                 = 0x00070038
	IOCTL_DISK_REQUEST_STRUCTURE                               = 0x0007003C
	IOCTL_DISK_REQUEST_DATA                                    = 0x00070040
	IOCTL_DISK_PERFORMANCE_OFF                                 = 0x00070060
	IOCTL_DISK_CONTROLLER_NUMBER                               = 0x00070044
	SMART_GET_VERSION                                          = 0x00074080
	SMART_SEND_DRIVE_COMMAND                                   = 0x0007C084
	SMART_RCV_DRIVE_DATA                                       = 0x0007C088
	IOCTL_DISK_GET_PARTITION_INFO_EX                           = 0x00070048
	IOCTL_DISK_SET_PARTITION_INFO_EX                           = 0x0007C04C
	IOCTL_DISK_GET_DRIVE_LAYOUT_EX                             = 0x00070050
	IOCTL_DISK_SET_DRIVE_LAYOUT_EX                             = 0x0007C054
	IOCTL_DISK_CREATE_DISK                                     = 0x0007C058
	IOCTL_DISK_GET_LENGTH_INFO                                 = 0x0007405C
	IOCTL_DISK_GET_DRIVE_GEOMETRY_EX                           = 0x000700A0
	IOCTL_DISK_REASSIGN_BLOCKS_EX                              = 0x0007C0A4
	IOCTL_DISK_UPDATE_DRIVE_SIZE                               = 0x0007C0C8
	IOCTL_DISK_GROW_PARTITION                                  = 0x0007C0D0
	IOCTL_DISK_GET_CACHE_INFORMATION                           = 0x000740D4
	IOCTL_DISK_SET_CACHE_INFORMATION                           = 0x0007C0D8
	IOCTL_DISK_GET_WRITE_CACHE_STATE                           = 0x000740DC
	OBSOLETE_DISK_GET_WRITE_CACHE_STATE                        = 0x000740DC
	IOCTL_DISK_DELETE_DRIVE_LAYOUT                             = 0x0007C100
	IOCTL_DISK_UPDATE_PROPERTIES                               = 0x00070140
	IOCTL_DISK_FORMAT_DRIVE                                    = 0x0007C3CC
	IOCTL_DISK_SENSE_DEVICE                                    = 0x000703E0
	IOCTL_DISK_GET_CACHE_SETTING                               = 0x000740E0
	IOCTL_DISK_SET_CACHE_SETTING                               = 0x0007C0E4
	IOCTL_DISK_COPY_DATA                                       = 0x0007C064
	IOCTL_DISK_INTERNAL_SET_VERIFY                             = 0x00070403
	IOCTL_DISK_INTERNAL_CLEAR_VERIFY                           = 0x00070407
	IOCTL_DISK_INTERNAL_SET_NOTIFY                             = 0x00070408
	IOCTL_DISK_CHECK_VERIFY                                    = 0x00074800
	IOCTL_DISK_MEDIA_REMOVAL                                   = 0x00074804
	IOCTL_DISK_EJECT_MEDIA                                     = 0x00074808
	IOCTL_DISK_LOAD_MEDIA                                      = 0x0007480C
	IOCTL_DISK_RESERVE                                         = 0x00074810
	IOCTL_DISK_RELEASE                                         = 0x00074814
	IOCTL_DISK_FIND_NEW_DEVICES                                = 0x00074818
	IOCTL_DISK_GET_MEDIA_TYPES                                 = 0x00070C00
	IOCTL_DISK_GET_PARTITION_ATTRIBUTES                        = 0x000700E8
	IOCTL_DISK_SET_PARTITION_ATTRIBUTES                        = 0x0007C0EC
	IOCTL_DISK_GET_DISK_ATTRIBUTES                             = 0x000700F0
	IOCTL_DISK_SET_DISK_ATTRIBUTES                             = 0x0007C0F4
	IOCTL_DISK_IS_CLUSTERED                                    = 0x000700F8
	IOCTL_DISK_GET_SAN_SETTINGS                                = 0x00074200
	IOCTL_DISK_SET_SAN_SETTINGS                                = 0x0007C204
	IOCTL_DISK_GET_SNAPSHOT_INFO                               = 0x00074208
	IOCTL_DISK_SET_SNAPSHOT_INFO                               = 0x0007C20C
	IOCTL_DISK_RESET_SNAPSHOT_INFO                             = 0x0007C210
	IOCTL_DISK_SIMBAD                                          = 0x0007D000
	FT_SECONDARY_READ                                          = 0x00664012
	FT_PRIMARY_READ                                            = 0x00664016
	FT_BALANCED_READ_MODE                                      = 0x0066001B
	FT_SYNC_REDUNDANT_COPY                                     = 0x0066001C
	FT_INITIALIZE_SET                                          = 0x00660000
	FT_REGENERATE                                              = 0x00660004
	FT_CONFIGURE                                               = 0x0066000B
	FT_VERIFY                                                  = 0x0066000C
	FT_SEQUENTIAL_WRITE_MODE                                   = 0x00660023
	FT_PARALLEL_WRITE_MODE                                     = 0x00660027
	FT_QUERY_SET_STATE                                         = 0x00660028
	FT_CLUSTER_SET_MEMBER_STATE                                = 0x0066002C
	FT_CLUSTER_GET_MEMBER_STATE                                = 0x00660030
	FT_ENUMERATE_LOGICAL_DISKS                                 = 0x00674008
	FT_CREATE_LOGICAL_DISK                                     = 0x0067C000
	FT_BREAK_LOGICAL_DISK                                      = 0x0067C004
	FT_QUERY_LOGICAL_DISK_INFORMATION                          = 0x0067400C
	FT_ORPHAN_LOGICAL_DISK_MEMBER                              = 0x0067C010
	FT_REPLACE_LOGICAL_DISK_MEMBER                             = 0x0067C014
	FT_QUERY_NT_DEVICE_NAME_FOR_LOGICAL_DISK                   = 0x00674018
	FT_INITIALIZE_LOGICAL_DISK                                 = 0x0067C01C
	FT_QUERY_DRIVE_LETTER_FOR_LOGICAL_DISK                     = 0x00674020
	FT_CHECK_IO                                                = 0x00674024
	FT_SET_DRIVE_LETTER_FOR_LOGICAL_DISK                       = 0x0067C028
	FT_QUERY_NT_DEVICE_NAME_FOR_PARTITION                      = 0x00674030
	FT_CHANGE_NOTIFY                                           = 0x00674034
	FT_STOP_SYNC_OPERATIONS                                    = 0x0067C038
	FT_QUERY_LOGICAL_DISK_ID                                   = 0x00674190
	FT_CREATE_PARTITION_LOGICAL_DISK                           = 0x0067C194
	IOCTL_KEYBOARD_QUERY_ATTRIBUTES                            = 0x000B0000
	IOCTL_KEYBOARD_SET_TYPEMATIC                               = 0x000B0004
	IOCTL_KEYBOARD_SET_INDICATORS                              = 0x000B0008
	IOCTL_KEYBOARD_QUERY_TYPEMATIC                             = 0x000B0020
	IOCTL_KEYBOARD_QUERY_INDICATORS                            = 0x000B0040
	IOCTL_KEYBOARD_QUERY_INDICATOR_TRANSLATION                 = 0x000B0080
	IOCTL_KEYBOARD_INSERT_DATA                                 = 0x000B0100
	IOCTL_KEYBOARD_QUERY_IME_STATUS                            = 0x000B1000
	IOCTL_KEYBOARD_SET_IME_STATUS                              = 0x000B1004
	IOCTL_MODEM_GET_PASSTHROUGH                                = 0x002B0004
	IOCTL_MODEM_SET_PASSTHROUGH                                = 0x002B0008
	IOCTL_MODEM_SET_DLE_MONITORING                             = 0x002B000C
	IOCTL_MODEM_GET_DLE                                        = 0x002B0010
	IOCTL_MODEM_SET_DLE_SHIELDING                              = 0x002B0014
	IOCTL_MODEM_STOP_WAVE_RECEIVE                              = 0x002B0018
	IOCTL_MODEM_SEND_MESSAGE                                   = 0x002B001C
	IOCTL_MODEM_GET_MESSAGE                                    = 0x002B0020
	IOCTL_MODEM_SEND_GET_MESSAGE                               = 0x002B0024
	IOCTL_MODEM_SEND_LOOPBACK_MESSAGE                          = 0x002B0028
	IOCTL_MODEM_CHECK_FOR_MODEM                                = 0x002B002C
	IOCTL_MODEM_SET_MIN_POWER                                  = 0x002B0030
	IOCTL_MODEM_WATCH_FOR_RESUME                               = 0x002B0034
	IOCTL_CANCEL_GET_SEND_MESSAGE                              = 0x002B0038
	IOCTL_SET_SERVER_STATE                                     = 0x002B003C
	IOCTL_MOUSE_QUERY_ATTRIBUTES                               = 0x000F0000
	IOCTL_MOUSE_INSERT_DATA                                    = 0x000F0004
	IOCTL_NDIS_RESERVED5                                       = 0x00170034
	IOCTL_NDIS_RESERVED6                                       = 0x00178038
	IOCTL_PAR_QUERY_INFORMATION                                = 0x00160004
	IOCTL_PAR_SET_INFORMATION                                  = 0x00160008
	IOCTL_PAR_QUERY_DEVICE_ID                                  = 0x0016000C
	IOCTL_PAR_QUERY_DEVICE_ID_SIZE                             = 0x00160010
	IOCTL_IEEE1284_GET_MODE                                    = 0x00160014
	IOCTL_IEEE1284_NEGOTIATE                                   = 0x00160018
	IOCTL_PAR_SET_WRITE_ADDRESS                                = 0x0016001C
	IOCTL_PAR_SET_READ_ADDRESS                                 = 0x00160020
	IOCTL_PAR_GET_DEVICE_CAPS                                  = 0x00160024
	IOCTL_PAR_GET_DEFAULT_MODES                                = 0x00160028
	IOCTL_PAR_PING                                             = 0x0016002C
	IOCTL_PAR_QUERY_RAW_DEVICE_ID                              = 0x00160030
	IOCTL_PAR_ECP_HOST_RECOVERY                                = 0x00160034
	IOCTL_PAR_GET_READ_ADDRESS                                 = 0x00160038
	IOCTL_PAR_GET_WRITE_ADDRESS                                = 0x0016003C
	IOCTL_PAR_TEST                                             = 0x00160050
	IOCTL_PAR_IS_PORT_FREE                                     = 0x00160054
	IOCTL_PAR_QUERY_LOCATION                                   = 0x00160058
	IOCTL_SCSI_PASS_THROUGH                                    = 0x0004D004
	IOCTL_SCSI_MINIPORT                                        = 0x0004D008
	IOCTL_SCSI_GET_INQUIRY_DATA                                = 0x0004100C
	IOCTL_SCSI_GET_CAPABILITIES                                = 0x00041010
	IOCTL_SCSI_PASS_THROUGH_DIRECT                             = 0x0004D014
	IOCTL_SCSI_GET_ADDRESS                                     = 0x00041018
	IOCTL_SCSI_RESCAN_BUS                                      = 0x0004101C
	IOCTL_SCSI_GET_DUMP_POINTERS                               = 0x00041020
	IOCTL_SCSI_FREE_DUMP_POINTERS                              = 0x00041024
	IOCTL_IDE_PASS_THROUGH                                     = 0x0004D028
	IOCTL_ATA_PASS_THROUGH                                     = 0x0004D02C
	IOCTL_ATA_PASS_THROUGH_DIRECT                              = 0x0004D030
	IOCTL_ATA_MINIPORT                                         = 0x0004D034
	IOCTL_MINIPORT_PROCESS_SERVICE_IRP                         = 0x0004D038
	IOCTL_MPIO_PASS_THROUGH_PATH                               = 0x0004D03C
	IOCTL_MPIO_PASS_THROUGH_PATH_DIRECT                        = 0x0004D040
	IOCTL_SERIAL_SET_BAUD_RATE                                 = 0x001B0004
	IOCTL_SERIAL_SET_QUEUE_SIZE                                = 0x001B0008
	IOCTL_SERIAL_SET_LINE_CONTROL                              = 0x001B000C
	IOCTL_SERIAL_SET_BREAK_ON                                  = 0x001B0010
	IOCTL_SERIAL_SET_BREAK_OFF                                 = 0x001B0014
	IOCTL_SERIAL_IMMEDIATE_CHAR                                = 0x001B0018
	IOCTL_SERIAL_SET_TIMEOUTS                                  = 0x001B001C
	IOCTL_SERIAL_GET_TIMEOUTS                                  = 0x001B0020
	IOCTL_SERIAL_SET_DTR                                       = 0x001B0024
	IOCTL_SERIAL_CLR_DTR                                       = 0x001B0028
	IOCTL_SERIAL_RESET_DEVICE                                  = 0x001B002C
	IOCTL_SERIAL_SET_RTS                                       = 0x001B0030
	IOCTL_SERIAL_CLR_RTS                                       = 0x001B0034
	IOCTL_SERIAL_SET_XOFF                                      = 0x001B0038
	IOCTL_SERIAL_SET_XON                                       = 0x001B003C
	IOCTL_SERIAL_GET_WAIT_MASK                                 = 0x001B0040
	IOCTL_SERIAL_SET_WAIT_MASK                                 = 0x001B0044
	IOCTL_SERIAL_WAIT_ON_MASK                                  = 0x001B0048
	IOCTL_SERIAL_PURGE                                         = 0x001B004C
	IOCTL_SERIAL_GET_BAUD_RATE                                 = 0x001B0050
	IOCTL_SERIAL_GET_LINE_CONTROL                              = 0x001B0054
	IOCTL_SERIAL_GET_CHARS                                     = 0x001B0058
	IOCTL_SERIAL_SET_CHARS                                     = 0x001B005C
	IOCTL_SERIAL_GET_HANDFLOW                                  = 0x001B0060
	IOCTL_SERIAL_SET_HANDFLOW                                  = 0x001B0064
	IOCTL_SERIAL_GET_MODEMSTATUS                               = 0x001B0068
	IOCTL_SERIAL_GET_COMMSTATUS                                = 0x001B006C
	IOCTL_SERIAL_XOFF_COUNTER                                  = 0x001B0070
	IOCTL_SERIAL_GET_PROPERTIES                                = 0x001B0074
	IOCTL_SERIAL_GET_DTRRTS                                    = 0x001B0078
	IOCTL_SERIAL_LSRMST_INSERT                                 = 0x001B007C
	IOCTL_SERENUM_EXPOSE_HARDWARE                              = 0x00370200
	IOCTL_SERENUM_REMOVE_HARDWARE                              = 0x00370204
	IOCTL_SERENUM_PORT_DESC                                    = 0x00370208
	IOCTL_SERENUM_GET_PORT_NAME                                = 0x0037020C
	IOCTL_SERIAL_CONFIG_SIZE                                   = 0x001B0080
	IOCTL_SERIAL_GET_COMMCONFIG                                = 0x001B0084
	IOCTL_SERIAL_SET_COMMCONFIG                                = 0x001B0088
	IOCTL_SERIAL_GET_STATS                                     = 0x001B008C
	IOCTL_SERIAL_CLEAR_STATS                                   = 0x001B0090
	IOCTL_SERIAL_GET_MODEM_CONTROL                             = 0x001B0094
	IOCTL_SERIAL_SET_MODEM_CONTROL                             = 0x001B0098
	IOCTL_SERIAL_SET_FIFO_CONTROL                              = 0x001B009C
	IOCTL_SERIAL_INTERNAL_DO_WAIT_WAKE                         = 0x001B0004
	IOCTL_SERIAL_INTERNAL_CANCEL_WAIT_WAKE                     = 0x001B0008
	IOCTL_SERIAL_INTERNAL_BASIC_SETTINGS                       = 0x001B000C
	IOCTL_SERIAL_INTERNAL_RESTORE_SETTINGS                     = 0x001B0010
	IOCTL_STORAGE_CHECK_VERIFY                                 = 0x002D4800
	IOCTL_STORAGE_CHECK_VERIFY2                                = 0x002D0800
	IOCTL_STORAGE_MEDIA_REMOVAL                                = 0x002D4804
	IOCTL_STORAGE_EJECT_MEDIA                                  = 0x002D4808
	IOCTL_STORAGE_LOAD_MEDIA                                   = 0x002D480C
	IOCTL_STORAGE_LOAD_MEDIA2                                  = 0x002D080C
	IOCTL_STORAGE_RESERVE                                      = 0x002D4810
	IOCTL_STORAGE_RELEASE                                      = 0x002D4814
	IOCTL_STORAGE_FIND_NEW_DEVICES                             = 0x002D4818
	IOCTL_STORAGE_EJECTION_CONTROL                             = 0x002D0940
	IOCTL_STORAGE_MCN_CONTROL                                  = 0x002D0944
	IOCTL_STORAGE_GET_MEDIA_TYPES                              = 0x002D0C00
	IOCTL_STORAGE_GET_MEDIA_TYPES_EX                           = 0x002D0C04
	IOCTL_STORAGE_GET_MEDIA_SERIAL_NUMBER                      = 0x002D0C10
	IOCTL_STORAGE_GET_HOTPLUG_INFO                             = 0x002D0C14
	IOCTL_STORAGE_SET_HOTPLUG_INFO                             = 0x002DCC18
	IOCTL_STORAGE_RESET_BUS                                    = 0x002D5000
	IOCTL_STORAGE_RESET_DEVICE                                 = 0x002D5004
	IOCTL_STORAGE_BREAK_RESERVATION                            = 0x002D5014
	IOCTL_STORAGE_PERSISTENT_RESERVE_IN                        = 0x002D5018
	IOCTL_STORAGE_PERSISTENT_RESERVE_OUT                       = 0x002DD01C
	IOCTL_STORAGE_GET_DEVICE_NUMBER                            = 0x002D1080
	IOCTL_STORAGE_PREDICT_FAILURE                              = 0x002D1100
	IOCTL_STORAGE_READ_CAPACITY                                = 0x002D5140
	IOCTL_STORAGE_QUERY_PROPERTY                               = 0x002D1400
	IOCTL_STORAGE_MANAGE_DATA_SET_ATTRIBUTES                   = 0x002D9404
	IOCTL_STORAGE_GET_BC_PROPERTIES                            = 0x002D5800
	IOCTL_STORAGE_ALLOCATE_BC_STREAM                           = 0x002DD804
	IOCTL_STORAGE_FREE_BC_STREAM                               = 0x002DD808
	IOCTL_STORAGE_CHECK_PRIORITY_HINT_SUPPORT                  = 0x002D1880
	OBSOLETE_IOCTL_STORAGE_RESET_BUS                           = 0x002DD000
	OBSOLETE_IOCTL_STORAGE_RESET_DEVICE                        = 0x002DD004
	IOCTL_TAPE_ERASE                                           = 0x001FC000
	IOCTL_TAPE_PREPARE                                         = 0x001F4004
	IOCTL_TAPE_WRITE_MARKS                                     = 0x001FC008
	IOCTL_TAPE_GET_POSITION                                    = 0x001F400C
	IOCTL_TAPE_SET_POSITION                                    = 0x001F4010
	IOCTL_TAPE_GET_DRIVE_PARAMS                                = 0x001F4014
	IOCTL_TAPE_SET_DRIVE_PARAMS                                = 0x001FC018
	IOCTL_TAPE_GET_MEDIA_PARAMS                                = 0x001F401C
	IOCTL_TAPE_SET_MEDIA_PARAMS                                = 0x001F4020
	IOCTL_TAPE_GET_STATUS                                      = 0x001F4024
	IOCTL_TAPE_CREATE_PARTITION                                = 0x001FC028
	IOCTL_TAPE_MEDIA_REMOVAL                                   = 0x001F4804
	IOCTL_TAPE_EJECT_MEDIA                                     = 0x001F4808
	IOCTL_TAPE_LOAD_MEDIA                                      = 0x001F480C
	IOCTL_TAPE_RESERVE                                         = 0x001F4810
	IOCTL_TAPE_RELEASE                                         = 0x001F4814
	IOCTL_TAPE_CHECK_VERIFY                                    = 0x001F4800
	IOCTL_TAPE_FIND_NEW_DEVICES                                = 0x00074818
	IOCTL_VOLUME_GET_VOLUME_DISK_EXTENTS                       = 0x00560000
	IOCTL_VOLUME_ONLINE                                        = 0x0056C008
	IOCTL_VOLUME_OFFLINE                                       = 0x0056C00C
	IOCTL_VOLUME_IS_CLUSTERED                                  = 0x00560030
	IOCTL_VOLUME_GET_GPT_ATTRIBUTES                            = 0x00560038
	IOCTL_VOLUME_SUPPORTS_ONLINE_OFFLINE                       = 0x00560004
	IOCTL_VOLUME_IS_OFFLINE                                    = 0x00560010
	IOCTL_VOLUME_IS_IO_CAPABLE                                 = 0x00560014
	IOCTL_VOLUME_QUERY_FAILOVER_SET                            = 0x00560018
	IOCTL_VOLUME_QUERY_VOLUME_NUMBER                           = 0x0056001C
	IOCTL_VOLUME_LOGICAL_TO_PHYSICAL                           = 0x00560020
	IOCTL_VOLUME_PHYSICAL_TO_LOGICAL                           = 0x00560024
	IOCTL_VOLUME_IS_PARTITION                                  = 0x00560028
	IOCTL_VOLUME_READ_PLEX                                     = 0x0056402E
	IOCTL_VOLUME_SET_GPT_ATTRIBUTES                            = 0x00560034
	IOCTL_VOLUME_GET_BC_PROPERTIES                             = 0x0056403C
	IOCTL_VOLUME_ALLOCATE_BC_STREAM                            = 0x0056C040
	IOCTL_VOLUME_FREE_BC_STREAM                                = 0x0056C044
	IOCTL_VOLUME_IS_DYNAMIC                                    = 0x00560048
	IOCTL_VOLUME_PREPARE_FOR_CRITICAL_IO                       = 0x0056C04C
	IOCTL_VOLUME_QUERY_ALLOCATION_HINT                         = 0x00564052
	IOCTL_VOLUME_UPDATE_PROPERTIES                             = 0x00560054
	IOCTL_VOLUME_QUERY_MINIMUM_SHRINK_SIZE                     = 0x00564058
	IOCTL_VOLUME_PREPARE_FOR_SHRINK                            = 0x0056C05C
	IOCTL_PMI_GET_CAPABILITIES                                 = 0x00454000
	IOCTL_PMI_GET_CONFIGURATION                                = 0x00454004
	IOCTL_PMI_SET_CONFIGURATION                                = 0x00458008
	IOCTL_PMI_GET_MEASUREMENT                                  = 0x0045400C
	IOCTL_PMI_REGISTER_EVENT_NOTIFY                            = 0x0045C010
	IOCTL_BIOMETRIC_VENDOR                                     = 0x00442000
)

func (k IoctlKind) String() string {
	switch k {
	case MONITOR_IOCTL_ENABLE_MONITOR:
		return "0x00120004"
	case MONITOR_IOCTL_DISABLE_MONITOR:
		return "0x00120008"
	case IOCTL_INTERNAL_KEYBOARD_CONNECT:
		return "0x000B0203"
	case IOCTL_INTERNAL_KEYBOARD_DISCONNECT:
		return "0x000B0403"
	case IOCTL_INTERNAL_KEYBOARD_ENABLE:
		return "0x000B0803"
	case IOCTL_INTERNAL_KEYBOARD_DISABLE:
		return "0x000B1003"
	case IOCTL_INTERNAL_MOUSE_CONNECT:
		return "0x000F0203"
	case IOCTL_INTERNAL_MOUSE_DISCONNECT:
		return "0x000F0403"
	case IOCTL_INTERNAL_MOUSE_ENABLE:
		return "0x000F0803"
	case IOCTL_INTERNAL_MOUSE_DISABLE:
		return "0x000F1003"
	case IOCTL_DO_KERNELMODE_SAMPLES:
		return "0x00222000"
	case IOCTL_REGISTER_CALLBACK:
		return "0x00222004"
	case IOCTL_UNREGISTER_CALLBACK:
		return "0x00222008"
	case IOCTL_GET_CALLBACK_VERSION:
		return "0x0022200C"
	case IOCTL_GET_VERSION:
		return "0x00222000"
	case IOCTL_RESET:
		return "0x00222007"
	case IOCTL_READWRITE_MAILBOX:
		return "0x00222008"
	case IOCTL_MAILBOX_WAIT:
		return "0x0022200C"
	case IOCTL_READ_DMA:
		return "0x00226012"
	case IOCTL_WRITE_DMA:
		return "0x0022A015"
	case IOCTL_ACPI_ASYNC_EVAL_METHOD:
		return "0x0032C000"
	case IOCTL_ACPI_EVAL_METHOD:
		return "0x0032C004"
	case IOCTL_ACPI_ACQUIRE_GLOBAL_LOCK:
		return "0x0032C010"
	case IOCTL_ACPI_RELEASE_GLOBAL_LOCK:
		return "0x0032C014"
	case IOCTL_ACPI_EVAL_METHOD_EX:
		return "0x0032C018"
	case IOCTL_ACPI_ASYNC_EVAL_METHOD_EX:
		return "0x0032C01C"
	case IOCTL_ACPI_ENUM_CHILDREN:
		return "0x0032C020"
	case IOCTL_AVC_UPDATE_VIRTUAL_SUBUNIT_INFO:
		return "0x002A0000"
	case IOCTL_AVC_REMOVE_VIRTUAL_SUBUNIT_INFO:
		return "0x002A0004"
	case IOCTL_AVC_BUS_RESET:
		return "0x002A0008"
	case IOCTL_DOT4_CREATE_SOCKET:
		return "0x003A2022"
	case IOCTL_DOT4_DESTROY_SOCKET:
		return "0x003A202A"
	case IOCTL_DOT4_WAIT_FOR_CHANNEL:
		return "0x003A2026"
	case IOCTL_DOT4_OPEN_CHANNEL:
		return "0x003A2006"
	case IOCTL_DOT4_CLOSE_CHANNEL:
		return "0x003A2008"
	case IOCTL_DOT4_READ:
		return "0x003A200E"
	case IOCTL_DOT4_WRITE:
		return "0x003A2011"
	case IOCTL_DOT4_ADD_ACTIVITY_BROADCAST:
		return "0x003A2014"
	case IOCTL_DOT4_REMOVE_ACTIVITY_BROADCAST:
		return "0x003A2018"
	case IOCTL_DOT4_WAIT_ACTIVITY_BROADCAST:
		return "0x003A201E"
	case IOCTL_MPDSM_REGISTER:
		return "0x736DC004"
	case IOCTL_MPDSM_DEREGISTER:
		return "0x736DC008"
	case IOCTL_MOUNTDEV_QUERY_UNIQUE_ID:
		return "0x004D0000"
	case IOCTL_MOUNTDEV_QUERY_SUGGESTED_LINK_NAME:
		return "0x004D000C"
	case IOCTL_MOUNTDEV_LINK_CREATED:
		return "0x004DC010"
	case IOCTL_MOUNTDEV_LINK_DELETED:
		return "0x004DC014"
	case IOCTL_MOUNTDEV_QUERY_STABLE_GUID:
		return "0x004D0018"
	case IOCTL_MOUNTMGR_CREATE_POINT:
		return "0x006DC000"
	case IOCTL_MOUNTMGR_DELETE_POINTS:
		return "0x006DC004"
	case IOCTL_MOUNTMGR_QUERY_POINTS:
		return "0x006D0008"
	case IOCTL_MOUNTMGR_DELETE_POINTS_DBONLY:
		return "0x006DC00C"
	case IOCTL_MOUNTMGR_NEXT_DRIVE_LETTER:
		return "0x006DC010"
	case IOCTL_MOUNTMGR_AUTO_DL_ASSIGNMENTS:
		return "0x006DC014"
	case IOCTL_MOUNTMGR_VOLUME_MOUNT_POINT_CREATED:
		return "0x006DC018"
	case IOCTL_MOUNTMGR_VOLUME_MOUNT_POINT_DELETED:
		return "0x006DC01C"
	case IOCTL_MOUNTMGR_CHANGE_NOTIFY:
		return "0x006D4020"
	case IOCTL_MOUNTMGR_KEEP_LINKS_WHEN_OFFLINE:
		return "0x006DC024"
	case IOCTL_MOUNTMGR_CHECK_UNPROCESSED_VOLUMES:
		return "0x006D4028"
	case IOCTL_MOUNTMGR_VOLUME_ARRIVAL_NOTIFICATION:
		return "0x006D402C"
	case IOCTL_MOUNTDEV_QUERY_DEVICE_NAME:
		return "0x004D0008"
	case IOCTL_MOUNTMGR_QUERY_DOS_VOLUME_PATH:
		return "0x006D0030"
	case IOCTL_MOUNTMGR_QUERY_DOS_VOLUME_PATHS:
		return "0x006D0034"
	case IOCTL_MOUNTMGR_SCRUB_REGISTRY:
		return "0x006DC038"
	case IOCTL_MOUNTMGR_QUERY_AUTO_MOUNT:
		return "0x006D003C"
	case IOCTL_MOUNTMGR_SET_AUTO_MOUNT:
		return "0x006DC040"
	case IOCTL_MOUNTMGR_BOOT_DL_ASSIGNMENT:
		return "0x006DC044"
	case IOCTL_MOUNTMGR_TRACELOG_CACHE:
		return "0x006D4048"
	case IOCTL_AVIO_ALLOCATE_STREAM:
		return "0x00000000"
	case IOCTL_AVIO_FREE_STREAM:
		return "0x00000000"
	case IOCTL_AVIO_MODIFY_STREAM:
		return "0x00000000"
	case NLB_IOCTL_REGISTER_HOOK:
		return "0xC0C08048"
	case NLB_PUBLIC_IOCTL_CLIENT_STICKINESS_NOTIFY:
		return "0xC0C08054"
	case IOCTL_GET_TUPLE_DATA:
		return "0x00042EE0"
	case IOCTL_SOCKET_INFORMATION:
		return "0x00042EF0"
	case IOCTL_PCMCIA_HIDE_DEVICE:
		return "0x0004AF08"
	case IOCTL_PCMCIA_REVEAL_DEVICE:
		return "0x0004AF0C"
	case IOCTL_SD_SUBMIT_REQUEST:
		return "0x00043073"
	case FSCTL_REQUEST_OPLOCK_LEVEL_1:
		return "0x00090000"
	case FSCTL_REQUEST_OPLOCK_LEVEL_2:
		return "0x00090004"
	case FSCTL_REQUEST_BATCH_OPLOCK:
		return "0x00090008"
	case FSCTL_OPLOCK_BREAK_ACKNOWLEDGE:
		return "0x0009000C"
	case FSCTL_OPBATCH_ACK_CLOSE_PENDING:
		return "0x00090010"
	case FSCTL_OPLOCK_BREAK_NOTIFY:
		return "0x00090014"
	case FSCTL_LOCK_VOLUME:
		return "0x00090018"
	case FSCTL_UNLOCK_VOLUME:
		return "0x0009001C"
	case FSCTL_DISMOUNT_VOLUME:
		return "0x00090020"
	case FSCTL_IS_VOLUME_MOUNTED:
		return "0x00090028"
	case FSCTL_IS_PATHNAME_VALID:
		return "0x0009002C"
	case FSCTL_MARK_VOLUME_DIRTY:
		return "0x00090030"
	case FSCTL_QUERY_RETRIEVAL_POINTERS:
		return "0x0009003B"
	case FSCTL_GET_COMPRESSION:
		return "0x0009003C"
	case FSCTL_SET_COMPRESSION:
		return "0x0009C040"
	case FSCTL_SET_BOOTLOADER_ACCESSED:
		return "0x0009004F"
	case FSCTL_OPLOCK_BREAK_ACK_NO_2:
		return "0x00090050"
	case FSCTL_INVALIDATE_VOLUMES:
		return "0x00090054"
	case FSCTL_QUERY_FAT_BPB:
		return "0x00090058"
	case FSCTL_REQUEST_FILTER_OPLOCK:
		return "0x0009005C"
	case FSCTL_FILESYSTEM_GET_STATISTICS:
		return "0x00090060"
	case FSCTL_GET_NTFS_VOLUME_DATA:
		return "0x00090064"
	case FSCTL_GET_NTFS_FILE_RECORD:
		return "0x00090068"
	case FSCTL_GET_VOLUME_BITMAP:
		return "0x0009006F"
	case FSCTL_GET_RETRIEVAL_POINTERS:
		return "0x00090073"
	case FSCTL_MOVE_FILE:
		return "0x00090074"
	case FSCTL_IS_VOLUME_DIRTY:
		return "0x00090078"
	case FSCTL_ALLOW_EXTENDED_DASD_IO:
		return "0x00090083"
	case FSCTL_FIND_FILES_BY_SID:
		return "0x0009008F"
	case FSCTL_SET_OBJECT_ID:
		return "0x00090098"
	case FSCTL_GET_OBJECT_ID:
		return "0x0009009C"
	case FSCTL_DELETE_OBJECT_ID:
		return "0x000900A0"
	case FSCTL_SET_REPARSE_POINT:
		return "0x000900A4"
	case FSCTL_GET_REPARSE_POINT:
		return "0x000900A8"
	case FSCTL_DELETE_REPARSE_POINT:
		return "0x000900AC"
	case FSCTL_ENUM_USN_DATA:
		return "0x000900B3"
	case FSCTL_SECURITY_ID_CHECK:
		return "0x000940B7"
	case FSCTL_READ_USN_JOURNAL:
		return "0x000900BB"
	case FSCTL_SET_OBJECT_ID_EXTENDED:
		return "0x000900BC"
	case FSCTL_CREATE_OR_GET_OBJECT_ID:
		return "0x000900C0"
	case FSCTL_SET_SPARSE:
		return "0x000900C4"
	case FSCTL_SET_ZERO_DATA:
		return "0x000980C8"
	case FSCTL_QUERY_ALLOCATED_RANGES:
		return "0x000940CF"
	case FSCTL_ENABLE_UPGRADE:
		return "0x000980D0"
	case FSCTL_SET_ENCRYPTION:
		return "0x000900D7"
	case FSCTL_ENCRYPTION_FSCTL_IO:
		return "0x000900DB"
	case FSCTL_WRITE_RAW_ENCRYPTED:
		return "0x000900DF"
	case FSCTL_READ_RAW_ENCRYPTED:
		return "0x000900E3"
	case FSCTL_CREATE_USN_JOURNAL:
		return "0x000900E7"
	case FSCTL_READ_FILE_USN_DATA:
		return "0x000900EB"
	case FSCTL_WRITE_USN_CLOSE_RECORD:
		return "0x000900EF"
	case FSCTL_EXTEND_VOLUME:
		return "0x000900F0"
	case FSCTL_QUERY_USN_JOURNAL:
		return "0x000900F4"
	case FSCTL_DELETE_USN_JOURNAL:
		return "0x000900F8"
	case FSCTL_MARK_HANDLE:
		return "0x000900FC"
	case FSCTL_SIS_COPYFILE:
		return "0x00090100"
	case FSCTL_SIS_LINK_FILES:
		return "0x0009C104"
	case FSCTL_RECALL_FILE:
		return "0x00090117"
	case FSCTL_READ_FROM_PLEX:
		return "0x0009411E"
	case FSCTL_FILE_PREFETCH:
		return "0x00090120"
	case FSCTL_MAKE_MEDIA_COMPATIBLE:
		return "0x00098130"
	case FSCTL_SET_DEFECT_MANAGEMENT:
		return "0x00098134"
	case FSCTL_QUERY_SPARING_INFO:
		return "0x00090138"
	case FSCTL_QUERY_ON_DISK_VOLUME_INFO:
		return "0x0009013C"
	case FSCTL_SET_VOLUME_COMPRESSION_STATE:
		return "0x00090140"
	case FSCTL_TXFS_MODIFY_RM:
		return "0x00098144"
	case FSCTL_TXFS_QUERY_RM_INFORMATION:
		return "0x00094148"
	case FSCTL_TXFS_ROLLFORWARD_REDO:
		return "0x00098150"
	case FSCTL_TXFS_ROLLFORWARD_UNDO:
		return "0x00098154"
	case FSCTL_TXFS_START_RM:
		return "0x00098158"
	case FSCTL_TXFS_SHUTDOWN_RM:
		return "0x0009815C"
	case FSCTL_TXFS_READ_BACKUP_INFORMATION:
		return "0x00094160"
	case FSCTL_TXFS_WRITE_BACKUP_INFORMATION:
		return "0x00098164"
	case FSCTL_TXFS_CREATE_SECONDARY_RM:
		return "0x00098168"
	case FSCTL_TXFS_GET_METADATA_INFO:
		return "0x0009416C"
	case FSCTL_TXFS_GET_TRANSACTED_VERSION:
		return "0x00094170"
	case FSCTL_TXFS_SAVEPOINT_INFORMATION:
		return "0x00098178"
	case FSCTL_TXFS_CREATE_MINIVERSION:
		return "0x0009817C"
	case FSCTL_TXFS_TRANSACTION_ACTIVE:
		return "0x0009418C"
	case FSCTL_SET_ZERO_ON_DEALLOCATION:
		return "0x00090194"
	case FSCTL_SET_REPAIR:
		return "0x00090198"
	case FSCTL_GET_REPAIR:
		return "0x0009019C"
	case FSCTL_WAIT_FOR_REPAIR:
		return "0x000901A0"
	case FSCTL_INITIATE_REPAIR:
		return "0x000901A8"
	case FSCTL_CSC_INTERNAL:
		return "0x000901AF"
	case FSCTL_SHRINK_VOLUME:
		return "0x000901B0"
	case FSCTL_SET_SHORT_NAME_BEHAVIOR:
		return "0x000901B4"
	case FSCTL_DFSR_SET_GHOST_HANDLE_STATE:
		return "0x000901B8"
	case FSCTL_TXFS_LIST_TRANSACTIONS:
		return "0x000941E4"
	case FSCTL_QUERY_PAGEFILE_ENCRYPTION:
		return "0x000901E8"
	case FSCTL_RESET_VOLUME_ALLOCATION_HINTS:
		return "0x000901EC"
	case FSCTL_QUERY_DEPENDENT_VOLUME:
		return "0x000901F0"
	case FSCTL_SD_GLOBAL_CHANGE:
		return "0x000901F4"
	case FSCTL_TXFS_READ_BACKUP_INFORMATION2:
		return "0x000901F8"
	case FSCTL_LOOKUP_STREAM_FROM_CLUSTER:
		return "0x000901FC"
	case FSCTL_TXFS_WRITE_BACKUP_INFORMATION2:
		return "0x00090200"
	case FSCTL_FILE_TYPE_NOTIFICATION:
		return "0x00090204"
	case FSCTL_GET_BOOT_AREA_INFO:
		return "0x00090230"
	case FSCTL_GET_RETRIEVAL_POINTER_BASE:
		return "0x00090234"
	case FSCTL_SET_PERSISTENT_VOLUME_STATE:
		return "0x00090238"
	case FSCTL_QUERY_PERSISTENT_VOLUME_STATE:
		return "0x0009023C"
	case FSCTL_REQUEST_OPLOCK:
		return "0x00090240"
	case FSCTL_CSV_TUNNEL_REQUEST:
		return "0x00090244"
	case FSCTL_IS_CSV_FILE:
		return "0x00090248"
	case FSCTL_QUERY_FILE_SYSTEM_RECOGNITION:
		return "0x0009024C"
	case FSCTL_CSV_GET_VOLUME_PATH_NAME:
		return "0x00090250"
	case FSCTL_CSV_GET_VOLUME_NAME_FOR_VOLUME_MOUNT_POINT:
		return "0x00090254"
	case FSCTL_CSV_GET_VOLUME_PATH_NAMES_FOR_VOLUME_NAME:
		return "0x00090258"
	case FSCTL_IS_FILE_ON_CSV_VOLUME:
		return "0x0009025C"
	case FSCTL_LMR_GET_LINK_TRACKING_INFORMATION:
		return "0x001400E8"
	case FSCTL_LMR_SET_LINK_TRACKING_INFORMATION:
		return "0x001400EC"
	case IOCTL_LMR_ARE_FILE_OBJECTS_ON_SAME_SERVER:
		return "0x001400F0"
	case FSCTL_PIPE_ASSIGN_EVENT:
		return "0x00110000"
	case FSCTL_PIPE_DISCONNECT:
		return "0x00110004"
	case FSCTL_PIPE_LISTEN:
		return "0x00110008"
	case FSCTL_PIPE_PEEK:
		return "0x0011400C"
	case FSCTL_PIPE_QUERY_EVENT:
		return "0x00110010"
	case FSCTL_PIPE_TRANSCEIVE:
		return "0x0011C017"
	case FSCTL_PIPE_WAIT:
		return "0x00110018"
	case FSCTL_PIPE_IMPERSONATE:
		return "0x0011001C"
	case FSCTL_PIPE_SET_CLIENT_PROCESS:
		return "0x00110020"
	case FSCTL_PIPE_QUERY_CLIENT_PROCESS:
		return "0x00110024"
	case FSCTL_PIPE_GET_PIPE_ATTRIBUTE:
		return "0x00110028"
	case FSCTL_PIPE_SET_PIPE_ATTRIBUTE:
		return "0x0011002C"
	case FSCTL_PIPE_GET_CONNECTION_ATTRIBUTE:
		return "0x00110030"
	case FSCTL_PIPE_SET_CONNECTION_ATTRIBUTE:
		return "0x00110034"
	case FSCTL_PIPE_GET_HANDLE_ATTRIBUTE:
		return "0x00110038"
	case FSCTL_PIPE_SET_HANDLE_ATTRIBUTE:
		return "0x0011003C"
	case FSCTL_PIPE_FLUSH:
		return "0x00118040"
	case FSCTL_PIPE_INTERNAL_READ:
		return "0x00115FF4"
	case FSCTL_PIPE_INTERNAL_WRITE:
		return "0x00119FF8"
	case FSCTL_PIPE_INTERNAL_TRANSCEIVE:
		return "0x0011DFFF"
	case FSCTL_PIPE_INTERNAL_READ_OVFLOW:
		return "0x00116000"
	case FSCTL_MAILSLOT_PEEK:
		return "0x000C4003"
	case IOCTL_REDIR_QUERY_PATH:
		return "0x0014018F"
	case IOCTL_REDIR_QUERY_PATH_EX:
		return "0x00140193"
	case IOCTL_VOLSNAP_FLUSH_AND_HOLD_WRITES:
		return "0x0053C000"
	case IOCTL_INTERNAL_PARALLEL_PORT_ALLOCATE:
		return "0x0016002C"
	case IOCTL_INTERNAL_GET_PARALLEL_PORT_INFO:
		return "0x00160030"
	case IOCTL_INTERNAL_PARALLEL_CONNECT_INTERRUPT:
		return "0x00160034"
	case IOCTL_INTERNAL_PARALLEL_DISCONNECT_INTERRUPT:
		return "0x00160038"
	case IOCTL_INTERNAL_RELEASE_PARALLEL_PORT_INFO:
		return "0x0016003C"
	case IOCTL_INTERNAL_GET_MORE_PARALLEL_PORT_INFO:
		return "0x00160044"
	case IOCTL_INTERNAL_PARCHIP_CONNECT:
		return "0x00160048"
	case IOCTL_INTERNAL_PARALLEL_SET_CHIP_MODE:
		return "0x0016004C"
	case IOCTL_INTERNAL_PARALLEL_CLEAR_CHIP_MODE:
		return "0x00160050"
	case IOCTL_INTERNAL_GET_PARALLEL_PNP_INFO:
		return "0x00160054"
	case IOCTL_INTERNAL_INIT_1284_3_BUS:
		return "0x00160058"
	case IOCTL_INTERNAL_SELECT_DEVICE:
		return "0x0016005C"
	case IOCTL_INTERNAL_DESELECT_DEVICE:
		return "0x00160060"
	case IOCTL_INTERNAL_GET_PARPORT_FDO:
		return "0x00160074"
	case IOCTL_INTERNAL_PARCLASS_CONNECT:
		return "0x00160078"
	case IOCTL_INTERNAL_PARCLASS_DISCONNECT:
		return "0x0016007C"
	case IOCTL_INTERNAL_DISCONNECT_IDLE:
		return "0x00160080"
	case IOCTL_INTERNAL_LOCK_PORT:
		return "0x00160094"
	case IOCTL_INTERNAL_UNLOCK_PORT:
		return "0x00160098"
	case IOCTL_INTERNAL_PARALLEL_PORT_FREE:
		return "0x001600A0"
	case IOCTL_INTERNAL_PARDOT3_CONNECT:
		return "0x001600A4"
	case IOCTL_INTERNAL_PARDOT3_DISCONNECT:
		return "0x001600A8"
	case IOCTL_INTERNAL_PARDOT3_RESET:
		return "0x001600AC"
	case IOCTL_INTERNAL_PARDOT3_SIGNAL:
		return "0x001600B0"
	case IOCTL_INTERNAL_REGISTER_FOR_REMOVAL_RELATIONS:
		return "0x001600C8"
	case IOCTL_INTERNAL_UNREGISTER_FOR_REMOVAL_RELATIONS:
		return "0x001600CC"
	case IOCTL_INTERNAL_LOCK_PORT_NO_SELECT:
		return "0x001600D0"
	case IOCTL_INTERNAL_UNLOCK_PORT_NO_DESELECT:
		return "0x001600D4"
	case IOCTL_INTERNAL_DISABLE_END_OF_CHAIN_BUS_RESCAN:
		return "0x001600D8"
	case IOCTL_INTERNAL_ENABLE_END_OF_CHAIN_BUS_RESCAN:
		return "0x001600DC"
	case IOCTL_WPD_MESSAGE_READWRITE_ACCESS:
		return "0x0040C108"
	case IOCTL_WPD_MESSAGE_READ_ACCESS:
		return "0x00404108"
	case IOCTL_SCSISCAN_CMD:
		return "0x00190012"
	case IOCTL_SCSISCAN_LOCKDEVICE:
		return "0x00190016"
	case IOCTL_SCSISCAN_UNLOCKDEVICE:
		return "0x0019001A"
	case IOCTL_SCSISCAN_SET_TIMEOUT:
		return "0x0019001C"
	case IOCTL_SCSISCAN_GET_INFO:
		return "0x00190022"
	case IOCTL_SWENUM_INSTALL_INTERFACE:
		return "0x002A0000"
	case IOCTL_SWENUM_REMOVE_INTERFACE:
		return "0x002A0004"
	case IOCTL_SWENUM_GET_BUS_ID:
		return "0x002A400B"
	case IOCTL_CANCEL_IO:
		return "0x80002004"
	case IOCTL_WAIT_ON_DEVICE_EVENT:
		return "0x80002008"
	case IOCTL_READ_REGISTERS:
		return "0x8000200C"
	case IOCTL_WRITE_REGISTERS:
		return "0x80002010"
	case IOCTL_GET_CHANNEL_ALIGN_RQST:
		return "0x80002014"
	case IOCTL_GET_DEVICE_DESCRIPTOR:
		return "0x80002018"
	case IOCTL_RESET_PIPE:
		return "0x8000201C"
	case IOCTL_GET_USB_DESCRIPTOR:
		return "0x80002020"
	case IOCTL_SEND_USB_REQUEST:
		return "0x80002024"
	case IOCTL_GET_PIPE_CONFIGURATION:
		return "0x80002028"
	case IOCTL_SET_TIMEOUT:
		return "0x8000202C"
	case IOCTL_EHSTOR_DEVICE_SET_AUTHZ_STATE:
		return "0x002D1404"
	case IOCTL_EHSTOR_DEVICE_GET_AUTHZ_STATE:
		return "0x002D1408"
	case IOCTL_EHSTOR_DEVICE_SILO_COMMAND:
		return "0x002D140C"
	case IOCTL_EHSTOR_DEVICE_ENUMERATE_PDOS:
		return "0x002D1410"
	case IOCTL_KS_PROPERTY:
		return "0x002F0003"
	case IOCTL_KS_ENABLE_EVENT:
		return "0x002F0007"
	case IOCTL_KS_DISABLE_EVENT:
		return "0x002F000B"
	case IOCTL_KS_METHOD:
		return "0x002F000F"
	case IOCTL_KS_WRITE_STREAM:
		return "0x002F8013"
	case IOCTL_KS_READ_STREAM:
		return "0x002F4017"
	case IOCTL_KS_RESET_STATE:
		return "0x002F001B"
	case IOCTL_KS_HANDSHAKE:
		return "0x002F001F"
	case IOCTL_INTERNAL_I8042_HOOK_KEYBOARD:
		return "0x000B3FC3"
	case IOCTL_INTERNAL_I8042_HOOK_MOUSE:
		return "0x000F3FC3"
	case IOCTL_INTERNAL_I8042_KEYBOARD_WRITE_BUFFER:
		return "0x000B3FC7"
	case IOCTL_INTERNAL_I8042_MOUSE_WRITE_BUFFER:
		return "0x000F3FC7"
	case IOCTL_INTERNAL_I8042_CONTROLLER_WRITE_BUFFER:
		return "0x000B3FCB"
	case IOCTL_INTERNAL_I8042_KEYBOARD_START_INFORMATION:
		return "0x000B3FCF"
	case IOCTL_INTERNAL_I8042_MOUSE_START_INFORMATION:
		return "0x000F3FCF"
	case IOCTL_BEEP_SET:
		return "0x00010000"
	case IOCTL_CDROM_UNLOAD_DRIVER:
		return "0x00025008"
	case IOCTL_CDROM_READ_TOC:
		return "0x00024000"
	case IOCTL_CDROM_SEEK_AUDIO_MSF:
		return "0x00024004"
	case IOCTL_CDROM_STOP_AUDIO:
		return "0x00024008"
	case IOCTL_CDROM_PAUSE_AUDIO:
		return "0x0002400C"
	case IOCTL_CDROM_RESUME_AUDIO:
		return "0x00024010"
	case IOCTL_CDROM_GET_VOLUME:
		return "0x00024014"
	case IOCTL_CDROM_PLAY_AUDIO_MSF:
		return "0x00024018"
	case IOCTL_CDROM_SET_VOLUME:
		return "0x00024028"
	case IOCTL_CDROM_READ_Q_CHANNEL:
		return "0x0002402C"
	case IOCTL_CDROM_GET_CONTROL:
		return "0x00024034"
	case OBSOLETE_IOCTL_CDROM_GET_CONTROL:
		return "0x00024034"
	case IOCTL_CDROM_GET_LAST_SESSION:
		return "0x00024038"
	case IOCTL_CDROM_RAW_READ:
		return "0x0002403E"
	case IOCTL_CDROM_DISK_TYPE:
		return "0x00020040"
	case IOCTL_CDROM_GET_DRIVE_GEOMETRY:
		return "0x0002404C"
	case IOCTL_CDROM_GET_DRIVE_GEOMETRY_EX:
		return "0x00024050"
	case IOCTL_CDROM_READ_TOC_EX:
		return "0x00024054"
	case IOCTL_CDROM_GET_CONFIGURATION:
		return "0x00024058"
	case IOCTL_CDROM_EXCLUSIVE_ACCESS:
		return "0x0002C05C"
	case IOCTL_CDROM_SET_SPEED:
		return "0x00024060"
	case IOCTL_CDROM_GET_INQUIRY_DATA:
		return "0x00024064"
	case IOCTL_CDROM_ENABLE_STREAMING:
		return "0x00024068"
	case IOCTL_CDROM_SEND_OPC_INFORMATION:
		return "0x0002C06C"
	case IOCTL_CDROM_GET_PERFORMANCE:
		return "0x00024070"
	case IOCTL_CDROM_CHECK_VERIFY:
		return "0x00024800"
	case IOCTL_CDROM_MEDIA_REMOVAL:
		return "0x00024804"
	case IOCTL_CDROM_EJECT_MEDIA:
		return "0x00024808"
	case IOCTL_CDROM_LOAD_MEDIA:
		return "0x0002480C"
	case IOCTL_CDROM_RESERVE:
		return "0x00024810"
	case IOCTL_CDROM_RELEASE:
		return "0x00024814"
	case IOCTL_CDROM_FIND_NEW_DEVICES:
		return "0x00024818"
	case IOCTL_CDROM_SIMBAD:
		return "0x0002400C"
	case IOCTL_DVD_START_SESSION:
		return "0x00335000"
	case IOCTL_DVD_READ_KEY:
		return "0x00335004"
	case IOCTL_DVD_SEND_KEY:
		return "0x00335008"
	case IOCTL_DVD_END_SESSION:
		return "0x0033500C"
	case IOCTL_DVD_SET_READ_AHEAD:
		return "0x00335010"
	case IOCTL_DVD_GET_REGION:
		return "0x00335014"
	case IOCTL_DVD_SEND_KEY2:
		return "0x0033D018"
	case IOCTL_AACS_READ_MEDIA_KEY_BLOCK_SIZE:
		return "0x003350C0"
	case IOCTL_AACS_READ_MEDIA_KEY_BLOCK:
		return "0x003350C4"
	case IOCTL_AACS_START_SESSION:
		return "0x003350C8"
	case IOCTL_AACS_END_SESSION:
		return "0x003350CC"
	case IOCTL_AACS_SEND_CERTIFICATE:
		return "0x003350D0"
	case IOCTL_AACS_GET_CERTIFICATE:
		return "0x003350D4"
	case IOCTL_AACS_GET_CHALLENGE_KEY:
		return "0x003350D8"
	case IOCTL_AACS_SEND_CHALLENGE_KEY:
		return "0x003350DC"
	case IOCTL_AACS_READ_VOLUME_ID:
		return "0x003350E0"
	case IOCTL_AACS_READ_SERIAL_NUMBER:
		return "0x003350E4"
	case IOCTL_AACS_READ_MEDIA_ID:
		return "0x003350E8"
	case IOCTL_AACS_READ_BINDING_NONCE:
		return "0x003350EC"
	case IOCTL_AACS_GENERATE_BINDING_NONCE:
		return "0x0033D0F0"
	case IOCTL_DVD_READ_STRUCTURE:
		return "0x00335140"
	case IOCTL_STORAGE_SET_READ_AHEAD:
		return "0x002D4400"
	case IOCTL_CHANGER_GET_PARAMETERS:
		return "0x00304000"
	case IOCTL_CHANGER_GET_STATUS:
		return "0x00304004"
	case IOCTL_CHANGER_GET_PRODUCT_DATA:
		return "0x00304008"
	case IOCTL_CHANGER_SET_ACCESS:
		return "0x0030C010"
	case IOCTL_CHANGER_GET_ELEMENT_STATUS:
		return "0x0030C014"
	case IOCTL_CHANGER_INITIALIZE_ELEMENT_STATUS:
		return "0x00304018"
	case IOCTL_CHANGER_SET_POSITION:
		return "0x0030401C"
	case IOCTL_CHANGER_EXCHANGE_MEDIUM:
		return "0x00304020"
	case IOCTL_CHANGER_MOVE_MEDIUM:
		return "0x00304024"
	case IOCTL_CHANGER_REINITIALIZE_TRANSPORT:
		return "0x00304028"
	case IOCTL_CHANGER_QUERY_VOLUME_TAGS:
		return "0x0030C02C"
	case IOCTL_DISK_GET_DRIVE_GEOMETRY:
		return "0x00070000"
	case IOCTL_DISK_GET_PARTITION_INFO:
		return "0x00074004"
	case IOCTL_DISK_SET_PARTITION_INFO:
		return "0x0007C008"
	case IOCTL_DISK_GET_DRIVE_LAYOUT:
		return "0x0007400C"
	case IOCTL_DISK_SET_DRIVE_LAYOUT:
		return "0x0007C010"
	case IOCTL_DISK_VERIFY:
		return "0x00070014"
	case IOCTL_DISK_FORMAT_TRACKS:
		return "0x0007C018"
	case IOCTL_DISK_REASSIGN_BLOCKS:
		return "0x0007C01C"
	case IOCTL_DISK_PERFORMANCE:
		return "0x00070020"
	case IOCTL_DISK_IS_WRITABLE:
		return "0x00070024"
	case IOCTL_DISK_LOGGING:
		return "0x00070028"
	case IOCTL_DISK_FORMAT_TRACKS_EX:
		return "0x0007C02C"
	case IOCTL_DISK_HISTOGRAM_STRUCTURE:
		return "0x00070030"
	case IOCTL_DISK_HISTOGRAM_DATA:
		return "0x00070034"
	case IOCTL_DISK_HISTOGRAM_RESET:
		return "0x00070038"
	case IOCTL_DISK_REQUEST_STRUCTURE:
		return "0x0007003C"
	case IOCTL_DISK_REQUEST_DATA:
		return "0x00070040"
	case IOCTL_DISK_PERFORMANCE_OFF:
		return "0x00070060"
	case IOCTL_DISK_CONTROLLER_NUMBER:
		return "0x00070044"
	case SMART_GET_VERSION:
		return "0x00074080"
	case SMART_SEND_DRIVE_COMMAND:
		return "0x0007C084"
	case SMART_RCV_DRIVE_DATA:
		return "0x0007C088"
	case IOCTL_DISK_GET_PARTITION_INFO_EX:
		return "0x00070048"
	case IOCTL_DISK_SET_PARTITION_INFO_EX:
		return "0x0007C04C"
	case IOCTL_DISK_GET_DRIVE_LAYOUT_EX:
		return "0x00070050"
	case IOCTL_DISK_SET_DRIVE_LAYOUT_EX:
		return "0x0007C054"
	case IOCTL_DISK_CREATE_DISK:
		return "0x0007C058"
	case IOCTL_DISK_GET_LENGTH_INFO:
		return "0x0007405C"
	case IOCTL_DISK_GET_DRIVE_GEOMETRY_EX:
		return "0x000700A0"
	case IOCTL_DISK_REASSIGN_BLOCKS_EX:
		return "0x0007C0A4"
	case IOCTL_DISK_UPDATE_DRIVE_SIZE:
		return "0x0007C0C8"
	case IOCTL_DISK_GROW_PARTITION:
		return "0x0007C0D0"
	case IOCTL_DISK_GET_CACHE_INFORMATION:
		return "0x000740D4"
	case IOCTL_DISK_SET_CACHE_INFORMATION:
		return "0x0007C0D8"
	case IOCTL_DISK_GET_WRITE_CACHE_STATE:
		return "0x000740DC"
	case OBSOLETE_DISK_GET_WRITE_CACHE_STATE:
		return "0x000740DC"
	case IOCTL_DISK_DELETE_DRIVE_LAYOUT:
		return "0x0007C100"
	case IOCTL_DISK_UPDATE_PROPERTIES:
		return "0x00070140"
	case IOCTL_DISK_FORMAT_DRIVE:
		return "0x0007C3CC"
	case IOCTL_DISK_SENSE_DEVICE:
		return "0x000703E0"
	case IOCTL_DISK_GET_CACHE_SETTING:
		return "0x000740E0"
	case IOCTL_DISK_SET_CACHE_SETTING:
		return "0x0007C0E4"
	case IOCTL_DISK_COPY_DATA:
		return "0x0007C064"
	case IOCTL_DISK_INTERNAL_SET_VERIFY:
		return "0x00070403"
	case IOCTL_DISK_INTERNAL_CLEAR_VERIFY:
		return "0x00070407"
	case IOCTL_DISK_INTERNAL_SET_NOTIFY:
		return "0x00070408"
	case IOCTL_DISK_CHECK_VERIFY:
		return "0x00074800"
	case IOCTL_DISK_MEDIA_REMOVAL:
		return "0x00074804"
	case IOCTL_DISK_EJECT_MEDIA:
		return "0x00074808"
	case IOCTL_DISK_LOAD_MEDIA:
		return "0x0007480C"
	case IOCTL_DISK_RESERVE:
		return "0x00074810"
	case IOCTL_DISK_RELEASE:
		return "0x00074814"
	case IOCTL_DISK_FIND_NEW_DEVICES:
		return "0x00074818"
	case IOCTL_DISK_GET_MEDIA_TYPES:
		return "0x00070C00"
	case IOCTL_DISK_GET_PARTITION_ATTRIBUTES:
		return "0x000700E8"
	case IOCTL_DISK_SET_PARTITION_ATTRIBUTES:
		return "0x0007C0EC"
	case IOCTL_DISK_GET_DISK_ATTRIBUTES:
		return "0x000700F0"
	case IOCTL_DISK_SET_DISK_ATTRIBUTES:
		return "0x0007C0F4"
	case IOCTL_DISK_IS_CLUSTERED:
		return "0x000700F8"
	case IOCTL_DISK_GET_SAN_SETTINGS:
		return "0x00074200"
	case IOCTL_DISK_SET_SAN_SETTINGS:
		return "0x0007C204"
	case IOCTL_DISK_GET_SNAPSHOT_INFO:
		return "0x00074208"
	case IOCTL_DISK_SET_SNAPSHOT_INFO:
		return "0x0007C20C"
	case IOCTL_DISK_RESET_SNAPSHOT_INFO:
		return "0x0007C210"
	case IOCTL_DISK_SIMBAD:
		return "0x0007D000"
	case FT_SECONDARY_READ:
		return "0x00664012"
	case FT_PRIMARY_READ:
		return "0x00664016"
	case FT_BALANCED_READ_MODE:
		return "0x0066001B"
	case FT_SYNC_REDUNDANT_COPY:
		return "0x0066001C"
	case FT_INITIALIZE_SET:
		return "0x00660000"
	case FT_REGENERATE:
		return "0x00660004"
	case FT_CONFIGURE:
		return "0x0066000B"
	case FT_VERIFY:
		return "0x0066000C"
	case FT_SEQUENTIAL_WRITE_MODE:
		return "0x00660023"
	case FT_PARALLEL_WRITE_MODE:
		return "0x00660027"
	case FT_QUERY_SET_STATE:
		return "0x00660028"
	case FT_CLUSTER_SET_MEMBER_STATE:
		return "0x0066002C"
	case FT_CLUSTER_GET_MEMBER_STATE:
		return "0x00660030"
	case FT_ENUMERATE_LOGICAL_DISKS:
		return "0x00674008"
	case FT_CREATE_LOGICAL_DISK:
		return "0x0067C000"
	case FT_BREAK_LOGICAL_DISK:
		return "0x0067C004"
	case FT_QUERY_LOGICAL_DISK_INFORMATION:
		return "0x0067400C"
	case FT_ORPHAN_LOGICAL_DISK_MEMBER:
		return "0x0067C010"
	case FT_REPLACE_LOGICAL_DISK_MEMBER:
		return "0x0067C014"
	case FT_QUERY_NT_DEVICE_NAME_FOR_LOGICAL_DISK:
		return "0x00674018"
	case FT_INITIALIZE_LOGICAL_DISK:
		return "0x0067C01C"
	case FT_QUERY_DRIVE_LETTER_FOR_LOGICAL_DISK:
		return "0x00674020"
	case FT_CHECK_IO:
		return "0x00674024"
	case FT_SET_DRIVE_LETTER_FOR_LOGICAL_DISK:
		return "0x0067C028"
	case FT_QUERY_NT_DEVICE_NAME_FOR_PARTITION:
		return "0x00674030"
	case FT_CHANGE_NOTIFY:
		return "0x00674034"
	case FT_STOP_SYNC_OPERATIONS:
		return "0x0067C038"
	case FT_QUERY_LOGICAL_DISK_ID:
		return "0x00674190"
	case FT_CREATE_PARTITION_LOGICAL_DISK:
		return "0x0067C194"
	case IOCTL_KEYBOARD_QUERY_ATTRIBUTES:
		return "0x000B0000"
	case IOCTL_KEYBOARD_SET_TYPEMATIC:
		return "0x000B0004"
	case IOCTL_KEYBOARD_SET_INDICATORS:
		return "0x000B0008"
	case IOCTL_KEYBOARD_QUERY_TYPEMATIC:
		return "0x000B0020"
	case IOCTL_KEYBOARD_QUERY_INDICATORS:
		return "0x000B0040"
	case IOCTL_KEYBOARD_QUERY_INDICATOR_TRANSLATION:
		return "0x000B0080"
	case IOCTL_KEYBOARD_INSERT_DATA:
		return "0x000B0100"
	case IOCTL_KEYBOARD_QUERY_IME_STATUS:
		return "0x000B1000"
	case IOCTL_KEYBOARD_SET_IME_STATUS:
		return "0x000B1004"
	case IOCTL_MODEM_GET_PASSTHROUGH:
		return "0x002B0004"
	case IOCTL_MODEM_SET_PASSTHROUGH:
		return "0x002B0008"
	case IOCTL_MODEM_SET_DLE_MONITORING:
		return "0x002B000C"
	case IOCTL_MODEM_GET_DLE:
		return "0x002B0010"
	case IOCTL_MODEM_SET_DLE_SHIELDING:
		return "0x002B0014"
	case IOCTL_MODEM_STOP_WAVE_RECEIVE:
		return "0x002B0018"
	case IOCTL_MODEM_SEND_MESSAGE:
		return "0x002B001C"
	case IOCTL_MODEM_GET_MESSAGE:
		return "0x002B0020"
	case IOCTL_MODEM_SEND_GET_MESSAGE:
		return "0x002B0024"
	case IOCTL_MODEM_SEND_LOOPBACK_MESSAGE:
		return "0x002B0028"
	case IOCTL_MODEM_CHECK_FOR_MODEM:
		return "0x002B002C"
	case IOCTL_MODEM_SET_MIN_POWER:
		return "0x002B0030"
	case IOCTL_MODEM_WATCH_FOR_RESUME:
		return "0x002B0034"
	case IOCTL_CANCEL_GET_SEND_MESSAGE:
		return "0x002B0038"
	case IOCTL_SET_SERVER_STATE:
		return "0x002B003C"
	case IOCTL_MOUSE_QUERY_ATTRIBUTES:
		return "0x000F0000"
	case IOCTL_MOUSE_INSERT_DATA:
		return "0x000F0004"
	case IOCTL_NDIS_RESERVED5:
		return "0x00170034"
	case IOCTL_NDIS_RESERVED6:
		return "0x00178038"
	case IOCTL_PAR_QUERY_INFORMATION:
		return "0x00160004"
	case IOCTL_PAR_SET_INFORMATION:
		return "0x00160008"
	case IOCTL_PAR_QUERY_DEVICE_ID:
		return "0x0016000C"
	case IOCTL_PAR_QUERY_DEVICE_ID_SIZE:
		return "0x00160010"
	case IOCTL_IEEE1284_GET_MODE:
		return "0x00160014"
	case IOCTL_IEEE1284_NEGOTIATE:
		return "0x00160018"
	case IOCTL_PAR_SET_WRITE_ADDRESS:
		return "0x0016001C"
	case IOCTL_PAR_SET_READ_ADDRESS:
		return "0x00160020"
	case IOCTL_PAR_GET_DEVICE_CAPS:
		return "0x00160024"
	case IOCTL_PAR_GET_DEFAULT_MODES:
		return "0x00160028"
	case IOCTL_PAR_PING:
		return "0x0016002C"
	case IOCTL_PAR_QUERY_RAW_DEVICE_ID:
		return "0x00160030"
	case IOCTL_PAR_ECP_HOST_RECOVERY:
		return "0x00160034"
	case IOCTL_PAR_GET_READ_ADDRESS:
		return "0x00160038"
	case IOCTL_PAR_GET_WRITE_ADDRESS:
		return "0x0016003C"
	case IOCTL_PAR_TEST:
		return "0x00160050"
	case IOCTL_PAR_IS_PORT_FREE:
		return "0x00160054"
	case IOCTL_PAR_QUERY_LOCATION:
		return "0x00160058"
	case IOCTL_SCSI_PASS_THROUGH:
		return "0x0004D004"
	case IOCTL_SCSI_MINIPORT:
		return "0x0004D008"
	case IOCTL_SCSI_GET_INQUIRY_DATA:
		return "0x0004100C"
	case IOCTL_SCSI_GET_CAPABILITIES:
		return "0x00041010"
	case IOCTL_SCSI_PASS_THROUGH_DIRECT:
		return "0x0004D014"
	case IOCTL_SCSI_GET_ADDRESS:
		return "0x00041018"
	case IOCTL_SCSI_RESCAN_BUS:
		return "0x0004101C"
	case IOCTL_SCSI_GET_DUMP_POINTERS:
		return "0x00041020"
	case IOCTL_SCSI_FREE_DUMP_POINTERS:
		return "0x00041024"
	case IOCTL_IDE_PASS_THROUGH:
		return "0x0004D028"
	case IOCTL_ATA_PASS_THROUGH:
		return "0x0004D02C"
	case IOCTL_ATA_PASS_THROUGH_DIRECT:
		return "0x0004D030"
	case IOCTL_ATA_MINIPORT:
		return "0x0004D034"
	case IOCTL_MINIPORT_PROCESS_SERVICE_IRP:
		return "0x0004D038"
	case IOCTL_MPIO_PASS_THROUGH_PATH:
		return "0x0004D03C"
	case IOCTL_MPIO_PASS_THROUGH_PATH_DIRECT:
		return "0x0004D040"
	case IOCTL_SERIAL_SET_BAUD_RATE:
		return "0x001B0004"
	case IOCTL_SERIAL_SET_QUEUE_SIZE:
		return "0x001B0008"
	case IOCTL_SERIAL_SET_LINE_CONTROL:
		return "0x001B000C"
	case IOCTL_SERIAL_SET_BREAK_ON:
		return "0x001B0010"
	case IOCTL_SERIAL_SET_BREAK_OFF:
		return "0x001B0014"
	case IOCTL_SERIAL_IMMEDIATE_CHAR:
		return "0x001B0018"
	case IOCTL_SERIAL_SET_TIMEOUTS:
		return "0x001B001C"
	case IOCTL_SERIAL_GET_TIMEOUTS:
		return "0x001B0020"
	case IOCTL_SERIAL_SET_DTR:
		return "0x001B0024"
	case IOCTL_SERIAL_CLR_DTR:
		return "0x001B0028"
	case IOCTL_SERIAL_RESET_DEVICE:
		return "0x001B002C"
	case IOCTL_SERIAL_SET_RTS:
		return "0x001B0030"
	case IOCTL_SERIAL_CLR_RTS:
		return "0x001B0034"
	case IOCTL_SERIAL_SET_XOFF:
		return "0x001B0038"
	case IOCTL_SERIAL_SET_XON:
		return "0x001B003C"
	case IOCTL_SERIAL_GET_WAIT_MASK:
		return "0x001B0040"
	case IOCTL_SERIAL_SET_WAIT_MASK:
		return "0x001B0044"
	case IOCTL_SERIAL_WAIT_ON_MASK:
		return "0x001B0048"
	case IOCTL_SERIAL_PURGE:
		return "0x001B004C"
	case IOCTL_SERIAL_GET_BAUD_RATE:
		return "0x001B0050"
	case IOCTL_SERIAL_GET_LINE_CONTROL:
		return "0x001B0054"
	case IOCTL_SERIAL_GET_CHARS:
		return "0x001B0058"
	case IOCTL_SERIAL_SET_CHARS:
		return "0x001B005C"
	case IOCTL_SERIAL_GET_HANDFLOW:
		return "0x001B0060"
	case IOCTL_SERIAL_SET_HANDFLOW:
		return "0x001B0064"
	case IOCTL_SERIAL_GET_MODEMSTATUS:
		return "0x001B0068"
	case IOCTL_SERIAL_GET_COMMSTATUS:
		return "0x001B006C"
	case IOCTL_SERIAL_XOFF_COUNTER:
		return "0x001B0070"
	case IOCTL_SERIAL_GET_PROPERTIES:
		return "0x001B0074"
	case IOCTL_SERIAL_GET_DTRRTS:
		return "0x001B0078"
	case IOCTL_SERIAL_LSRMST_INSERT:
		return "0x001B007C"
	case IOCTL_SERENUM_EXPOSE_HARDWARE:
		return "0x00370200"
	case IOCTL_SERENUM_REMOVE_HARDWARE:
		return "0x00370204"
	case IOCTL_SERENUM_PORT_DESC:
		return "0x00370208"
	case IOCTL_SERENUM_GET_PORT_NAME:
		return "0x0037020C"
	case IOCTL_SERIAL_CONFIG_SIZE:
		return "0x001B0080"
	case IOCTL_SERIAL_GET_COMMCONFIG:
		return "0x001B0084"
	case IOCTL_SERIAL_SET_COMMCONFIG:
		return "0x001B0088"
	case IOCTL_SERIAL_GET_STATS:
		return "0x001B008C"
	case IOCTL_SERIAL_CLEAR_STATS:
		return "0x001B0090"
	case IOCTL_SERIAL_GET_MODEM_CONTROL:
		return "0x001B0094"
	case IOCTL_SERIAL_SET_MODEM_CONTROL:
		return "0x001B0098"
	case IOCTL_SERIAL_SET_FIFO_CONTROL:
		return "0x001B009C"
	case IOCTL_SERIAL_INTERNAL_DO_WAIT_WAKE:
		return "0x001B0004"
	case IOCTL_SERIAL_INTERNAL_CANCEL_WAIT_WAKE:
		return "0x001B0008"
	case IOCTL_SERIAL_INTERNAL_BASIC_SETTINGS:
		return "0x001B000C"
	case IOCTL_SERIAL_INTERNAL_RESTORE_SETTINGS:
		return "0x001B0010"
	case IOCTL_STORAGE_CHECK_VERIFY:
		return "0x002D4800"
	case IOCTL_STORAGE_CHECK_VERIFY2:
		return "0x002D0800"
	case IOCTL_STORAGE_MEDIA_REMOVAL:
		return "0x002D4804"
	case IOCTL_STORAGE_EJECT_MEDIA:
		return "0x002D4808"
	case IOCTL_STORAGE_LOAD_MEDIA:
		return "0x002D480C"
	case IOCTL_STORAGE_LOAD_MEDIA2:
		return "0x002D080C"
	case IOCTL_STORAGE_RESERVE:
		return "0x002D4810"
	case IOCTL_STORAGE_RELEASE:
		return "0x002D4814"
	case IOCTL_STORAGE_FIND_NEW_DEVICES:
		return "0x002D4818"
	case IOCTL_STORAGE_EJECTION_CONTROL:
		return "0x002D0940"
	case IOCTL_STORAGE_MCN_CONTROL:
		return "0x002D0944"
	case IOCTL_STORAGE_GET_MEDIA_TYPES:
		return "0x002D0C00"
	case IOCTL_STORAGE_GET_MEDIA_TYPES_EX:
		return "0x002D0C04"
	case IOCTL_STORAGE_GET_MEDIA_SERIAL_NUMBER:
		return "0x002D0C10"
	case IOCTL_STORAGE_GET_HOTPLUG_INFO:
		return "0x002D0C14"
	case IOCTL_STORAGE_SET_HOTPLUG_INFO:
		return "0x002DCC18"
	case IOCTL_STORAGE_RESET_BUS:
		return "0x002D5000"
	case IOCTL_STORAGE_RESET_DEVICE:
		return "0x002D5004"
	case IOCTL_STORAGE_BREAK_RESERVATION:
		return "0x002D5014"
	case IOCTL_STORAGE_PERSISTENT_RESERVE_IN:
		return "0x002D5018"
	case IOCTL_STORAGE_PERSISTENT_RESERVE_OUT:
		return "0x002DD01C"
	case IOCTL_STORAGE_GET_DEVICE_NUMBER:
		return "0x002D1080"
	case IOCTL_STORAGE_PREDICT_FAILURE:
		return "0x002D1100"
	case IOCTL_STORAGE_READ_CAPACITY:
		return "0x002D5140"
	case IOCTL_STORAGE_QUERY_PROPERTY:
		return "0x002D1400"
	case IOCTL_STORAGE_MANAGE_DATA_SET_ATTRIBUTES:
		return "0x002D9404"
	case IOCTL_STORAGE_GET_BC_PROPERTIES:
		return "0x002D5800"
	case IOCTL_STORAGE_ALLOCATE_BC_STREAM:
		return "0x002DD804"
	case IOCTL_STORAGE_FREE_BC_STREAM:
		return "0x002DD808"
	case IOCTL_STORAGE_CHECK_PRIORITY_HINT_SUPPORT:
		return "0x002D1880"
	case OBSOLETE_IOCTL_STORAGE_RESET_BUS:
		return "0x002DD000"
	case OBSOLETE_IOCTL_STORAGE_RESET_DEVICE:
		return "0x002DD004"
	case IOCTL_TAPE_ERASE:
		return "0x001FC000"
	case IOCTL_TAPE_PREPARE:
		return "0x001F4004"
	case IOCTL_TAPE_WRITE_MARKS:
		return "0x001FC008"
	case IOCTL_TAPE_GET_POSITION:
		return "0x001F400C"
	case IOCTL_TAPE_SET_POSITION:
		return "0x001F4010"
	case IOCTL_TAPE_GET_DRIVE_PARAMS:
		return "0x001F4014"
	case IOCTL_TAPE_SET_DRIVE_PARAMS:
		return "0x001FC018"
	case IOCTL_TAPE_GET_MEDIA_PARAMS:
		return "0x001F401C"
	case IOCTL_TAPE_SET_MEDIA_PARAMS:
		return "0x001F4020"
	case IOCTL_TAPE_GET_STATUS:
		return "0x001F4024"
	case IOCTL_TAPE_CREATE_PARTITION:
		return "0x001FC028"
	case IOCTL_TAPE_MEDIA_REMOVAL:
		return "0x001F4804"
	case IOCTL_TAPE_EJECT_MEDIA:
		return "0x001F4808"
	case IOCTL_TAPE_LOAD_MEDIA:
		return "0x001F480C"
	case IOCTL_TAPE_RESERVE:
		return "0x001F4810"
	case IOCTL_TAPE_RELEASE:
		return "0x001F4814"
	case IOCTL_TAPE_CHECK_VERIFY:
		return "0x001F4800"
	case IOCTL_TAPE_FIND_NEW_DEVICES:
		return "0x00074818"
	case IOCTL_VOLUME_GET_VOLUME_DISK_EXTENTS:
		return "0x00560000"
	case IOCTL_VOLUME_ONLINE:
		return "0x0056C008"
	case IOCTL_VOLUME_OFFLINE:
		return "0x0056C00C"
	case IOCTL_VOLUME_IS_CLUSTERED:
		return "0x00560030"
	case IOCTL_VOLUME_GET_GPT_ATTRIBUTES:
		return "0x00560038"
	case IOCTL_VOLUME_SUPPORTS_ONLINE_OFFLINE:
		return "0x00560004"
	case IOCTL_VOLUME_IS_OFFLINE:
		return "0x00560010"
	case IOCTL_VOLUME_IS_IO_CAPABLE:
		return "0x00560014"
	case IOCTL_VOLUME_QUERY_FAILOVER_SET:
		return "0x00560018"
	case IOCTL_VOLUME_QUERY_VOLUME_NUMBER:
		return "0x0056001C"
	case IOCTL_VOLUME_LOGICAL_TO_PHYSICAL:
		return "0x00560020"
	case IOCTL_VOLUME_PHYSICAL_TO_LOGICAL:
		return "0x00560024"
	case IOCTL_VOLUME_IS_PARTITION:
		return "0x00560028"
	case IOCTL_VOLUME_READ_PLEX:
		return "0x0056402E"
	case IOCTL_VOLUME_SET_GPT_ATTRIBUTES:
		return "0x00560034"
	case IOCTL_VOLUME_GET_BC_PROPERTIES:
		return "0x0056403C"
	case IOCTL_VOLUME_ALLOCATE_BC_STREAM:
		return "0x0056C040"
	case IOCTL_VOLUME_FREE_BC_STREAM:
		return "0x0056C044"
	case IOCTL_VOLUME_IS_DYNAMIC:
		return "0x00560048"
	case IOCTL_VOLUME_PREPARE_FOR_CRITICAL_IO:
		return "0x0056C04C"
	case IOCTL_VOLUME_QUERY_ALLOCATION_HINT:
		return "0x00564052"
	case IOCTL_VOLUME_UPDATE_PROPERTIES:
		return "0x00560054"
	case IOCTL_VOLUME_QUERY_MINIMUM_SHRINK_SIZE:
		return "0x00564058"
	case IOCTL_VOLUME_PREPARE_FOR_SHRINK:
		return "0x0056C05C"
	case IOCTL_PMI_GET_CAPABILITIES:
		return "0x00454000"
	case IOCTL_PMI_GET_CONFIGURATION:
		return "0x00454004"
	case IOCTL_PMI_SET_CONFIGURATION:
		return "0x00458008"
	case IOCTL_PMI_GET_MEASUREMENT:
		return "0x0045400C"
	case IOCTL_PMI_REGISTER_EVENT_NOTIFY:
		return "0x0045C010"
	case IOCTL_BIOMETRIC_VENDOR:
		return "0x00442000"
	default:
		return "unknown"
	}
}

func (k IoctlKind) Elements() []IoctlKind {
	return []IoctlKind{
		MONITOR_IOCTL_ENABLE_MONITOR,
		MONITOR_IOCTL_DISABLE_MONITOR,
		IOCTL_INTERNAL_KEYBOARD_CONNECT,
		IOCTL_INTERNAL_KEYBOARD_DISCONNECT,
		IOCTL_INTERNAL_KEYBOARD_ENABLE,
		IOCTL_INTERNAL_KEYBOARD_DISABLE,
		IOCTL_INTERNAL_MOUSE_CONNECT,
		IOCTL_INTERNAL_MOUSE_DISCONNECT,
		IOCTL_INTERNAL_MOUSE_ENABLE,
		IOCTL_INTERNAL_MOUSE_DISABLE,
		IOCTL_DO_KERNELMODE_SAMPLES,
		IOCTL_REGISTER_CALLBACK,
		IOCTL_UNREGISTER_CALLBACK,
		IOCTL_GET_CALLBACK_VERSION,
		IOCTL_GET_VERSION,
		IOCTL_RESET,
		IOCTL_READWRITE_MAILBOX,
		IOCTL_MAILBOX_WAIT,
		IOCTL_READ_DMA,
		IOCTL_WRITE_DMA,
		IOCTL_ACPI_ASYNC_EVAL_METHOD,
		IOCTL_ACPI_EVAL_METHOD,
		IOCTL_ACPI_ACQUIRE_GLOBAL_LOCK,
		IOCTL_ACPI_RELEASE_GLOBAL_LOCK,
		IOCTL_ACPI_EVAL_METHOD_EX,
		IOCTL_ACPI_ASYNC_EVAL_METHOD_EX,
		IOCTL_ACPI_ENUM_CHILDREN,
		IOCTL_AVC_UPDATE_VIRTUAL_SUBUNIT_INFO,
		IOCTL_AVC_REMOVE_VIRTUAL_SUBUNIT_INFO,
		IOCTL_AVC_BUS_RESET,
		IOCTL_DOT4_CREATE_SOCKET,
		IOCTL_DOT4_DESTROY_SOCKET,
		IOCTL_DOT4_WAIT_FOR_CHANNEL,
		IOCTL_DOT4_OPEN_CHANNEL,
		IOCTL_DOT4_CLOSE_CHANNEL,
		IOCTL_DOT4_READ,
		IOCTL_DOT4_WRITE,
		IOCTL_DOT4_ADD_ACTIVITY_BROADCAST,
		IOCTL_DOT4_REMOVE_ACTIVITY_BROADCAST,
		IOCTL_DOT4_WAIT_ACTIVITY_BROADCAST,
		IOCTL_MPDSM_REGISTER,
		IOCTL_MPDSM_DEREGISTER,
		IOCTL_MOUNTDEV_QUERY_UNIQUE_ID,
		IOCTL_MOUNTDEV_QUERY_SUGGESTED_LINK_NAME,
		IOCTL_MOUNTDEV_LINK_CREATED,
		IOCTL_MOUNTDEV_LINK_DELETED,
		IOCTL_MOUNTDEV_QUERY_STABLE_GUID,
		IOCTL_MOUNTMGR_CREATE_POINT,
		IOCTL_MOUNTMGR_DELETE_POINTS,
		IOCTL_MOUNTMGR_QUERY_POINTS,
		IOCTL_MOUNTMGR_DELETE_POINTS_DBONLY,
		IOCTL_MOUNTMGR_NEXT_DRIVE_LETTER,
		IOCTL_MOUNTMGR_AUTO_DL_ASSIGNMENTS,
		IOCTL_MOUNTMGR_VOLUME_MOUNT_POINT_CREATED,
		IOCTL_MOUNTMGR_VOLUME_MOUNT_POINT_DELETED,
		IOCTL_MOUNTMGR_CHANGE_NOTIFY,
		IOCTL_MOUNTMGR_KEEP_LINKS_WHEN_OFFLINE,
		IOCTL_MOUNTMGR_CHECK_UNPROCESSED_VOLUMES,
		IOCTL_MOUNTMGR_VOLUME_ARRIVAL_NOTIFICATION,
		IOCTL_MOUNTDEV_QUERY_DEVICE_NAME,
		IOCTL_MOUNTMGR_QUERY_DOS_VOLUME_PATH,
		IOCTL_MOUNTMGR_QUERY_DOS_VOLUME_PATHS,
		IOCTL_MOUNTMGR_SCRUB_REGISTRY,
		IOCTL_MOUNTMGR_QUERY_AUTO_MOUNT,
		IOCTL_MOUNTMGR_SET_AUTO_MOUNT,
		IOCTL_MOUNTMGR_BOOT_DL_ASSIGNMENT,
		IOCTL_MOUNTMGR_TRACELOG_CACHE,
		IOCTL_AVIO_ALLOCATE_STREAM,
		IOCTL_AVIO_FREE_STREAM,
		IOCTL_AVIO_MODIFY_STREAM,
		NLB_IOCTL_REGISTER_HOOK,
		NLB_PUBLIC_IOCTL_CLIENT_STICKINESS_NOTIFY,
		IOCTL_GET_TUPLE_DATA,
		IOCTL_SOCKET_INFORMATION,
		IOCTL_PCMCIA_HIDE_DEVICE,
		IOCTL_PCMCIA_REVEAL_DEVICE,
		IOCTL_SD_SUBMIT_REQUEST,
		FSCTL_REQUEST_OPLOCK_LEVEL_1,
		FSCTL_REQUEST_OPLOCK_LEVEL_2,
		FSCTL_REQUEST_BATCH_OPLOCK,
		FSCTL_OPLOCK_BREAK_ACKNOWLEDGE,
		FSCTL_OPBATCH_ACK_CLOSE_PENDING,
		FSCTL_OPLOCK_BREAK_NOTIFY,
		FSCTL_LOCK_VOLUME,
		FSCTL_UNLOCK_VOLUME,
		FSCTL_DISMOUNT_VOLUME,
		FSCTL_IS_VOLUME_MOUNTED,
		FSCTL_IS_PATHNAME_VALID,
		FSCTL_MARK_VOLUME_DIRTY,
		FSCTL_QUERY_RETRIEVAL_POINTERS,
		FSCTL_GET_COMPRESSION,
		FSCTL_SET_COMPRESSION,
		FSCTL_SET_BOOTLOADER_ACCESSED,
		FSCTL_OPLOCK_BREAK_ACK_NO_2,
		FSCTL_INVALIDATE_VOLUMES,
		FSCTL_QUERY_FAT_BPB,
		FSCTL_REQUEST_FILTER_OPLOCK,
		FSCTL_FILESYSTEM_GET_STATISTICS,
		FSCTL_GET_NTFS_VOLUME_DATA,
		FSCTL_GET_NTFS_FILE_RECORD,
		FSCTL_GET_VOLUME_BITMAP,
		FSCTL_GET_RETRIEVAL_POINTERS,
		FSCTL_MOVE_FILE,
		FSCTL_IS_VOLUME_DIRTY,
		FSCTL_ALLOW_EXTENDED_DASD_IO,
		FSCTL_FIND_FILES_BY_SID,
		FSCTL_SET_OBJECT_ID,
		FSCTL_GET_OBJECT_ID,
		FSCTL_DELETE_OBJECT_ID,
		FSCTL_SET_REPARSE_POINT,
		FSCTL_GET_REPARSE_POINT,
		FSCTL_DELETE_REPARSE_POINT,
		FSCTL_ENUM_USN_DATA,
		FSCTL_SECURITY_ID_CHECK,
		FSCTL_READ_USN_JOURNAL,
		FSCTL_SET_OBJECT_ID_EXTENDED,
		FSCTL_CREATE_OR_GET_OBJECT_ID,
		FSCTL_SET_SPARSE,
		FSCTL_SET_ZERO_DATA,
		FSCTL_QUERY_ALLOCATED_RANGES,
		FSCTL_ENABLE_UPGRADE,
		FSCTL_SET_ENCRYPTION,
		FSCTL_ENCRYPTION_FSCTL_IO,
		FSCTL_WRITE_RAW_ENCRYPTED,
		FSCTL_READ_RAW_ENCRYPTED,
		FSCTL_CREATE_USN_JOURNAL,
		FSCTL_READ_FILE_USN_DATA,
		FSCTL_WRITE_USN_CLOSE_RECORD,
		FSCTL_EXTEND_VOLUME,
		FSCTL_QUERY_USN_JOURNAL,
		FSCTL_DELETE_USN_JOURNAL,
		FSCTL_MARK_HANDLE,
		FSCTL_SIS_COPYFILE,
		FSCTL_SIS_LINK_FILES,
		FSCTL_RECALL_FILE,
		FSCTL_READ_FROM_PLEX,
		FSCTL_FILE_PREFETCH,
		FSCTL_MAKE_MEDIA_COMPATIBLE,
		FSCTL_SET_DEFECT_MANAGEMENT,
		FSCTL_QUERY_SPARING_INFO,
		FSCTL_QUERY_ON_DISK_VOLUME_INFO,
		FSCTL_SET_VOLUME_COMPRESSION_STATE,
		FSCTL_TXFS_MODIFY_RM,
		FSCTL_TXFS_QUERY_RM_INFORMATION,
		FSCTL_TXFS_ROLLFORWARD_REDO,
		FSCTL_TXFS_ROLLFORWARD_UNDO,
		FSCTL_TXFS_START_RM,
		FSCTL_TXFS_SHUTDOWN_RM,
		FSCTL_TXFS_READ_BACKUP_INFORMATION,
		FSCTL_TXFS_WRITE_BACKUP_INFORMATION,
		FSCTL_TXFS_CREATE_SECONDARY_RM,
		FSCTL_TXFS_GET_METADATA_INFO,
		FSCTL_TXFS_GET_TRANSACTED_VERSION,
		FSCTL_TXFS_SAVEPOINT_INFORMATION,
		FSCTL_TXFS_CREATE_MINIVERSION,
		FSCTL_TXFS_TRANSACTION_ACTIVE,
		FSCTL_SET_ZERO_ON_DEALLOCATION,
		FSCTL_SET_REPAIR,
		FSCTL_GET_REPAIR,
		FSCTL_WAIT_FOR_REPAIR,
		FSCTL_INITIATE_REPAIR,
		FSCTL_CSC_INTERNAL,
		FSCTL_SHRINK_VOLUME,
		FSCTL_SET_SHORT_NAME_BEHAVIOR,
		FSCTL_DFSR_SET_GHOST_HANDLE_STATE,
		FSCTL_TXFS_LIST_TRANSACTIONS,
		FSCTL_QUERY_PAGEFILE_ENCRYPTION,
		FSCTL_RESET_VOLUME_ALLOCATION_HINTS,
		FSCTL_QUERY_DEPENDENT_VOLUME,
		FSCTL_SD_GLOBAL_CHANGE,
		FSCTL_TXFS_READ_BACKUP_INFORMATION2,
		FSCTL_LOOKUP_STREAM_FROM_CLUSTER,
		FSCTL_TXFS_WRITE_BACKUP_INFORMATION2,
		FSCTL_FILE_TYPE_NOTIFICATION,
		FSCTL_GET_BOOT_AREA_INFO,
		FSCTL_GET_RETRIEVAL_POINTER_BASE,
		FSCTL_SET_PERSISTENT_VOLUME_STATE,
		FSCTL_QUERY_PERSISTENT_VOLUME_STATE,
		FSCTL_REQUEST_OPLOCK,
		FSCTL_CSV_TUNNEL_REQUEST,
		FSCTL_IS_CSV_FILE,
		FSCTL_QUERY_FILE_SYSTEM_RECOGNITION,
		FSCTL_CSV_GET_VOLUME_PATH_NAME,
		FSCTL_CSV_GET_VOLUME_NAME_FOR_VOLUME_MOUNT_POINT,
		FSCTL_CSV_GET_VOLUME_PATH_NAMES_FOR_VOLUME_NAME,
		FSCTL_IS_FILE_ON_CSV_VOLUME,
		FSCTL_LMR_GET_LINK_TRACKING_INFORMATION,
		FSCTL_LMR_SET_LINK_TRACKING_INFORMATION,
		IOCTL_LMR_ARE_FILE_OBJECTS_ON_SAME_SERVER,
		FSCTL_PIPE_ASSIGN_EVENT,
		FSCTL_PIPE_DISCONNECT,
		FSCTL_PIPE_LISTEN,
		FSCTL_PIPE_PEEK,
		FSCTL_PIPE_QUERY_EVENT,
		FSCTL_PIPE_TRANSCEIVE,
		FSCTL_PIPE_WAIT,
		FSCTL_PIPE_IMPERSONATE,
		FSCTL_PIPE_SET_CLIENT_PROCESS,
		FSCTL_PIPE_QUERY_CLIENT_PROCESS,
		FSCTL_PIPE_GET_PIPE_ATTRIBUTE,
		FSCTL_PIPE_SET_PIPE_ATTRIBUTE,
		FSCTL_PIPE_GET_CONNECTION_ATTRIBUTE,
		FSCTL_PIPE_SET_CONNECTION_ATTRIBUTE,
		FSCTL_PIPE_GET_HANDLE_ATTRIBUTE,
		FSCTL_PIPE_SET_HANDLE_ATTRIBUTE,
		FSCTL_PIPE_FLUSH,
		FSCTL_PIPE_INTERNAL_READ,
		FSCTL_PIPE_INTERNAL_WRITE,
		FSCTL_PIPE_INTERNAL_TRANSCEIVE,
		FSCTL_PIPE_INTERNAL_READ_OVFLOW,
		FSCTL_MAILSLOT_PEEK,
		IOCTL_REDIR_QUERY_PATH,
		IOCTL_REDIR_QUERY_PATH_EX,
		IOCTL_VOLSNAP_FLUSH_AND_HOLD_WRITES,
		IOCTL_INTERNAL_PARALLEL_PORT_ALLOCATE,
		IOCTL_INTERNAL_GET_PARALLEL_PORT_INFO,
		IOCTL_INTERNAL_PARALLEL_CONNECT_INTERRUPT,
		IOCTL_INTERNAL_PARALLEL_DISCONNECT_INTERRUPT,
		IOCTL_INTERNAL_RELEASE_PARALLEL_PORT_INFO,
		IOCTL_INTERNAL_GET_MORE_PARALLEL_PORT_INFO,
		IOCTL_INTERNAL_PARCHIP_CONNECT,
		IOCTL_INTERNAL_PARALLEL_SET_CHIP_MODE,
		IOCTL_INTERNAL_PARALLEL_CLEAR_CHIP_MODE,
		IOCTL_INTERNAL_GET_PARALLEL_PNP_INFO,
		IOCTL_INTERNAL_INIT_1284_3_BUS,
		IOCTL_INTERNAL_SELECT_DEVICE,
		IOCTL_INTERNAL_DESELECT_DEVICE,
		IOCTL_INTERNAL_GET_PARPORT_FDO,
		IOCTL_INTERNAL_PARCLASS_CONNECT,
		IOCTL_INTERNAL_PARCLASS_DISCONNECT,
		IOCTL_INTERNAL_DISCONNECT_IDLE,
		IOCTL_INTERNAL_LOCK_PORT,
		IOCTL_INTERNAL_UNLOCK_PORT,
		IOCTL_INTERNAL_PARALLEL_PORT_FREE,
		IOCTL_INTERNAL_PARDOT3_CONNECT,
		IOCTL_INTERNAL_PARDOT3_DISCONNECT,
		IOCTL_INTERNAL_PARDOT3_RESET,
		IOCTL_INTERNAL_PARDOT3_SIGNAL,
		IOCTL_INTERNAL_REGISTER_FOR_REMOVAL_RELATIONS,
		IOCTL_INTERNAL_UNREGISTER_FOR_REMOVAL_RELATIONS,
		IOCTL_INTERNAL_LOCK_PORT_NO_SELECT,
		IOCTL_INTERNAL_UNLOCK_PORT_NO_DESELECT,
		IOCTL_INTERNAL_DISABLE_END_OF_CHAIN_BUS_RESCAN,
		IOCTL_INTERNAL_ENABLE_END_OF_CHAIN_BUS_RESCAN,
		IOCTL_WPD_MESSAGE_READWRITE_ACCESS,
		IOCTL_WPD_MESSAGE_READ_ACCESS,
		IOCTL_SCSISCAN_CMD,
		IOCTL_SCSISCAN_LOCKDEVICE,
		IOCTL_SCSISCAN_UNLOCKDEVICE,
		IOCTL_SCSISCAN_SET_TIMEOUT,
		IOCTL_SCSISCAN_GET_INFO,
		IOCTL_SWENUM_INSTALL_INTERFACE,
		IOCTL_SWENUM_REMOVE_INTERFACE,
		IOCTL_SWENUM_GET_BUS_ID,
		IOCTL_CANCEL_IO,
		IOCTL_WAIT_ON_DEVICE_EVENT,
		IOCTL_READ_REGISTERS,
		IOCTL_WRITE_REGISTERS,
		IOCTL_GET_CHANNEL_ALIGN_RQST,
		IOCTL_GET_DEVICE_DESCRIPTOR,
		IOCTL_RESET_PIPE,
		IOCTL_GET_USB_DESCRIPTOR,
		IOCTL_SEND_USB_REQUEST,
		IOCTL_GET_PIPE_CONFIGURATION,
		IOCTL_SET_TIMEOUT,
		IOCTL_EHSTOR_DEVICE_SET_AUTHZ_STATE,
		IOCTL_EHSTOR_DEVICE_GET_AUTHZ_STATE,
		IOCTL_EHSTOR_DEVICE_SILO_COMMAND,
		IOCTL_EHSTOR_DEVICE_ENUMERATE_PDOS,
		IOCTL_KS_PROPERTY,
		IOCTL_KS_ENABLE_EVENT,
		IOCTL_KS_DISABLE_EVENT,
		IOCTL_KS_METHOD,
		IOCTL_KS_WRITE_STREAM,
		IOCTL_KS_READ_STREAM,
		IOCTL_KS_RESET_STATE,
		IOCTL_KS_HANDSHAKE,
		IOCTL_INTERNAL_I8042_HOOK_KEYBOARD,
		IOCTL_INTERNAL_I8042_HOOK_MOUSE,
		IOCTL_INTERNAL_I8042_KEYBOARD_WRITE_BUFFER,
		IOCTL_INTERNAL_I8042_MOUSE_WRITE_BUFFER,
		IOCTL_INTERNAL_I8042_CONTROLLER_WRITE_BUFFER,
		IOCTL_INTERNAL_I8042_KEYBOARD_START_INFORMATION,
		IOCTL_INTERNAL_I8042_MOUSE_START_INFORMATION,
		IOCTL_BEEP_SET,
		IOCTL_CDROM_UNLOAD_DRIVER,
		IOCTL_CDROM_READ_TOC,
		IOCTL_CDROM_SEEK_AUDIO_MSF,
		IOCTL_CDROM_STOP_AUDIO,
		IOCTL_CDROM_PAUSE_AUDIO,
		IOCTL_CDROM_RESUME_AUDIO,
		IOCTL_CDROM_GET_VOLUME,
		IOCTL_CDROM_PLAY_AUDIO_MSF,
		IOCTL_CDROM_SET_VOLUME,
		IOCTL_CDROM_READ_Q_CHANNEL,
		IOCTL_CDROM_GET_CONTROL,
		OBSOLETE_IOCTL_CDROM_GET_CONTROL,
		IOCTL_CDROM_GET_LAST_SESSION,
		IOCTL_CDROM_RAW_READ,
		IOCTL_CDROM_DISK_TYPE,
		IOCTL_CDROM_GET_DRIVE_GEOMETRY,
		IOCTL_CDROM_GET_DRIVE_GEOMETRY_EX,
		IOCTL_CDROM_READ_TOC_EX,
		IOCTL_CDROM_GET_CONFIGURATION,
		IOCTL_CDROM_EXCLUSIVE_ACCESS,
		IOCTL_CDROM_SET_SPEED,
		IOCTL_CDROM_GET_INQUIRY_DATA,
		IOCTL_CDROM_ENABLE_STREAMING,
		IOCTL_CDROM_SEND_OPC_INFORMATION,
		IOCTL_CDROM_GET_PERFORMANCE,
		IOCTL_CDROM_CHECK_VERIFY,
		IOCTL_CDROM_MEDIA_REMOVAL,
		IOCTL_CDROM_EJECT_MEDIA,
		IOCTL_CDROM_LOAD_MEDIA,
		IOCTL_CDROM_RESERVE,
		IOCTL_CDROM_RELEASE,
		IOCTL_CDROM_FIND_NEW_DEVICES,
		IOCTL_CDROM_SIMBAD,
		IOCTL_DVD_START_SESSION,
		IOCTL_DVD_READ_KEY,
		IOCTL_DVD_SEND_KEY,
		IOCTL_DVD_END_SESSION,
		IOCTL_DVD_SET_READ_AHEAD,
		IOCTL_DVD_GET_REGION,
		IOCTL_DVD_SEND_KEY2,
		IOCTL_AACS_READ_MEDIA_KEY_BLOCK_SIZE,
		IOCTL_AACS_READ_MEDIA_KEY_BLOCK,
		IOCTL_AACS_START_SESSION,
		IOCTL_AACS_END_SESSION,
		IOCTL_AACS_SEND_CERTIFICATE,
		IOCTL_AACS_GET_CERTIFICATE,
		IOCTL_AACS_GET_CHALLENGE_KEY,
		IOCTL_AACS_SEND_CHALLENGE_KEY,
		IOCTL_AACS_READ_VOLUME_ID,
		IOCTL_AACS_READ_SERIAL_NUMBER,
		IOCTL_AACS_READ_MEDIA_ID,
		IOCTL_AACS_READ_BINDING_NONCE,
		IOCTL_AACS_GENERATE_BINDING_NONCE,
		IOCTL_DVD_READ_STRUCTURE,
		IOCTL_STORAGE_SET_READ_AHEAD,
		IOCTL_CHANGER_GET_PARAMETERS,
		IOCTL_CHANGER_GET_STATUS,
		IOCTL_CHANGER_GET_PRODUCT_DATA,
		IOCTL_CHANGER_SET_ACCESS,
		IOCTL_CHANGER_GET_ELEMENT_STATUS,
		IOCTL_CHANGER_INITIALIZE_ELEMENT_STATUS,
		IOCTL_CHANGER_SET_POSITION,
		IOCTL_CHANGER_EXCHANGE_MEDIUM,
		IOCTL_CHANGER_MOVE_MEDIUM,
		IOCTL_CHANGER_REINITIALIZE_TRANSPORT,
		IOCTL_CHANGER_QUERY_VOLUME_TAGS,
		IOCTL_DISK_GET_DRIVE_GEOMETRY,
		IOCTL_DISK_GET_PARTITION_INFO,
		IOCTL_DISK_SET_PARTITION_INFO,
		IOCTL_DISK_GET_DRIVE_LAYOUT,
		IOCTL_DISK_SET_DRIVE_LAYOUT,
		IOCTL_DISK_VERIFY,
		IOCTL_DISK_FORMAT_TRACKS,
		IOCTL_DISK_REASSIGN_BLOCKS,
		IOCTL_DISK_PERFORMANCE,
		IOCTL_DISK_IS_WRITABLE,
		IOCTL_DISK_LOGGING,
		IOCTL_DISK_FORMAT_TRACKS_EX,
		IOCTL_DISK_HISTOGRAM_STRUCTURE,
		IOCTL_DISK_HISTOGRAM_DATA,
		IOCTL_DISK_HISTOGRAM_RESET,
		IOCTL_DISK_REQUEST_STRUCTURE,
		IOCTL_DISK_REQUEST_DATA,
		IOCTL_DISK_PERFORMANCE_OFF,
		IOCTL_DISK_CONTROLLER_NUMBER,
		SMART_GET_VERSION,
		SMART_SEND_DRIVE_COMMAND,
		SMART_RCV_DRIVE_DATA,
		IOCTL_DISK_GET_PARTITION_INFO_EX,
		IOCTL_DISK_SET_PARTITION_INFO_EX,
		IOCTL_DISK_GET_DRIVE_LAYOUT_EX,
		IOCTL_DISK_SET_DRIVE_LAYOUT_EX,
		IOCTL_DISK_CREATE_DISK,
		IOCTL_DISK_GET_LENGTH_INFO,
		IOCTL_DISK_GET_DRIVE_GEOMETRY_EX,
		IOCTL_DISK_REASSIGN_BLOCKS_EX,
		IOCTL_DISK_UPDATE_DRIVE_SIZE,
		IOCTL_DISK_GROW_PARTITION,
		IOCTL_DISK_GET_CACHE_INFORMATION,
		IOCTL_DISK_SET_CACHE_INFORMATION,
		IOCTL_DISK_GET_WRITE_CACHE_STATE,
		OBSOLETE_DISK_GET_WRITE_CACHE_STATE,
		IOCTL_DISK_DELETE_DRIVE_LAYOUT,
		IOCTL_DISK_UPDATE_PROPERTIES,
		IOCTL_DISK_FORMAT_DRIVE,
		IOCTL_DISK_SENSE_DEVICE,
		IOCTL_DISK_GET_CACHE_SETTING,
		IOCTL_DISK_SET_CACHE_SETTING,
		IOCTL_DISK_COPY_DATA,
		IOCTL_DISK_INTERNAL_SET_VERIFY,
		IOCTL_DISK_INTERNAL_CLEAR_VERIFY,
		IOCTL_DISK_INTERNAL_SET_NOTIFY,
		IOCTL_DISK_CHECK_VERIFY,
		IOCTL_DISK_MEDIA_REMOVAL,
		IOCTL_DISK_EJECT_MEDIA,
		IOCTL_DISK_LOAD_MEDIA,
		IOCTL_DISK_RESERVE,
		IOCTL_DISK_RELEASE,
		IOCTL_DISK_FIND_NEW_DEVICES,
		IOCTL_DISK_GET_MEDIA_TYPES,
		IOCTL_DISK_GET_PARTITION_ATTRIBUTES,
		IOCTL_DISK_SET_PARTITION_ATTRIBUTES,
		IOCTL_DISK_GET_DISK_ATTRIBUTES,
		IOCTL_DISK_SET_DISK_ATTRIBUTES,
		IOCTL_DISK_IS_CLUSTERED,
		IOCTL_DISK_GET_SAN_SETTINGS,
		IOCTL_DISK_SET_SAN_SETTINGS,
		IOCTL_DISK_GET_SNAPSHOT_INFO,
		IOCTL_DISK_SET_SNAPSHOT_INFO,
		IOCTL_DISK_RESET_SNAPSHOT_INFO,
		IOCTL_DISK_SIMBAD,
		FT_SECONDARY_READ,
		FT_PRIMARY_READ,
		FT_BALANCED_READ_MODE,
		FT_SYNC_REDUNDANT_COPY,
		FT_INITIALIZE_SET,
		FT_REGENERATE,
		FT_CONFIGURE,
		FT_VERIFY,
		FT_SEQUENTIAL_WRITE_MODE,
		FT_PARALLEL_WRITE_MODE,
		FT_QUERY_SET_STATE,
		FT_CLUSTER_SET_MEMBER_STATE,
		FT_CLUSTER_GET_MEMBER_STATE,
		FT_ENUMERATE_LOGICAL_DISKS,
		FT_CREATE_LOGICAL_DISK,
		FT_BREAK_LOGICAL_DISK,
		FT_QUERY_LOGICAL_DISK_INFORMATION,
		FT_ORPHAN_LOGICAL_DISK_MEMBER,
		FT_REPLACE_LOGICAL_DISK_MEMBER,
		FT_QUERY_NT_DEVICE_NAME_FOR_LOGICAL_DISK,
		FT_INITIALIZE_LOGICAL_DISK,
		FT_QUERY_DRIVE_LETTER_FOR_LOGICAL_DISK,
		FT_CHECK_IO,
		FT_SET_DRIVE_LETTER_FOR_LOGICAL_DISK,
		FT_QUERY_NT_DEVICE_NAME_FOR_PARTITION,
		FT_CHANGE_NOTIFY,
		FT_STOP_SYNC_OPERATIONS,
		FT_QUERY_LOGICAL_DISK_ID,
		FT_CREATE_PARTITION_LOGICAL_DISK,
		IOCTL_KEYBOARD_QUERY_ATTRIBUTES,
		IOCTL_KEYBOARD_SET_TYPEMATIC,
		IOCTL_KEYBOARD_SET_INDICATORS,
		IOCTL_KEYBOARD_QUERY_TYPEMATIC,
		IOCTL_KEYBOARD_QUERY_INDICATORS,
		IOCTL_KEYBOARD_QUERY_INDICATOR_TRANSLATION,
		IOCTL_KEYBOARD_INSERT_DATA,
		IOCTL_KEYBOARD_QUERY_IME_STATUS,
		IOCTL_KEYBOARD_SET_IME_STATUS,
		IOCTL_MODEM_GET_PASSTHROUGH,
		IOCTL_MODEM_SET_PASSTHROUGH,
		IOCTL_MODEM_SET_DLE_MONITORING,
		IOCTL_MODEM_GET_DLE,
		IOCTL_MODEM_SET_DLE_SHIELDING,
		IOCTL_MODEM_STOP_WAVE_RECEIVE,
		IOCTL_MODEM_SEND_MESSAGE,
		IOCTL_MODEM_GET_MESSAGE,
		IOCTL_MODEM_SEND_GET_MESSAGE,
		IOCTL_MODEM_SEND_LOOPBACK_MESSAGE,
		IOCTL_MODEM_CHECK_FOR_MODEM,
		IOCTL_MODEM_SET_MIN_POWER,
		IOCTL_MODEM_WATCH_FOR_RESUME,
		IOCTL_CANCEL_GET_SEND_MESSAGE,
		IOCTL_SET_SERVER_STATE,
		IOCTL_MOUSE_QUERY_ATTRIBUTES,
		IOCTL_MOUSE_INSERT_DATA,
		IOCTL_NDIS_RESERVED5,
		IOCTL_NDIS_RESERVED6,
		IOCTL_PAR_QUERY_INFORMATION,
		IOCTL_PAR_SET_INFORMATION,
		IOCTL_PAR_QUERY_DEVICE_ID,
		IOCTL_PAR_QUERY_DEVICE_ID_SIZE,
		IOCTL_IEEE1284_GET_MODE,
		IOCTL_IEEE1284_NEGOTIATE,
		IOCTL_PAR_SET_WRITE_ADDRESS,
		IOCTL_PAR_SET_READ_ADDRESS,
		IOCTL_PAR_GET_DEVICE_CAPS,
		IOCTL_PAR_GET_DEFAULT_MODES,
		IOCTL_PAR_PING,
		IOCTL_PAR_QUERY_RAW_DEVICE_ID,
		IOCTL_PAR_ECP_HOST_RECOVERY,
		IOCTL_PAR_GET_READ_ADDRESS,
		IOCTL_PAR_GET_WRITE_ADDRESS,
		IOCTL_PAR_TEST,
		IOCTL_PAR_IS_PORT_FREE,
		IOCTL_PAR_QUERY_LOCATION,
		IOCTL_SCSI_PASS_THROUGH,
		IOCTL_SCSI_MINIPORT,
		IOCTL_SCSI_GET_INQUIRY_DATA,
		IOCTL_SCSI_GET_CAPABILITIES,
		IOCTL_SCSI_PASS_THROUGH_DIRECT,
		IOCTL_SCSI_GET_ADDRESS,
		IOCTL_SCSI_RESCAN_BUS,
		IOCTL_SCSI_GET_DUMP_POINTERS,
		IOCTL_SCSI_FREE_DUMP_POINTERS,
		IOCTL_IDE_PASS_THROUGH,
		IOCTL_ATA_PASS_THROUGH,
		IOCTL_ATA_PASS_THROUGH_DIRECT,
		IOCTL_ATA_MINIPORT,
		IOCTL_MINIPORT_PROCESS_SERVICE_IRP,
		IOCTL_MPIO_PASS_THROUGH_PATH,
		IOCTL_MPIO_PASS_THROUGH_PATH_DIRECT,
		IOCTL_SERIAL_SET_BAUD_RATE,
		IOCTL_SERIAL_SET_QUEUE_SIZE,
		IOCTL_SERIAL_SET_LINE_CONTROL,
		IOCTL_SERIAL_SET_BREAK_ON,
		IOCTL_SERIAL_SET_BREAK_OFF,
		IOCTL_SERIAL_IMMEDIATE_CHAR,
		IOCTL_SERIAL_SET_TIMEOUTS,
		IOCTL_SERIAL_GET_TIMEOUTS,
		IOCTL_SERIAL_SET_DTR,
		IOCTL_SERIAL_CLR_DTR,
		IOCTL_SERIAL_RESET_DEVICE,
		IOCTL_SERIAL_SET_RTS,
		IOCTL_SERIAL_CLR_RTS,
		IOCTL_SERIAL_SET_XOFF,
		IOCTL_SERIAL_SET_XON,
		IOCTL_SERIAL_GET_WAIT_MASK,
		IOCTL_SERIAL_SET_WAIT_MASK,
		IOCTL_SERIAL_WAIT_ON_MASK,
		IOCTL_SERIAL_PURGE,
		IOCTL_SERIAL_GET_BAUD_RATE,
		IOCTL_SERIAL_GET_LINE_CONTROL,
		IOCTL_SERIAL_GET_CHARS,
		IOCTL_SERIAL_SET_CHARS,
		IOCTL_SERIAL_GET_HANDFLOW,
		IOCTL_SERIAL_SET_HANDFLOW,
		IOCTL_SERIAL_GET_MODEMSTATUS,
		IOCTL_SERIAL_GET_COMMSTATUS,
		IOCTL_SERIAL_XOFF_COUNTER,
		IOCTL_SERIAL_GET_PROPERTIES,
		IOCTL_SERIAL_GET_DTRRTS,
		IOCTL_SERIAL_LSRMST_INSERT,
		IOCTL_SERENUM_EXPOSE_HARDWARE,
		IOCTL_SERENUM_REMOVE_HARDWARE,
		IOCTL_SERENUM_PORT_DESC,
		IOCTL_SERENUM_GET_PORT_NAME,
		IOCTL_SERIAL_CONFIG_SIZE,
		IOCTL_SERIAL_GET_COMMCONFIG,
		IOCTL_SERIAL_SET_COMMCONFIG,
		IOCTL_SERIAL_GET_STATS,
		IOCTL_SERIAL_CLEAR_STATS,
		IOCTL_SERIAL_GET_MODEM_CONTROL,
		IOCTL_SERIAL_SET_MODEM_CONTROL,
		IOCTL_SERIAL_SET_FIFO_CONTROL,
		IOCTL_SERIAL_INTERNAL_DO_WAIT_WAKE,
		IOCTL_SERIAL_INTERNAL_CANCEL_WAIT_WAKE,
		IOCTL_SERIAL_INTERNAL_BASIC_SETTINGS,
		IOCTL_SERIAL_INTERNAL_RESTORE_SETTINGS,
		IOCTL_STORAGE_CHECK_VERIFY,
		IOCTL_STORAGE_CHECK_VERIFY2,
		IOCTL_STORAGE_MEDIA_REMOVAL,
		IOCTL_STORAGE_EJECT_MEDIA,
		IOCTL_STORAGE_LOAD_MEDIA,
		IOCTL_STORAGE_LOAD_MEDIA2,
		IOCTL_STORAGE_RESERVE,
		IOCTL_STORAGE_RELEASE,
		IOCTL_STORAGE_FIND_NEW_DEVICES,
		IOCTL_STORAGE_EJECTION_CONTROL,
		IOCTL_STORAGE_MCN_CONTROL,
		IOCTL_STORAGE_GET_MEDIA_TYPES,
		IOCTL_STORAGE_GET_MEDIA_TYPES_EX,
		IOCTL_STORAGE_GET_MEDIA_SERIAL_NUMBER,
		IOCTL_STORAGE_GET_HOTPLUG_INFO,
		IOCTL_STORAGE_SET_HOTPLUG_INFO,
		IOCTL_STORAGE_RESET_BUS,
		IOCTL_STORAGE_RESET_DEVICE,
		IOCTL_STORAGE_BREAK_RESERVATION,
		IOCTL_STORAGE_PERSISTENT_RESERVE_IN,
		IOCTL_STORAGE_PERSISTENT_RESERVE_OUT,
		IOCTL_STORAGE_GET_DEVICE_NUMBER,
		IOCTL_STORAGE_PREDICT_FAILURE,
		IOCTL_STORAGE_READ_CAPACITY,
		IOCTL_STORAGE_QUERY_PROPERTY,
		IOCTL_STORAGE_MANAGE_DATA_SET_ATTRIBUTES,
		IOCTL_STORAGE_GET_BC_PROPERTIES,
		IOCTL_STORAGE_ALLOCATE_BC_STREAM,
		IOCTL_STORAGE_FREE_BC_STREAM,
		IOCTL_STORAGE_CHECK_PRIORITY_HINT_SUPPORT,
		OBSOLETE_IOCTL_STORAGE_RESET_BUS,
		OBSOLETE_IOCTL_STORAGE_RESET_DEVICE,
		IOCTL_TAPE_ERASE,
		IOCTL_TAPE_PREPARE,
		IOCTL_TAPE_WRITE_MARKS,
		IOCTL_TAPE_GET_POSITION,
		IOCTL_TAPE_SET_POSITION,
		IOCTL_TAPE_GET_DRIVE_PARAMS,
		IOCTL_TAPE_SET_DRIVE_PARAMS,
		IOCTL_TAPE_GET_MEDIA_PARAMS,
		IOCTL_TAPE_SET_MEDIA_PARAMS,
		IOCTL_TAPE_GET_STATUS,
		IOCTL_TAPE_CREATE_PARTITION,
		IOCTL_TAPE_MEDIA_REMOVAL,
		IOCTL_TAPE_EJECT_MEDIA,
		IOCTL_TAPE_LOAD_MEDIA,
		IOCTL_TAPE_RESERVE,
		IOCTL_TAPE_RELEASE,
		IOCTL_TAPE_CHECK_VERIFY,
		IOCTL_TAPE_FIND_NEW_DEVICES,
		IOCTL_VOLUME_GET_VOLUME_DISK_EXTENTS,
		IOCTL_VOLUME_ONLINE,
		IOCTL_VOLUME_OFFLINE,
		IOCTL_VOLUME_IS_CLUSTERED,
		IOCTL_VOLUME_GET_GPT_ATTRIBUTES,
		IOCTL_VOLUME_SUPPORTS_ONLINE_OFFLINE,
		IOCTL_VOLUME_IS_OFFLINE,
		IOCTL_VOLUME_IS_IO_CAPABLE,
		IOCTL_VOLUME_QUERY_FAILOVER_SET,
		IOCTL_VOLUME_QUERY_VOLUME_NUMBER,
		IOCTL_VOLUME_LOGICAL_TO_PHYSICAL,
		IOCTL_VOLUME_PHYSICAL_TO_LOGICAL,
		IOCTL_VOLUME_IS_PARTITION,
		IOCTL_VOLUME_READ_PLEX,
		IOCTL_VOLUME_SET_GPT_ATTRIBUTES,
		IOCTL_VOLUME_GET_BC_PROPERTIES,
		IOCTL_VOLUME_ALLOCATE_BC_STREAM,
		IOCTL_VOLUME_FREE_BC_STREAM,
		IOCTL_VOLUME_IS_DYNAMIC,
		IOCTL_VOLUME_PREPARE_FOR_CRITICAL_IO,
		IOCTL_VOLUME_QUERY_ALLOCATION_HINT,
		IOCTL_VOLUME_UPDATE_PROPERTIES,
		IOCTL_VOLUME_QUERY_MINIMUM_SHRINK_SIZE,
		IOCTL_VOLUME_PREPARE_FOR_SHRINK,
		IOCTL_PMI_GET_CAPABILITIES,
		IOCTL_PMI_GET_CONFIGURATION,
		IOCTL_PMI_SET_CONFIGURATION,
		IOCTL_PMI_GET_MEASUREMENT,
		IOCTL_PMI_REGISTER_EVENT_NOTIFY,
		IOCTL_BIOMETRIC_VENDOR,
	}
}
