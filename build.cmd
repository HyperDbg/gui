@REM (cmake -B Release -G "Ninja" -DCMAKE_BUILD_TYPE=Release . && cmake --build Release --config Release) 2>&1 | powershell -NoProfile -Command "$input | Tee-Object -FilePath build.Release.log"
(cmake -B Debug -G "Ninja" -DCMAKE_BUILD_TYPE=Debug . && cmake --build Debug --config Debug) 2>&1 | powershell -NoProfile -Command "$input | Tee-Object -FilePath build.Debug.log"
