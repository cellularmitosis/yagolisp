package main

import (
	"fmt"
	"os"
	"regexp"
)

// TokenType is a structure describing a particular type of token.
type TokenType struct {
	ID      uint
	Name    string
	Pattern string
	Regex   *regexp.Regexp
}

// The Token type identifiers.
const (
	TOK_OPAREN = iota
	TOK_CPAREN
	TOK_OBRACK
	TOK_CBRACK
	TOK_OBRACE
	TOK_CBRACE
	TOK_OHASHBRACE
	TOK_WSPACE
	TOK_STRING
	TOK_REAL
	TOK_INT
	TOK_SYMBOL
)

// The table to token definitions.
var tokTypes = []TokenType{
	{ID: TOK_OPAREN, Name: "TOK_OPAREN", Pattern: `^\(`, Regex: nil},
	{ID: TOK_CPAREN, Name: "TOK_CPAREN", Pattern: `^\)`, Regex: nil},
	{ID: TOK_OBRACK, Name: "TOK_OBRACK", Pattern: `^\[`, Regex: nil},
	{ID: TOK_CBRACK, Name: "TOK_CBRACK", Pattern: `^]`, Regex: nil},
	{ID: TOK_OBRACE, Name: "TOK_OBRACE", Pattern: `^{`, Regex: nil},
	{ID: TOK_CBRACE, Name: "TOK_CBRACE", Pattern: `^}`, Regex: nil},
	{ID: TOK_OHASHBRACE, Name: "TOK_OHASHBRACE", Pattern: `^#{`, Regex: nil},
	{ID: TOK_WSPACE, Name: "TOK_WSPACE", Pattern: `^[\s]+`, Regex: nil},
	{ID: TOK_STRING, Name: "TOK_STRING", Pattern: `^"([^"\\]|\\[\s\S])*"`, Regex: nil},
	{ID: TOK_REAL, Name: "TOK_REAL", Pattern: `^-?\d+\.\d+`, Regex: nil},
	{ID: TOK_INT, Name: "TOK_INT", Pattern: `^-?\d+`, Regex: nil},
	{ID: TOK_SYMBOL, Name: "TOK_SYMBOL", Pattern: `^[^\(\)\[\]{}#:"\s]+`, Regex: nil},
}

// Token is a structure describing a lexeme.
type Token struct {
	TypeID uint
	Bytes  []byte
}

// LexedFile is a structure describing a file which has been lexed.
type LexedFile struct {
	Filename string
	Tokens   []Token
}

// Loads a file into RAM.  Aborts on failure.
func mustLoadFile(fname string) []byte {
	file, err := os.Open(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: can't open file %s: %s\n", fname, err)
		os.Exit(EXITCODE_FOPEN)
	}
	defer file.Close()
	finfo, err := file.Stat()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: can't stat file %s: %s\n", fname, err)
		os.Exit(EXITCODE_FSTAT)
	}
	flen := finfo.Size()
	buff := make([]byte, flen)
	_, err = file.Read(buff)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: can't read file %s: %s\n", fname, err)
		os.Exit(EXITCODE_FREAD)
	}
	_ = buff
	return buff
}

// Compiles the table to token regexes.  Aborts on failure.
func mustCompileRegexes() {
	for i := 0; i < len(tokTypes); i++ {
		pattern := tokTypes[i].Pattern
		regex := regexp.MustCompile(pattern)
		tokTypes[i].Regex = regex
	}
}

// Tokenizes the loaded file.  Aborts on failure.
func mustLex(bytes []byte) []Token {
	tokens := make([]Token, 0, 1024)
	// iterate over the bytes of the file
	for i := 0; i < len(bytes); {
		didMatch := false
		// try each of the regexes until you find a match
		for j := 0; j < len(tokTypes); j++ {
			tokType := tokTypes[j]
			regex := tokType.Regex
			cursor := bytes[i:]
			match := regex.FindIndex(cursor)
			if match == nil {
				continue
			} else {
				matchStart := i + match[0]
				matchEnd := i + match[1]
				// drop any whitespace tokens
				if tokType.ID != TOK_WSPACE {
					matchedBytes := bytes[matchStart:matchEnd]
					token := Token{
						TypeID: tokType.ID,
						Bytes:  matchedBytes,
					}
					tokens = append(tokens, token)
				}
				length := matchEnd - matchStart
				i += length
				didMatch = true
				break
			}
		}

		if !didMatch {
			context := safeSlice(bytes, i, i+16)
			fmt.Fprintf(
				os.Stderr,
				"Error: can't match any token at index %d: \"%s\"",
				i,
				context,
			)
			os.Exit(EXITCODE_LEX)
		}
	}
	return tokens
}
