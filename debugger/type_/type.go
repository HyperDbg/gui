package type_

import (
	"fmt"
	"maps"

	"gioui.org/layout"
	"github.com/ddkwork/HyperDbg/debugger/api"
	"github.com/ddkwork/golibrary/std/safemap"
)

type EncodeType int

const (
	EncodeUnknown EncodeType = iota
	EncodeByte
	EncodeWord
	EncodeDword
	EncodeFword
	EncodeQword
	EncodeTbyte
	EncodeOword
	EncodeMmword
	EncodeXmmword
	EncodeYmmword
	EncodeZmmword
	EncodeReal4
	EncodeReal8
	EncodeReal10
	EncodeAscii
	EncodeUnicode
	EncodeCode
	EncodeJunk
	EncodeMiddle
)

type TypeDefinition struct {
	Name   string
	Size   uint32
	Encode EncodeType
	Fields []FieldDefinition
	Auto   bool
}

type FieldDefinition struct {
	Name     string
	Offset   uint32
	Size     uint32
	Encode   EncodeType
	TypeName string
}

type TypeManager struct {
	types       *safemap.M[string, *TypeDefinition]
	encodeTypes *safemap.M[uint64, EncodeType]
}

func New() api.Interface {
	return &TypeManager{
		types:       safemap.New[string, *TypeDefinition](),
		encodeTypes: safemap.New[uint64, EncodeType](),
	}
}

func (m *TypeManager) AddType(name string, size uint32, encode EncodeType, auto bool) error {
	if name == "" {
		return fmt.Errorf("type name cannot be empty")
	}

	if _, exists := m.types.Get(name); exists {
		return fmt.Errorf("type %s already exists", name)
	}

	m.types.Update(name, &TypeDefinition{
		Name:   name,
		Size:   size,
		Encode: encode,
		Auto:   auto,
	})

	return nil
}

func (m *TypeManager) GetType(name string) *TypeDefinition {
	typ, _ := m.types.Get(name)
	return typ
}

func (m *TypeManager) DeleteType(name string) {
	m.types.Delete(name)
}

func (m *TypeManager) GetAllTypes() []*TypeDefinition {
	result := make([]*TypeDefinition, 0)
	for _, typ := range m.types.Range() {
		result = append(result, typ)
	}
	return result
}

func (m *TypeManager) GetAutoTypes() []*TypeDefinition {
	result := make([]*TypeDefinition, 0)
	for _, typ := range m.types.Range() {
		if typ.Auto {
			result = append(result, typ)
		}
	}
	return result
}

func (m *TypeManager) GetUserTypes() []*TypeDefinition {
	result := make([]*TypeDefinition, 0)
	for _, typ := range m.types.Range() {
		if !typ.Auto {
			result = append(result, typ)
		}
	}
	return result
}

func (m *TypeManager) HasType(name string) bool {
	_, exists := m.types.Get(name)
	return exists
}

func (m *TypeManager) SetEncodeType(address uint64, encode EncodeType) {
	m.encodeTypes.Update(address, encode)
}

func (m *TypeManager) GetEncodeType(address uint64) EncodeType {
	if encode, exists := m.encodeTypes.Get(address); exists {
		return encode
	}
	return EncodeUnknown
}

func (m *TypeManager) DeleteEncodeType(address uint64) {
	m.encodeTypes.Delete(address)
}

func (m *TypeManager) DeleteEncodeTypeRange(start, end uint64) {
	for addr := range m.encodeTypes.Range() {
		if addr >= start && addr <= end {
			m.encodeTypes.Delete(addr)
		}
	}
}

func (m *TypeManager) GetEncodeTypeBuffer() map[uint64]EncodeType {
	result := maps.Collect(m.encodeTypes.Range())
	return result
}

func (m *TypeManager) GetEncodeSize(address uint64, codeSize uint32) uint32 {
	if encode, exists := m.encodeTypes.Get(address); exists {
		return m.GetEncodeTypeSize(encode)
	}
	return codeSize
}

func (m *TypeManager) GetEncodeTypeSize(encode EncodeType) uint32 {
	switch encode {
	case EncodeByte:
		return 1
	case EncodeWord:
		return 2
	case EncodeDword, EncodeReal4:
		return 4
	case EncodeFword:
		return 6
	case EncodeQword, EncodeReal8, EncodeMmword:
		return 8
	case EncodeTbyte, EncodeReal10:
		return 10
	case EncodeOword, EncodeXmmword:
		return 16
	case EncodeYmmword:
		return 32
	case EncodeZmmword:
		return 64
	default:
		return 0
	}
}

