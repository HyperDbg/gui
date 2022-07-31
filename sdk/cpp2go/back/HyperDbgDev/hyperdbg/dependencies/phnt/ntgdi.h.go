package phnt


const(
_NTGDI_H =  //col:13
GDI_MAX_HANDLE_COUNT = 0xFFFF // 0x4000 //col:15
GDI_HANDLE_INDEX_SHIFT = 0 //col:17
GDI_HANDLE_INDEX_BITS = 16 //col:18
GDI_HANDLE_INDEX_MASK = 0xffff //col:19
GDI_HANDLE_TYPE_SHIFT = 16 //col:21
GDI_HANDLE_TYPE_BITS = 5 //col:22
GDI_HANDLE_TYPE_MASK = 0x1f //col:23
GDI_HANDLE_ALTTYPE_SHIFT = 21 //col:25
GDI_HANDLE_ALTTYPE_BITS = 2 //col:26
GDI_HANDLE_ALTTYPE_MASK = 0x3 //col:27
GDI_HANDLE_STOCK_SHIFT = 23 //col:29
GDI_HANDLE_STOCK_BITS = 1 //col:30
GDI_HANDLE_STOCK_MASK = 0x1 //col:31
GDI_HANDLE_UNIQUE_SHIFT = 24 //col:33
GDI_HANDLE_UNIQUE_BITS = 8 //col:34
GDI_HANDLE_UNIQUE_MASK = 0xff //col:35
GDI_HANDLE_INDEX(Handle) = ((ULONG)(Handle) & GDI_HANDLE_INDEX_MASK) //col:37
GDI_HANDLE_TYPE(Handle) = (((ULONG)(Handle) >> GDI_HANDLE_TYPE_SHIFT) & GDI_HANDLE_TYPE_MASK) //col:38
GDI_HANDLE_ALTTYPE(Handle) = (((ULONG)(Handle) >> GDI_HANDLE_ALTTYPE_SHIFT) & GDI_HANDLE_ALTTYPE_MASK) //col:39
GDI_HANDLE_STOCK(Handle) = (((ULONG)(Handle) >> GDI_HANDLE_STOCK_SHIFT)) & GDI_HANDLE_STOCK_MASK) //col:40
GDI_MAKE_HANDLE(Index, = Unique) ((ULONG)(((ULONG)(Unique) << GDI_HANDLE_INDEX_BITS) | (ULONG)(Index))) //col:42
GDI_DEF_TYPE = 0 // invalid handle //col:46
GDI_DC_TYPE = 1 //col:47
GDI_DD_DIRECTDRAW_TYPE = 2 //col:48
GDI_DD_SURFACE_TYPE = 3 //col:49
GDI_RGN_TYPE = 4 //col:50
GDI_SURF_TYPE = 5 //col:51
GDI_CLIENTOBJ_TYPE = 6 //col:52
GDI_PATH_TYPE = 7 //col:53
GDI_PAL_TYPE = 8 //col:54
GDI_ICMLCS_TYPE = 9 //col:55
GDI_LFONT_TYPE = 10 //col:56
GDI_RFONT_TYPE = 11 //col:57
GDI_PFE_TYPE = 12 //col:58
GDI_PFT_TYPE = 13 //col:59
GDI_ICMCXF_TYPE = 14 //col:60
GDI_ICMDLL_TYPE = 15 //col:61
GDI_BRUSH_TYPE = 16 //col:62
GDI_PFF_TYPE = 17 // unused //col:63
GDI_CACHE_TYPE = 18 // unused //col:64
GDI_SPACE_TYPE = 19 //col:65
GDI_DBRUSH_TYPE = 20 // unused //col:66
GDI_META_TYPE = 21 //col:67
GDI_EFSTATE_TYPE = 22 //col:68
GDI_BMFD_TYPE = 23 // unused //col:69
GDI_VTFD_TYPE = 24 // unused //col:70
GDI_TTFD_TYPE = 25 // unused //col:71
GDI_RC_TYPE = 26 // unused //col:72
GDI_TEMP_TYPE = 27 // unused //col:73
GDI_DRVOBJ_TYPE = 28 //col:74
GDI_DCIOBJ_TYPE = 29 // unused //col:75
GDI_SPOOL_TYPE = 30 //col:76
GDI_CLIENT_TYPE_FROM_HANDLE(Handle) ((ULONG)(Handle) & ((GDI_HANDLE_ALTTYPE_MASK << GDI_HANDLE_ALTTYPE_SHIFT) | = (GDI_HANDLE_TYPE_MASK << GDI_HANDLE_TYPE_SHIFT))) //col:80
GDI_CLIENT_TYPE_FROM_UNIQUE(Unique) = GDI_CLIENT_TYPE_FROM_HANDLE((ULONG)(Unique) << 16) //col:82
GDI_ALTTYPE_1 = (1 << GDI_HANDLE_ALTTYPE_SHIFT) //col:84
GDI_ALTTYPE_2 = (2 << GDI_HANDLE_ALTTYPE_SHIFT) //col:85
GDI_ALTTYPE_3 = (3 << GDI_HANDLE_ALTTYPE_SHIFT) //col:86
GDI_CLIENT_BITMAP_TYPE = (GDI_SURF_TYPE << GDI_HANDLE_TYPE_SHIFT) //col:88
GDI_CLIENT_BRUSH_TYPE = (GDI_BRUSH_TYPE << GDI_HANDLE_TYPE_SHIFT) //col:89
GDI_CLIENT_CLIENTOBJ_TYPE = (GDI_CLIENTOBJ_TYPE << GDI_HANDLE_TYPE_SHIFT) //col:90
GDI_CLIENT_DC_TYPE = (GDI_DC_TYPE << GDI_HANDLE_TYPE_SHIFT) //col:91
GDI_CLIENT_FONT_TYPE = (GDI_LFONT_TYPE << GDI_HANDLE_TYPE_SHIFT) //col:92
GDI_CLIENT_PALETTE_TYPE = (GDI_PAL_TYPE << GDI_HANDLE_TYPE_SHIFT) //col:93
GDI_CLIENT_REGION_TYPE = (GDI_RGN_TYPE << GDI_HANDLE_TYPE_SHIFT) //col:94
GDI_CLIENT_ALTDC_TYPE = (GDI_CLIENT_DC_TYPE | GDI_ALTTYPE_1) //col:96
GDI_CLIENT_DIBSECTION_TYPE = (GDI_CLIENT_BITMAP_TYPE | GDI_ALTTYPE_1) //col:97
GDI_CLIENT_EXTPEN_TYPE = (GDI_CLIENT_BRUSH_TYPE | GDI_ALTTYPE_2) //col:98
GDI_CLIENT_METADC16_TYPE = (GDI_CLIENT_CLIENTOBJ_TYPE | GDI_ALTTYPE_3) //col:99
GDI_CLIENT_METAFILE_TYPE = (GDI_CLIENT_CLIENTOBJ_TYPE | GDI_ALTTYPE_2) //col:100
GDI_CLIENT_METAFILE16_TYPE = (GDI_CLIENT_CLIENTOBJ_TYPE | GDI_ALTTYPE_1) //col:101
GDI_CLIENT_PEN_TYPE = (GDI_CLIENT_BRUSH_TYPE | GDI_ALTTYPE_1) //col:102
)

type (
Ntgdi interface{
    ()(ok bool)//col:125
}






)

func NewNtgdi() { return & ntgdi{} }

func (n *ntgdi)    ()(ok bool){//col:125









































return true
}









