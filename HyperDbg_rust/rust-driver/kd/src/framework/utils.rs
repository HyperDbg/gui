extern crate alloc;

use alloc::vec::Vec;
use wdk_sys::UNICODE_STRING;

pub trait ToUnicodeString {
    fn to_unicode_string(&self) -> Option<UNICODE_STRING>;
}

impl ToUnicodeString for Vec<u16> {
    fn to_unicode_string(&self) -> Option<UNICODE_STRING> {
        create_unicode_string(self)
    }
}

pub fn create_unicode_string(s: &[u16]) -> Option<UNICODE_STRING> {
    let len = if !s.is_empty() {
        s.len()
    } else {
        return None;
    };

    let len_checked = if len > 0 && s[len - 1] == 0 {
        len - 1
    } else {
        len
    };

    Some(UNICODE_STRING {
        Length: (len_checked * 2) as u16,
        MaximumLength: (len * 2) as u16,
        Buffer: s.as_ptr() as *mut u16,
    })
}

pub trait ToU16Vec {
    fn to_u16_vec(&self) -> Vec<u16>;
}

impl ToU16Vec for &str {
    fn to_u16_vec(&self) -> Vec<u16> {
        let mut buf = Vec::with_capacity(self.len() + 1);

        for c in self.chars() {
            let mut c_buf = [0; 2];
            let encoded = c.encode_utf16(&mut c_buf);
            buf.extend_from_slice(encoded);
        }

        buf.push(0);
        buf
    }
}

#[macro_export]
macro_rules! utf16 {
    ($s:expr) => {
        $crate::utils::ToU16Vec::to_u16_vec(&$s)
    };
}
