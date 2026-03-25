fn main() {
    let _ = wdk_build::configure_wdk_binary_build();
    println!("cargo:rustc-link-lib=netio");
}
