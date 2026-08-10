package gjson

import (
	"fmt"
	"strings"
	"testing"
)

func TestIterBasic(t *testing.T) {
	json := `{"name":"Tom","age":37}`
	res := Parse(json)
	var keys []string
	var values []string
	for k, v := range res.Range() {
		keys = append(keys, k.String())
		values = append(values, v.String())
	}
	if len(keys) != 2 {
		t.Fatalf("expected 2 keys, got %d", len(keys))
	}
	if keys[0] != "name" || keys[1] != "age" {
		t.Fatalf("expected keys [name,age], got %v", keys)
	}
	if values[0] != "Tom" || values[1] != "37" {
		t.Fatalf("expected values [Tom,37], got %v", values)
	}
}

func TestIterArray(t *testing.T) {
	json := `[1,2,3,4]`
	res := Parse(json)
	var values []float64
	for v := range res.Array() {
		values = append(values, v.Num)
	}
	if len(values) != 4 {
		t.Fatalf("expected 4 values, got %d", len(values))
	}
	for i, v := range values {
		if v != float64(i+1) {
			t.Fatalf("expected %d, got %f", i+1, v)
		}
	}
}

func TestIterMap(t *testing.T) {
	json := `{"first":"Tom","last":"Smith"}`
	res := Parse(json)
	m := make(map[string]string)
	for k, v := range res.Map() {
		m[k] = v.String()
	}
	if m["first"] != "Tom" {
		t.Fatalf("expected Tom, got %s", m["first"])
	}
	if m["last"] != "Smith" {
		t.Fatalf("expected Smith, got %s", m["last"])
	}
}

func TestMany(t *testing.T) {
	json := `{"name":"Tom","age":37,"city":"NYC"}`
	paths := []string{"name", "age", "city"}
	var results []Result
	for _, r := range Many(json, paths...) {
		results = append(results, r)
	}
	if len(results) != 3 {
		t.Fatalf("expected 3 results, got %d", len(results))
	}
	if results[0].String() != "Tom" {
		t.Fatalf("expected Tom, got %s", results[0].String())
	}
	if results[1].Num != 37 {
		t.Fatalf("expected 37, got %f", results[1].Num)
	}
	if results[2].String() != "NYC" {
		t.Fatalf("expected NYC, got %s", results[2].String())
	}
}

func TestManyBytes(t *testing.T) {
	json := []byte(`{"name":"Tom","age":37}`)
	paths := []string{"name", "age"}
	var results []Result
	for _, r := range ManyBytes(json, paths...) {
		results = append(results, r)
	}
	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}
	if results[0].String() != "Tom" {
		t.Fatalf("expected Tom, got %s", results[0].String())
	}
}

func TestLines(t *testing.T) {
	json := `{"name":"Tom"}
{"name":"Jane"}
{"name":"Bob"}`
	var names []string
	for line := range Lines(json) {
		names = append(names, line.Get("name").String())
	}
	if len(names) != 3 {
		t.Fatalf("expected 3 names, got %d", len(names))
	}
	if names[0] != "Tom" || names[1] != "Jane" || names[2] != "Bob" {
		t.Fatalf("expected [Tom,Jane,Bob], got %v", names)
	}
}

func TestIterRecursiveDescent(t *testing.T) {
	json := `{"a":{"b":{"c":1}},"d":2}`
	parent := Parse(json)
	var values []float64
	for r := range iterRecursiveDescent(parent, "c") {
		values = append(values, r.Num)
	}
	if len(values) != 1 {
		t.Fatalf("expected 1 result, got %d", len(values))
	}
	if values[0] != 1 {
		t.Fatalf("expected 1, got %f", values[0])
	}
}

func TestPaths(t *testing.T) {
	json := `{"friends":[{"first":"Dale","last":"Murphy"},{"first":"Roger","last":"Craig"}]}`
	res := Get(json, "friends.#.first")
	var paths []string
	for p := range res.Paths(json) {
		paths = append(paths, p)
	}
	if len(paths) != 2 {
		t.Fatalf("expected 2 paths, got %d: %v", len(paths), paths)
	}
}

