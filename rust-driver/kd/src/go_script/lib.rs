//! Go script parser and analyzer for HyperDbg
//!
//! This module provides Go code parsing capabilities using goscript-parser,
//! which is a port of the Go standard library parser.

pub mod analyzer;
pub mod executor;
pub mod extractor;
pub mod generator;
pub mod mapper;
pub mod parser;
pub mod visitor;

// Re-export goscript parser
pub use goscript_parser::{
    parse_file, AstObjects, FileSet, ErrorList, ast, visitor,
};