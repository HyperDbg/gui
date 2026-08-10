"""
移除 HyperDbg 平台层的 Linux/BSD 条件编译代码。

处理模式:
1. #if defined(__linux__) ... #endif  →  整个删除（包含 #include）
2. #if defined(_WIN32) || defined(_WIN64) [WIN] #elif defined(__linux__) [LINUX] #else ... #endif  →  只保留 [WIN]
3. #ifdef HYPERDBG_ENV_LINUX ... #endif  →  整个删除
"""

import os
import re
import sys

PROJECT_ROOT = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))

# 平台相关目录
TARGET_DIRS = [
    os.path.join(PROJECT_ROOT, "HyperDbg", "hyperdbg", "include", "platform"),
]


def remove_linux_include_guards(text):
    """移除 #if defined(__linux__) ... #include ... #endif 块"""
    pattern = re.compile(
        r"#if\s+defined\(__linux__\)\s*\n"
        r"\s*#\s*include.*\n"
        r"#endif\s*(//.*)?\n?",
        re.MULTILINE,
    )
    return pattern.sub("", text)


def remove_linux_ifdef_blocks(text):
    """移除 #ifdef HYPERDBG_ENV_LINUX ... #endif 块"""
    pattern = re.compile(
        r"#ifdef\s+HYPERDBG_ENV_LINUX\s*\n.*?#endif\s*\n?",
        re.DOTALL | re.MULTILINE,
    )
    return pattern.sub("", text)


def simplify_platform_conditionals(text):
    """
    处理三路条件编译:
    #if defined(_WIN32) || defined(_WIN64)
        [WIN CODE]
    #elif defined(__linux__)
        [LINUX CODE]
    #else
        #error ...
    #endif
    =>
        [WIN CODE]  (无 #if/#endif 包裹)
    """
    # 匹配完整的 #if _WIN32 ... #elif __linux__ ... [#else ...] #endif 块
    pattern = re.compile(
        r"#if\s+defined\(_WIN32\)\s*\|\|\s*defined\(_WIN64\)\s*\n"
        r"(.*?)"  # Windows 代码 (group 1)
        r"#elif\s+defined\(__linux__\)\s*\n"
        r".*?"  # Linux 代码（不捕获）
        r"(?:\#else\s*\n.*?)?"
        r"#endif\s*\n?",
        re.DOTALL,
    )

    def replace_with_windows_code(match):
        win_code = match.group(1)
        # 清理 Windows 代码块的缩进
        lines = win_code.split("\n")
        # 去掉每行开头的多余缩进（4空格）
        cleaned = []
        for line in lines:
            cleaned.append(line)
        return "\n".join(cleaned) + "\n"

    return pattern.sub(replace_with_windows_code, text)


def remove_single_linux_guards(text):
    """移除独立的 #if defined(__linux__) ... #endif（非条件编译块）"""
    pattern = re.compile(
        r"#if\s+defined\(__linux__\)\s*\n.*?#endif\s*\n?",
        re.DOTALL,
    )
    return pattern.sub("", text)


def process_file(filepath, dry_run=False):
    """处理单个文件"""
    with open(filepath, "r", encoding="utf-8", errors="replace") as f:
        original = f.read()

    text = original
    text = remove_linux_include_guards(text)
    text = remove_linux_ifdef_blocks(text)
    text = simplify_platform_conditionals(text)
    text = remove_single_linux_guards(text)

    if text == original:
        return True  # 无变化

    if dry_run:
        removed = len(original) - len(text)
        print(f"  [DRY-RUN] {os.path.relpath(filepath, PROJECT_ROOT)}: 移除 {removed} 字节")
        return True

    with open(filepath, "w", encoding="utf-8") as f:
        f.write(text)
    print(f"  ✓ {os.path.relpath(filepath, PROJECT_ROOT)}")
    return True


def main():
    dry_run = "--dry-run" in sys.argv

    print(f"移除 Linux/BSD 条件编译代码")
    print(f"模式: {'DRY-RUN' if dry_run else '执行'}")
    print()

    count = 0
    errors = []

    for target_dir in TARGET_DIRS:
        if not os.path.isdir(target_dir):
            print(f"  跳过（目录不存在）: {target_dir}")
            continue

        for root, dirs, files in os.walk(target_dir):
            for fname in files:
                ext = os.path.splitext(fname)[1].lower()
                if ext not in {".c", ".h", ".cpp"}:
                    continue

                filepath = os.path.join(root, fname)
                try:
                    if process_file(filepath, dry_run):
                        count += 1
                except Exception as e:
                    errors.append(f"{filepath}: {e}")

    print()
    print(f"扫描: {count} 个文件")
    if errors:
        print(f"错误: {len(errors)} 个")
        for e in errors[:5]:
            print(f"  - {e}")
    if dry_run:
        print("提示: 去掉 --dry-run 参数执行实际修改")


if __name__ == "__main__":
    main()