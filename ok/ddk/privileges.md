# Windows Privileges (特权)

> 头文件来源：`winnt.h` (Windows SDK)
> 本项目引用路径：`d:\todo\fork\fakeWindows\MiniSDK\inc\sdk\winnt.h#L4995-L5024`

Windows 特权是访问令牌（Access Token）中的权限项，决定了进程能执行哪些超出普通权限的系统操作。所有特权名宏均使用 `TEXT()` 包裹，UNICODE 编译时展开为 `L"..."`，否则展开为 `"..."`。

## 特权列表

| 宏 | 特权名 | 功能 | 危险等级 |
|---|--------|------|----------|
| `SE_CREATE_TOKEN_NAME` | SeCreateTokenPrivilege | 创建任意访问令牌（可伪造身份） | ⚠️⚠️⚠️ |
| `SE_ASSIGNPRIMARYTOKEN_NAME` | SeAssignPrimaryTokenPrivilege | 替换进程的主令牌 | ⚠️⚠️⚠️ |
| `SE_LOCK_MEMORY_NAME` | SeLockMemoryPrivilege | 锁定物理内存页面（防止被换页到磁盘） | ⚠️⚠️ |
| `SE_INCREASE_QUOTA_NAME` | SeIncreaseQuotaPrivilege | 提高进程的 CPU/内存配额限制 | ⚠️ |
| `SE_UNSOLICITED_INPUT_NAME` | SeUnsolicitedInputPrivilege | 向其他进程窗口发送未请求的输入（模拟键鼠） | ⚠️⚠️ |
| `SE_MACHINE_ACCOUNT_NAME` | SeMachineAccountPrivilege | 将计算机加入域 | ⚠️ |
| `SE_TCB_NAME` | SeTcbPrivilege | 操作系统本身权限（最高特权，可信计算基，等同 SYSTEM） | ⚠️⚠️⚠️ |
| `SE_SECURITY_NAME` | SeSecurityPrivilege | 管理安全日志、审计策略、对象 SACL | ⚠️⚠️ |
| `SE_TAKE_OWNERSHIP_NAME` | SeTakeOwnershipPrivilege | 取得任意对象的所有权（绕过 DACL） | ⚠️⚠️⚠️ |
| `SE_LOAD_DRIVER_NAME` | SeLoadDriverPrivilege | 加载/卸载内核驱动 | ⚠️⚠️⚠️ |
| `SE_SYSTEM_PROFILE_NAME` | SeSystemProfilePrivilege | 系统级性能分析（收集整个系统的 profiling 数据） | ⚠️ |
| `SE_SYSTEMTIME_NAME` | SeSystemtimePrivilege | 修改系统时间 | ⚠️ |
| `SE_PROF_SINGLE_PROCESS_NAME` | SeProfileSingleProcessPrivilege | 单进程性能分析 | ⚠️ |
| `SE_INC_BASE_PRIORITY_NAME` | SeIncreaseBasePriorityPrivilege | 提高进程调度优先级 | ⚠️ |
| `SE_CREATE_PAGEFILE_NAME` | SeCreatePagefilePrivilege | 创建/修改页面文件 | ⚠️ |
| `SE_CREATE_PERMANENT_NAME` | SeCreatePermanentPrivilege | 创建永久对象（目录对象，不会被自动删除） | ⚠️⚠️ |
| `SE_BACKUP_NAME` | SeBackupPrivilege | 备份文件（绕过读取权限检查） | ⚠️⚠️ |
| `SE_RESTORE_NAME` | SeRestorePrivilege | 还原文件（绕过写入权限检查） | ⚠️⚠️ |
| `SE_SHUTDOWN_NAME` | SeShutdownPrivilege | 关机/重启 | ⚠️ |
| `SE_DEBUG_NAME` | SeDebugPrivilege | 调试任意进程（读写进程内存、注入线程） | ⚠️⚠️⚠️ |
| `SE_AUDIT_NAME` | SeAuditPrivilege | 生成安全审计日志条目 | ⚠️ |
| `SE_SYSTEM_ENVIRONMENT_NAME` | SeSystemEnvironmentPrivilege | 修改 NVRAM/EFI 变量 | ⚠️⚠️ |
| `SE_CHANGE_NOTIFY_NAME` | SeChangeNotifyPrivilege | 接收文件/注册表变更通知（默认所有用户都有） | - |
| `SE_REMOTE_SHUTDOWN_NAME` | SeRemoteShutdownPrivilege | 远程关机 | ⚠️ |
| `SE_UNDOCK_NAME` | SeUndockPrivilege | 弹出笔记本电脑（热插拔） | - |
| `SE_SYNC_AGENT_NAME` | SeSyncAgentPrivilege | 目录服务同步代理 | ⚠️ |
| `SE_ENABLE_DELEGATION_NAME` | SeEnableDelegationPrivilege | 委派认证（Kerberos 委派） | ⚠️⚠️ |
| `SE_MANAGE_VOLUME_NAME` | SeManageVolumePrivilege | 执行卷管理（格式化、磁盘检查等） | ⚠️⚠️ |
| `SE_IMPERSONATE_NAME` | SeImpersonatePrivilege | 模拟客户端令牌（服务冒充用户身份） | ⚠️⚠️⚠️ |
| `SE_CREATE_GLOBAL_NAME` | SeCreateGlobalPrivilege | 在全局命名空间创建文件映射/符号链接 | ⚠️ |

## 本项目用到的特权

| 特权 | 用途 |
|------|------|
| **SeDebugPrivilege** | 调试任意进程，读取内核模块地址（核心！没它 `NtQuerySystemInformation(SystemModuleInformation)` 返回的 `ImageBase` 会被 Windows 清零） |
| **SeProfileSingleProcessPrivilege** | 单进程性能采样 |
| **SeSystemProfilePrivilege** | 系统级性能分析 |
| **SeIncreaseQuotaPrivilege** | 提高内存/CPU 配额 |
| **SeIncreaseBasePriorityPrivilege** | 提高调度优先级 |

## 常见攻击面

| 特权 | 攻击方式 |
|------|----------|
| **SeDebugPrivilege** | 读写任意进程内存，注入代码到高权限进程 |
| **SeLoadDriverPrivilege** | 加载带漏洞的合法驱动（BYOVD 攻击），获取 Ring 0 任意读写 |
| **SeImpersonatePrivilege** | 模拟高权限用户令牌，提权到 SYSTEM |
| **SeTakeOwnershipPrivilege** | 取得 SAM 数据库文件所有权，提取密码哈希 |
| **SeBackupPrivilege** | 绕过文件权限读取 SAM/SYSTEM 注册表 hive |
| **SeCreateTokenPrivilege** | 伪造包含任意特权的访问令牌 |

## 编码注意事项

`TEXT()` 宏在 UNICODE 编译时展开为 `L"..."`，否则展开为 `"..."`。调用 `LookupPrivilegeValueW` 等 W 后缀 API 时，必须传入宽字符串：

```c
// ❌ 错误：非 UNICODE 下 SE_DEBUG_NAME 展开为 char*，传给 LPCWSTR 类型不匹配
LookupPrivilegeValueW(NULL, SE_DEBUG_NAME, &luid);

// ✅ 正确：显式使用宽字符串字面量
LookupPrivilegeValueW(NULL, L"SeDebugPrivilege", &luid);

// ✅ 正确：使用显式 W 后缀宏
LookupPrivilegeValueW(NULL, SE_DEBUG_NAMEW, &luid);
```