func TestIterPathsThree(t *testing.T) {
	json := `{
  "name": {"first": "Tom", "last": "Anderson"},
  "age":37,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
    {"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
    {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
  ]
}
`
	res := Get(json, "friends.#.first")
	var count int
	for v := range res.Array() {
		_ = v
		count++
	}
	if count != 3 {
		t.Fatalf("expected 3 array elements, got %d (IsArray=%v, Raw=%q)", count, res.IsArray(), res.Raw)
	}
	var paths []string
	for p := range res.Paths(json) {
		paths = append(paths, p)
	}
	if len(paths) != 3 {
		t.Fatalf("expected 3 paths, got %d: %v", len(paths), paths)
	}
	expected := []string{"friends.0.first", "friends.1.first", "friends.2.first"}
	for i, p := range paths {
		if p != expected[i] {
			t.Fatalf("path[%d]: expected %q, got %q", i, expected[i], p)
		}
	}
}

func TestIterEarlyBreak(t *testing.T) {
	json := `[1,2,3,4,5,6,7,8,9,10]`
	res := Parse(json)
	var values []float64
	for v := range res.Array() {
		values = append(values, v.Num)
		if v.Num == 3 {
			break
		}
	}
	if len(values) != 3 {
		t.Fatalf("expected 3 values (early break), got %d", len(values))
	}
}

func TestIterEmptyResult(t *testing.T) {
	res := Result{}
	count := 0
	for range res.Range() {
		count++
	}
	if count != 0 {
		t.Fatalf("expected 0 iterations for empty result, got %d", count)
	}
}

func TestIterNonJSONArray(t *testing.T) {
	json := `"hello"`
	res := Parse(json)
	var values []string
	for k, v := range res.Range() {
		_ = k
		values = append(values, v.String())
	}
	if len(values) != 1 || values[0] != "hello" {
		t.Fatalf("expected [hello], got %v", values)
	}
}

func TestIterManyEarlyBreak(t *testing.T) {
	json := `{"a":1,"b":2,"c":3,"d":4}`
	paths := []string{"a", "b", "c", "d"}
	var results []Result
	for _, r := range Many(json, paths...) {
		results = append(results, r)
		if r.Num == 2 {
			break
		}
	}
	if len(results) != 2 {
		t.Fatalf("expected 2 results (early break), got %d", len(results))
	}
}

func TestIterLinesEarlyBreak(t *testing.T) {
	json := `{"n":1}
{"n":2}
{"n":3}`
	var results []float64
	for line := range Lines(json) {
		results = append(results, line.Get("n").Num)
		if len(results) == 2 {
			break
		}
	}
	if len(results) != 2 {
		t.Fatalf("expected 2 results (early break), got %d", len(results))
	}
}

func BenchmarkIterIter(b *testing.B) {
	json := `{"name":"Tom","age":37,"city":"NYC","active":true,"score":95.5}`
	res := Parse(json)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for range res.Range() {
		}
	}
}

func BenchmarkIterIterArray(b *testing.B) {
	json := `[1,2,3,4,5,6,7,8,9,10]`
	res := Parse(json)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for range res.Array() {
		}
	}
}

func BenchmarkIterIterMap(b *testing.B) {
	json := `{"name":"Tom","age":37,"city":"NYC","active":true,"score":95.5}`
	res := Parse(json)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for range res.Map() {
		}
	}
}

func BenchmarkMany(b *testing.B) {
	json := `{"a":1,"b":2,"c":3,"d":4,"e":5}`
	paths := []string{"a", "b", "c", "d", "e"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for range Many(json, paths...) {
		}
	}
}

func BenchmarkLines(b *testing.B) {
	var json strings.Builder
	for i := range 100 {
		json.WriteString(fmt.Sprintf(`{"id":%d,"name":"user%d"}`, i, i))
		if i < 99 {
			json.WriteString("\n")
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for range Lines(json.String()) {
		}
	}
}

func BenchmarkIterLargeArray(b *testing.B) {
	var json strings.Builder
	json.WriteString("[")
	for i := range 1000 {
		if i > 0 {
			json.WriteString(",")
		}
		json.WriteString(fmt.Sprintf(`{"id":%d,"value":%d}`, i, i*10))
	}
	json.WriteString("]")
	res := Parse(json.String())
	b.Run("IterArray", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for range res.Array() {
			}
		}
	})
	b.Run("IterArray_Consume", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var count int
			for v := range res.Array() {
				_ = v
				count++
			}
			_ = count
		}
	})
}

func BenchmarkIterLargeObject(b *testing.B) {
	var json strings.Builder
	json.WriteString("{")
	for i := range 1000 {
		if i > 0 {
			json.WriteString(",")
		}
		json.WriteString(fmt.Sprintf(`"key%d":"value%d"`, i, i))
	}
	json.WriteString("}")
	res := Parse(json.String())
	b.Run("IterMap", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for range res.Map() {
			}
		}
	})
	b.Run("Iter", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for range res.Range() {
			}
		}
	})
}
