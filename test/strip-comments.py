"""
HyperDbg 精简工具 — 移除代码注释和多余空行

用法:
    python scripts/strip-comments.py [--dry-run]

选项:
    --dry-run  只显示将删除哪些内容，不实际写文件

处理规则:
    1. .c/.h/.cpp/.asm 文件:
       - 移除所有 // 行注释
       - 移除所有 /* */ 块注释
       - 连续空行合并为最多一个空行
    2. CMakeLists.txt: 跳过，保持原样
    3. 二进制文件/第三方依赖: 跳过
"""

import os
import re
import sys

# 项目根目录
PROJECT_ROOT = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))

# 需要跳过的目录（第三方依赖）
SKIP_DIRS = {
    "dependencies",
    "libraries",
    ".git",
    ".trae",
    "out",
    "build",
    "__pycache__",
}

# 只处理这些后缀
SOURCE_EXTENSIONS = {".c", ".h", ".cpp", ".asm", ".bat", ".ps1"}

# 跳过 CMakeLists.txt（保留注释）
KEEP_COMMENTS_FILES = {"CMakeLists.txt"}


def should_skip_dir(path_parts):
    """检查路径是否包含需要跳过的目录"""
    return any(skip in path_parts for skip in SKIP_DIRS)


def strip_line_comments(text):
    """移除 // 行注释（在字符串字面量中的除外）"""
    result = []
    for line in text.splitlines():
        # 简单处理：找到第一个不在字符串中的 //
        in_string = False
        string_char = None
        i = 0
        while i < len(line):
            if not in_string:
                if line[i] in ('"', "'"):
                    in_string = True
                    string_char = line[i]
                elif line[i : i + 2] == "//":
                    # 注释开始，截断
                    line = line[:i]
                    break
            else:
                if line[i] == "\\":
                    i += 1  # skip escaped char
                elif line[i] == string_char:
                    in_string = False
            i += 1

        result.append(line)
    return "\n".join(result)


def strip_block_comments(text):
    """移除 /* */ 块注释"""
    # 使用非贪婪匹配
    return re.sub(r"/\*.*?\*/", "", text, flags=re.DOTALL)


def collapse_blank_lines(text):
    """将连续空行合并为最多一个空行，并移除只包含空白字符的行"""
    # 1. 移除只包含空白字符（空格/制表符）的行
    text = re.sub(r"[ \t]*\n", "\n", text)
    # 2. 将 3 个以上连续换行替换为 2 个（保留一个空行）
    text = re.sub(r"\n{3,}", "\n\n", text)
    # 3. 去除文件开头和结尾的空白行
    text = text.strip()
    if text:
        text = text + "\n"
    return text


def strip_frontmatter_comment(text):
    """
    移除文件开头的许可证/文件信息注释块（/** ... */ 或 /* ... */）
    保留 #pragma once 和 #include 之前的空行
    """
    # 匹配文件开头的注释块
    pattern = re.compile(r"^(/\*.*?\*/\s*)+", re.DOTALL)
    text = pattern.sub("", text)
    return text


def process_source_file(filepath, dry_run=False):
    """处理单个源文件"""
    with open(filepath, "r", encoding="utf-8", errors="replace") as f:
        original = f.read()

    text = original

    # 只处理源文件
    ext = os.path.splitext(filepath)[1].lower()
    basename = os.path.basename(filepath)

    if basename in KEEP_COMMENTS_FILES:
        # 跳过 CMakeLists.txt
        return True

    if ext in SOURCE_EXTENSIONS:
        # 移除块注释
        text = strip_block_comments(text)
        # 移除行注释
        text = strip_line_comments(text)
        # 合并空行
        text = collapse_blank_lines(text)

    if text == original:
        return True

    if dry_run:
        removed = len(original) - len(text)
        print(f"  [DRY-RUN] {os.path.relpath(filepath, PROJECT_ROOT)}: 移除 {removed} 字节")
        return True

    with open(filepath, "w", encoding="utf-8") as f:
        f.write(text)
    return True


def main():
    dry_run = "--dry-run" in sys.argv

    print(f"HyperDbg 注释移除工具")
    print(f"项目根目录: {PROJECT_ROOT}")
    print(f"模式: {'DRY-RUN（只预览不修改）' if dry_run else '执行'}")
    print()

    count = 0
    errors = []

    for root, dirs, files in os.walk(PROJECT_ROOT):
        # 跳过不需要的目录
        rel_root = os.path.relpath(root, PROJECT_ROOT)
        path_parts = rel_root.split(os.sep)
        if should_skip_dir(path_parts):
            continue

        # 只处理 HyperDbg 目录下的源文件
        if "HyperDbg" not in path_parts and root != PROJECT_ROOT:
            # 检查上层路径
            if not any("HyperDbg" in p for p in path_parts):
                continue

        for fname in files:
            # 跳过二进制文件
            if fname in KEEP_COMMENTS_FILES:
                ext = ".txt"
            else:
                ext = os.path.splitext(fname)[1].lower()

            if ext not in SOURCE_EXTENSIONS and fname not in KEEP_COMMENTS_FILES:
                continue

            filepath = os.path.join(root, fname)
            try:
                if process_source_file(filepath, dry_run):
                    count += 1
                else:
                    errors.append(filepath)
            except Exception as e:
                errors.append(f"{filepath}: {e}")

    print()
    print(f"处理完成: {count} 个文件")
    if errors:
        print(f"错误: {len(errors)} 个")
        for e in errors[:10]:
            print(f"  - {e}")
    if dry_run:
        print("提示: 去掉 --dry-run 参数执行实际修改")


if __name__ == "__main__":
    main()