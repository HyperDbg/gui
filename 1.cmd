set RELEASE_ZIP_FILE_NAME=hyperdbgui
powershell -Command "Get-ChildItem -Recurse | Where-Object { $_.FullName -notmatch '\.git' -and $_.Extension -notmatch '\.dll|\.sys|\.exe' } | Compress-Archive -DestinationPath \"$env:RELEASE_ZIP_FILE_NAME-src.zip\" -CompressionLevel Optimal"
powershell -Command "Get-ChildItem -Recurse -Include '*.dll', '*.sys', '*.exe' | Compress-Archive -DestinationPath binaries-bin.zip"
