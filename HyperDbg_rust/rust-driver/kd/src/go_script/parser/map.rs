#[cfg(feature = "btree_map")]
pub type Map<K, V> = alloc::collections::BTreeMap<K, V>;
#[cfg(not(feature = "btree_map"))]
pub type Map<K, V> = alloc::collections::BTreeMap<K, V>;

#[cfg(feature = "btree_map")]
pub type MapIter<'a, K, V> = alloc::collections::btree_map::Iter<'a, K, V>;
#[cfg(not(feature = "btree_map"))]
pub type MapIter<'a, K, V> = alloc::collections::btree_map::Iter<'a, K, V>;
