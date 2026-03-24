fn main() -> Result<(), wdk_build::ConfigError> {
    println!("Starting build process...");
    #[cfg(feature = "secure-device")]
    println!("cargo:rustc-link-lib=Wdmsec");
    wdk_build::configure_wdk_binary_build()
}
