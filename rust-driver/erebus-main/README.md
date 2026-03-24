# Erebus

**Erebus** is a Proof-of-Concept (PoC) Windows Kernel-Mode driver written in Rust. It utilizes the [windows-drivers-rs](https://github.com/microsoft/windows-drivers-rs) crate to interact with the Windows Kernel.

The project consists of two main components:

1.  **Kernel-Mode Driver (`km`)**: A driver that exposes an IOCTL interface to read and write arbitrary process memory using `MmCopyVirtualMemory`.
2.  **User-Mode Client (`um`)**: A CLI application that communicates with the driver to perform memory operations on a target process.

> **Disclaimer**: This project is for educational purposes only.

## Features

- **Read Memory**: Read data from a specific virtual address of a target process.
- **Write Memory**: Write data to a specific virtual address of a target process.
- **Safe Abstractions**: Uses Rust's safety features and `windows-drivers-rs` bindings where possible.
- **Process Lookup**: Resolves target processes by ID using `PsLookupProcessByProcessId`.

## Prerequisites

- [Rust](https://www.rust-lang.org/tools/install)
- [Windows Driver Kit (WDK)](https://learn.microsoft.com/en-us/windows-hardware/drivers/download-the-wdk)
- [cargo-make](https://github.com/sagiegurari/cargo-make) (`cargo install cargo-make`)

## Building

### Kernel Driver

The driver build process is managed by `cargo-make` and `wdk-build`.

```powershell
cd km
cargo make
```

This will build the driver, sign it with a self-signed test certificate, and generate the necessary `.sys`, `.inf`, and `.cat` files in `km/target/debug/erebus_package` (or similar).

To build a version of the driver that enforces security descriptors (restricting access to SYSTEM and Administrators only):

```powershell
cd km
cargo make build-secure
```


### User-Mode Client

```powershell
cd um
cargo build --release
```

## Usage

### 1. Enable Test Signing

Since the driver is signed with a self-generated test certificate, you must enable test signing mode on Windows.

```powershell
bcdedit /set testsigning on
# Restart your computer for changes to take effect
shutdown /r /t 0
```

### 2. Install & Start the Driver

You can use the `sc` command or a tool like [OSR Driver Loader](https://www.osr.com/osr-loader/) to load the driver.

```powershell
# Create the service (Run as Administrator)
sc create Erebus type= kernel binPath= "C:\full\path\to\erebus.sys"

# Start the service
sc start Erebus
```

### 3. Run the Client

Use the CLI client to interact with the driver.

```powershell
# Syntax: um.exe <process_name> <address_in_hex>
./target/release/um.exe notepad.exe 0x12345678
```

## License

MIT - See [LICENSE.md](LICENSE.md)

## Acknowledgements

- [windows-drivers-rs](https://github.com/microsoft/windows-drivers-rs)
- [0xflux](https://fluxsec.red/rust-windows-driver)
