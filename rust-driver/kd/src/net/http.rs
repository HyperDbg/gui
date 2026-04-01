#![allow(non_snake_case)]

use core::ptr;
use alloc::string::String;
use alloc::format;

use serde::Serialize;

#[repr(C)]
pub struct Request {
    pub Method: *const u8,
    pub URL: *const u8,
    pub Header: *const u8,
    pub Body: *const u8,
    pub BodyLength: usize,
}

#[repr(C)]
#[derive(Clone, Copy)]
pub struct ResponseWriter {
    pub buffer: *mut u8,
    pub length: usize,
    pub capacity: usize,
}

impl ResponseWriter {
    pub unsafe fn Write(&mut self, data: *const u8, len: usize) -> isize {
        if self.length + len > self.capacity { return -1; }
        ptr::copy_nonoverlapping(data, self.buffer.add(self.length), len);
        self.length += len;
        len as isize
    }

    pub unsafe fn WriteHeader(&mut self, status_code: u16, content_type: &str) {
        let header = format!("HTTP/1.1 {}\r\nContent-Type: {}\r\nContent-Length: {}\r\n\r\n", 
            status_code, content_type, self.length);
        let header_bytes = header.as_bytes();
        let header_len = header_bytes.len();
        
        if header_len > self.capacity { return; }
        ptr::copy_nonoverlapping(header_bytes.as_ptr(), self.buffer, header_len);
        self.length = header_len;
    }

    pub unsafe fn WriteJSON<T: Serialize>(&mut self, obj: &T) -> isize {
        // 必须先发送 HTTP 响应头，否则客户端等待 "HTTP/1.1 200 OK\r\n..." 超时
        let bytes = super::json::Marshal(obj).unwrap();
        self.WriteJSONBytes(&bytes)
    }

    pub unsafe fn WriteJSONBytes(&mut self, bytes: &[u8]) -> isize {
        // 直接写入已序列化的 JSON 字节
        let header = format!(
            "HTTP/1.1 200 OK\r\nContent-Type: application/json\r\nConnection: close\r\nContent-Length: {}\r\n\r\n",
            bytes.len()
        );
        let header_bytes = header.as_bytes();
        self.Write(header_bytes.as_ptr(), header_bytes.len());
        self.Write(bytes.as_ptr(), bytes.len())
    }
}

pub unsafe fn ReadRequest(data: *const u8, len: usize) -> Option<Request> {
    if len < 16 { return None; }
    
    let bytes = core::slice::from_raw_parts(data, len);
    
    let mut method_end = 0;
    while method_end < len && bytes[method_end] != b' ' {
        method_end += 1;
    }
    if method_end == 0 || method_end >= len { return None; }
    
    let mut url_start = method_end + 1;
    while url_start < len && bytes[url_start] == b' ' {
        url_start += 1;
    }
    if url_start >= len { return None; }
    
    let mut url_end = url_start;
    while url_end < len && bytes[url_end] != b' ' && bytes[url_end] != b'\r' && bytes[url_end] != b'\n' {
        url_end += 1;
    }
    if url_end == url_start { return None; }
    
    let mut header_start = url_end;
    while header_start < len && bytes[header_start] != b'\n' {
        header_start += 1;
    }
    if header_start >= len { return None; }
    header_start += 1;
    
    let mut body_start = 0;
    for i in 0..len.saturating_sub(3) {
        if bytes[i] == b'\r' && bytes[i + 1] == b'\n' && bytes[i + 2] == b'\r' && bytes[i + 3] == b'\n' {
            body_start = i + 4;
            break;
        }
    }
    
    if body_start == 0 {
        body_start = len;
    }
    
    let body_len = if body_start < len { len - body_start } else { 0 };
    
    Some(Request {
        Method: data.add(0),
        URL: data.add(url_start),
        Header: if header_start < body_start.saturating_sub(4) { data.add(header_start) } else { ptr::null() },
        Body: if body_len > 0 { data.add(body_start) } else { ptr::null() },
        BodyLength: body_len,
    })
}

impl Request {
    pub unsafe fn Path(&self) -> &str {
        let url_bytes = core::slice::from_raw_parts(self.URL, 256);
        let mut len = 0;
        while len < url_bytes.len() && url_bytes[len] != 0 && url_bytes[len] != b' ' && url_bytes[len] != b'?' {
            len += 1;
        }
        core::str::from_utf8_unchecked(&url_bytes[..len])
    }

    pub unsafe fn MethodStr(&self) -> &str {
        let method_bytes = core::slice::from_raw_parts(self.Method, 16);
        let mut len = 0;
        while len < method_bytes.len() && method_bytes[len] != 0 && method_bytes[len] != b' ' {
            len += 1;
        }
        core::str::from_utf8_unchecked(&method_bytes[..len])
    }

    pub unsafe fn GetHeader(&self, name: &str) -> Option<String> {
        if self.Header.is_null() { return None; }
        
        let header_bytes = core::slice::from_raw_parts(self.Header, 1024);
        let header_str = String::from_utf8_lossy(header_bytes);
        
        let search = format!("{}:", name);
        if let Some(start) = header_str.find(&search) {
            let start = start + search.len();
            let mut end = start;
            while end < header_str.len() && header_str.as_bytes()[end] != b'\r' && header_str.as_bytes()[end] != b'\n' {
                end += 1;
            }
            let value = header_str[start..end].trim();
            Some(String::from(value))
        } else {
            None
        }
    }

    pub unsafe fn ContentType(&self) -> Option<String> {
        self.GetHeader("Content-Type")
    }

    pub unsafe fn ContentLength(&self) -> Option<usize> {
        if let Some(len_str) = self.GetHeader("Content-Length") {
            len_str.parse::<usize>().ok()
        } else {
            None
        }
    }
}

fn format(s: &str, args: &[&str]) -> String {
    let mut result = String::from(s);
    for arg in args {
        result = result.replacen("{}", arg, 1);
    }
    result
}
