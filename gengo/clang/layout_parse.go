package clang

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ddkwork/golibrary/mylog"
)

type layout struct {
	indent int // internal use
	Offset int
	Name   string
	Type   string
	Fields []*layout
}

func (n layout) UnmarshalString(data string) error {
	return nil
}

func (n *layout) tostring(wr *strings.Builder, pad string) {
	fmt.Fprintf(wr, "%s %s %s @ +0x%x\n", pad, n.Type, n.Name, n.Offset)
	for _, field := range n.Fields {
		field.tostring(wr, "  "+pad)
	}
}

func (n *layout) String() string {
	wr := &strings.Builder{}
	n.tostring(wr, "-")
	return wr.String()
}

func (n *layout) regroup() {
	res := make([]*layout, 0, len(n.Fields))

	for i := 0; i < len(n.Fields); i++ {
		field := n.Fields[i]

		// Find all children of the field, merge them into the field.
		for j := i + 1; j < len(n.Fields); j++ {
			if n.Fields[j].indent <= field.indent {
				break
			}
			field.Fields = append(field.Fields, n.Fields[j])
			i++
		}

		// Recursively regroup the field.
		field.regroup()

		// Append the field to the result.
		res = append(res, field)
	}
	n.Fields = res
}

type RecordLayout struct {
	layout
	Size, Align int
}

func (r *RecordLayout) UnmarshalString(data string) error {
	// mylog.Check(errors.New("improperly terminated layout"))

	switch {
	case strings.Contains(data, "__NSConstantString_tag"):
		//mylog.Warning("skip unmarshal RecordLayout", data)
		//return nil
	}

	first := true
	offset := 0

	for _, line := range strings.Split(data, "\n") {
		if line == "" {
			continue
		}
		before, after, found := strings.Cut(line, "|")
		if !found {
			continue
		}

		// If before is empty, then it is the size and align.
		before = strings.TrimSpace(before)
		if before == "" {
			after = strings.TrimSpace(after)
			//mylog.Check2(fmt.Sscanf(after, "[sizeof=%d, align=%d]", &r.Size, &r.Align))
			mylog.Check2(fmt.Sscanf(after, "[sizeof=%d, align=%d", &r.Size, &r.Align))
			break
		}

		// Parse offset
		if strings.Contains(before, ":") && strings.Contains(before, "-") {
			split := strings.Split(before, ":")
			bitRange := strings.Split(split[1], "-")
			start, end := mylog.Check2(strconv.Atoi(bitRange[0])), mylog.Check2(strconv.Atoi(bitRange[1]))
			offset += end - start + 1
		} else {
			offset = mylog.Check2Ignore(strconv.Atoi(strings.TrimSpace(before)))
		}

		// Determine indentation level
		indent := len(after)
		after = strings.TrimLeft(after, " \t")
		indent -= len(after)
		after = strings.TrimSpace(after)

		// Parse name and type
		name := "" //todo test

		typen := after
		//save strut type todo test
		if strings.HasPrefix(typen, "struct ") {
			//typen = "struct "
		}
		if lastSpace := strings.LastIndex(after, " "); lastSpace != -1 {
			// If the last space is followed by a closing parenthesis, then it is part of the type.
			if !strings.Contains(after[lastSpace+1:], ")") {
				// If type name is "struct" or "union", then the name is the real type.
				if after[:lastSpace] != "struct" && after[:lastSpace] != "union" {
					typen = after[:lastSpace]
					name = after[lastSpace+1:]
				}
			}
		}

		if name == "" {
			//continue
		}

		// Create node
		if first {
			r.Offset = offset
			r.Name = name
			r.Type = typen
			r.indent = indent
			first = false
		} else {
			r.Fields = append(r.Fields, &layout{
				Offset: offset,
				Name:   name,
				Type:   typen,
				indent: indent,
			})
		}
	}
	mylog.Json("layout", r.layout.Fields)
	// Group fields
	r.regroup()
	return nil
}

type LayoutMap struct {
	Records []*RecordLayout
	Map     map[string]*RecordLayout
}

const layoutMarker = "*** Dumping AST Record Layout"

func (l *LayoutMap) UnmarshalString(data string) error {
	mylog.Todo("bug _CR3_TYPE not decode into type map")
	data = strings.TrimSpace(data)
	data, found := strings.CutPrefix(data, layoutMarker)
	if !found {
		return fmt.Errorf("invalid layout data")
	}
	layout := strings.Split(data, layoutMarker)

	for _, part := range layout {
		record := &RecordLayout{}
		mylog.Check(record.UnmarshalString(part))
		l.Records = append(l.Records, record)
		l.Map[record.Type] = record
	}
	return nil
}

func ParseLayoutMap(data []byte) (*LayoutMap, error) {
	l := &LayoutMap{
		Map: make(map[string]*RecordLayout),
	}
	return l, l.UnmarshalString(string(data))
}
