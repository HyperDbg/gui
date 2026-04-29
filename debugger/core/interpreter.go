package core

import (
	"strings"
	"unicode"
)

type CommandToken struct {
	Value string
	Type  TokenType
}

type TokenType int

const (
	TokenTypeCommand TokenType = iota
	TokenTypeArgument
	TokenTypeQuotedString
	TokenTypeComment
)

type Interpreter struct {
	previousCommandShouldBeContinued bool
	isInterpreterOnString            bool
	isPreviousCharBackSlash          bool
	countOfOpenCurlyBrackets         int
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		countOfOpenCurlyBrackets: 0,
	}
}

func (interp *Interpreter) Parse(input string) []CommandToken {
	var tokens []CommandToken
	var current strings.Builder
	inQuotes := false

	input = interp.removeComments(input)

	for i, c := range input {
		if c == '\\' && i+1 < len(input) {
			next := input[i+1]
			if next == '"' || next == '\\' || next == 'n' || next == 't' || next == 'r' {
				current.WriteRune(c)
				current.WriteRune(rune(next))
				interp.isPreviousCharBackSlash = true
				continue
			}
		}

		if c == '"' && !interp.isPreviousCharBackSlash {
			if inQuotes {
				if current.Len() > 0 {
					tokens = append(tokens, CommandToken{
						Value: current.String(),
						Type:  TokenTypeQuotedString,
					})
				}
				current.Reset()
				inQuotes = false
			} else {
				inQuotes = true
			}
			interp.isPreviousCharBackSlash = false
			continue
		}

		if !inQuotes {
			if unicode.IsSpace(c) {
				if current.Len() > 0 {
					tokens = append(tokens, CommandToken{
						Value: current.String(),
						Type:  TokenTypeArgument,
					})
					current.Reset()
				}
				interp.isPreviousCharBackSlash = false
				continue
			}

			if c == '{' {
				interp.countOfOpenCurlyBrackets++
			} else if c == '}' {
				if interp.countOfOpenCurlyBrackets > 0 {
					interp.countOfOpenCurlyBrackets--
				}
			}
		}

		current.WriteRune(c)
		interp.isPreviousCharBackSlash = false
	}

	if inQuotes {
	} else if current.Len() > 0 {
		tokens = append(tokens, CommandToken{
			Value: current.String(),
			Type:  TokenTypeArgument,
		})
	}

	if len(tokens) > 0 {
		tokens[0].Type = TokenTypeCommand
	}

	return tokens
}

func (interp *Interpreter) removeComments(input string) string {
	var result strings.Builder
	inSingleLineComment := false

	for i := 0; i < len(input); i++ {
		if i+1 < len(input) && input[i] == '/' && input[i+1] == '/' {
			inSingleLineComment = true
			i++
			continue
		}

		if inSingleLineComment {
			if input[i] == '\n' {
				inSingleLineComment = false
				result.WriteByte('\n')
			}
			continue
		}

		result.WriteByte(input[i])
	}

	return result.String()
}

func (interp *Interpreter) ShouldPreviousCommandBeContinued() bool {
	return interp.previousCommandShouldBeContinued
}

func (interp *Interpreter) SetPreviousCommandShouldBeContinued(should bool) {
	interp.previousCommandShouldBeContinued = should
}

func (interp *Interpreter) GetCountOfOpenCurlyBrackets() int {
	return interp.countOfOpenCurlyBrackets
}

func (interp *Interpreter) IsInScriptBlock() bool {
	return interp.countOfOpenCurlyBrackets > 0
}

type CommandParseResult struct {
	Tokens    []CommandToken
	Command   string
	Arguments []string
	IsValid   bool
	Error     string
}

func (interp *Interpreter) ParseCommand(input string) CommandParseResult {
	result := CommandParseResult{}

	trimmed := strings.TrimSpace(input)
	if len(trimmed) == 0 {
		result.IsValid = false
		result.Error = "empty command"
		return result
	}

	tokens := interp.Parse(trimmed)
	if len(tokens) == 0 {
		result.IsValid = false
		result.Error = "no tokens parsed"
		return result
	}

	result.Tokens = tokens
	result.Command = strings.ToLower(tokens[0].Value)
	result.IsValid = true

	for i := 1; i < len(tokens); i++ {
		result.Arguments = append(result.Arguments, tokens[i].Value)
	}

	return result
}
