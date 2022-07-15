# Copyright 2019 John Papandriopoulos.  All rights reserved.
# Use of this source code is governed by a MIT-style
# license that can be found in the LICENSE file.

# This Makefile was designed to run on macOS, using a cross-compiler to
# build the Linux and Windows versions of the zydis library.
#   1. Xcode
#   2. brew install mingw-w64
#   3. brew install musl-cross --with-aarch64

CC_DARWIN_AMD64=clang -arch x86_64
CC_DARWIN_ARM64=clang -arch arm64
CC_LINUX_AMD64=x86_64-linux-musl-gcc
CC_LINUX_ARM64=aarch64-linux-musl-gcc
CC_WINDOWS_386=i686-w64-mingw32-gcc
CC_WINDOWS_AMD64=x86_64-w64-mingw32-gcc

CXX_DARWIN_AMD64=clang++ -arch x86_64
CXX_DARWIN_ARM64=clang++ -arch arm64
CXX_LINUX_AMD64=x86_64-linux-musl-g++
CXX_LINUX_ARM64=aarch64-linux-musl-g++
CXX_WINDOWS_386=i686-w64-mingw32-g++
CXX_WINDOWS_AMD64=x86_64-w64-mingw32-g++

LD_DARWIN_AMD64=ld -arch x86_64
LD_DARWIN_ARM64=ld -arch arm64
LD_LINUX_AMD64=x86_64-linux-musl-ld
LD_LINUX_ARM64=aarch64-linux-musl-ld
LD_WINDOWS_386=i686-w64-mingw32-ld
LD_WINDOWS_AMD64=x86_64-w64-mingw32-ld

.PHONY: all prepare clean install_headers darwin_amd64 darwin_arm64 linux_amd64 linux_arm64 windows_386 windows_amd64

all: darwin_amd64 darwin_arm64 linux_amd64 linux_arm64 windows_386 windows_amd64 install_headers 

# Prepare the build
prepare: clean
	git clone https://github.com/zyantific/zydis.git zydis
	cd zydis && \
	git submodule update --init

darwin_amd64: prepare
	mkdir -p out/darwin_amd64
	env \
		CC="$(CC_DARWIN_AMD64)" \
		CXX="$(CXX_DARWIN_AMD64)" \
	cmake \
		-DCMAKE_INSTALL_PREFIX=./out/install \
		-S zydis -B out/darwin_amd64 && \
	cmake --build out/darwin_amd64
	$(LD_DARWIN_AMD64) -r -o ../libzydis_darwin_amd64.syso -force_load out/darwin_amd64/libZydis.a

darwin_arm64: prepare
	mkdir -p out/darwin_arm64
	env \
		CC="$(CC_DARWIN_RM64)" \
		CXX="$(CXX_DARWIN_RM64)" \
	cmake -S zydis -B out/darwin_arm64 && \
	cmake --build out/darwin_arm64
	$(LD_DARWIN_ARM64) -r -o ../libzydis_darwin_arm64.syso -force_load out/darwin_arm64/libZydis.a

linux_amd64: prepare
	mkdir -p out/linux_amd64
	env \
		CC=$(CC_LINUX_AMD64) \
		CXX=$(CXX_LINUX_AMD64) \
	cmake \
		-DCMAKE_SYSTEM_NAME=linux \
		-DZYDIS_BUILD_EXAMPLES=OFF \
		-DCMAKE_POSITION_INDEPENDENT_CODE=ON \
		-S zydis -B out/linux_amd64 && \
	cmake --build out/linux_amd64
	$(LD_LINUX_AMD64) -r -o ../libzydis_linux_amd64.syso --whole-archive out/linux_amd64/libZydis.a

linux_arm64: prepare
	mkdir -p out/linux_arm64
	env \
		CC=$(CC_LINUX_ARM64) \
		CXX=$(CXX_LINUX_ARM64) \
	cmake \
		-DCMAKE_SYSTEM_NAME=linux \
		-DZYDIS_BUILD_EXAMPLES=OFF \
		-DCMAKE_POSITION_INDEPENDENT_CODE=ON \
		-S zydis -B out/linux_arm64 && \
	cmake --build out/linux_arm64
	$(LD_LINUX_ARM64) -r -o ../libzydis_linux_arm64.syso --whole-archive out/linux_arm64/libZydis.a

windows_amd64: prepare
	mkdir -p out/windows_amd64
	env \
		CC=$(CC_WINDOWS_AMD64) \
		CXX=$(CXX_WINDOWS_AMD64) \
	cmake \
		-DCMAKE_SYSTEM_NAME=windows \
		-DZYDIS_BUILD_EXAMPLES=OFF \
		-DCMAKE_POSITION_INDEPENDENT_CODE=ON \
		-S zydis -B out/windows_amd64 && \
	cmake --build out/windows_amd64
	$(LD_WINDOWS_AMD64) -r -o ../libzydis_windows_amd64.syso --whole-archive out/windows_amd64/libZydis.a

windows_386: prepare
	mkdir -p out/windows_386
	env \
		CC=$(CC_WINDOWS_386) \
		CXX=$(CXX_WINDOWS_386) \
	cmake \
		-DCMAKE_SYSTEM_NAME=windows \
		-DZYDIS_BUILD_EXAMPLES=OFF \
		-DCMAKE_POSITION_INDEPENDENT_CODE=ON \
		-S zydis -B out/windows_386 && \
	cmake --build out/windows_386
	$(LD_WINDOWS_386) -r -o ../libzydis_windows_386.syso --whole-archive out/windows_386/libZydis.a

# Clean the build
clean:
	rm -rf zydis out

install_headers: darwin_amd64
	# Copy include folder
	mkdir -p include
	mkdir -p out/install
	cmake --install out/darwin_amd64
	cmake --install out/darwin_amd64/zycore
	cp -Rvp out/install/include/* include/
