#![allow(unexpected_cfgs)]

pub type Map<K, V> = alloc::collections::BTreeMap<K, V>;
pub type MapIter<'a, K, V> = alloc::collections::btree_map::Iter<'a, K, V>;
