// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type NaslLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var nasllexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func nasllexerLexerInit() {
	staticData := &nasllexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "", "'['", "']'", "'('", "')'", "'{'", "'}'", "';'", "','", "'='",
		"':'", "'.'", "'++'", "'--'", "'+'", "'-'", "'~'", "'&'", "'^'", "'|'",
		"'>>'", "'<<'", "'>>>'", "'!'", "'*'", "'**'", "'/'", "'%'", "'<'",
		"'>'", "'<='", "'>='", "'=='", "'=~'", "'!='", "'!~'", "'>!<'", "'><'",
		"'&&'", "'||'", "'*='", "'/='", "'%='", "'+='", "'-='", "'x'", "'break'",
		"'local_var'", "'global_var'", "'else'", "'return'", "'continue'", "'for'",
		"'foreach'", "'if'", "'function'", "'repeat'", "'while'", "'until'",
		"'exit'", "", "", "", "", "", "", "'NULL'",
	}
	staticData.symbolicNames = []string{
		"", "SingleLineComment", "OpenBracket", "CloseBracket", "OpenParen",
		"CloseParen", "OpenBrace", "CloseBrace", "SemiColon", "Comma", "Assign",
		"Colon", "Dot", "PlusPlus", "MinusMinus", "Plus", "Minus", "BitNot",
		"BitAnd", "BitXOr", "BitOr", "RightShiftArithmetic", "LeftShiftArithmetic",
		"RightShiftLogical", "Not", "Multiply", "Pow", "Divide", "Modulus",
		"LessThan", "MoreThan", "LessThanEquals", "GreaterThanEquals", "Equals_",
		"EqualsRe", "NotEquals", "NotLong", "MTNotLT", "MTLT", "And", "Or",
		"MultiplyAssign", "DivideAssign", "ModulusAssign", "PlusAssign", "MinusAssign",
		"X", "Break", "LocalVar", "GlobalVar", "Else", "Return", "Continue",
		"For", "ForEach", "If", "Function_", "Repeat", "While", "Until", "Exit",
		"StringLiteral", "BooleanLiteral", "IntegerLiteral", "FloatLiteral",
		"IpLiteral", "HexLiteral", "NULLLiteral", "Identifier", "WhiteSpaces",
		"LineTerminator",
	}
	staticData.ruleNames = []string{
		"SingleLineComment", "OpenBracket", "CloseBracket", "OpenParen", "CloseParen",
		"OpenBrace", "CloseBrace", "SemiColon", "Comma", "Assign", "Colon",
		"Dot", "PlusPlus", "MinusMinus", "Plus", "Minus", "BitNot", "BitAnd",
		"BitXOr", "BitOr", "RightShiftArithmetic", "LeftShiftArithmetic", "RightShiftLogical",
		"Not", "Multiply", "Pow", "Divide", "Modulus", "LessThan", "MoreThan",
		"LessThanEquals", "GreaterThanEquals", "Equals_", "EqualsRe", "NotEquals",
		"NotLong", "MTNotLT", "MTLT", "And", "Or", "MultiplyAssign", "DivideAssign",
		"ModulusAssign", "PlusAssign", "MinusAssign", "X", "Break", "LocalVar",
		"GlobalVar", "Else", "Return", "Continue", "For", "ForEach", "If", "Function_",
		"Repeat", "While", "Until", "Exit", "StringLiteral", "SingleStringCharacter",
		"BooleanLiteral", "IntegerLiteral", "FloatLiteral", "IpLiteral", "HexLiteral",
		"NULLLiteral", "Identifier", "WhiteSpaces", "LineTerminator",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 70, 474, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67,
		2, 68, 7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 1, 0, 1, 0, 5, 0, 146, 8, 0,
		10, 0, 12, 0, 149, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1,
		14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19,
		1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27,
		1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1,
		31, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 35,
		1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 38, 1,
		38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41,
		1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 45, 1,
		45, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 1, 47,
		1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1, 48, 1, 48, 1,
		48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 49, 1, 49, 1, 49, 1, 49,
		1, 49, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 51, 1, 51, 1,
		51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 52, 1, 52, 1, 52, 1, 52,
		1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 54, 1, 54, 1,
		54, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 56,
		1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 57, 1, 57, 1, 57, 1, 57, 1,
		57, 1, 57, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 59, 1, 59, 1, 59,
		1, 59, 1, 59, 1, 60, 1, 60, 5, 60, 364, 8, 60, 10, 60, 12, 60, 367, 9,
		60, 1, 60, 1, 60, 1, 60, 5, 60, 372, 8, 60, 10, 60, 12, 60, 375, 9, 60,
		1, 60, 3, 60, 378, 8, 60, 1, 61, 1, 61, 1, 61, 3, 61, 383, 8, 61, 1, 62,
		1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1,
		62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 3, 62, 403, 8, 62, 1, 63,
		4, 63, 406, 8, 63, 11, 63, 12, 63, 407, 1, 64, 1, 64, 1, 64, 5, 64, 413,
		8, 64, 10, 64, 12, 64, 416, 9, 64, 1, 65, 5, 65, 419, 8, 65, 10, 65, 12,
		65, 422, 9, 65, 1, 65, 1, 65, 5, 65, 426, 8, 65, 10, 65, 12, 65, 429, 9,
		65, 1, 65, 1, 65, 5, 65, 433, 8, 65, 10, 65, 12, 65, 436, 9, 65, 1, 65,
		1, 65, 5, 65, 440, 8, 65, 10, 65, 12, 65, 443, 9, 65, 1, 66, 1, 66, 1,
		66, 4, 66, 448, 8, 66, 11, 66, 12, 66, 449, 1, 67, 1, 67, 1, 67, 1, 67,
		1, 67, 1, 68, 1, 68, 5, 68, 459, 8, 68, 10, 68, 12, 68, 462, 9, 68, 1,
		69, 4, 69, 465, 8, 69, 11, 69, 12, 69, 466, 1, 69, 1, 69, 1, 70, 1, 70,
		1, 70, 1, 70, 0, 0, 71, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
		35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26,
		53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35,
		71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44,
		89, 45, 91, 46, 93, 47, 95, 48, 97, 49, 99, 50, 101, 51, 103, 52, 105,
		53, 107, 54, 109, 55, 111, 56, 113, 57, 115, 58, 117, 59, 119, 60, 121,
		61, 123, 0, 125, 62, 127, 63, 129, 64, 131, 65, 133, 66, 135, 67, 137,
		68, 139, 69, 141, 70, 1, 0, 9, 3, 0, 10, 10, 13, 13, 8232, 8233, 1, 0,
		34, 34, 2, 0, 39, 39, 92, 92, 1, 0, 48, 57, 2, 0, 88, 88, 120, 120, 3,
		0, 48, 57, 65, 70, 97, 102, 4, 0, 36, 36, 65, 90, 95, 95, 97, 122, 5, 0,
		36, 36, 48, 57, 65, 90, 95, 95, 97, 122, 4, 0, 9, 9, 11, 12, 32, 32, 160,
		160, 489, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0,
		23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0,
		0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0,
		0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0,
		0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1,
		0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61,
		1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0,
		69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0,
		0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0,
		0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0,
		0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1,
		0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0,
		107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 111, 1, 0, 0, 0, 0, 113, 1, 0,
		0, 0, 0, 115, 1, 0, 0, 0, 0, 117, 1, 0, 0, 0, 0, 119, 1, 0, 0, 0, 0, 121,
		1, 0, 0, 0, 0, 125, 1, 0, 0, 0, 0, 127, 1, 0, 0, 0, 0, 129, 1, 0, 0, 0,
		0, 131, 1, 0, 0, 0, 0, 133, 1, 0, 0, 0, 0, 135, 1, 0, 0, 0, 0, 137, 1,
		0, 0, 0, 0, 139, 1, 0, 0, 0, 0, 141, 1, 0, 0, 0, 1, 143, 1, 0, 0, 0, 3,
		152, 1, 0, 0, 0, 5, 154, 1, 0, 0, 0, 7, 156, 1, 0, 0, 0, 9, 158, 1, 0,
		0, 0, 11, 160, 1, 0, 0, 0, 13, 162, 1, 0, 0, 0, 15, 164, 1, 0, 0, 0, 17,
		166, 1, 0, 0, 0, 19, 168, 1, 0, 0, 0, 21, 170, 1, 0, 0, 0, 23, 172, 1,
		0, 0, 0, 25, 174, 1, 0, 0, 0, 27, 177, 1, 0, 0, 0, 29, 180, 1, 0, 0, 0,
		31, 182, 1, 0, 0, 0, 33, 184, 1, 0, 0, 0, 35, 186, 1, 0, 0, 0, 37, 188,
		1, 0, 0, 0, 39, 190, 1, 0, 0, 0, 41, 192, 1, 0, 0, 0, 43, 195, 1, 0, 0,
		0, 45, 198, 1, 0, 0, 0, 47, 202, 1, 0, 0, 0, 49, 204, 1, 0, 0, 0, 51, 206,
		1, 0, 0, 0, 53, 209, 1, 0, 0, 0, 55, 211, 1, 0, 0, 0, 57, 213, 1, 0, 0,
		0, 59, 215, 1, 0, 0, 0, 61, 217, 1, 0, 0, 0, 63, 220, 1, 0, 0, 0, 65, 223,
		1, 0, 0, 0, 67, 226, 1, 0, 0, 0, 69, 229, 1, 0, 0, 0, 71, 232, 1, 0, 0,
		0, 73, 235, 1, 0, 0, 0, 75, 239, 1, 0, 0, 0, 77, 242, 1, 0, 0, 0, 79, 245,
		1, 0, 0, 0, 81, 248, 1, 0, 0, 0, 83, 251, 1, 0, 0, 0, 85, 254, 1, 0, 0,
		0, 87, 257, 1, 0, 0, 0, 89, 260, 1, 0, 0, 0, 91, 263, 1, 0, 0, 0, 93, 265,
		1, 0, 0, 0, 95, 271, 1, 0, 0, 0, 97, 281, 1, 0, 0, 0, 99, 292, 1, 0, 0,
		0, 101, 297, 1, 0, 0, 0, 103, 304, 1, 0, 0, 0, 105, 313, 1, 0, 0, 0, 107,
		317, 1, 0, 0, 0, 109, 325, 1, 0, 0, 0, 111, 328, 1, 0, 0, 0, 113, 337,
		1, 0, 0, 0, 115, 344, 1, 0, 0, 0, 117, 350, 1, 0, 0, 0, 119, 356, 1, 0,
		0, 0, 121, 377, 1, 0, 0, 0, 123, 382, 1, 0, 0, 0, 125, 402, 1, 0, 0, 0,
		127, 405, 1, 0, 0, 0, 129, 409, 1, 0, 0, 0, 131, 420, 1, 0, 0, 0, 133,
		444, 1, 0, 0, 0, 135, 451, 1, 0, 0, 0, 137, 456, 1, 0, 0, 0, 139, 464,
		1, 0, 0, 0, 141, 470, 1, 0, 0, 0, 143, 147, 5, 35, 0, 0, 144, 146, 8, 0,
		0, 0, 145, 144, 1, 0, 0, 0, 146, 149, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0,
		147, 148, 1, 0, 0, 0, 148, 150, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 150,
		151, 6, 0, 0, 0, 151, 2, 1, 0, 0, 0, 152, 153, 5, 91, 0, 0, 153, 4, 1,
		0, 0, 0, 154, 155, 5, 93, 0, 0, 155, 6, 1, 0, 0, 0, 156, 157, 5, 40, 0,
		0, 157, 8, 1, 0, 0, 0, 158, 159, 5, 41, 0, 0, 159, 10, 1, 0, 0, 0, 160,
		161, 5, 123, 0, 0, 161, 12, 1, 0, 0, 0, 162, 163, 5, 125, 0, 0, 163, 14,
		1, 0, 0, 0, 164, 165, 5, 59, 0, 0, 165, 16, 1, 0, 0, 0, 166, 167, 5, 44,
		0, 0, 167, 18, 1, 0, 0, 0, 168, 169, 5, 61, 0, 0, 169, 20, 1, 0, 0, 0,
		170, 171, 5, 58, 0, 0, 171, 22, 1, 0, 0, 0, 172, 173, 5, 46, 0, 0, 173,
		24, 1, 0, 0, 0, 174, 175, 5, 43, 0, 0, 175, 176, 5, 43, 0, 0, 176, 26,
		1, 0, 0, 0, 177, 178, 5, 45, 0, 0, 178, 179, 5, 45, 0, 0, 179, 28, 1, 0,
		0, 0, 180, 181, 5, 43, 0, 0, 181, 30, 1, 0, 0, 0, 182, 183, 5, 45, 0, 0,
		183, 32, 1, 0, 0, 0, 184, 185, 5, 126, 0, 0, 185, 34, 1, 0, 0, 0, 186,
		187, 5, 38, 0, 0, 187, 36, 1, 0, 0, 0, 188, 189, 5, 94, 0, 0, 189, 38,
		1, 0, 0, 0, 190, 191, 5, 124, 0, 0, 191, 40, 1, 0, 0, 0, 192, 193, 5, 62,
		0, 0, 193, 194, 5, 62, 0, 0, 194, 42, 1, 0, 0, 0, 195, 196, 5, 60, 0, 0,
		196, 197, 5, 60, 0, 0, 197, 44, 1, 0, 0, 0, 198, 199, 5, 62, 0, 0, 199,
		200, 5, 62, 0, 0, 200, 201, 5, 62, 0, 0, 201, 46, 1, 0, 0, 0, 202, 203,
		5, 33, 0, 0, 203, 48, 1, 0, 0, 0, 204, 205, 5, 42, 0, 0, 205, 50, 1, 0,
		0, 0, 206, 207, 5, 42, 0, 0, 207, 208, 5, 42, 0, 0, 208, 52, 1, 0, 0, 0,
		209, 210, 5, 47, 0, 0, 210, 54, 1, 0, 0, 0, 211, 212, 5, 37, 0, 0, 212,
		56, 1, 0, 0, 0, 213, 214, 5, 60, 0, 0, 214, 58, 1, 0, 0, 0, 215, 216, 5,
		62, 0, 0, 216, 60, 1, 0, 0, 0, 217, 218, 5, 60, 0, 0, 218, 219, 5, 61,
		0, 0, 219, 62, 1, 0, 0, 0, 220, 221, 5, 62, 0, 0, 221, 222, 5, 61, 0, 0,
		222, 64, 1, 0, 0, 0, 223, 224, 5, 61, 0, 0, 224, 225, 5, 61, 0, 0, 225,
		66, 1, 0, 0, 0, 226, 227, 5, 61, 0, 0, 227, 228, 5, 126, 0, 0, 228, 68,
		1, 0, 0, 0, 229, 230, 5, 33, 0, 0, 230, 231, 5, 61, 0, 0, 231, 70, 1, 0,
		0, 0, 232, 233, 5, 33, 0, 0, 233, 234, 5, 126, 0, 0, 234, 72, 1, 0, 0,
		0, 235, 236, 5, 62, 0, 0, 236, 237, 5, 33, 0, 0, 237, 238, 5, 60, 0, 0,
		238, 74, 1, 0, 0, 0, 239, 240, 5, 62, 0, 0, 240, 241, 5, 60, 0, 0, 241,
		76, 1, 0, 0, 0, 242, 243, 5, 38, 0, 0, 243, 244, 5, 38, 0, 0, 244, 78,
		1, 0, 0, 0, 245, 246, 5, 124, 0, 0, 246, 247, 5, 124, 0, 0, 247, 80, 1,
		0, 0, 0, 248, 249, 5, 42, 0, 0, 249, 250, 5, 61, 0, 0, 250, 82, 1, 0, 0,
		0, 251, 252, 5, 47, 0, 0, 252, 253, 5, 61, 0, 0, 253, 84, 1, 0, 0, 0, 254,
		255, 5, 37, 0, 0, 255, 256, 5, 61, 0, 0, 256, 86, 1, 0, 0, 0, 257, 258,
		5, 43, 0, 0, 258, 259, 5, 61, 0, 0, 259, 88, 1, 0, 0, 0, 260, 261, 5, 45,
		0, 0, 261, 262, 5, 61, 0, 0, 262, 90, 1, 0, 0, 0, 263, 264, 5, 120, 0,
		0, 264, 92, 1, 0, 0, 0, 265, 266, 5, 98, 0, 0, 266, 267, 5, 114, 0, 0,
		267, 268, 5, 101, 0, 0, 268, 269, 5, 97, 0, 0, 269, 270, 5, 107, 0, 0,
		270, 94, 1, 0, 0, 0, 271, 272, 5, 108, 0, 0, 272, 273, 5, 111, 0, 0, 273,
		274, 5, 99, 0, 0, 274, 275, 5, 97, 0, 0, 275, 276, 5, 108, 0, 0, 276, 277,
		5, 95, 0, 0, 277, 278, 5, 118, 0, 0, 278, 279, 5, 97, 0, 0, 279, 280, 5,
		114, 0, 0, 280, 96, 1, 0, 0, 0, 281, 282, 5, 103, 0, 0, 282, 283, 5, 108,
		0, 0, 283, 284, 5, 111, 0, 0, 284, 285, 5, 98, 0, 0, 285, 286, 5, 97, 0,
		0, 286, 287, 5, 108, 0, 0, 287, 288, 5, 95, 0, 0, 288, 289, 5, 118, 0,
		0, 289, 290, 5, 97, 0, 0, 290, 291, 5, 114, 0, 0, 291, 98, 1, 0, 0, 0,
		292, 293, 5, 101, 0, 0, 293, 294, 5, 108, 0, 0, 294, 295, 5, 115, 0, 0,
		295, 296, 5, 101, 0, 0, 296, 100, 1, 0, 0, 0, 297, 298, 5, 114, 0, 0, 298,
		299, 5, 101, 0, 0, 299, 300, 5, 116, 0, 0, 300, 301, 5, 117, 0, 0, 301,
		302, 5, 114, 0, 0, 302, 303, 5, 110, 0, 0, 303, 102, 1, 0, 0, 0, 304, 305,
		5, 99, 0, 0, 305, 306, 5, 111, 0, 0, 306, 307, 5, 110, 0, 0, 307, 308,
		5, 116, 0, 0, 308, 309, 5, 105, 0, 0, 309, 310, 5, 110, 0, 0, 310, 311,
		5, 117, 0, 0, 311, 312, 5, 101, 0, 0, 312, 104, 1, 0, 0, 0, 313, 314, 5,
		102, 0, 0, 314, 315, 5, 111, 0, 0, 315, 316, 5, 114, 0, 0, 316, 106, 1,
		0, 0, 0, 317, 318, 5, 102, 0, 0, 318, 319, 5, 111, 0, 0, 319, 320, 5, 114,
		0, 0, 320, 321, 5, 101, 0, 0, 321, 322, 5, 97, 0, 0, 322, 323, 5, 99, 0,
		0, 323, 324, 5, 104, 0, 0, 324, 108, 1, 0, 0, 0, 325, 326, 5, 105, 0, 0,
		326, 327, 5, 102, 0, 0, 327, 110, 1, 0, 0, 0, 328, 329, 5, 102, 0, 0, 329,
		330, 5, 117, 0, 0, 330, 331, 5, 110, 0, 0, 331, 332, 5, 99, 0, 0, 332,
		333, 5, 116, 0, 0, 333, 334, 5, 105, 0, 0, 334, 335, 5, 111, 0, 0, 335,
		336, 5, 110, 0, 0, 336, 112, 1, 0, 0, 0, 337, 338, 5, 114, 0, 0, 338, 339,
		5, 101, 0, 0, 339, 340, 5, 112, 0, 0, 340, 341, 5, 101, 0, 0, 341, 342,
		5, 97, 0, 0, 342, 343, 5, 116, 0, 0, 343, 114, 1, 0, 0, 0, 344, 345, 5,
		119, 0, 0, 345, 346, 5, 104, 0, 0, 346, 347, 5, 105, 0, 0, 347, 348, 5,
		108, 0, 0, 348, 349, 5, 101, 0, 0, 349, 116, 1, 0, 0, 0, 350, 351, 5, 117,
		0, 0, 351, 352, 5, 110, 0, 0, 352, 353, 5, 116, 0, 0, 353, 354, 5, 105,
		0, 0, 354, 355, 5, 108, 0, 0, 355, 118, 1, 0, 0, 0, 356, 357, 5, 101, 0,
		0, 357, 358, 5, 120, 0, 0, 358, 359, 5, 105, 0, 0, 359, 360, 5, 116, 0,
		0, 360, 120, 1, 0, 0, 0, 361, 365, 5, 34, 0, 0, 362, 364, 8, 1, 0, 0, 363,
		362, 1, 0, 0, 0, 364, 367, 1, 0, 0, 0, 365, 363, 1, 0, 0, 0, 365, 366,
		1, 0, 0, 0, 366, 368, 1, 0, 0, 0, 367, 365, 1, 0, 0, 0, 368, 378, 5, 34,
		0, 0, 369, 373, 5, 39, 0, 0, 370, 372, 3, 123, 61, 0, 371, 370, 1, 0, 0,
		0, 372, 375, 1, 0, 0, 0, 373, 371, 1, 0, 0, 0, 373, 374, 1, 0, 0, 0, 374,
		376, 1, 0, 0, 0, 375, 373, 1, 0, 0, 0, 376, 378, 5, 39, 0, 0, 377, 361,
		1, 0, 0, 0, 377, 369, 1, 0, 0, 0, 378, 122, 1, 0, 0, 0, 379, 383, 8, 2,
		0, 0, 380, 381, 5, 92, 0, 0, 381, 383, 9, 0, 0, 0, 382, 379, 1, 0, 0, 0,
		382, 380, 1, 0, 0, 0, 383, 124, 1, 0, 0, 0, 384, 385, 5, 116, 0, 0, 385,
		386, 5, 114, 0, 0, 386, 387, 5, 117, 0, 0, 387, 403, 5, 101, 0, 0, 388,
		389, 5, 102, 0, 0, 389, 390, 5, 97, 0, 0, 390, 391, 5, 108, 0, 0, 391,
		392, 5, 115, 0, 0, 392, 403, 5, 101, 0, 0, 393, 394, 5, 70, 0, 0, 394,
		395, 5, 65, 0, 0, 395, 396, 5, 76, 0, 0, 396, 397, 5, 83, 0, 0, 397, 403,
		5, 69, 0, 0, 398, 399, 5, 84, 0, 0, 399, 400, 5, 82, 0, 0, 400, 401, 5,
		85, 0, 0, 401, 403, 5, 69, 0, 0, 402, 384, 1, 0, 0, 0, 402, 388, 1, 0,
		0, 0, 402, 393, 1, 0, 0, 0, 402, 398, 1, 0, 0, 0, 403, 126, 1, 0, 0, 0,
		404, 406, 7, 3, 0, 0, 405, 404, 1, 0, 0, 0, 406, 407, 1, 0, 0, 0, 407,
		405, 1, 0, 0, 0, 407, 408, 1, 0, 0, 0, 408, 128, 1, 0, 0, 0, 409, 410,
		3, 127, 63, 0, 410, 414, 5, 46, 0, 0, 411, 413, 7, 3, 0, 0, 412, 411, 1,
		0, 0, 0, 413, 416, 1, 0, 0, 0, 414, 412, 1, 0, 0, 0, 414, 415, 1, 0, 0,
		0, 415, 130, 1, 0, 0, 0, 416, 414, 1, 0, 0, 0, 417, 419, 7, 3, 0, 0, 418,
		417, 1, 0, 0, 0, 419, 422, 1, 0, 0, 0, 420, 418, 1, 0, 0, 0, 420, 421,
		1, 0, 0, 0, 421, 423, 1, 0, 0, 0, 422, 420, 1, 0, 0, 0, 423, 427, 5, 46,
		0, 0, 424, 426, 7, 3, 0, 0, 425, 424, 1, 0, 0, 0, 426, 429, 1, 0, 0, 0,
		427, 425, 1, 0, 0, 0, 427, 428, 1, 0, 0, 0, 428, 430, 1, 0, 0, 0, 429,
		427, 1, 0, 0, 0, 430, 434, 5, 46, 0, 0, 431, 433, 7, 3, 0, 0, 432, 431,
		1, 0, 0, 0, 433, 436, 1, 0, 0, 0, 434, 432, 1, 0, 0, 0, 434, 435, 1, 0,
		0, 0, 435, 437, 1, 0, 0, 0, 436, 434, 1, 0, 0, 0, 437, 441, 5, 46, 0, 0,
		438, 440, 7, 3, 0, 0, 439, 438, 1, 0, 0, 0, 440, 443, 1, 0, 0, 0, 441,
		439, 1, 0, 0, 0, 441, 442, 1, 0, 0, 0, 442, 132, 1, 0, 0, 0, 443, 441,
		1, 0, 0, 0, 444, 445, 5, 48, 0, 0, 445, 447, 7, 4, 0, 0, 446, 448, 7, 5,
		0, 0, 447, 446, 1, 0, 0, 0, 448, 449, 1, 0, 0, 0, 449, 447, 1, 0, 0, 0,
		449, 450, 1, 0, 0, 0, 450, 134, 1, 0, 0, 0, 451, 452, 5, 78, 0, 0, 452,
		453, 5, 85, 0, 0, 453, 454, 5, 76, 0, 0, 454, 455, 5, 76, 0, 0, 455, 136,
		1, 0, 0, 0, 456, 460, 7, 6, 0, 0, 457, 459, 7, 7, 0, 0, 458, 457, 1, 0,
		0, 0, 459, 462, 1, 0, 0, 0, 460, 458, 1, 0, 0, 0, 460, 461, 1, 0, 0, 0,
		461, 138, 1, 0, 0, 0, 462, 460, 1, 0, 0, 0, 463, 465, 7, 8, 0, 0, 464,
		463, 1, 0, 0, 0, 465, 466, 1, 0, 0, 0, 466, 464, 1, 0, 0, 0, 466, 467,
		1, 0, 0, 0, 467, 468, 1, 0, 0, 0, 468, 469, 6, 69, 0, 0, 469, 140, 1, 0,
		0, 0, 470, 471, 7, 0, 0, 0, 471, 472, 1, 0, 0, 0, 472, 473, 6, 70, 0, 0,
		473, 142, 1, 0, 0, 0, 16, 0, 147, 365, 373, 377, 382, 402, 407, 414, 420,
		427, 434, 441, 449, 460, 466, 1, 0, 1, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// NaslLexerInit initializes any static state used to implement NaslLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewNaslLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func NaslLexerInit() {
	staticData := &nasllexerLexerStaticData
	staticData.once.Do(nasllexerLexerInit)
}

// NewNaslLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewNaslLexer(input antlr.CharStream) *NaslLexer {
	NaslLexerInit()
	l := new(NaslLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &nasllexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "NaslLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// NaslLexer tokens.
const (
	NaslLexerSingleLineComment    = 1
	NaslLexerOpenBracket          = 2
	NaslLexerCloseBracket         = 3
	NaslLexerOpenParen            = 4
	NaslLexerCloseParen           = 5
	NaslLexerOpenBrace            = 6
	NaslLexerCloseBrace           = 7
	NaslLexerSemiColon            = 8
	NaslLexerComma                = 9
	NaslLexerAssign               = 10
	NaslLexerColon                = 11
	NaslLexerDot                  = 12
	NaslLexerPlusPlus             = 13
	NaslLexerMinusMinus           = 14
	NaslLexerPlus                 = 15
	NaslLexerMinus                = 16
	NaslLexerBitNot               = 17
	NaslLexerBitAnd               = 18
	NaslLexerBitXOr               = 19
	NaslLexerBitOr                = 20
	NaslLexerRightShiftArithmetic = 21
	NaslLexerLeftShiftArithmetic  = 22
	NaslLexerRightShiftLogical    = 23
	NaslLexerNot                  = 24
	NaslLexerMultiply             = 25
	NaslLexerPow                  = 26
	NaslLexerDivide               = 27
	NaslLexerModulus              = 28
	NaslLexerLessThan             = 29
	NaslLexerMoreThan             = 30
	NaslLexerLessThanEquals       = 31
	NaslLexerGreaterThanEquals    = 32
	NaslLexerEquals_              = 33
	NaslLexerEqualsRe             = 34
	NaslLexerNotEquals            = 35
	NaslLexerNotLong              = 36
	NaslLexerMTNotLT              = 37
	NaslLexerMTLT                 = 38
	NaslLexerAnd                  = 39
	NaslLexerOr                   = 40
	NaslLexerMultiplyAssign       = 41
	NaslLexerDivideAssign         = 42
	NaslLexerModulusAssign        = 43
	NaslLexerPlusAssign           = 44
	NaslLexerMinusAssign          = 45
	NaslLexerX                    = 46
	NaslLexerBreak                = 47
	NaslLexerLocalVar             = 48
	NaslLexerGlobalVar            = 49
	NaslLexerElse                 = 50
	NaslLexerReturn               = 51
	NaslLexerContinue             = 52
	NaslLexerFor                  = 53
	NaslLexerForEach              = 54
	NaslLexerIf                   = 55
	NaslLexerFunction_            = 56
	NaslLexerRepeat               = 57
	NaslLexerWhile                = 58
	NaslLexerUntil                = 59
	NaslLexerExit                 = 60
	NaslLexerStringLiteral        = 61
	NaslLexerBooleanLiteral       = 62
	NaslLexerIntegerLiteral       = 63
	NaslLexerFloatLiteral         = 64
	NaslLexerIpLiteral            = 65
	NaslLexerHexLiteral           = 66
	NaslLexerNULLLiteral          = 67
	NaslLexerIdentifier           = 68
	NaslLexerWhiteSpaces          = 69
	NaslLexerLineTerminator       = 70
)
