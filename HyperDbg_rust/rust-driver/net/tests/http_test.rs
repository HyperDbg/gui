#![allow(non_snake_case)]

use std::ptr;

#[repr(C)]
pub struct Request {
    pub Method: *const u8,
    pub URL: *const u8,
    pub Header: *const u8,
    pub Body: *const u8,
    pub BodyLength: usize,
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
        Header: if body_start > 4 { data.add(url_end) } else { ptr::null() },
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
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_read_request_get() {
        let request = b"GET /ping HTTP/1.1\r\nHost: localhost\r\n\r\n";
        let result = unsafe { ReadRequest(request.as_ptr(), request.len()) };
        assert!(result.is_some());
        let req = result.unwrap();
        assert_eq!(unsafe { req.MethodStr() }, "GET");
        assert_eq!(unsafe { req.Path() }, "/ping");
    }

    #[test]
    fn test_read_request_post() {
        let request = b"POST /api HTTP/1.1\r\nContent-Type: application/json\r\nContent-Length: 17\r\n\r\n{\"action\":\"ping\"}";
        let result = unsafe { ReadRequest(request.as_ptr(), request.len()) };
        assert!(result.is_some());
        let req = result.unwrap();
        assert_eq!(unsafe { req.MethodStr() }, "POST");
        assert_eq!(unsafe { req.Path() }, "/api");
        assert_eq!(req.BodyLength, 17);
    }

    #[test]
    fn test_read_request_with_body() {
        let request = b"POST /test HTTP/1.1\r\nContent-Type: application/json\r\nContent-Length: 45\r\n\r\n{\"action\":\"read\",\"target\":\"0x1000\",\"size\":64}";
        let result = unsafe { ReadRequest(request.as_ptr(), request.len()) };
        assert!(result.is_some());
        let req = result.unwrap();
        assert_eq!(unsafe { req.MethodStr() }, "POST");
        assert_eq!(unsafe { req.Path() }, "/test");
        assert_eq!(req.BodyLength, 45);
    }

    #[test]
    fn test_invalid_request_too_short() {
        let request = b"GET";
        let result = unsafe { ReadRequest(request.as_ptr(), request.len()) };
        assert!(result.is_none());
    }

    #[test]
    fn test_invalid_request_no_space() {
        let request = b"GETHTTP/1.1\r\n\r\n";
        let result = unsafe { ReadRequest(request.as_ptr(), request.len()) };
        assert!(result.is_none());
    }

    #[test]
    fn test_empty_body() {
        let request = b"GET /status HTTP/1.1\r\nHost: localhost\r\n\r\n";
        let result = unsafe { ReadRequest(request.as_ptr(), request.len()) };
        assert!(result.is_some());
        let req = result.unwrap();
        assert_eq!(unsafe { req.MethodStr() }, "GET");
        assert_eq!(unsafe { req.Path() }, "/status");
        assert_eq!(req.BodyLength, 0);
    }

    #[test]
    fn test_delete_request() {
        let request = b"DELETE /resource/123 HTTP/1.1\r\nHost: localhost\r\n\r\n";
        let result = unsafe { ReadRequest(request.as_ptr(), request.len()) };
        assert!(result.is_some());
        let req = result.unwrap();
        assert_eq!(unsafe { req.MethodStr() }, "DELETE");
        assert_eq!(unsafe { req.Path() }, "/resource/123");
    }

    #[test]
    fn test_put_request_with_body() {
        let request = b"PUT /update HTTP/1.1\r\nContent-Type: text/plain\r\nContent-Length: 11\r\n\r\nHello World";
        let result = unsafe { ReadRequest(request.as_ptr(), request.len()) };
        assert!(result.is_some());
        let req = result.unwrap();
        assert_eq!(unsafe { req.MethodStr() }, "PUT");
        assert_eq!(unsafe { req.Path() }, "/update");
        assert_eq!(req.BodyLength, 11);
    }

    #[test]
    fn test_request_with_query_params() {
        let request = b"GET /search?q=test&page=1 HTTP/1.1\r\nHost: localhost\r\n\r\n";
        let result = unsafe { ReadRequest(request.as_ptr(), request.len()) };
        assert!(result.is_some());
        let req = result.unwrap();
        assert_eq!(unsafe { req.MethodStr() }, "GET");
        assert_eq!(unsafe { req.Path() }, "/search");
    }

    #[test]
    fn test_request_with_multiple_headers() {
        let request = b"POST /api HTTP/1.1\r\nHost: localhost\r\nContent-Type: application/json\r\nAuthorization: Bearer token123\r\nContent-Length: 5\r\n\r\nhello";
        let result = unsafe { ReadRequest(request.as_ptr(), request.len()) };
        assert!(result.is_some());
        let req = result.unwrap();
        assert_eq!(unsafe { req.MethodStr() }, "POST");
        assert_eq!(unsafe { req.Path() }, "/api");
        assert_eq!(req.BodyLength, 5);
    }
}
