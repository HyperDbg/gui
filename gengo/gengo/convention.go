package gengo

import "strings"

type NameConvention int

const (
	ConventionCamelCase  NameConvention = iota // camelCase
	ConventionSnakeCase                        // snake_case
	ConventionConstCase                        // CONST_CASE
	ConventionPascalCase                       // PascalCase
)

const (
	upperRunes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerRunes = "abcdefghijklmnopqrstuvwxyz"
)

func isupper(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func islower(c byte) bool {
	return c >= 'a' && c <= 'z'
}

func isAllLower(s string) bool {
	for i := 0; i < len(s); i++ {
		if isupper(s[i]) {
			return false
		}
	}
	return true
}

func isAllUpper(s string) bool {
	for i := 0; i < len(s); i++ {
		if islower(s[i]) {
			return false
		}
	}
	return true
}

func detectConvention(s string) NameConvention {
	if isAllLower(s) {
		// snake_case or camelCase
		if strings.IndexByte(s, '_') >= 0 {
			return ConventionSnakeCase
		} else {
			return ConventionCamelCase
		}
	} else {
		// CONST_CASE or PascalCase
		if isAllUpper(s) {
			return ConventionConstCase
		} else {
			return ConventionPascalCase
		}
	}
}

func toCapital(s string) string {
	// "" -> ""
	if len(s) == 0 {
		return ""
	}
	firstLower := strings.IndexAny(s, lowerRunes)

	// "ABCD" -> "Abcd"
	if firstLower < 0 {
		return s[:1] + strings.ToLower(s[1:])
	}

	// "ABcd" -> "ABcd"
	if firstLower != 0 && !strings.ContainsAny(s[firstLower:], upperRunes) {
		return s
	}

	// "abcd"/"abCD"/"AbCd" -> "Abcd"
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}

func toPascalCase(s string, conv NameConvention) string {
	// PascalCase -> PascalCase
	if conv == ConventionPascalCase {
		return s
	}

	// snake_case -> SnakeCase
	// CONST_CASE -> ConstCase
	if conv == ConventionConstCase || conv == ConventionSnakeCase {
		words := strings.Split(s, "_")
		builder := strings.Builder{}
		for _, word := range words {
			builder.WriteString(toCapital(word))
		}
		return builder.String()
	}

	// camelCase -> CamelCase
	splitFrom := strings.IndexAny(s, upperRunes)
	if splitFrom < 0 {
		return toCapital(s)
	}
	return toCapital(s[:splitFrom]) + s[splitFrom:]
}

func toCamelCase(s string, conv NameConvention) string {
	// camelCase -> camelCase
	if conv == ConventionCamelCase {
		return s
	}

	// snake_case -> snakeCase
	// CONST_CASE -> constCase
	if conv == ConventionConstCase || conv == ConventionSnakeCase {
		words := strings.Split(s, "_")
		builder := strings.Builder{}
		builder.WriteString(strings.ToLower(words[0]))
		for _, word := range words[1:] {
			builder.WriteString(toCapital(word))
		}
		return builder.String()
	}

	// PascalCase -> pascalCase
	splitFrom := strings.IndexAny(s, lowerRunes)
	if splitFrom < 0 {
		return strings.ToLower(s)
	}
	return strings.ToLower(s[:splitFrom]) + s[splitFrom:]
}

func toSnakeCase(s string, conv NameConvention) string {
	// snake_case -> snake_case
	if conv == ConventionSnakeCase {
		return s
	}
	// CONST_CASE -> const_case
	if conv == ConventionConstCase {
		return strings.ToLower(s)
	}

	// camelCase -> camel_case
	// PascalCase -> pascal_case
	// PAscalCAse -> pascal_case
	builder := strings.Builder{}
	split := false
	for i, c := range s {
		if i > 0 && isupper(byte(c)) {
			if !split {
				builder.WriteByte('_')
			}
			split = true
		} else {
			split = false
		}
		if 'A' <= c && c <= 'Z' {
			c += 'a' - 'A'
		}
		builder.WriteByte(byte(c))
	}
	return builder.String()
}

func toConstCase(s string, conv NameConvention) string {
	// CONST_CASE -> CONST_CASE
	if conv == ConventionConstCase {
		return s
	}
	// snake_case -> SNAKE_CASE
	if conv == ConventionSnakeCase {
		return strings.ToUpper(s)
	}

	// camelCase -> CAMEL_CASE
	// PascalCase -> PASCAL_CASE
	// PAscalCAse -> PASCAL_CASE
	builder := strings.Builder{}
	split := false
	for i, c := range s {
		if i > 0 && isupper(byte(c)) {
			if !split {
				builder.WriteByte('_')
			}
			split = true
		} else {
			split = false
		}
		builder.WriteByte(byte(c))
	}
	return strings.ToUpper(builder.String())
}

func ConvertCase(s string, into NameConvention) string {
	s, pfx := strings.CutPrefix(s, "_")
	s, sfx := strings.CutSuffix(s, "_")
	conv := detectConvention(s)

	pfxv, sfxv := "", ""
	if pfx {
		pfxv = "_"
	}
	if sfx {
		sfxv = "_"
	}

	switch into {
	case ConventionCamelCase:
		return pfxv + toCamelCase(s, conv) + sfxv
	case ConventionSnakeCase:
		return pfxv + toSnakeCase(s, conv) + sfxv
	case ConventionConstCase:
		return pfxv + toConstCase(s, conv) + sfxv
	case ConventionPascalCase:
		return pfxv + toPascalCase(s, conv) + sfxv
	default:
		return pfxv + s + sfxv
	}
}
