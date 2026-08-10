package scriptengine

import (
	"crypto/sha256"
	"encoding/hex"
	"sync"
)

// AstCache memoises binary AST payloads produced by Compile. The same Go
// callback source is frequently registered against multiple hooks (e.g. the
// same RtlAllocateHeap callback attached to several processes); caching
// avoids re-running go/parser + the subset validator + the wire-format
// serialiser on every registration, which dominates Compile latency.
//
// The cache is keyed by the SHA-256 hash of the source string. SHA-256 is
// chosen over a faster non-crypto hash (FNV-1a etc.) for two reasons:
//  1. collision resistance — a collision would silently deliver the wrong
//     AST to the kernel, which is a security-relevant correctness bug for
//     a debugger; the ~1µs/KB SHA-256 cost is dwarfed by the avoided
//     parse + validate pass on a hit;
//  2. determinism across processes — the same source compiles to the same
//     key in every Wrapper instance, so a future cross-Debugger cache could
//     reuse these digests.
//
// The cache is safe for concurrent use by any number of goroutines.
type AstCache struct {
	mu    sync.RWMutex
	store map[string][]byte
}

// NewAstCache returns an empty cache.
func NewAstCache() *AstCache {
	return &AstCache{store: make(map[string][]byte)}
}

// Get returns the cached AST for src, if present. The returned slice is a
// copy of the cached bytes, so callers may freely mutate it without affecting
// the cache. The bool result is false when src has never been Put.
func (c *AstCache) Get(src string) ([]byte, bool) {
	key := hashSource(src)
	c.mu.RLock()
	defer c.mu.RUnlock()
	ast, ok := c.store[key]
	if !ok {
		return nil, false
	}
	// Copy on read: the cache must remain immutable from the caller's
	// perspective, even though the cached bytes themselves never change.
	out := make([]byte, len(ast))
	copy(out, ast)
	return out, true
}

// Put stores a copy of ast keyed by src's SHA-256 hash. Subsequent mutations
// to the caller's slice do not affect the cached value. If src is already
// cached, the previous value is overwritten (callers should treat the AST
// for a given source as canonical — there is only one valid encoding).
func (c *AstCache) Put(src string, ast []byte) {
	key := hashSource(src)
	cp := make([]byte, len(ast))
	copy(cp, ast)
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = cp
}

// Clear removes all cached payloads. Use this to force re-compilation after
// the go-bridge encoder is upgraded (e.g. the binary protocol changes), or
// to reclaim memory when the cache has grown large during long debugging
// sessions.
func (c *AstCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store = make(map[string][]byte)
}

// Len returns the number of cached payloads. Primarily useful for diagnostics
// and tests; production callers should not branch on cache size.
func (c *AstCache) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.store)
}

// hashSource returns the lowercase hex SHA-256 digest of src. Hex encoding
// is used (rather than the raw 32-byte digest) so the key is a string, which
// is the natural map key type in Go and is comparable with the standard
// equality operator.
func hashSource(src string) string {
	sum := sha256.Sum256([]byte(src))
	return hex.EncodeToString(sum[:])
}
