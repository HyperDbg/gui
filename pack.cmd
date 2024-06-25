set RELEASE_ZIP_FILE_NAME ="hyperdbgui"
powershell -c "Compress-Archive -Path * -DestinationPath $env:RELEASE_ZIP_FILE_NAME-src.zip -CompressionLevel Optimal -force '.git/**','*.dll','*.sys','*.exe'"
powershell -c "Compress-Archive -Path *.dll, *.sys, *.exe -DestinationPath binaries-bin.zip"