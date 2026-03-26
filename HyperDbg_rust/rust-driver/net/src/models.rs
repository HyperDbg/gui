#![allow(non_snake_case)]

use alloc::string::String;
use serde::{Serialize, Deserialize};

#[derive(Serialize, Deserialize)]
pub struct Command {
    pub action: String,
    pub target: Option<String>,
    pub size: Option<i32>,
    pub data: Option<String>,
}

#[derive(Serialize, Deserialize)]
pub struct Response {
    pub success: bool,
    pub message: String,
}
