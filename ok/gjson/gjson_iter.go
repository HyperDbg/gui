package gjson

import "iter"

func (t Result) Range() iter.Seq2[Result, Result] {
	return func(yield func(Result, Result) bool) {
		if !t.Exists() {
			return
		}
		if t.Type != JSON {
			yield(Result{}, t)
			return
		}
		json := t.Raw
		var obj bool
		var i int
		var key, value Result
		for ; i < len(json); i++ {
			if json[i] == '{' {
				i++
				key.Type = String
				obj = true
				break
			} else if json[i] == '[' {
				i++
				key.Type = Number
				key.Num = -1
				break
			}
			if json[i] > ' ' {
				return
			}
		}
		var str string
		var vesc bool
		var ok bool
		var idx int
		for ; i < len(json); i++ {
			if obj {
				if json[i] != '"' {
					continue
				}
				s := i
				i, str, vesc, ok = parseString(json, i+1)
				if !ok {
					return
				}
				if vesc {
					key.Str = unescape(str[1 : len(str)-1])
				} else {
					key.Str = str[1 : len(str)-1]
				}
				key.Raw = str
				key.Index = s + t.Index
			} else {
				key.Num += 1
			}
			for ; i < len(json); i++ {
				if json[i] <= ' ' || json[i] == ',' || json[i] == ':' {
					continue
				}
				break
			}
			s := i
			i, value, ok = parseAny(json, i, true)
			if !ok {
				return
			}
			if t.Indexes != nil {
				if idx < len(t.Indexes) {
					value.Index = t.Indexes[idx]
				}
			} else {
				value.Index = s + t.Index
			}
			if !yield(key, value) {
				return
			}
			idx++
		}
	}
}

func (t Result) Range2() iter.Seq2[int, Result] {
	return func(yield func(int, Result) bool) {
		idx := 0
		for _, v := range t.Range() {
			if !yield(idx, v) {
				return
			}
			idx++
		}
	}
}

func (t Result) Array() iter.Seq[Result] {
	return func(yield func(Result) bool) {
		if t.Type == Null {
			return
		}
		if !t.IsArray() {
			yield(t)
			return
		}
		for _, v := range t.Range() {
			if !yield(v) {
				return
			}
		}
	}
}

func (t Result) Array2() iter.Seq2[int, Result] {
	return func(yield func(int, Result) bool) {
		if t.Type == Null {
			return
		}
		if !t.IsArray() {
			yield(0, t)
			return
		}
		for k, v := range t.Range2() {
			if !yield(k, v) {
				return
			}
		}
	}
}

func (t Result) Map() iter.Seq2[string, Result] {
	return func(yield func(string, Result) bool) {
		for k, v := range t.Range() {
			if k.Type == String {
				if !yield(k.Str, v) {
					return
				}
			}
		}
	}
}

func Many(json string, path ...string) iter.Seq2[string, Result] {
	return func(yield func(string, Result) bool) {
		for _, p := range path {
			if !yield(p, Get(json, p)) {
				return
			}
		}
	}
}

func Many2(json string, path ...string) iter.Seq2[int, Result] {
	return func(yield func(int, Result) bool) {
		for i, p := range path {
			if !yield(i, Get(json, p)) {
				return
			}
		}
	}
}

func ManyBytes(json []byte, path ...string) iter.Seq2[string, Result] {
	return func(yield func(string, Result) bool) {
		for _, p := range path {
			if !yield(p, GetBytes(json, p)) {
				return
			}
		}
	}
}

func ManyBytes2(json []byte, path ...string) iter.Seq2[int, Result] {
	return func(yield func(int, Result) bool) {
		for i, p := range path {
			if !yield(i, GetBytes(json, p)) {
				return
			}
		}
	}
}

func Lines(json string) iter.Seq[Result] {
	return func(yield func(Result) bool) {
		var res Result
		var i int
		for {
			i, res, _ = parseAny(json, i, true)
			if !res.Exists() {
				break
			}
			if !yield(res) {
				return
			}
		}
	}
}

func Lines2(json string) iter.Seq2[int, Result] {
	return func(yield func(int, Result) bool) {
		var res Result
		var i int
		var lineNum int
		for {
			i, res, _ = parseAny(json, i, true)
			if !res.Exists() {
				break
			}
			if !yield(lineNum, res) {
				return
			}
			lineNum++
		}
	}
}

func iterRecursiveDescent(parent Result, path string) iter.Seq[Result] {
	return func(yield func(Result) bool) {
		if res := parent.Get(path); res.Exists() {
			if !yield(res) {
				return
			}
		}
		if parent.IsArray() || parent.IsObject() {
			for _, val := range parent.Range() {
				for r := range iterRecursiveDescent(val, path) {
					if !yield(r) {
						return
					}
				}
			}
		}
	}
}

func (t Result) Paths(json string) iter.Seq[string] {
	return func(yield func(string) bool) {
		if t.Indexes == nil {
			return
		}
		for value := range t.Array() {
			if !yield(value.Path(json)) {
				return
			}
		}
	}
}

func (t Result) Paths2(json string) iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		if t.Indexes == nil {
			return
		}
		for i, value := range t.Array2() {
			if !yield(i, value.Path(json)) {
				return
			}
		}
	}
}
