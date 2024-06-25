@echo off

:: 设置控制台为UTF-8编码
chcp 65001 >nul

:: 定义版本变量
set "VERSION=v0.0.1"

:: 定义脚本操作
setlocal

echo 正在运行 go mod tidy ...
go mod tidy

echo 正在添加更改到 git ...
git add .

echo 正在提交更改到 git ...
git commit -m "%VERSION%"

echo 正在创建标签 %VERSION% ...
git tag %VERSION%

echo 正在推送到远程仓库（包括标签）...
git push origin master --tags

echo 发布流程完成。

endlocal
pause
