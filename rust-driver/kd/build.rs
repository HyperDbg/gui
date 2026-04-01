fn main() {
    let _ = wdk_build::configure_wdk_binary_build();

    println!("cargo:rustc-link-lib=netio");
    println!("cargo:rustc-link-lib=wdm");

    if let Some(sdk_lib_path) = find_sdk_lib_path() {
        println!("cargo:rustc-link-search=native={}", sdk_lib_path);
        println!("cargo:rustc-link-lib=uuid");
        println!("cargo:rerun-if-env-changed=WindowsSdkDir");
    }
}

fn find_sdk_lib_path() -> Option<String> {
    if let Ok(sdk_dir) = std::env::var("WindowsSdkDir") {
        let lib_path = format!("{}Lib\\{}\\um\\x64", sdk_dir, find_sdk_version(&sdk_dir));
        if std::path::Path::new(&lib_path).exists() {
            return Some(lib_path);
        }
    }

    let common_paths = [
        "C:\\Program Files (x86)\\Windows Kits\\10\\Lib",
        "E:\\Program Files\\Windows Kits\\10\\Lib",
    ];

    for base_path in &common_paths {
        if let Ok(entries) = std::fs::read_dir(base_path) {
            for entry in entries.flatten() {
                let lib_path = format!("{}\\um\\x64", entry.path().display());
                if std::path::Path::new(&lib_path).join("uuid.lib").exists() {
                    return Some(lib_path);
                }
            }
        }
    }

    None
}

fn find_sdk_version(sdk_dir: &str) -> String {
    let lib_dir = format!("{}Lib", sdk_dir);
    if let Ok(entries) = std::fs::read_dir(&lib_dir) {
        let mut versions: Vec<String> = entries
            .flatten()
            .filter_map(|e| {
                let name = e.file_name();
                let name_str = name.to_string_lossy();
                if name_str.starts_with("10.0") {
                    Some(name_str.to_string())
                } else {
                    None
                }
            })
            .collect();
        versions.sort();
        versions.pop().unwrap_or_else(|| "10.0.28000.0".to_string())
    } else {
        "10.0.28000.0".to_string()
    }
}
