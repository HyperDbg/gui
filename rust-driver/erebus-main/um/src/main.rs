mod driver;
mod utils;

use crate::{
    driver::Driver,
    utils::{get_process_id, str_to_address},
};
use shared::constants::DRIVER_UM_NAME;
use std::path::Path;

fn main() {
    if let Err(err) = run() {
        eprintln!("{err}");
        std::process::exit(1);
    }
}

fn parse_args(args: &[String]) -> Result<(&str, &str), String> {
    if args.len() < 3 {
        let filename = Path::new(&args[0])
            .file_name()
            .and_then(|name| name.to_str())
            .unwrap_or("unknown");

        return Err(format!(
            "Usage: {filename} <process_name> <address>\n\
            Example: {filename} test-binary.exe 0x12345678"
        ));
    }

    Ok((&args[1], &args[2]))
}

fn run() -> Result<(), String> {
    let args: Vec<String> = std::env::args().collect();
    let (process_name, address_str) = parse_args(&args)?;

    // Retrieve the process ID using the process name.
    let process_id = get_process_id(process_name)
        .map_err(|err| format!("Failed to find process id for {process_name}! Error: {err}"))?;

    // Open the driver.
    let driver = Driver::new(DRIVER_UM_NAME)
        .map_err(|_| format!("Failed to open driver device {DRIVER_UM_NAME}!"))?;

    // Parse the address string argument into a `usize`, then cast it to a mutable pointer
    // of the required type for reading or writing.
    let address: *mut i32 = str_to_address(address_str)? as _;

    // Read a value from the process memory at the specified address.
    let read_value = driver
        .read_process_memory(process_id, address)
        .map_err(|err| format!("Could not read process memory: {err}"))?;
    println!("Value at {address:p} = {read_value:?}");

    // Write a value to the process memory at the specified address.
    let write_value = 1337;
    driver
        .write_process_memory(process_id, address, &write_value)
        .map_err(|err| format!("Could not write process memory: {err}"))?;
    println!("Finished writing, read again.");

    // Read the value from the process memory at the specified address again after writing,
    // to verify if it has been updated.
    let read_values = driver
        .read_process_memory(process_id, address)
        .map_err(|err| format!("Could not read process memory: {err}"))?;
    println!("Value at {address:p} = {read_values:?}");

    Ok(())
}
