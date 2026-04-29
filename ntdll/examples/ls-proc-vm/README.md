# Example: List process memory map

This program opens a process and outputs its virtual memory regions
including type, state, protection flags, and backing file (if
applicable).

Output from a freshly started `notepad.exe` on a Windows 10 system:

``` console
PS C:\Users\user> get-process notepad

Handles  NPM(K)    PM(K)      WS(K)     CPU(s)     Id  SI ProcessName
-------  ------    -----      -----     ------     --  -- -----------
    235      13     2536      16224       0.19   6912   1 notepad


PS C:\Users\user> .\ls-proc-vm.exe -pid  6912
    2abe0000     2abe0000 +         1000 ty:-p- st:c-- pr:rw--
    2abf0000     2abf0000 +         1000 ty:-p- st:c-- pr:rw--
    7ffe0000     7ffe0000 +         1000 ty:-p- st:c-- pr:r---
    7ffe2000     7ffe2000 +         1000 ty:-p- st:c-- pr:r---
  177e650000   177e650000 +        6c000 ty:-p- st:-r- pr:----
  177e650000   177e6bc000 +         3000 ty:-p- st:c-- pr:rw-G
  177e650000   177e6bf000 +        11000 ty:-p- st:c-- pr:rw--
  177e800000   177e800000 +        1c000 ty:-p- st:-r- pr:----
  177e800000   177e81c000 +         3000 ty:-p- st:c-- pr:rw--
  177e800000   177e81f000 +       1e1000 ty:-p- st:-r- pr:----
 145eabe0000  145eabe0000 +         1000 ty:m-- st:c-- pr:r---
 145eabf0000  145eabf0000 +         1000 ty:m-- st:c-- pr:r---
 145eac00000  145eac00000 +        1d000 ty:m-- st:c-- pr:r---
 145eac20000  145eac20000 +         4000 ty:m-- st:c-- pr:r---
 145eac30000  145eac30000 +         3000 ty:m-- st:c-- pr:r---
 145eac40000  145eac40000 +         2000 ty:-p- st:c-- pr:rw--
 145eac50000  145eac50000 +         1000 ty:m-- st:c-- pr:r---
 145eac60000  145eac60000 +        10000 ty:m-- st:c-- pr:rw--
 145eac70000  145eac70000 +        c9000 ty:m-- st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\locale.nls
 145ead40000  145ead40000 +         2000 ty:-p- st:c-- pr:rw--
 145ead40000  145ead42000 +         b000 ty:-p- st:-r- pr:----
 145ead50000  145ead50000 +         4000 ty:m-- st:c-- pr:r---
 145ead50000  145ead54000 +         4000 ty:m-- st:-r- pr:----
 145ead60000  145ead60000 +         3000 ty:m-- st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\en-US\notepad.exe.mui
 145ead70000  145ead70000 +         1000 ty:-p- st:c-- pr:rw--
 145ead70000  145ead71000 +         c000 ty:-p- st:-r- pr:----
 145ead80000  145ead80000 +        77000 ty:-p- st:c-- pr:rw--
 145ead80000  145eadf7000 +        89000 ty:-p- st:-r- pr:----
 145eae80000  145eae80000 +        12000 ty:m-- st:c-- pr:r---
 145eae80000  145eae92000 +       1ee000 ty:m-- st:-r- pr:----
 145eb080000  145eb080000 +       181000 ty:m-- st:c-- pr:r---
 145eb210000  145eb210000 +        b4000 ty:m-- st:c-- pr:r---
 145eb210000  145eb2c4000 +      134d000 ty:m-- st:-r- pr:----
 145ec620000  145ec620000 +         1000 ty:m-- st:c-- pr:r---
 145ec630000  145ec630000 +         2000 ty:m-- st:c-- pr:r---
 145ec640000  145ec640000 +        47000 ty:m-- st:c-- pr:-w-c
 145ec690000  145ec690000 +         1000 ty:-p- st:c-- pr:rw--
 145ec6a0000  145ec6a0000 +        1a000 ty:m-- st:c-- pr:r--- \Device\HarddiskVolume2\Windows\SystemResources\notepad.exe.mun
 145ec6c0000  145ec6c0000 +         1000 ty:-p- st:c-- pr:rw--
 145ec6d0000  145ec6d0000 +         1000 ty:m-- st:c-- pr:rw--
 145ec6e0000  145ec6e0000 +         1000 ty:-p- st:c-- pr:rw--
 145ec6f0000  145ec6f0000 +         4000 ty:m-- st:c-- pr:r---
 145ec700000  145ec700000 +         1000 ty:-p- st:c-- pr:rw--
 145ec700000  145ec701000 +         c000 ty:-p- st:-r- pr:----
 145ec710000  145ec710000 +         1000 ty:m-- st:c-- pr:r---
 145ec720000  145ec720000 +         6000 ty:m-- st:c-- pr:r--- \Device\HarddiskVolume2\Windows\Registration\R000000000006.clb
 145ec730000  145ec730000 +         1000 ty:-p- st:c-- pr:rw--
 145ec730000  145ec731000 +         f000 ty:-p- st:-r- pr:----
 145ec740000  145ec740000 +         2000 ty:m-- st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\oleaccrc.dll
 145ec750000  145ec750000 +         8000 ty:-p- st:c-- pr:rw--
 145ec750000  145ec758000 +         8000 ty:-p- st:-r- pr:----
 145ec800000  145ec800000 +         f000 ty:-p- st:c-- pr:rw--
 145ec800000  145ec80f000 +         1000 ty:-p- st:-r- pr:----
 145ec810000  145ec810000 +       338000 ty:m-- st:c-- pr:r--- \Device\HarddiskVolume2\Windows\Globalization\Sorting\SortDefault.nls
 145ecb50000  145ecb50000 +        e2000 ty:m-- st:c-- pr:r---
 145eda50000  145eda50000 +      1260000 ty:m-- st:c-- pr:r--- \Device\HarddiskVolume2\Windows\Fonts\StaticCache.dat
 145eecb0000  145eecb0000 +         1000 ty:-p- st:c-- pr:rw--
 145eecb0000  145eecb1000 +        ff000 ty:-p- st:-r- pr:----
 145eedb0000  145eedb0000 +       18b000 ty:m-- st:c-- pr:rw--
 145eef40000  145eef40000 +         1000 ty:-p- st:c-- pr:rw--
 145eef40000  145eef41000 +       7ff000 ty:-p- st:-r- pr:----
 145ef740000  145ef740000 +         7000 ty:-p- st:c-- pr:rw--
 145ef740000  145ef747000 +        f9000 ty:-p- st:-r- pr:----
7df45c150000 7df45c150000 +         5000 ty:m-- st:c-- pr:r---
7df45c150000 7df45c155000 +        fb000 ty:m-- st:-r- pr:----
7df45c250000 7df45c250000 +    100020000 ty:-p- st:-r- pr:----
7df55c270000 7df55c270000 +      2000000 ty:-p- st:-r- pr:----
7df55c270000 7df55e270000 +         1000 ty:-p- st:c-- pr:rw--
7df55e280000 7df55e280000 +         1000 ty:m-- st:c-- pr:r---
7df55e290000 7df55e290000 +        23000 ty:m-- st:c-- pr:r---
7df55e2c0000 7df55e2c0000 +      1b4e000 ty:m-- st:-r- pr:----
7df55e2c0000 7df55fe0e000 +       222000 ty:m-- st:c-- pr:---!
7df55e2c0000 7df560030000 +        6e000 ty:m-- st:-r- pr:----
7df55e2c0000 7df56009e000 +         1000 ty:m-- st:c-- pr:---!
7df55e2c0000 7df56009f000 +  1ffdd87f000 ty:m-- st:-r- pr:----
7df55e2c0000 7ff53d91e000 +         1000 ty:m-- st:c-- pr:r---
7df55e2c0000 7ff53d91f000 +     12245000 ty:m-- st:-r- pr:----
7df55e2c0000 7ff54fb64000 +      1184000 ty:m-- st:c-- pr:---!
7df55e2c0000 7ff550ce8000 +         c000 ty:m-- st:c-- pr:r---
7df55e2c0000 7ff550cf4000 +       134000 ty:m-- st:c-- pr:---!
7df55e2c0000 7ff550e28000 +         3000 ty:m-- st:c-- pr:r---
7df55e2c0000 7ff550e2b000 +        61000 ty:m-- st:c-- pr:---!
7df55e2c0000 7ff550e8c000 +         3000 ty:m-- st:c-- pr:r---
7df55e2c0000 7ff550e8f000 +        b4000 ty:m-- st:c-- pr:---!
7df55e2c0000 7ff550f43000 +         5000 ty:m-- st:c-- pr:r---
7df55e2c0000 7ff550f48000 +        48000 ty:m-- st:c-- pr:---!
7df55e2c0000 7ff550f90000 +         5000 ty:m-- st:c-- pr:r---
7df55e2c0000 7ff550f95000 +        17000 ty:m-- st:c-- pr:---!
7df55e2c0000 7ff550fac000 +         5000 ty:m-- st:c-- pr:r---
7df55e2c0000 7ff550fb1000 +        7f000 ty:m-- st:c-- pr:---!
7df55e2c0000 7ff551030000 +         1000 ty:m-- st:c-- pr:r---
7df55e2c0000 7ff551031000 +        f0000 ty:m-- st:c-- pr:---!
7df55e2c0000 7ff551121000 +         9000 ty:m-- st:c-- pr:r---
7df55e2c0000 7ff55112a000 +        4e000 ty:m-- st:c-- pr:---!
7df55e2c0000 7ff551178000 +         6000 ty:m-- st:c-- pr:r---
7df55e2c0000 7ff55117e000 +        20000 ty:m-- st:c-- pr:---!
7df55e2c0000 7ff55119e000 +        12000 ty:m-- st:c-- pr:r---
7df55e2c0000 7ff5511b0000 +         b000 ty:m-- st:c-- pr:---!
7df55e2c0000 7ff5511bb000 +         3000 ty:m-- st:c-- pr:r---
7df55e2c0000 7ff5511be000 +        11000 ty:m-- st:c-- pr:---!
7df55e2c0000 7ff5511cf000 +         1000 ty:m-- st:c-- pr:r---
7df55e2c0000 7ff5511d0000 +         7000 ty:m-- st:c-- pr:---!
7df55e2c0000 7ff5511d7000 +        1f000 ty:m-- st:c-- pr:r---
7df55e2c0000 7ff5511f6000 +        10000 ty:m-- st:c-- pr:---!
7df55e2c0000 7ff551206000 +         2000 ty:m-- st:c-- pr:r---
7df55e2c0000 7ff551208000 +        31000 ty:m-- st:c-- pr:---!
7df55e2c0000 7ff551239000 +         1000 ty:m-- st:c-- pr:r---
7df55e2c0000 7ff55123a000 +        1b000 ty:m-- st:c-- pr:---!
7df55e2c0000 7ff551255000 +         6000 ty:m-- st:c-- pr:r---
7df55e2c0000 7ff55125b000 +         1000 ty:m-- st:c-- pr:---!
7df55e2c0000 7ff55125c000 +         c000 ty:m-- st:c-- pr:r---
7df55e2c0000 7ff551268000 +         1000 ty:m-- st:c-- pr:---!
7df55e2c0000 7ff551269000 +         1000 ty:m-- st:c-- pr:r---
7df55e2c0000 7ff55126a000 +         5000 ty:m-- st:c-- pr:---!
7df55e2c0000 7ff55126f000 +         a000 ty:m-- st:c-- pr:r---
7df55e2c0000 7ff551279000 +         4000 ty:m-- st:c-- pr:---!
7df55e2c0000 7ff55127d000 +        21000 ty:m-- st:c-- pr:r---
7df55e2c0000 7ff55129e000 +         3000 ty:m-- st:c-- pr:---!
7df55e2c0000 7ff5512a1000 +        19000 ty:m-- st:c-- pr:r---
7df55e2c0000 7ff5512ba000 +         2000 ty:m-- st:c-- pr:---!
7df55e2c0000 7ff5512bc000 +         3000 ty:m-- st:c-- pr:r---
7df55e2c0000 7ff5512bf000 +         3000 ty:m-- st:c-- pr:---!
7df55e2c0000 7ff5512c2000 +         3000 ty:m-- st:c-- pr:r---
7df55e2c0000 7ff5512c5000 +         1000 ty:m-- st:c-- pr:---!
7df55e2c0000 7ff5512c6000 +         3000 ty:m-- st:c-- pr:r---
7df55e2c0000 7ff5512c9000 +         6000 ty:m-- st:c-- pr:---!
7df55e2c0000 7ff5512cf000 +         c000 ty:m-- st:c-- pr:r---
7df55e2c0000 7ff5512db000 +        11000 ty:m-- st:c-- pr:---!
7df55e2c0000 7ff5512ec000 +         5000 ty:m-- st:c-- pr:r---
7df55e2c0000 7ff5512f1000 +         1000 ty:m-- st:c-- pr:---!
7df55e2c0000 7ff5512f2000 +         d000 ty:m-- st:c-- pr:r---
7df55e2c0000 7ff5512ff000 +      cfc1000 ty:m-- st:-r- pr:----
7ff7d9780000 7ff7d9780000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\notepad.exe
7ff7d9780000 7ff7d9781000 +        25000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\notepad.exe
7ff7d9780000 7ff7d97a6000 +         a000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\notepad.exe
7ff7d9780000 7ff7d97b0000 +         3000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\notepad.exe
7ff7d9780000 7ff7d97b3000 +         5000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\notepad.exe
7ffca8a30000 7ffca8a30000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\WinSxS\amd64_microsoft.windows.common-controls_6595b64144ccf1df_6.0.19041.1110_none_60b5254171f9507e\comctl32.dll
7ffca8a30000 7ffca8a31000 +       1e3000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\WinSxS\amd64_microsoft.windows.common-controls_6595b64144ccf1df_6.0.19041.1110_none_60b5254171f9507e\comctl32.dll
7ffca8a30000 7ffca8c14000 +        49000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\WinSxS\amd64_microsoft.windows.common-controls_6595b64144ccf1df_6.0.19041.1110_none_60b5254171f9507e\comctl32.dll
7ffca8a30000 7ffca8c5d000 +         4000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\WinSxS\amd64_microsoft.windows.common-controls_6595b64144ccf1df_6.0.19041.1110_none_60b5254171f9507e\comctl32.dll
7ffca8a30000 7ffca8c61000 +        69000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\WinSxS\amd64_microsoft.windows.common-controls_6595b64144ccf1df_6.0.19041.1110_none_60b5254171f9507e\comctl32.dll
7ffcada30000 7ffcada30000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\oleacc.dll
7ffcada30000 7ffcada31000 +        40000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\oleacc.dll
7ffcada30000 7ffcada71000 +        17000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\oleacc.dll
7ffcada30000 7ffcada88000 +         1000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\oleacc.dll
7ffcada30000 7ffcada89000 +         d000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\oleacc.dll
7ffcaf310000 7ffcaf310000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\TextShaping.dll
7ffcaf310000 7ffcaf311000 +        4b000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\TextShaping.dll
7ffcaf310000 7ffcaf35c000 +        5b000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\TextShaping.dll
7ffcaf310000 7ffcaf3b7000 +         1000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\TextShaping.dll
7ffcaf310000 7ffcaf3b8000 +         4000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\TextShaping.dll
7ffcb20f0000 7ffcb20f0000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\efswrt.dll
7ffcb20f0000 7ffcb20f1000 +        87000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\efswrt.dll
7ffcb20f0000 7ffcb2178000 +        46000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\efswrt.dll
7ffcb20f0000 7ffcb21be000 +         2000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\efswrt.dll
7ffcb20f0000 7ffcb21c0000 +         e000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\efswrt.dll
7ffcb3420000 7ffcb3420000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\TextInputFramework.dll
7ffcb3420000 7ffcb3421000 +        b1000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\TextInputFramework.dll
7ffcb3420000 7ffcb34d2000 +        34000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\TextInputFramework.dll
7ffcb3420000 7ffcb3506000 +         3000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\TextInputFramework.dll
7ffcb3420000 7ffcb3509000 +        10000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\TextInputFramework.dll
7ffcb3b10000 7ffcb3b10000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\MrmCoreR.dll
7ffcb3b10000 7ffcb3b11000 +        a8000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\MrmCoreR.dll
7ffcb3b10000 7ffcb3bb9000 +        35000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\MrmCoreR.dll
7ffcb3b10000 7ffcb3bee000 +         1000 ty:--i st:c-- pr:-w-c \Device\HarddiskVolume2\Windows\System32\MrmCoreR.dll
7ffcb3b10000 7ffcb3bef000 +         2000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\MrmCoreR.dll
7ffcb3b10000 7ffcb3bf1000 +        13000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\MrmCoreR.dll
7ffcb5c20000 7ffcb5c20000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\mpr.dll
7ffcb5c20000 7ffcb5c21000 +        12000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\mpr.dll
7ffcb5c20000 7ffcb5c33000 +         5000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\mpr.dll
7ffcb5c20000 7ffcb5c38000 +         1000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\mpr.dll
7ffcb5c20000 7ffcb5c39000 +         4000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\mpr.dll
7ffcb9860000 7ffcb9860000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\twinapi.appcore.dll
7ffcb9860000 7ffcb9861000 +       162000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\twinapi.appcore.dll
7ffcb9860000 7ffcb99c3000 +        73000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\twinapi.appcore.dll
7ffcb9860000 7ffcb9a36000 +         4000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\twinapi.appcore.dll
7ffcb9860000 7ffcb9a3a000 +        26000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\twinapi.appcore.dll
7ffcbae10000 7ffcbae10000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\WinTypes.dll
7ffcbae10000 7ffcbae11000 +        77000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\WinTypes.dll
7ffcbae10000 7ffcbae88000 +        b7000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\WinTypes.dll
7ffcbae10000 7ffcbaf3f000 +         2000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\WinTypes.dll
7ffcbae10000 7ffcbaf41000 +        23000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\WinTypes.dll
7ffcbb780000 7ffcbb780000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\CoreUIComponents.dll
7ffcbb780000 7ffcbb781000 +       1b9000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\CoreUIComponents.dll
7ffcbb780000 7ffcbb93a000 +       128000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\CoreUIComponents.dll
7ffcbb780000 7ffcbba62000 +         2000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\CoreUIComponents.dll
7ffcbb780000 7ffcbba64000 +         1000 ty:--i st:c-- pr:-w-c \Device\HarddiskVolume2\Windows\System32\CoreUIComponents.dll
7ffcbb780000 7ffcbba65000 +         1000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\CoreUIComponents.dll
7ffcbb780000 7ffcbba66000 +        78000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\CoreUIComponents.dll
7ffcbbae0000 7ffcbbae0000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\CoreMessaging.dll
7ffcbbae0000 7ffcbbae1000 +        95000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\CoreMessaging.dll
7ffcbbae0000 7ffcbbb76000 +        39000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\CoreMessaging.dll
7ffcbbae0000 7ffcbbbaf000 +         3000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\CoreMessaging.dll
7ffcbbae0000 7ffcbbbb2000 +        20000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\CoreMessaging.dll
7ffcbbed0000 7ffcbbed0000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\uxtheme.dll
7ffcbbed0000 7ffcbbed1000 +        5e000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\uxtheme.dll
7ffcbbed0000 7ffcbbf2f000 +        32000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\uxtheme.dll
7ffcbbed0000 7ffcbbf61000 +         3000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\uxtheme.dll
7ffcbbed0000 7ffcbbf64000 +         a000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\uxtheme.dll
7ffcbc3d0000 7ffcbc3d0000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\kernel.appcore.dll
7ffcbc3d0000 7ffcbc3d1000 +         4000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\kernel.appcore.dll
7ffcbc3d0000 7ffcbc3d5000 +         8000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\kernel.appcore.dll
7ffcbc3d0000 7ffcbc3dd000 +         1000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\kernel.appcore.dll
7ffcbc3d0000 7ffcbc3de000 +         4000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\kernel.appcore.dll
7ffcbc5d0000 7ffcbc5d0000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\windows.storage.dll
7ffcbc5d0000 7ffcbc5d1000 +       58e000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\windows.storage.dll
7ffcbc5d0000 7ffcbcb5f000 +       184000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\windows.storage.dll
7ffcbc5d0000 7ffcbcce3000 +         7000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\windows.storage.dll
7ffcbc5d0000 7ffcbccea000 +         2000 ty:--i st:c-- pr:-w-c \Device\HarddiskVolume2\Windows\System32\windows.storage.dll
7ffcbc5d0000 7ffcbccec000 +        78000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\windows.storage.dll
7ffcbd1b0000 7ffcbd1b0000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\ntmarta.dll
7ffcbd1b0000 7ffcbd1b1000 +        23000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\ntmarta.dll
7ffcbd1b0000 7ffcbd1d4000 +         8000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\ntmarta.dll
7ffcbd1b0000 7ffcbd1dc000 +         2000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\ntmarta.dll
7ffcbd1b0000 7ffcbd1de000 +         5000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\ntmarta.dll
7ffcbde40000 7ffcbde40000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\wldp.dll
7ffcbde40000 7ffcbde41000 +        1a000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\wldp.dll
7ffcbde40000 7ffcbde5b000 +         e000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\wldp.dll
7ffcbde40000 7ffcbde69000 +         1000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\wldp.dll
7ffcbde40000 7ffcbde6a000 +         6000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\wldp.dll
7ffcbe560000 7ffcbe560000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\msvcp_win.dll
7ffcbe560000 7ffcbe561000 +        54000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\msvcp_win.dll
7ffcbe560000 7ffcbe5b5000 +        3c000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\msvcp_win.dll
7ffcbe560000 7ffcbe5f1000 +         1000 ty:--i st:c-- pr:-w-c \Device\HarddiskVolume2\Windows\System32\msvcp_win.dll
7ffcbe560000 7ffcbe5f2000 +         3000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\msvcp_win.dll
7ffcbe560000 7ffcbe5f5000 +         8000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\msvcp_win.dll
7ffcbe600000 7ffcbe600000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\bcryptprimitives.dll
7ffcbe600000 7ffcbe601000 +        64000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\bcryptprimitives.dll
7ffcbe600000 7ffcbe665000 +        16000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\bcryptprimitives.dll
7ffcbe600000 7ffcbe67b000 +         1000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\bcryptprimitives.dll
7ffcbe600000 7ffcbe67c000 +         6000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\bcryptprimitives.dll
7ffcbe710000 7ffcbe710000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\KernelBase.dll
7ffcbe710000 7ffcbe711000 +       115000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\KernelBase.dll
7ffcbe710000 7ffcbe826000 +       17a000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\KernelBase.dll
7ffcbe710000 7ffcbe9a0000 +         4000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\KernelBase.dll
7ffcbe710000 7ffcbe9a4000 +         1000 ty:--i st:c-- pr:-w-c \Device\HarddiskVolume2\Windows\System32\KernelBase.dll
7ffcbe710000 7ffcbe9a5000 +        39000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\KernelBase.dll
7ffcbea50000 7ffcbea50000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\win32u.dll
7ffcbea50000 7ffcbea51000 +         b000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\win32u.dll
7ffcbea50000 7ffcbea5c000 +         f000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\win32u.dll
7ffcbea50000 7ffcbea6b000 +         1000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\win32u.dll
7ffcbea50000 7ffcbea6c000 +         6000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\win32u.dll
7ffcbebe0000 7ffcbebe0000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\ucrtbase.dll
7ffcbebe0000 7ffcbebe1000 +        b4000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\ucrtbase.dll
7ffcbebe0000 7ffcbec95000 +        3a000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\ucrtbase.dll
7ffcbebe0000 7ffcbeccf000 +         3000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\ucrtbase.dll
7ffcbebe0000 7ffcbecd2000 +         e000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\ucrtbase.dll
7ffcbece0000 7ffcbece0000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\gdi32full.dll
7ffcbece0000 7ffcbece1000 +        9c000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\gdi32full.dll
7ffcbece0000 7ffcbed7d000 +        4d000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\gdi32full.dll
7ffcbece0000 7ffcbedca000 +         4000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\gdi32full.dll
7ffcbece0000 7ffcbedce000 +         1000 ty:--i st:c-- pr:-w-c \Device\HarddiskVolume2\Windows\System32\gdi32full.dll
7ffcbece0000 7ffcbedcf000 +        1c000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\gdi32full.dll
7ffcbedf0000 7ffcbedf0000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\imm32.dll
7ffcbedf0000 7ffcbedf1000 +        1e000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\imm32.dll
7ffcbedf0000 7ffcbee0f000 +         7000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\imm32.dll
7ffcbedf0000 7ffcbee16000 +         1000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\imm32.dll
7ffcbedf0000 7ffcbee17000 +         9000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\imm32.dll
7ffcbef50000 7ffcbef50000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\msvcrt.dll
7ffcbef50000 7ffcbef51000 +        75000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\msvcrt.dll
7ffcbef50000 7ffcbefc6000 +        19000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\msvcrt.dll
7ffcbef50000 7ffcbefdf000 +         2000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\msvcrt.dll
7ffcbef50000 7ffcbefe1000 +         3000 ty:--i st:c-- pr:-w-c \Device\HarddiskVolume2\Windows\System32\msvcrt.dll
7ffcbef50000 7ffcbefe4000 +         2000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\msvcrt.dll
7ffcbef50000 7ffcbefe6000 +         1000 ty:--i st:c-- pr:-w-c \Device\HarddiskVolume2\Windows\System32\msvcrt.dll
7ffcbef50000 7ffcbefe7000 +         7000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\msvcrt.dll
7ffcbf000000 7ffcbf000000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\shell32.dll
7ffcbf000000 7ffcbf001000 +       589000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\shell32.dll
7ffcbf000000 7ffcbf58a000 +       149000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\shell32.dll
7ffcbf000000 7ffcbf6d3000 +         8000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\shell32.dll
7ffcbf000000 7ffcbf6db000 +         2000 ty:--i st:c-- pr:-w-c \Device\HarddiskVolume2\Windows\System32\shell32.dll
7ffcbf000000 7ffcbf6dd000 +        67000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\shell32.dll
7ffcbf840000 7ffcbf840000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\shlwapi.dll
7ffcbf840000 7ffcbf841000 +        2d000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\shlwapi.dll
7ffcbf840000 7ffcbf86e000 +        20000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\shlwapi.dll
7ffcbf840000 7ffcbf88e000 +         1000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\shlwapi.dll
7ffcbf840000 7ffcbf88f000 +         6000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\shlwapi.dll
7ffcbf8a0000 7ffcbf8a0000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\user32.dll
7ffcbf8a0000 7ffcbf8a1000 +        90000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\user32.dll
7ffcbf8a0000 7ffcbf931000 +        21000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\user32.dll
7ffcbf8a0000 7ffcbf952000 +         2000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\user32.dll
7ffcbf8a0000 7ffcbf954000 +        ec000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\user32.dll
7ffcbfa40000 7ffcbfa40000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\oleaut32.dll
7ffcbfa40000 7ffcbfa41000 +        95000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\oleaut32.dll
7ffcbfa40000 7ffcbfad6000 +        26000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\oleaut32.dll
7ffcbfa40000 7ffcbfafc000 +         3000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\oleaut32.dll
7ffcbfa40000 7ffcbfaff000 +         e000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\oleaut32.dll
7ffcbfb10000 7ffcbfb10000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\combase.dll
7ffcbfb10000 7ffcbfb11000 +       239000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\combase.dll
7ffcbfb10000 7ffcbfd4a000 +        c6000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\combase.dll
7ffcbfb10000 7ffcbfe10000 +         6000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\combase.dll
7ffcbfb10000 7ffcbfe16000 +        4e000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\combase.dll
7ffcbff30000 7ffcbff30000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\ws2_32.dll
7ffcbff30000 7ffcbff31000 +        44000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\ws2_32.dll
7ffcbff30000 7ffcbff75000 +         d000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\ws2_32.dll
7ffcbff30000 7ffcbff82000 +         1000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\ws2_32.dll
7ffcbff30000 7ffcbff83000 +        18000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\ws2_32.dll
7ffcc0080000 7ffcc0080000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\SHCore.dll
7ffcc0080000 7ffcc0081000 +        75000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\SHCore.dll
7ffcc0080000 7ffcc00f6000 +        26000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\SHCore.dll
7ffcc0080000 7ffcc011c000 +         2000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\SHCore.dll
7ffcc0080000 7ffcc011e000 +         f000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\SHCore.dll
7ffcc0190000 7ffcc0190000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\sechost.dll
7ffcc0190000 7ffcc0191000 +        65000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\sechost.dll
7ffcc0190000 7ffcc01f6000 +        28000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\sechost.dll
7ffcc0190000 7ffcc021e000 +         4000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\sechost.dll
7ffcc0190000 7ffcc0222000 +         a000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\sechost.dll
7ffcc03e0000 7ffcc03e0000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\kernel32.dll
7ffcc03e0000 7ffcc03e1000 +        7e000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\kernel32.dll
7ffcc03e0000 7ffcc045f000 +        33000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\kernel32.dll
7ffcc03e0000 7ffcc0492000 +         1000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\kernel32.dll
7ffcc03e0000 7ffcc0493000 +         1000 ty:--i st:c-- pr:-w-c \Device\HarddiskVolume2\Windows\System32\kernel32.dll
7ffcc03e0000 7ffcc0494000 +         9000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\kernel32.dll
7ffcc04a0000 7ffcc04a0000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\gdi32.dll
7ffcc04a0000 7ffcc04a1000 +         e000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\gdi32.dll
7ffcc04a0000 7ffcc04af000 +        14000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\gdi32.dll
7ffcc04a0000 7ffcc04c3000 +         1000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\gdi32.dll
7ffcc04a0000 7ffcc04c4000 +         6000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\gdi32.dll
7ffcc04d0000 7ffcc04d0000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\clbcatq.dll
7ffcc04d0000 7ffcc04d1000 +        6b000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\clbcatq.dll
7ffcc04d0000 7ffcc053c000 +        2d000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\clbcatq.dll
7ffcc04d0000 7ffcc0569000 +         4000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\clbcatq.dll
7ffcc04d0000 7ffcc056d000 +         1000 ty:--i st:c-- pr:-w-c \Device\HarddiskVolume2\Windows\System32\clbcatq.dll
7ffcc04d0000 7ffcc056e000 +         1000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\clbcatq.dll
7ffcc04d0000 7ffcc056f000 +         a000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\clbcatq.dll
7ffcc0580000 7ffcc0580000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\msctf.dll
7ffcc0580000 7ffcc0581000 +        d4000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\msctf.dll
7ffcc0580000 7ffcc0655000 +        2c000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\msctf.dll
7ffcc0580000 7ffcc0681000 +         3000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\msctf.dll
7ffcc0580000 7ffcc0684000 +        11000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\msctf.dll
7ffcc0b10000 7ffcc0b10000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\rpcrt4.dll
7ffcc0b10000 7ffcc0b11000 +        e1000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\rpcrt4.dll
7ffcc0b10000 7ffcc0bf2000 +        2c000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\rpcrt4.dll
7ffcc0b10000 7ffcc0c1e000 +         2000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\rpcrt4.dll
7ffcc0b10000 7ffcc0c20000 +        15000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\rpcrt4.dll
7ffcc0ca0000 7ffcc0ca0000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\advapi32.dll
7ffcc0ca0000 7ffcc0ca1000 +        68000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\advapi32.dll
7ffcc0ca0000 7ffcc0d09000 +        37000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\advapi32.dll
7ffcc0ca0000 7ffcc0d40000 +         1000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\advapi32.dll
7ffcc0ca0000 7ffcc0d41000 +         1000 ty:--i st:c-- pr:-w-c \Device\HarddiskVolume2\Windows\System32\advapi32.dll
7ffcc0ca0000 7ffcc0d42000 +         2000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\advapi32.dll
7ffcc0ca0000 7ffcc0d44000 +         1000 ty:--i st:c-- pr:-w-c \Device\HarddiskVolume2\Windows\System32\advapi32.dll
7ffcc0ca0000 7ffcc0d45000 +         9000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\advapi32.dll
7ffcc0d90000 7ffcc0d90000 +         1000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\ntdll.dll
7ffcc0d90000 7ffcc0d91000 +       11c000 ty:--i st:c-- pr:r-x- \Device\HarddiskVolume2\Windows\System32\ntdll.dll
7ffcc0d90000 7ffcc0ead000 +        49000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\ntdll.dll
7ffcc0d90000 7ffcc0ef6000 +         1000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\ntdll.dll
7ffcc0d90000 7ffcc0ef7000 +         2000 ty:--i st:c-- pr:-w-c \Device\HarddiskVolume2\Windows\System32\ntdll.dll
7ffcc0d90000 7ffcc0ef9000 +         9000 ty:--i st:c-- pr:rw-- \Device\HarddiskVolume2\Windows\System32\ntdll.dll
7ffcc0d90000 7ffcc0f02000 +        86000 ty:--i st:c-- pr:r--- \Device\HarddiskVolume2\Windows\System32\ntdll.dll
PS C:\Users\user>
```