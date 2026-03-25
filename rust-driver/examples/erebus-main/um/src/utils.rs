use sysinfo::System;

pub(crate) fn get_process_id(process_name: &str) -> Result<u32, String> {
    System::new_all()
        .processes_by_name(process_name.as_ref())
        .next()
        .map(|process| process.pid().as_u32())
        .ok_or_else(|| format!("No process found with name '{process_name}'"))
}

pub(crate) fn vec_to_usize(vec: Vec<u8>) -> Result<usize, String> {
    // Ensure the value doesn't exceed the size of `usize`
    if vec.len() > size_of::<usize>() {
        return Err("Value is too large to fit in usize".to_string());
    }

    // Perform the conversion
    let result = vec
        .into_iter()
        .rev() // Reverse to ensure little-endian order
        .enumerate()
        .map(|(i, byte)| (byte as usize) << (i * 8))
        .sum();

    Ok(result)
}

pub(crate) fn str_to_address(str: &str) -> Result<usize, String> {
    let bytes = hex::decode(str.trim_start_matches("0x"))
        .map_err(|err| format!("Invalid hexadecimal string: {err}"))?;

    vec_to_usize(bytes).map_err(|err| format!("Address out of bounds: {err}"))
}