func (m *TypeManager) AddField(typeName, fieldName string, offset, size uint32, encode EncodeType) error {
	typ, exists := m.types.Get(typeName)
	if !exists {
		return fmt.Errorf("type %s does not exist", typeName)
	}

	for _, field := range typ.Fields {
		if field.Name == fieldName {
			return fmt.Errorf("field %s already exists in type %s", fieldName, typeName)
		}
	}

	typ.Fields = append(typ.Fields, FieldDefinition{
		Name:     fieldName,
		Offset:   offset,
		Size:     size,
		Encode:   encode,
		TypeName: typeName,
	})

	return nil
}

func (m *TypeManager) GetFields(typeName string) []FieldDefinition {
	if typ, exists := m.types.Get(typeName); exists {
		result := make([]FieldDefinition, len(typ.Fields))
		copy(result, typ.Fields)
		return result
	}
	return nil
}

func (m *TypeManager) GetField(typeName, fieldName string) *FieldDefinition {
	if typ, exists := m.types.Get(typeName); exists {
		for _, field := range typ.Fields {
			if field.Name == fieldName {
				return &field
			}
		}
	}
	return nil
}

func (m *TypeManager) DeleteField(typeName, fieldName string) {
	if typ, exists := m.types.Get(typeName); exists {
		for i, field := range typ.Fields {
			if field.Name == fieldName {
				typ.Fields = append(typ.Fields[:i], typ.Fields[i+1:]...)
				break
			}
		}
	}
}

func (m *TypeManager) GetTypeNameByAddress(address uint64) string {
	for _, typ := range m.types.Range() {
		if typ.Encode != EncodeUnknown {
			return typ.Name
		}
	}
	return ""
}

func (m *TypeManager) GetTypeCount() int {
	count := 0
	for range m.types.Range() {
		count++
	}
	return count
}

func (m *TypeManager) GetAutoTypeCount() int {
	count := 0
	for _, typ := range m.types.Range() {
		if typ.Auto {
			count++
		}
	}
	return count
}

func (m *TypeManager) GetUserTypeCount() int {
	count := 0
	for _, typ := range m.types.Range() {
		if !typ.Auto {
			count++
		}
	}
	return count
}

func (m *TypeManager) ClearAutoTypes() {
	for name, typ := range m.types.Range() {
		if typ.Auto {
			m.types.Delete(name)
		}
	}
}

func (m *TypeManager) FindTypesBySize(size uint32) []*TypeDefinition {
	result := make([]*TypeDefinition, 0)
	for _, typ := range m.types.Range() {
		if typ.Size == size {
			result = append(result, typ)
		}
	}
	return result
}

func (m *TypeManager) FindTypesByEncode(encode EncodeType) []*TypeDefinition {
	result := make([]*TypeDefinition, 0)
	for _, typ := range m.types.Range() {
		if typ.Encode == encode {
			result = append(result, typ)
		}
	}
	return result
}

func (m *TypeManager) SetTypeName(oldName, newName string) error {
	if _, exists := m.types.Get(oldName); !exists {
		return fmt.Errorf("type %s does not exist", oldName)
	}

	if _, exists := m.types.Get(newName); exists {
		return fmt.Errorf("type %s already exists", newName)
	}

	oldTyp, _ := m.types.Get(oldName)
	m.types.Update(newName, oldTyp)
	newTyp, _ := m.types.Get(newName)
	newTyp.Name = newName
	m.types.Delete(oldName)

	return nil
}

func (m *TypeManager) SetTypeSize(name string, size uint32) error {
	if typ, exists := m.types.Get(name); exists {
		typ.Size = size
		return nil
	}
	return fmt.Errorf("type %s does not exist", name)
}

func (m *TypeManager) SetTypeEncode(name string, encode EncodeType) error {
	if typ, exists := m.types.Get(name); exists {
		typ.Encode = encode
		return nil
	}
	return fmt.Errorf("type %s does not exist", name)
}

func (m *TypeManager) Layout() layout.Widget {
	return nil
}

func (m *TypeManager) Update() error {
	return nil
}

func (m *TypeManager) Clear() {
	m.types.Reset()
	m.encodeTypes.Reset()
}

func (m *TypeManager) Self() any {
	return m
}
