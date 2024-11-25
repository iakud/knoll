// Code generated from kds.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // kds

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type kdsParser struct {
	*antlr.BaseParser
}

var KdsParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func kdsParserInit() {
	staticData := &KdsParserStaticData
	staticData.LiteralNames = []string{
		"", "'import'", "'weak'", "'public'", "'package'", "'option'", "'repeated'",
		"'map'", "'int32'", "'int64'", "'uint32'", "'uint64'", "'sint32'", "'sint64'",
		"'fixed32'", "'fixed64'", "'sfixed32'", "'sfixed64'", "'bool'", "'string'",
		"'double'", "'float'", "'bytes'", "'enum'", "'entity'", "'component'",
		"'\"proto3\"'", "''proto3''", "';'", "'='", "'('", "')'", "'['", "']'",
		"'{'", "'}'", "'<'", "'>'", "'.'", "','", "':'", "'+'", "'-'",
	}
	staticData.SymbolicNames = []string{
		"", "IMPORT", "WEAK", "PUBLIC", "PACKAGE", "OPTION", "REPEATED", "MAP",
		"INT32", "INT64", "UINT32", "UINT64", "SINT32", "SINT64", "FIXED32",
		"FIXED64", "SFIXED32", "SFIXED64", "BOOL", "STRING", "DOUBLE", "FLOAT",
		"BYTES", "ENUM", "ENTITY", "COMPONENT", "PROTO3_LIT_SINGLE", "PROTO3_LIT_DOBULE",
		"SEMI", "EQ", "LP", "RP", "LB", "RB", "LC", "RC", "LT", "GT", "DOT",
		"COMMA", "COLON", "PLUS", "MINUS", "STR_LIT", "BOOL_LIT", "FLOAT_LIT",
		"INT_LIT", "IDENTIFIER", "WS", "LINE_COMMENT", "COMMENT",
	}
	staticData.RuleNames = []string{
		"kds", "importStatement", "packageStatement", "optionStatement", "optionName",
		"field", "fieldLabel", "fieldNumber", "mapField", "keyType", "type_",
		"topLevelDef", "enumDef", "enumBody", "enumElement", "enumField", "entityDef",
		"entityName", "entityBody", "entityElement", "componentDef", "componentName",
		"componentBody", "componentElement", "constant", "blockLit", "emptyStatement_",
		"ident", "fullIdent", "fieldName", "messageName", "enumName", "mapName",
		"messageType", "enumType", "intLit", "strLit", "boolLit", "floatLit",
		"keywords",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 50, 317, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		5, 0, 86, 8, 0, 10, 0, 12, 0, 89, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 95,
		8, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 116, 8, 4, 3, 4,
		118, 8, 4, 1, 5, 3, 5, 121, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		3, 10, 163, 8, 10, 1, 11, 1, 11, 1, 11, 3, 11, 168, 8, 11, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 13, 1, 13, 5, 13, 176, 8, 13, 10, 13, 12, 13, 179, 9,
		13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 3, 15, 188, 8, 15,
		1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1,
		18, 5, 18, 201, 8, 18, 10, 18, 12, 18, 204, 9, 18, 1, 18, 1, 18, 1, 19,
		1, 19, 1, 19, 3, 19, 211, 8, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1,
		21, 1, 22, 1, 22, 5, 22, 221, 8, 22, 10, 22, 12, 22, 224, 9, 22, 1, 22,
		1, 22, 1, 23, 1, 23, 1, 23, 3, 23, 231, 8, 23, 1, 24, 1, 24, 3, 24, 235,
		8, 24, 1, 24, 1, 24, 3, 24, 239, 8, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3,
		24, 245, 8, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 5, 25, 252, 8, 25, 10,
		25, 12, 25, 255, 9, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 3, 27,
		263, 8, 27, 1, 28, 1, 28, 1, 28, 5, 28, 268, 8, 28, 10, 28, 12, 28, 271,
		9, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 3,
		33, 282, 8, 33, 1, 33, 1, 33, 1, 33, 5, 33, 287, 8, 33, 10, 33, 12, 33,
		290, 9, 33, 1, 33, 1, 33, 1, 34, 3, 34, 295, 8, 34, 1, 34, 1, 34, 1, 34,
		5, 34, 300, 8, 34, 10, 34, 12, 34, 303, 9, 34, 1, 34, 1, 34, 1, 35, 1,
		35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 0, 0,
		40, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70,
		72, 74, 76, 78, 0, 5, 1, 0, 2, 3, 1, 0, 8, 19, 1, 0, 41, 42, 2, 0, 26,
		27, 43, 43, 3, 0, 4, 4, 7, 25, 44, 44, 324, 0, 80, 1, 0, 0, 0, 2, 92, 1,
		0, 0, 0, 4, 99, 1, 0, 0, 0, 6, 103, 1, 0, 0, 0, 8, 117, 1, 0, 0, 0, 10,
		120, 1, 0, 0, 0, 12, 128, 1, 0, 0, 0, 14, 130, 1, 0, 0, 0, 16, 132, 1,
		0, 0, 0, 18, 143, 1, 0, 0, 0, 20, 162, 1, 0, 0, 0, 22, 167, 1, 0, 0, 0,
		24, 169, 1, 0, 0, 0, 26, 173, 1, 0, 0, 0, 28, 182, 1, 0, 0, 0, 30, 184,
		1, 0, 0, 0, 32, 192, 1, 0, 0, 0, 34, 196, 1, 0, 0, 0, 36, 198, 1, 0, 0,
		0, 38, 210, 1, 0, 0, 0, 40, 212, 1, 0, 0, 0, 42, 216, 1, 0, 0, 0, 44, 218,
		1, 0, 0, 0, 46, 230, 1, 0, 0, 0, 48, 244, 1, 0, 0, 0, 50, 246, 1, 0, 0,
		0, 52, 258, 1, 0, 0, 0, 54, 262, 1, 0, 0, 0, 56, 264, 1, 0, 0, 0, 58, 272,
		1, 0, 0, 0, 60, 274, 1, 0, 0, 0, 62, 276, 1, 0, 0, 0, 64, 278, 1, 0, 0,
		0, 66, 281, 1, 0, 0, 0, 68, 294, 1, 0, 0, 0, 70, 306, 1, 0, 0, 0, 72, 308,
		1, 0, 0, 0, 74, 310, 1, 0, 0, 0, 76, 312, 1, 0, 0, 0, 78, 314, 1, 0, 0,
		0, 80, 87, 3, 4, 2, 0, 81, 86, 3, 2, 1, 0, 82, 86, 3, 6, 3, 0, 83, 86,
		3, 22, 11, 0, 84, 86, 3, 52, 26, 0, 85, 81, 1, 0, 0, 0, 85, 82, 1, 0, 0,
		0, 85, 83, 1, 0, 0, 0, 85, 84, 1, 0, 0, 0, 86, 89, 1, 0, 0, 0, 87, 85,
		1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 90, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0,
		90, 91, 5, 0, 0, 1, 91, 1, 1, 0, 0, 0, 92, 94, 5, 1, 0, 0, 93, 95, 7, 0,
		0, 0, 94, 93, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 97,
		3, 72, 36, 0, 97, 98, 5, 28, 0, 0, 98, 3, 1, 0, 0, 0, 99, 100, 5, 4, 0,
		0, 100, 101, 3, 56, 28, 0, 101, 102, 5, 28, 0, 0, 102, 5, 1, 0, 0, 0, 103,
		104, 5, 5, 0, 0, 104, 105, 3, 8, 4, 0, 105, 106, 5, 29, 0, 0, 106, 107,
		3, 48, 24, 0, 107, 108, 5, 28, 0, 0, 108, 7, 1, 0, 0, 0, 109, 118, 3, 56,
		28, 0, 110, 111, 5, 30, 0, 0, 111, 112, 3, 56, 28, 0, 112, 115, 5, 31,
		0, 0, 113, 114, 5, 38, 0, 0, 114, 116, 3, 56, 28, 0, 115, 113, 1, 0, 0,
		0, 115, 116, 1, 0, 0, 0, 116, 118, 1, 0, 0, 0, 117, 109, 1, 0, 0, 0, 117,
		110, 1, 0, 0, 0, 118, 9, 1, 0, 0, 0, 119, 121, 3, 12, 6, 0, 120, 119, 1,
		0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 123, 3, 20, 10,
		0, 123, 124, 3, 58, 29, 0, 124, 125, 5, 29, 0, 0, 125, 126, 3, 14, 7, 0,
		126, 127, 5, 28, 0, 0, 127, 11, 1, 0, 0, 0, 128, 129, 5, 6, 0, 0, 129,
		13, 1, 0, 0, 0, 130, 131, 3, 70, 35, 0, 131, 15, 1, 0, 0, 0, 132, 133,
		5, 7, 0, 0, 133, 134, 5, 36, 0, 0, 134, 135, 3, 18, 9, 0, 135, 136, 5,
		39, 0, 0, 136, 137, 3, 20, 10, 0, 137, 138, 5, 37, 0, 0, 138, 139, 3, 64,
		32, 0, 139, 140, 5, 29, 0, 0, 140, 141, 3, 14, 7, 0, 141, 142, 5, 28, 0,
		0, 142, 17, 1, 0, 0, 0, 143, 144, 7, 1, 0, 0, 144, 19, 1, 0, 0, 0, 145,
		163, 5, 20, 0, 0, 146, 163, 5, 21, 0, 0, 147, 163, 5, 8, 0, 0, 148, 163,
		5, 9, 0, 0, 149, 163, 5, 10, 0, 0, 150, 163, 5, 11, 0, 0, 151, 163, 5,
		12, 0, 0, 152, 163, 5, 13, 0, 0, 153, 163, 5, 14, 0, 0, 154, 163, 5, 15,
		0, 0, 155, 163, 5, 16, 0, 0, 156, 163, 5, 17, 0, 0, 157, 163, 5, 18, 0,
		0, 158, 163, 5, 19, 0, 0, 159, 163, 5, 22, 0, 0, 160, 163, 3, 66, 33, 0,
		161, 163, 3, 68, 34, 0, 162, 145, 1, 0, 0, 0, 162, 146, 1, 0, 0, 0, 162,
		147, 1, 0, 0, 0, 162, 148, 1, 0, 0, 0, 162, 149, 1, 0, 0, 0, 162, 150,
		1, 0, 0, 0, 162, 151, 1, 0, 0, 0, 162, 152, 1, 0, 0, 0, 162, 153, 1, 0,
		0, 0, 162, 154, 1, 0, 0, 0, 162, 155, 1, 0, 0, 0, 162, 156, 1, 0, 0, 0,
		162, 157, 1, 0, 0, 0, 162, 158, 1, 0, 0, 0, 162, 159, 1, 0, 0, 0, 162,
		160, 1, 0, 0, 0, 162, 161, 1, 0, 0, 0, 163, 21, 1, 0, 0, 0, 164, 168, 3,
		32, 16, 0, 165, 168, 3, 40, 20, 0, 166, 168, 3, 24, 12, 0, 167, 164, 1,
		0, 0, 0, 167, 165, 1, 0, 0, 0, 167, 166, 1, 0, 0, 0, 168, 23, 1, 0, 0,
		0, 169, 170, 5, 23, 0, 0, 170, 171, 3, 62, 31, 0, 171, 172, 3, 26, 13,
		0, 172, 25, 1, 0, 0, 0, 173, 177, 5, 34, 0, 0, 174, 176, 3, 28, 14, 0,
		175, 174, 1, 0, 0, 0, 176, 179, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 177,
		178, 1, 0, 0, 0, 178, 180, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 180, 181,
		5, 35, 0, 0, 181, 27, 1, 0, 0, 0, 182, 183, 3, 30, 15, 0, 183, 29, 1, 0,
		0, 0, 184, 185, 3, 54, 27, 0, 185, 187, 5, 29, 0, 0, 186, 188, 5, 42, 0,
		0, 187, 186, 1, 0, 0, 0, 187, 188, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189,
		190, 3, 70, 35, 0, 190, 191, 5, 28, 0, 0, 191, 31, 1, 0, 0, 0, 192, 193,
		5, 24, 0, 0, 193, 194, 3, 34, 17, 0, 194, 195, 3, 36, 18, 0, 195, 33, 1,
		0, 0, 0, 196, 197, 3, 54, 27, 0, 197, 35, 1, 0, 0, 0, 198, 202, 5, 34,
		0, 0, 199, 201, 3, 38, 19, 0, 200, 199, 1, 0, 0, 0, 201, 204, 1, 0, 0,
		0, 202, 200, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 205, 1, 0, 0, 0, 204,
		202, 1, 0, 0, 0, 205, 206, 5, 35, 0, 0, 206, 37, 1, 0, 0, 0, 207, 211,
		3, 10, 5, 0, 208, 211, 3, 16, 8, 0, 209, 211, 3, 52, 26, 0, 210, 207, 1,
		0, 0, 0, 210, 208, 1, 0, 0, 0, 210, 209, 1, 0, 0, 0, 211, 39, 1, 0, 0,
		0, 212, 213, 5, 25, 0, 0, 213, 214, 3, 42, 21, 0, 214, 215, 3, 44, 22,
		0, 215, 41, 1, 0, 0, 0, 216, 217, 3, 54, 27, 0, 217, 43, 1, 0, 0, 0, 218,
		222, 5, 34, 0, 0, 219, 221, 3, 46, 23, 0, 220, 219, 1, 0, 0, 0, 221, 224,
		1, 0, 0, 0, 222, 220, 1, 0, 0, 0, 222, 223, 1, 0, 0, 0, 223, 225, 1, 0,
		0, 0, 224, 222, 1, 0, 0, 0, 225, 226, 5, 35, 0, 0, 226, 45, 1, 0, 0, 0,
		227, 231, 3, 10, 5, 0, 228, 231, 3, 16, 8, 0, 229, 231, 3, 52, 26, 0, 230,
		227, 1, 0, 0, 0, 230, 228, 1, 0, 0, 0, 230, 229, 1, 0, 0, 0, 231, 47, 1,
		0, 0, 0, 232, 245, 3, 56, 28, 0, 233, 235, 7, 2, 0, 0, 234, 233, 1, 0,
		0, 0, 234, 235, 1, 0, 0, 0, 235, 236, 1, 0, 0, 0, 236, 245, 3, 70, 35,
		0, 237, 239, 7, 2, 0, 0, 238, 237, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239,
		240, 1, 0, 0, 0, 240, 245, 3, 76, 38, 0, 241, 245, 3, 72, 36, 0, 242, 245,
		3, 74, 37, 0, 243, 245, 3, 50, 25, 0, 244, 232, 1, 0, 0, 0, 244, 234, 1,
		0, 0, 0, 244, 238, 1, 0, 0, 0, 244, 241, 1, 0, 0, 0, 244, 242, 1, 0, 0,
		0, 244, 243, 1, 0, 0, 0, 245, 49, 1, 0, 0, 0, 246, 253, 5, 34, 0, 0, 247,
		248, 3, 54, 27, 0, 248, 249, 5, 40, 0, 0, 249, 250, 3, 48, 24, 0, 250,
		252, 1, 0, 0, 0, 251, 247, 1, 0, 0, 0, 252, 255, 1, 0, 0, 0, 253, 251,
		1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254, 256, 1, 0, 0, 0, 255, 253, 1, 0,
		0, 0, 256, 257, 5, 35, 0, 0, 257, 51, 1, 0, 0, 0, 258, 259, 5, 28, 0, 0,
		259, 53, 1, 0, 0, 0, 260, 263, 5, 47, 0, 0, 261, 263, 3, 78, 39, 0, 262,
		260, 1, 0, 0, 0, 262, 261, 1, 0, 0, 0, 263, 55, 1, 0, 0, 0, 264, 269, 3,
		54, 27, 0, 265, 266, 5, 38, 0, 0, 266, 268, 3, 54, 27, 0, 267, 265, 1,
		0, 0, 0, 268, 271, 1, 0, 0, 0, 269, 267, 1, 0, 0, 0, 269, 270, 1, 0, 0,
		0, 270, 57, 1, 0, 0, 0, 271, 269, 1, 0, 0, 0, 272, 273, 3, 54, 27, 0, 273,
		59, 1, 0, 0, 0, 274, 275, 3, 54, 27, 0, 275, 61, 1, 0, 0, 0, 276, 277,
		3, 54, 27, 0, 277, 63, 1, 0, 0, 0, 278, 279, 3, 54, 27, 0, 279, 65, 1,
		0, 0, 0, 280, 282, 5, 38, 0, 0, 281, 280, 1, 0, 0, 0, 281, 282, 1, 0, 0,
		0, 282, 288, 1, 0, 0, 0, 283, 284, 3, 54, 27, 0, 284, 285, 5, 38, 0, 0,
		285, 287, 1, 0, 0, 0, 286, 283, 1, 0, 0, 0, 287, 290, 1, 0, 0, 0, 288,
		286, 1, 0, 0, 0, 288, 289, 1, 0, 0, 0, 289, 291, 1, 0, 0, 0, 290, 288,
		1, 0, 0, 0, 291, 292, 3, 60, 30, 0, 292, 67, 1, 0, 0, 0, 293, 295, 5, 38,
		0, 0, 294, 293, 1, 0, 0, 0, 294, 295, 1, 0, 0, 0, 295, 301, 1, 0, 0, 0,
		296, 297, 3, 54, 27, 0, 297, 298, 5, 38, 0, 0, 298, 300, 1, 0, 0, 0, 299,
		296, 1, 0, 0, 0, 300, 303, 1, 0, 0, 0, 301, 299, 1, 0, 0, 0, 301, 302,
		1, 0, 0, 0, 302, 304, 1, 0, 0, 0, 303, 301, 1, 0, 0, 0, 304, 305, 3, 62,
		31, 0, 305, 69, 1, 0, 0, 0, 306, 307, 5, 46, 0, 0, 307, 71, 1, 0, 0, 0,
		308, 309, 7, 3, 0, 0, 309, 73, 1, 0, 0, 0, 310, 311, 5, 44, 0, 0, 311,
		75, 1, 0, 0, 0, 312, 313, 5, 45, 0, 0, 313, 77, 1, 0, 0, 0, 314, 315, 7,
		4, 0, 0, 315, 79, 1, 0, 0, 0, 24, 85, 87, 94, 115, 117, 120, 162, 167,
		177, 187, 202, 210, 222, 230, 234, 238, 244, 253, 262, 269, 281, 288, 294,
		301,
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

// kdsParserInit initializes any static state used to implement kdsParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewkdsParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func KdsParserInit() {
	staticData := &KdsParserStaticData
	staticData.once.Do(kdsParserInit)
}

// NewkdsParser produces a new parser instance for the optional input antlr.TokenStream.
func NewkdsParser(input antlr.TokenStream) *kdsParser {
	KdsParserInit()
	this := new(kdsParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &KdsParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "kds.g4"

	return this
}

// kdsParser tokens.
const (
	kdsParserEOF               = antlr.TokenEOF
	kdsParserIMPORT            = 1
	kdsParserWEAK              = 2
	kdsParserPUBLIC            = 3
	kdsParserPACKAGE           = 4
	kdsParserOPTION            = 5
	kdsParserREPEATED          = 6
	kdsParserMAP               = 7
	kdsParserINT32             = 8
	kdsParserINT64             = 9
	kdsParserUINT32            = 10
	kdsParserUINT64            = 11
	kdsParserSINT32            = 12
	kdsParserSINT64            = 13
	kdsParserFIXED32           = 14
	kdsParserFIXED64           = 15
	kdsParserSFIXED32          = 16
	kdsParserSFIXED64          = 17
	kdsParserBOOL              = 18
	kdsParserSTRING            = 19
	kdsParserDOUBLE            = 20
	kdsParserFLOAT             = 21
	kdsParserBYTES             = 22
	kdsParserENUM              = 23
	kdsParserENTITY            = 24
	kdsParserCOMPONENT         = 25
	kdsParserPROTO3_LIT_SINGLE = 26
	kdsParserPROTO3_LIT_DOBULE = 27
	kdsParserSEMI              = 28
	kdsParserEQ                = 29
	kdsParserLP                = 30
	kdsParserRP                = 31
	kdsParserLB                = 32
	kdsParserRB                = 33
	kdsParserLC                = 34
	kdsParserRC                = 35
	kdsParserLT                = 36
	kdsParserGT                = 37
	kdsParserDOT               = 38
	kdsParserCOMMA             = 39
	kdsParserCOLON             = 40
	kdsParserPLUS              = 41
	kdsParserMINUS             = 42
	kdsParserSTR_LIT           = 43
	kdsParserBOOL_LIT          = 44
	kdsParserFLOAT_LIT         = 45
	kdsParserINT_LIT           = 46
	kdsParserIDENTIFIER        = 47
	kdsParserWS                = 48
	kdsParserLINE_COMMENT      = 49
	kdsParserCOMMENT           = 50
)

// kdsParser rules.
const (
	kdsParserRULE_kds              = 0
	kdsParserRULE_importStatement  = 1
	kdsParserRULE_packageStatement = 2
	kdsParserRULE_optionStatement  = 3
	kdsParserRULE_optionName       = 4
	kdsParserRULE_field            = 5
	kdsParserRULE_fieldLabel       = 6
	kdsParserRULE_fieldNumber      = 7
	kdsParserRULE_mapField         = 8
	kdsParserRULE_keyType          = 9
	kdsParserRULE_type_            = 10
	kdsParserRULE_topLevelDef      = 11
	kdsParserRULE_enumDef          = 12
	kdsParserRULE_enumBody         = 13
	kdsParserRULE_enumElement      = 14
	kdsParserRULE_enumField        = 15
	kdsParserRULE_entityDef        = 16
	kdsParserRULE_entityName       = 17
	kdsParserRULE_entityBody       = 18
	kdsParserRULE_entityElement    = 19
	kdsParserRULE_componentDef     = 20
	kdsParserRULE_componentName    = 21
	kdsParserRULE_componentBody    = 22
	kdsParserRULE_componentElement = 23
	kdsParserRULE_constant         = 24
	kdsParserRULE_blockLit         = 25
	kdsParserRULE_emptyStatement_  = 26
	kdsParserRULE_ident            = 27
	kdsParserRULE_fullIdent        = 28
	kdsParserRULE_fieldName        = 29
	kdsParserRULE_messageName      = 30
	kdsParserRULE_enumName         = 31
	kdsParserRULE_mapName          = 32
	kdsParserRULE_messageType      = 33
	kdsParserRULE_enumType         = 34
	kdsParserRULE_intLit           = 35
	kdsParserRULE_strLit           = 36
	kdsParserRULE_boolLit          = 37
	kdsParserRULE_floatLit         = 38
	kdsParserRULE_keywords         = 39
)

// IKdsContext is an interface to support dynamic dispatch.
type IKdsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PackageStatement() IPackageStatementContext
	EOF() antlr.TerminalNode
	AllImportStatement() []IImportStatementContext
	ImportStatement(i int) IImportStatementContext
	AllOptionStatement() []IOptionStatementContext
	OptionStatement(i int) IOptionStatementContext
	AllTopLevelDef() []ITopLevelDefContext
	TopLevelDef(i int) ITopLevelDefContext
	AllEmptyStatement_() []IEmptyStatement_Context
	EmptyStatement_(i int) IEmptyStatement_Context

	// IsKdsContext differentiates from other interfaces.
	IsKdsContext()
}

type KdsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKdsContext() *KdsContext {
	var p = new(KdsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_kds
	return p
}

func InitEmptyKdsContext(p *KdsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_kds
}

func (*KdsContext) IsKdsContext() {}

func NewKdsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KdsContext {
	var p = new(KdsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_kds

	return p
}

func (s *KdsContext) GetParser() antlr.Parser { return s.parser }

func (s *KdsContext) PackageStatement() IPackageStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPackageStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPackageStatementContext)
}

func (s *KdsContext) EOF() antlr.TerminalNode {
	return s.GetToken(kdsParserEOF, 0)
}

func (s *KdsContext) AllImportStatement() []IImportStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IImportStatementContext); ok {
			len++
		}
	}

	tst := make([]IImportStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IImportStatementContext); ok {
			tst[i] = t.(IImportStatementContext)
			i++
		}
	}

	return tst
}

func (s *KdsContext) ImportStatement(i int) IImportStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportStatementContext)
}

func (s *KdsContext) AllOptionStatement() []IOptionStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOptionStatementContext); ok {
			len++
		}
	}

	tst := make([]IOptionStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOptionStatementContext); ok {
			tst[i] = t.(IOptionStatementContext)
			i++
		}
	}

	return tst
}

func (s *KdsContext) OptionStatement(i int) IOptionStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionStatementContext)
}

func (s *KdsContext) AllTopLevelDef() []ITopLevelDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITopLevelDefContext); ok {
			len++
		}
	}

	tst := make([]ITopLevelDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITopLevelDefContext); ok {
			tst[i] = t.(ITopLevelDefContext)
			i++
		}
	}

	return tst
}

func (s *KdsContext) TopLevelDef(i int) ITopLevelDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITopLevelDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITopLevelDefContext)
}

func (s *KdsContext) AllEmptyStatement_() []IEmptyStatement_Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
			len++
		}
	}

	tst := make([]IEmptyStatement_Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEmptyStatement_Context); ok {
			tst[i] = t.(IEmptyStatement_Context)
			i++
		}
	}

	return tst
}

func (s *KdsContext) EmptyStatement_(i int) IEmptyStatement_Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmptyStatement_Context)
}

func (s *KdsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KdsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KdsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterKds(s)
	}
}

func (s *KdsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitKds(s)
	}
}

func (p *kdsParser) Kds() (localctx IKdsContext) {
	localctx = NewKdsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, kdsParserRULE_kds)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.PackageStatement()
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&327155746) != 0 {
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case kdsParserIMPORT:
			{
				p.SetState(81)
				p.ImportStatement()
			}

		case kdsParserOPTION:
			{
				p.SetState(82)
				p.OptionStatement()
			}

		case kdsParserENUM, kdsParserENTITY, kdsParserCOMPONENT:
			{
				p.SetState(83)
				p.TopLevelDef()
			}

		case kdsParserSEMI:
			{
				p.SetState(84)
				p.EmptyStatement_()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(90)
		p.Match(kdsParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImportStatementContext is an interface to support dynamic dispatch.
type IImportStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IMPORT() antlr.TerminalNode
	StrLit() IStrLitContext
	SEMI() antlr.TerminalNode
	WEAK() antlr.TerminalNode
	PUBLIC() antlr.TerminalNode

	// IsImportStatementContext differentiates from other interfaces.
	IsImportStatementContext()
}

type ImportStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportStatementContext() *ImportStatementContext {
	var p = new(ImportStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_importStatement
	return p
}

func InitEmptyImportStatementContext(p *ImportStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_importStatement
}

func (*ImportStatementContext) IsImportStatementContext() {}

func NewImportStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportStatementContext {
	var p = new(ImportStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_importStatement

	return p
}

func (s *ImportStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportStatementContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(kdsParserIMPORT, 0)
}

func (s *ImportStatementContext) StrLit() IStrLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStrLitContext)
}

func (s *ImportStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(kdsParserSEMI, 0)
}

func (s *ImportStatementContext) WEAK() antlr.TerminalNode {
	return s.GetToken(kdsParserWEAK, 0)
}

func (s *ImportStatementContext) PUBLIC() antlr.TerminalNode {
	return s.GetToken(kdsParserPUBLIC, 0)
}

func (s *ImportStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterImportStatement(s)
	}
}

func (s *ImportStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitImportStatement(s)
	}
}

func (p *kdsParser) ImportStatement() (localctx IImportStatementContext) {
	localctx = NewImportStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, kdsParserRULE_importStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Match(kdsParserIMPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == kdsParserWEAK || _la == kdsParserPUBLIC {
		{
			p.SetState(93)
			_la = p.GetTokenStream().LA(1)

			if !(_la == kdsParserWEAK || _la == kdsParserPUBLIC) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(96)
		p.StrLit()
	}
	{
		p.SetState(97)
		p.Match(kdsParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPackageStatementContext is an interface to support dynamic dispatch.
type IPackageStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PACKAGE() antlr.TerminalNode
	FullIdent() IFullIdentContext
	SEMI() antlr.TerminalNode

	// IsPackageStatementContext differentiates from other interfaces.
	IsPackageStatementContext()
}

type PackageStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackageStatementContext() *PackageStatementContext {
	var p = new(PackageStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_packageStatement
	return p
}

func InitEmptyPackageStatementContext(p *PackageStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_packageStatement
}

func (*PackageStatementContext) IsPackageStatementContext() {}

func NewPackageStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageStatementContext {
	var p = new(PackageStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_packageStatement

	return p
}

func (s *PackageStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageStatementContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(kdsParserPACKAGE, 0)
}

func (s *PackageStatementContext) FullIdent() IFullIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFullIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *PackageStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(kdsParserSEMI, 0)
}

func (s *PackageStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackageStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterPackageStatement(s)
	}
}

func (s *PackageStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitPackageStatement(s)
	}
}

func (p *kdsParser) PackageStatement() (localctx IPackageStatementContext) {
	localctx = NewPackageStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, kdsParserRULE_packageStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Match(kdsParserPACKAGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(100)
		p.FullIdent()
	}
	{
		p.SetState(101)
		p.Match(kdsParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOptionStatementContext is an interface to support dynamic dispatch.
type IOptionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPTION() antlr.TerminalNode
	OptionName() IOptionNameContext
	EQ() antlr.TerminalNode
	Constant() IConstantContext
	SEMI() antlr.TerminalNode

	// IsOptionStatementContext differentiates from other interfaces.
	IsOptionStatementContext()
}

type OptionStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionStatementContext() *OptionStatementContext {
	var p = new(OptionStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_optionStatement
	return p
}

func InitEmptyOptionStatementContext(p *OptionStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_optionStatement
}

func (*OptionStatementContext) IsOptionStatementContext() {}

func NewOptionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionStatementContext {
	var p = new(OptionStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_optionStatement

	return p
}

func (s *OptionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionStatementContext) OPTION() antlr.TerminalNode {
	return s.GetToken(kdsParserOPTION, 0)
}

func (s *OptionStatementContext) OptionName() IOptionNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *OptionStatementContext) EQ() antlr.TerminalNode {
	return s.GetToken(kdsParserEQ, 0)
}

func (s *OptionStatementContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *OptionStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(kdsParserSEMI, 0)
}

func (s *OptionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterOptionStatement(s)
	}
}

func (s *OptionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitOptionStatement(s)
	}
}

func (p *kdsParser) OptionStatement() (localctx IOptionStatementContext) {
	localctx = NewOptionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, kdsParserRULE_optionStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.Match(kdsParserOPTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(104)
		p.OptionName()
	}
	{
		p.SetState(105)
		p.Match(kdsParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
		p.Constant()
	}
	{
		p.SetState(107)
		p.Match(kdsParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOptionNameContext is an interface to support dynamic dispatch.
type IOptionNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFullIdent() []IFullIdentContext
	FullIdent(i int) IFullIdentContext
	LP() antlr.TerminalNode
	RP() antlr.TerminalNode
	DOT() antlr.TerminalNode

	// IsOptionNameContext differentiates from other interfaces.
	IsOptionNameContext()
}

type OptionNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionNameContext() *OptionNameContext {
	var p = new(OptionNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_optionName
	return p
}

func InitEmptyOptionNameContext(p *OptionNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_optionName
}

func (*OptionNameContext) IsOptionNameContext() {}

func NewOptionNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionNameContext {
	var p = new(OptionNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_optionName

	return p
}

func (s *OptionNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionNameContext) AllFullIdent() []IFullIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFullIdentContext); ok {
			len++
		}
	}

	tst := make([]IFullIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFullIdentContext); ok {
			tst[i] = t.(IFullIdentContext)
			i++
		}
	}

	return tst
}

func (s *OptionNameContext) FullIdent(i int) IFullIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFullIdentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *OptionNameContext) LP() antlr.TerminalNode {
	return s.GetToken(kdsParserLP, 0)
}

func (s *OptionNameContext) RP() antlr.TerminalNode {
	return s.GetToken(kdsParserRP, 0)
}

func (s *OptionNameContext) DOT() antlr.TerminalNode {
	return s.GetToken(kdsParserDOT, 0)
}

func (s *OptionNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterOptionName(s)
	}
}

func (s *OptionNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitOptionName(s)
	}
}

func (p *kdsParser) OptionName() (localctx IOptionNameContext) {
	localctx = NewOptionNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, kdsParserRULE_optionName)
	var _la int

	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case kdsParserPACKAGE, kdsParserMAP, kdsParserINT32, kdsParserINT64, kdsParserUINT32, kdsParserUINT64, kdsParserSINT32, kdsParserSINT64, kdsParserFIXED32, kdsParserFIXED64, kdsParserSFIXED32, kdsParserSFIXED64, kdsParserBOOL, kdsParserSTRING, kdsParserDOUBLE, kdsParserFLOAT, kdsParserBYTES, kdsParserENUM, kdsParserENTITY, kdsParserCOMPONENT, kdsParserBOOL_LIT, kdsParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(109)
			p.FullIdent()
		}

	case kdsParserLP:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(110)
			p.Match(kdsParserLP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.FullIdent()
		}
		{
			p.SetState(112)
			p.Match(kdsParserRP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == kdsParserDOT {
			{
				p.SetState(113)
				p.Match(kdsParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(114)
				p.FullIdent()
			}

		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() IType_Context
	FieldName() IFieldNameContext
	EQ() antlr.TerminalNode
	FieldNumber() IFieldNumberContext
	SEMI() antlr.TerminalNode
	FieldLabel() IFieldLabelContext

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_field
	return p
}

func InitEmptyFieldContext(p *FieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_field
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *FieldContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *FieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(kdsParserEQ, 0)
}

func (s *FieldContext) FieldNumber() IFieldNumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNumberContext)
}

func (s *FieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(kdsParserSEMI, 0)
}

func (s *FieldContext) FieldLabel() IFieldLabelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldLabelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldLabelContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *kdsParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, kdsParserRULE_field)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == kdsParserREPEATED {
		{
			p.SetState(119)
			p.FieldLabel()
		}

	}
	{
		p.SetState(122)
		p.Type_()
	}
	{
		p.SetState(123)
		p.FieldName()
	}
	{
		p.SetState(124)
		p.Match(kdsParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(125)
		p.FieldNumber()
	}
	{
		p.SetState(126)
		p.Match(kdsParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldLabelContext is an interface to support dynamic dispatch.
type IFieldLabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	REPEATED() antlr.TerminalNode

	// IsFieldLabelContext differentiates from other interfaces.
	IsFieldLabelContext()
}

type FieldLabelContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldLabelContext() *FieldLabelContext {
	var p = new(FieldLabelContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_fieldLabel
	return p
}

func InitEmptyFieldLabelContext(p *FieldLabelContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_fieldLabel
}

func (*FieldLabelContext) IsFieldLabelContext() {}

func NewFieldLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldLabelContext {
	var p = new(FieldLabelContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_fieldLabel

	return p
}

func (s *FieldLabelContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldLabelContext) REPEATED() antlr.TerminalNode {
	return s.GetToken(kdsParserREPEATED, 0)
}

func (s *FieldLabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldLabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldLabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterFieldLabel(s)
	}
}

func (s *FieldLabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitFieldLabel(s)
	}
}

func (p *kdsParser) FieldLabel() (localctx IFieldLabelContext) {
	localctx = NewFieldLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, kdsParserRULE_fieldLabel)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Match(kdsParserREPEATED)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldNumberContext is an interface to support dynamic dispatch.
type IFieldNumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IntLit() IIntLitContext

	// IsFieldNumberContext differentiates from other interfaces.
	IsFieldNumberContext()
}

type FieldNumberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNumberContext() *FieldNumberContext {
	var p = new(FieldNumberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_fieldNumber
	return p
}

func InitEmptyFieldNumberContext(p *FieldNumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_fieldNumber
}

func (*FieldNumberContext) IsFieldNumberContext() {}

func NewFieldNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNumberContext {
	var p = new(FieldNumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_fieldNumber

	return p
}

func (s *FieldNumberContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNumberContext) IntLit() IIntLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntLitContext)
}

func (s *FieldNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterFieldNumber(s)
	}
}

func (s *FieldNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitFieldNumber(s)
	}
}

func (p *kdsParser) FieldNumber() (localctx IFieldNumberContext) {
	localctx = NewFieldNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, kdsParserRULE_fieldNumber)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)
		p.IntLit()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMapFieldContext is an interface to support dynamic dispatch.
type IMapFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MAP() antlr.TerminalNode
	LT() antlr.TerminalNode
	KeyType() IKeyTypeContext
	COMMA() antlr.TerminalNode
	Type_() IType_Context
	GT() antlr.TerminalNode
	MapName() IMapNameContext
	EQ() antlr.TerminalNode
	FieldNumber() IFieldNumberContext
	SEMI() antlr.TerminalNode

	// IsMapFieldContext differentiates from other interfaces.
	IsMapFieldContext()
}

type MapFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapFieldContext() *MapFieldContext {
	var p = new(MapFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_mapField
	return p
}

func InitEmptyMapFieldContext(p *MapFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_mapField
}

func (*MapFieldContext) IsMapFieldContext() {}

func NewMapFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapFieldContext {
	var p = new(MapFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_mapField

	return p
}

func (s *MapFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *MapFieldContext) MAP() antlr.TerminalNode {
	return s.GetToken(kdsParserMAP, 0)
}

func (s *MapFieldContext) LT() antlr.TerminalNode {
	return s.GetToken(kdsParserLT, 0)
}

func (s *MapFieldContext) KeyType() IKeyTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyTypeContext)
}

func (s *MapFieldContext) COMMA() antlr.TerminalNode {
	return s.GetToken(kdsParserCOMMA, 0)
}

func (s *MapFieldContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *MapFieldContext) GT() antlr.TerminalNode {
	return s.GetToken(kdsParserGT, 0)
}

func (s *MapFieldContext) MapName() IMapNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapNameContext)
}

func (s *MapFieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(kdsParserEQ, 0)
}

func (s *MapFieldContext) FieldNumber() IFieldNumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNumberContext)
}

func (s *MapFieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(kdsParserSEMI, 0)
}

func (s *MapFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterMapField(s)
	}
}

func (s *MapFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitMapField(s)
	}
}

func (p *kdsParser) MapField() (localctx IMapFieldContext) {
	localctx = NewMapFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, kdsParserRULE_mapField)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.Match(kdsParserMAP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(133)
		p.Match(kdsParserLT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(134)
		p.KeyType()
	}
	{
		p.SetState(135)
		p.Match(kdsParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(136)
		p.Type_()
	}
	{
		p.SetState(137)
		p.Match(kdsParserGT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(138)
		p.MapName()
	}
	{
		p.SetState(139)
		p.Match(kdsParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(140)
		p.FieldNumber()
	}
	{
		p.SetState(141)
		p.Match(kdsParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IKeyTypeContext is an interface to support dynamic dispatch.
type IKeyTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT32() antlr.TerminalNode
	INT64() antlr.TerminalNode
	UINT32() antlr.TerminalNode
	UINT64() antlr.TerminalNode
	SINT32() antlr.TerminalNode
	SINT64() antlr.TerminalNode
	FIXED32() antlr.TerminalNode
	FIXED64() antlr.TerminalNode
	SFIXED32() antlr.TerminalNode
	SFIXED64() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsKeyTypeContext differentiates from other interfaces.
	IsKeyTypeContext()
}

type KeyTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyTypeContext() *KeyTypeContext {
	var p = new(KeyTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_keyType
	return p
}

func InitEmptyKeyTypeContext(p *KeyTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_keyType
}

func (*KeyTypeContext) IsKeyTypeContext() {}

func NewKeyTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyTypeContext {
	var p = new(KeyTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_keyType

	return p
}

func (s *KeyTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyTypeContext) INT32() antlr.TerminalNode {
	return s.GetToken(kdsParserINT32, 0)
}

func (s *KeyTypeContext) INT64() antlr.TerminalNode {
	return s.GetToken(kdsParserINT64, 0)
}

func (s *KeyTypeContext) UINT32() antlr.TerminalNode {
	return s.GetToken(kdsParserUINT32, 0)
}

func (s *KeyTypeContext) UINT64() antlr.TerminalNode {
	return s.GetToken(kdsParserUINT64, 0)
}

func (s *KeyTypeContext) SINT32() antlr.TerminalNode {
	return s.GetToken(kdsParserSINT32, 0)
}

func (s *KeyTypeContext) SINT64() antlr.TerminalNode {
	return s.GetToken(kdsParserSINT64, 0)
}

func (s *KeyTypeContext) FIXED32() antlr.TerminalNode {
	return s.GetToken(kdsParserFIXED32, 0)
}

func (s *KeyTypeContext) FIXED64() antlr.TerminalNode {
	return s.GetToken(kdsParserFIXED64, 0)
}

func (s *KeyTypeContext) SFIXED32() antlr.TerminalNode {
	return s.GetToken(kdsParserSFIXED32, 0)
}

func (s *KeyTypeContext) SFIXED64() antlr.TerminalNode {
	return s.GetToken(kdsParserSFIXED64, 0)
}

func (s *KeyTypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(kdsParserBOOL, 0)
}

func (s *KeyTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(kdsParserSTRING, 0)
}

func (s *KeyTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterKeyType(s)
	}
}

func (s *KeyTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitKeyType(s)
	}
}

func (p *kdsParser) KeyType() (localctx IKeyTypeContext) {
	localctx = NewKeyTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, kdsParserRULE_keyType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1048320) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IType_Context is an interface to support dynamic dispatch.
type IType_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOUBLE() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	INT32() antlr.TerminalNode
	INT64() antlr.TerminalNode
	UINT32() antlr.TerminalNode
	UINT64() antlr.TerminalNode
	SINT32() antlr.TerminalNode
	SINT64() antlr.TerminalNode
	FIXED32() antlr.TerminalNode
	FIXED64() antlr.TerminalNode
	SFIXED32() antlr.TerminalNode
	SFIXED64() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	STRING() antlr.TerminalNode
	BYTES() antlr.TerminalNode
	MessageType() IMessageTypeContext
	EnumType() IEnumTypeContext

	// IsType_Context differentiates from other interfaces.
	IsType_Context()
}

type Type_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_Context() *Type_Context {
	var p = new(Type_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_type_
	return p
}

func InitEmptyType_Context(p *Type_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_type_
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) DOUBLE() antlr.TerminalNode {
	return s.GetToken(kdsParserDOUBLE, 0)
}

func (s *Type_Context) FLOAT() antlr.TerminalNode {
	return s.GetToken(kdsParserFLOAT, 0)
}

func (s *Type_Context) INT32() antlr.TerminalNode {
	return s.GetToken(kdsParserINT32, 0)
}

func (s *Type_Context) INT64() antlr.TerminalNode {
	return s.GetToken(kdsParserINT64, 0)
}

func (s *Type_Context) UINT32() antlr.TerminalNode {
	return s.GetToken(kdsParserUINT32, 0)
}

func (s *Type_Context) UINT64() antlr.TerminalNode {
	return s.GetToken(kdsParserUINT64, 0)
}

func (s *Type_Context) SINT32() antlr.TerminalNode {
	return s.GetToken(kdsParserSINT32, 0)
}

func (s *Type_Context) SINT64() antlr.TerminalNode {
	return s.GetToken(kdsParserSINT64, 0)
}

func (s *Type_Context) FIXED32() antlr.TerminalNode {
	return s.GetToken(kdsParserFIXED32, 0)
}

func (s *Type_Context) FIXED64() antlr.TerminalNode {
	return s.GetToken(kdsParserFIXED64, 0)
}

func (s *Type_Context) SFIXED32() antlr.TerminalNode {
	return s.GetToken(kdsParserSFIXED32, 0)
}

func (s *Type_Context) SFIXED64() antlr.TerminalNode {
	return s.GetToken(kdsParserSFIXED64, 0)
}

func (s *Type_Context) BOOL() antlr.TerminalNode {
	return s.GetToken(kdsParserBOOL, 0)
}

func (s *Type_Context) STRING() antlr.TerminalNode {
	return s.GetToken(kdsParserSTRING, 0)
}

func (s *Type_Context) BYTES() antlr.TerminalNode {
	return s.GetToken(kdsParserBYTES, 0)
}

func (s *Type_Context) MessageType() IMessageTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageTypeContext)
}

func (s *Type_Context) EnumType() IEnumTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumTypeContext)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterType_(s)
	}
}

func (s *Type_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitType_(s)
	}
}

func (p *kdsParser) Type_() (localctx IType_Context) {
	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, kdsParserRULE_type_)
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(145)
			p.Match(kdsParserDOUBLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(146)
			p.Match(kdsParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(147)
			p.Match(kdsParserINT32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(148)
			p.Match(kdsParserINT64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(149)
			p.Match(kdsParserUINT32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(150)
			p.Match(kdsParserUINT64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(151)
			p.Match(kdsParserSINT32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(152)
			p.Match(kdsParserSINT64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(153)
			p.Match(kdsParserFIXED32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(154)
			p.Match(kdsParserFIXED64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(155)
			p.Match(kdsParserSFIXED32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(156)
			p.Match(kdsParserSFIXED64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(157)
			p.Match(kdsParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(158)
			p.Match(kdsParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(159)
			p.Match(kdsParserBYTES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(160)
			p.MessageType()
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(161)
			p.EnumType()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITopLevelDefContext is an interface to support dynamic dispatch.
type ITopLevelDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EntityDef() IEntityDefContext
	ComponentDef() IComponentDefContext
	EnumDef() IEnumDefContext

	// IsTopLevelDefContext differentiates from other interfaces.
	IsTopLevelDefContext()
}

type TopLevelDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopLevelDefContext() *TopLevelDefContext {
	var p = new(TopLevelDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_topLevelDef
	return p
}

func InitEmptyTopLevelDefContext(p *TopLevelDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_topLevelDef
}

func (*TopLevelDefContext) IsTopLevelDefContext() {}

func NewTopLevelDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopLevelDefContext {
	var p = new(TopLevelDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_topLevelDef

	return p
}

func (s *TopLevelDefContext) GetParser() antlr.Parser { return s.parser }

func (s *TopLevelDefContext) EntityDef() IEntityDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEntityDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEntityDefContext)
}

func (s *TopLevelDefContext) ComponentDef() IComponentDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComponentDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComponentDefContext)
}

func (s *TopLevelDefContext) EnumDef() IEnumDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumDefContext)
}

func (s *TopLevelDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopLevelDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterTopLevelDef(s)
	}
}

func (s *TopLevelDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitTopLevelDef(s)
	}
}

func (p *kdsParser) TopLevelDef() (localctx ITopLevelDefContext) {
	localctx = NewTopLevelDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, kdsParserRULE_topLevelDef)
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case kdsParserENTITY:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(164)
			p.EntityDef()
		}

	case kdsParserCOMPONENT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(165)
			p.ComponentDef()
		}

	case kdsParserENUM:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(166)
			p.EnumDef()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnumDefContext is an interface to support dynamic dispatch.
type IEnumDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ENUM() antlr.TerminalNode
	EnumName() IEnumNameContext
	EnumBody() IEnumBodyContext

	// IsEnumDefContext differentiates from other interfaces.
	IsEnumDefContext()
}

type EnumDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumDefContext() *EnumDefContext {
	var p = new(EnumDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_enumDef
	return p
}

func InitEmptyEnumDefContext(p *EnumDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_enumDef
}

func (*EnumDefContext) IsEnumDefContext() {}

func NewEnumDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumDefContext {
	var p = new(EnumDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_enumDef

	return p
}

func (s *EnumDefContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumDefContext) ENUM() antlr.TerminalNode {
	return s.GetToken(kdsParserENUM, 0)
}

func (s *EnumDefContext) EnumName() IEnumNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumNameContext)
}

func (s *EnumDefContext) EnumBody() IEnumBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumBodyContext)
}

func (s *EnumDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterEnumDef(s)
	}
}

func (s *EnumDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitEnumDef(s)
	}
}

func (p *kdsParser) EnumDef() (localctx IEnumDefContext) {
	localctx = NewEnumDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, kdsParserRULE_enumDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(kdsParserENUM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
		p.EnumName()
	}
	{
		p.SetState(171)
		p.EnumBody()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnumBodyContext is an interface to support dynamic dispatch.
type IEnumBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LC() antlr.TerminalNode
	RC() antlr.TerminalNode
	AllEnumElement() []IEnumElementContext
	EnumElement(i int) IEnumElementContext

	// IsEnumBodyContext differentiates from other interfaces.
	IsEnumBodyContext()
}

type EnumBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumBodyContext() *EnumBodyContext {
	var p = new(EnumBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_enumBody
	return p
}

func InitEmptyEnumBodyContext(p *EnumBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_enumBody
}

func (*EnumBodyContext) IsEnumBodyContext() {}

func NewEnumBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumBodyContext {
	var p = new(EnumBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_enumBody

	return p
}

func (s *EnumBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumBodyContext) LC() antlr.TerminalNode {
	return s.GetToken(kdsParserLC, 0)
}

func (s *EnumBodyContext) RC() antlr.TerminalNode {
	return s.GetToken(kdsParserRC, 0)
}

func (s *EnumBodyContext) AllEnumElement() []IEnumElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEnumElementContext); ok {
			len++
		}
	}

	tst := make([]IEnumElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEnumElementContext); ok {
			tst[i] = t.(IEnumElementContext)
			i++
		}
	}

	return tst
}

func (s *EnumBodyContext) EnumElement(i int) IEnumElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumElementContext)
}

func (s *EnumBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterEnumBody(s)
	}
}

func (s *EnumBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitEnumBody(s)
	}
}

func (p *kdsParser) EnumBody() (localctx IEnumBodyContext) {
	localctx = NewEnumBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, kdsParserRULE_enumBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(173)
		p.Match(kdsParserLC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&158329741508496) != 0 {
		{
			p.SetState(174)
			p.EnumElement()
		}

		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(180)
		p.Match(kdsParserRC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnumElementContext is an interface to support dynamic dispatch.
type IEnumElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EnumField() IEnumFieldContext

	// IsEnumElementContext differentiates from other interfaces.
	IsEnumElementContext()
}

type EnumElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumElementContext() *EnumElementContext {
	var p = new(EnumElementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_enumElement
	return p
}

func InitEmptyEnumElementContext(p *EnumElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_enumElement
}

func (*EnumElementContext) IsEnumElementContext() {}

func NewEnumElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumElementContext {
	var p = new(EnumElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_enumElement

	return p
}

func (s *EnumElementContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumElementContext) EnumField() IEnumFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumFieldContext)
}

func (s *EnumElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterEnumElement(s)
	}
}

func (s *EnumElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitEnumElement(s)
	}
}

func (p *kdsParser) EnumElement() (localctx IEnumElementContext) {
	localctx = NewEnumElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, kdsParserRULE_enumElement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		p.EnumField()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnumFieldContext is an interface to support dynamic dispatch.
type IEnumFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext
	EQ() antlr.TerminalNode
	IntLit() IIntLitContext
	SEMI() antlr.TerminalNode
	MINUS() antlr.TerminalNode

	// IsEnumFieldContext differentiates from other interfaces.
	IsEnumFieldContext()
}

type EnumFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumFieldContext() *EnumFieldContext {
	var p = new(EnumFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_enumField
	return p
}

func InitEmptyEnumFieldContext(p *EnumFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_enumField
}

func (*EnumFieldContext) IsEnumFieldContext() {}

func NewEnumFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumFieldContext {
	var p = new(EnumFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_enumField

	return p
}

func (s *EnumFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumFieldContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *EnumFieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(kdsParserEQ, 0)
}

func (s *EnumFieldContext) IntLit() IIntLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntLitContext)
}

func (s *EnumFieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(kdsParserSEMI, 0)
}

func (s *EnumFieldContext) MINUS() antlr.TerminalNode {
	return s.GetToken(kdsParserMINUS, 0)
}

func (s *EnumFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterEnumField(s)
	}
}

func (s *EnumFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitEnumField(s)
	}
}

func (p *kdsParser) EnumField() (localctx IEnumFieldContext) {
	localctx = NewEnumFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, kdsParserRULE_enumField)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(184)
		p.Ident()
	}
	{
		p.SetState(185)
		p.Match(kdsParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == kdsParserMINUS {
		{
			p.SetState(186)
			p.Match(kdsParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(189)
		p.IntLit()
	}
	{
		p.SetState(190)
		p.Match(kdsParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEntityDefContext is an interface to support dynamic dispatch.
type IEntityDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ENTITY() antlr.TerminalNode
	EntityName() IEntityNameContext
	EntityBody() IEntityBodyContext

	// IsEntityDefContext differentiates from other interfaces.
	IsEntityDefContext()
}

type EntityDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEntityDefContext() *EntityDefContext {
	var p = new(EntityDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_entityDef
	return p
}

func InitEmptyEntityDefContext(p *EntityDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_entityDef
}

func (*EntityDefContext) IsEntityDefContext() {}

func NewEntityDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EntityDefContext {
	var p = new(EntityDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_entityDef

	return p
}

func (s *EntityDefContext) GetParser() antlr.Parser { return s.parser }

func (s *EntityDefContext) ENTITY() antlr.TerminalNode {
	return s.GetToken(kdsParserENTITY, 0)
}

func (s *EntityDefContext) EntityName() IEntityNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEntityNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEntityNameContext)
}

func (s *EntityDefContext) EntityBody() IEntityBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEntityBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEntityBodyContext)
}

func (s *EntityDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EntityDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EntityDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterEntityDef(s)
	}
}

func (s *EntityDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitEntityDef(s)
	}
}

func (p *kdsParser) EntityDef() (localctx IEntityDefContext) {
	localctx = NewEntityDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, kdsParserRULE_entityDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Match(kdsParserENTITY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(193)
		p.EntityName()
	}
	{
		p.SetState(194)
		p.EntityBody()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEntityNameContext is an interface to support dynamic dispatch.
type IEntityNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext

	// IsEntityNameContext differentiates from other interfaces.
	IsEntityNameContext()
}

type EntityNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEntityNameContext() *EntityNameContext {
	var p = new(EntityNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_entityName
	return p
}

func InitEmptyEntityNameContext(p *EntityNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_entityName
}

func (*EntityNameContext) IsEntityNameContext() {}

func NewEntityNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EntityNameContext {
	var p = new(EntityNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_entityName

	return p
}

func (s *EntityNameContext) GetParser() antlr.Parser { return s.parser }

func (s *EntityNameContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *EntityNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EntityNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EntityNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterEntityName(s)
	}
}

func (s *EntityNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitEntityName(s)
	}
}

func (p *kdsParser) EntityName() (localctx IEntityNameContext) {
	localctx = NewEntityNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, kdsParserRULE_entityName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Ident()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEntityBodyContext is an interface to support dynamic dispatch.
type IEntityBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LC() antlr.TerminalNode
	RC() antlr.TerminalNode
	AllEntityElement() []IEntityElementContext
	EntityElement(i int) IEntityElementContext

	// IsEntityBodyContext differentiates from other interfaces.
	IsEntityBodyContext()
}

type EntityBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEntityBodyContext() *EntityBodyContext {
	var p = new(EntityBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_entityBody
	return p
}

func InitEmptyEntityBodyContext(p *EntityBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_entityBody
}

func (*EntityBodyContext) IsEntityBodyContext() {}

func NewEntityBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EntityBodyContext {
	var p = new(EntityBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_entityBody

	return p
}

func (s *EntityBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *EntityBodyContext) LC() antlr.TerminalNode {
	return s.GetToken(kdsParserLC, 0)
}

func (s *EntityBodyContext) RC() antlr.TerminalNode {
	return s.GetToken(kdsParserRC, 0)
}

func (s *EntityBodyContext) AllEntityElement() []IEntityElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEntityElementContext); ok {
			len++
		}
	}

	tst := make([]IEntityElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEntityElementContext); ok {
			tst[i] = t.(IEntityElementContext)
			i++
		}
	}

	return tst
}

func (s *EntityBodyContext) EntityElement(i int) IEntityElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEntityElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEntityElementContext)
}

func (s *EntityBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EntityBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EntityBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterEntityBody(s)
	}
}

func (s *EntityBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitEntityBody(s)
	}
}

func (p *kdsParser) EntityBody() (localctx IEntityBodyContext) {
	localctx = NewEntityBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, kdsParserRULE_entityBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)
		p.Match(kdsParserLC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&158604887850960) != 0 {
		{
			p.SetState(199)
			p.EntityElement()
		}

		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(205)
		p.Match(kdsParserRC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEntityElementContext is an interface to support dynamic dispatch.
type IEntityElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Field() IFieldContext
	MapField() IMapFieldContext
	EmptyStatement_() IEmptyStatement_Context

	// IsEntityElementContext differentiates from other interfaces.
	IsEntityElementContext()
}

type EntityElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEntityElementContext() *EntityElementContext {
	var p = new(EntityElementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_entityElement
	return p
}

func InitEmptyEntityElementContext(p *EntityElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_entityElement
}

func (*EntityElementContext) IsEntityElementContext() {}

func NewEntityElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EntityElementContext {
	var p = new(EntityElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_entityElement

	return p
}

func (s *EntityElementContext) GetParser() antlr.Parser { return s.parser }

func (s *EntityElementContext) Field() IFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *EntityElementContext) MapField() IMapFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapFieldContext)
}

func (s *EntityElementContext) EmptyStatement_() IEmptyStatement_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmptyStatement_Context)
}

func (s *EntityElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EntityElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EntityElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterEntityElement(s)
	}
}

func (s *EntityElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitEntityElement(s)
	}
}

func (p *kdsParser) EntityElement() (localctx IEntityElementContext) {
	localctx = NewEntityElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, kdsParserRULE_entityElement)
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(207)
			p.Field()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(208)
			p.MapField()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(209)
			p.EmptyStatement_()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComponentDefContext is an interface to support dynamic dispatch.
type IComponentDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMPONENT() antlr.TerminalNode
	ComponentName() IComponentNameContext
	ComponentBody() IComponentBodyContext

	// IsComponentDefContext differentiates from other interfaces.
	IsComponentDefContext()
}

type ComponentDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComponentDefContext() *ComponentDefContext {
	var p = new(ComponentDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_componentDef
	return p
}

func InitEmptyComponentDefContext(p *ComponentDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_componentDef
}

func (*ComponentDefContext) IsComponentDefContext() {}

func NewComponentDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComponentDefContext {
	var p = new(ComponentDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_componentDef

	return p
}

func (s *ComponentDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ComponentDefContext) COMPONENT() antlr.TerminalNode {
	return s.GetToken(kdsParserCOMPONENT, 0)
}

func (s *ComponentDefContext) ComponentName() IComponentNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComponentNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComponentNameContext)
}

func (s *ComponentDefContext) ComponentBody() IComponentBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComponentBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComponentBodyContext)
}

func (s *ComponentDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComponentDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComponentDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterComponentDef(s)
	}
}

func (s *ComponentDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitComponentDef(s)
	}
}

func (p *kdsParser) ComponentDef() (localctx IComponentDefContext) {
	localctx = NewComponentDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, kdsParserRULE_componentDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)
		p.Match(kdsParserCOMPONENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(213)
		p.ComponentName()
	}
	{
		p.SetState(214)
		p.ComponentBody()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComponentNameContext is an interface to support dynamic dispatch.
type IComponentNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext

	// IsComponentNameContext differentiates from other interfaces.
	IsComponentNameContext()
}

type ComponentNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComponentNameContext() *ComponentNameContext {
	var p = new(ComponentNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_componentName
	return p
}

func InitEmptyComponentNameContext(p *ComponentNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_componentName
}

func (*ComponentNameContext) IsComponentNameContext() {}

func NewComponentNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComponentNameContext {
	var p = new(ComponentNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_componentName

	return p
}

func (s *ComponentNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ComponentNameContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *ComponentNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComponentNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComponentNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterComponentName(s)
	}
}

func (s *ComponentNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitComponentName(s)
	}
}

func (p *kdsParser) ComponentName() (localctx IComponentNameContext) {
	localctx = NewComponentNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, kdsParserRULE_componentName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
		p.Ident()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComponentBodyContext is an interface to support dynamic dispatch.
type IComponentBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LC() antlr.TerminalNode
	RC() antlr.TerminalNode
	AllComponentElement() []IComponentElementContext
	ComponentElement(i int) IComponentElementContext

	// IsComponentBodyContext differentiates from other interfaces.
	IsComponentBodyContext()
}

type ComponentBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComponentBodyContext() *ComponentBodyContext {
	var p = new(ComponentBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_componentBody
	return p
}

func InitEmptyComponentBodyContext(p *ComponentBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_componentBody
}

func (*ComponentBodyContext) IsComponentBodyContext() {}

func NewComponentBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComponentBodyContext {
	var p = new(ComponentBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_componentBody

	return p
}

func (s *ComponentBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ComponentBodyContext) LC() antlr.TerminalNode {
	return s.GetToken(kdsParserLC, 0)
}

func (s *ComponentBodyContext) RC() antlr.TerminalNode {
	return s.GetToken(kdsParserRC, 0)
}

func (s *ComponentBodyContext) AllComponentElement() []IComponentElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IComponentElementContext); ok {
			len++
		}
	}

	tst := make([]IComponentElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IComponentElementContext); ok {
			tst[i] = t.(IComponentElementContext)
			i++
		}
	}

	return tst
}

func (s *ComponentBodyContext) ComponentElement(i int) IComponentElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComponentElementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComponentElementContext)
}

func (s *ComponentBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComponentBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComponentBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterComponentBody(s)
	}
}

func (s *ComponentBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitComponentBody(s)
	}
}

func (p *kdsParser) ComponentBody() (localctx IComponentBodyContext) {
	localctx = NewComponentBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, kdsParserRULE_componentBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(218)
		p.Match(kdsParserLC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&158604887850960) != 0 {
		{
			p.SetState(219)
			p.ComponentElement()
		}

		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(225)
		p.Match(kdsParserRC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComponentElementContext is an interface to support dynamic dispatch.
type IComponentElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Field() IFieldContext
	MapField() IMapFieldContext
	EmptyStatement_() IEmptyStatement_Context

	// IsComponentElementContext differentiates from other interfaces.
	IsComponentElementContext()
}

type ComponentElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComponentElementContext() *ComponentElementContext {
	var p = new(ComponentElementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_componentElement
	return p
}

func InitEmptyComponentElementContext(p *ComponentElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_componentElement
}

func (*ComponentElementContext) IsComponentElementContext() {}

func NewComponentElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComponentElementContext {
	var p = new(ComponentElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_componentElement

	return p
}

func (s *ComponentElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ComponentElementContext) Field() IFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *ComponentElementContext) MapField() IMapFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapFieldContext)
}

func (s *ComponentElementContext) EmptyStatement_() IEmptyStatement_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmptyStatement_Context)
}

func (s *ComponentElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComponentElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComponentElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterComponentElement(s)
	}
}

func (s *ComponentElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitComponentElement(s)
	}
}

func (p *kdsParser) ComponentElement() (localctx IComponentElementContext) {
	localctx = NewComponentElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, kdsParserRULE_componentElement)
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(227)
			p.Field()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(228)
			p.MapField()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(229)
			p.EmptyStatement_()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FullIdent() IFullIdentContext
	IntLit() IIntLitContext
	MINUS() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	FloatLit() IFloatLitContext
	StrLit() IStrLitContext
	BoolLit() IBoolLitContext
	BlockLit() IBlockLitContext

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_constant
	return p
}

func InitEmptyConstantContext(p *ConstantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_constant
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) FullIdent() IFullIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFullIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *ConstantContext) IntLit() IIntLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntLitContext)
}

func (s *ConstantContext) MINUS() antlr.TerminalNode {
	return s.GetToken(kdsParserMINUS, 0)
}

func (s *ConstantContext) PLUS() antlr.TerminalNode {
	return s.GetToken(kdsParserPLUS, 0)
}

func (s *ConstantContext) FloatLit() IFloatLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloatLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloatLitContext)
}

func (s *ConstantContext) StrLit() IStrLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStrLitContext)
}

func (s *ConstantContext) BoolLit() IBoolLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolLitContext)
}

func (s *ConstantContext) BlockLit() IBlockLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockLitContext)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (p *kdsParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, kdsParserRULE_constant)
	var _la int

	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(232)
			p.FullIdent()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == kdsParserPLUS || _la == kdsParserMINUS {
			{
				p.SetState(233)
				_la = p.GetTokenStream().LA(1)

				if !(_la == kdsParserPLUS || _la == kdsParserMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(236)
			p.IntLit()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(238)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == kdsParserPLUS || _la == kdsParserMINUS {
			{
				p.SetState(237)
				_la = p.GetTokenStream().LA(1)

				if !(_la == kdsParserPLUS || _la == kdsParserMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(240)
			p.FloatLit()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(241)
			p.StrLit()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(242)
			p.BoolLit()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(243)
			p.BlockLit()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockLitContext is an interface to support dynamic dispatch.
type IBlockLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LC() antlr.TerminalNode
	RC() antlr.TerminalNode
	AllIdent() []IIdentContext
	Ident(i int) IIdentContext
	AllCOLON() []antlr.TerminalNode
	COLON(i int) antlr.TerminalNode
	AllConstant() []IConstantContext
	Constant(i int) IConstantContext

	// IsBlockLitContext differentiates from other interfaces.
	IsBlockLitContext()
}

type BlockLitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockLitContext() *BlockLitContext {
	var p = new(BlockLitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_blockLit
	return p
}

func InitEmptyBlockLitContext(p *BlockLitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_blockLit
}

func (*BlockLitContext) IsBlockLitContext() {}

func NewBlockLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockLitContext {
	var p = new(BlockLitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_blockLit

	return p
}

func (s *BlockLitContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockLitContext) LC() antlr.TerminalNode {
	return s.GetToken(kdsParserLC, 0)
}

func (s *BlockLitContext) RC() antlr.TerminalNode {
	return s.GetToken(kdsParserRC, 0)
}

func (s *BlockLitContext) AllIdent() []IIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentContext); ok {
			len++
		}
	}

	tst := make([]IIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentContext); ok {
			tst[i] = t.(IIdentContext)
			i++
		}
	}

	return tst
}

func (s *BlockLitContext) Ident(i int) IIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *BlockLitContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(kdsParserCOLON)
}

func (s *BlockLitContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(kdsParserCOLON, i)
}

func (s *BlockLitContext) AllConstant() []IConstantContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstantContext); ok {
			len++
		}
	}

	tst := make([]IConstantContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstantContext); ok {
			tst[i] = t.(IConstantContext)
			i++
		}
	}

	return tst
}

func (s *BlockLitContext) Constant(i int) IConstantContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *BlockLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterBlockLit(s)
	}
}

func (s *BlockLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitBlockLit(s)
	}
}

func (p *kdsParser) BlockLit() (localctx IBlockLitContext) {
	localctx = NewBlockLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, kdsParserRULE_blockLit)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(246)
		p.Match(kdsParserLC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&158329741508496) != 0 {
		{
			p.SetState(247)
			p.Ident()
		}
		{
			p.SetState(248)
			p.Match(kdsParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(249)
			p.Constant()
		}

		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(256)
		p.Match(kdsParserRC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEmptyStatement_Context is an interface to support dynamic dispatch.
type IEmptyStatement_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SEMI() antlr.TerminalNode

	// IsEmptyStatement_Context differentiates from other interfaces.
	IsEmptyStatement_Context()
}

type EmptyStatement_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmptyStatement_Context() *EmptyStatement_Context {
	var p = new(EmptyStatement_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_emptyStatement_
	return p
}

func InitEmptyEmptyStatement_Context(p *EmptyStatement_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_emptyStatement_
}

func (*EmptyStatement_Context) IsEmptyStatement_Context() {}

func NewEmptyStatement_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmptyStatement_Context {
	var p = new(EmptyStatement_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_emptyStatement_

	return p
}

func (s *EmptyStatement_Context) GetParser() antlr.Parser { return s.parser }

func (s *EmptyStatement_Context) SEMI() antlr.TerminalNode {
	return s.GetToken(kdsParserSEMI, 0)
}

func (s *EmptyStatement_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyStatement_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmptyStatement_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterEmptyStatement_(s)
	}
}

func (s *EmptyStatement_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitEmptyStatement_(s)
	}
}

func (p *kdsParser) EmptyStatement_() (localctx IEmptyStatement_Context) {
	localctx = NewEmptyStatement_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, kdsParserRULE_emptyStatement_)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(258)
		p.Match(kdsParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentContext is an interface to support dynamic dispatch.
type IIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Keywords() IKeywordsContext

	// IsIdentContext differentiates from other interfaces.
	IsIdentContext()
}

type IdentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentContext() *IdentContext {
	var p = new(IdentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_ident
	return p
}

func InitEmptyIdentContext(p *IdentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_ident
}

func (*IdentContext) IsIdentContext() {}

func NewIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentContext {
	var p = new(IdentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_ident

	return p
}

func (s *IdentContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(kdsParserIDENTIFIER, 0)
}

func (s *IdentContext) Keywords() IKeywordsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeywordsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeywordsContext)
}

func (s *IdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterIdent(s)
	}
}

func (s *IdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitIdent(s)
	}
}

func (p *kdsParser) Ident() (localctx IIdentContext) {
	localctx = NewIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, kdsParserRULE_ident)
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case kdsParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(260)
			p.Match(kdsParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case kdsParserPACKAGE, kdsParserMAP, kdsParserINT32, kdsParserINT64, kdsParserUINT32, kdsParserUINT64, kdsParserSINT32, kdsParserSINT64, kdsParserFIXED32, kdsParserFIXED64, kdsParserSFIXED32, kdsParserSFIXED64, kdsParserBOOL, kdsParserSTRING, kdsParserDOUBLE, kdsParserFLOAT, kdsParserBYTES, kdsParserENUM, kdsParserENTITY, kdsParserCOMPONENT, kdsParserBOOL_LIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(261)
			p.Keywords()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFullIdentContext is an interface to support dynamic dispatch.
type IFullIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdent() []IIdentContext
	Ident(i int) IIdentContext
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsFullIdentContext differentiates from other interfaces.
	IsFullIdentContext()
}

type FullIdentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFullIdentContext() *FullIdentContext {
	var p = new(FullIdentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_fullIdent
	return p
}

func InitEmptyFullIdentContext(p *FullIdentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_fullIdent
}

func (*FullIdentContext) IsFullIdentContext() {}

func NewFullIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FullIdentContext {
	var p = new(FullIdentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_fullIdent

	return p
}

func (s *FullIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *FullIdentContext) AllIdent() []IIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentContext); ok {
			len++
		}
	}

	tst := make([]IIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentContext); ok {
			tst[i] = t.(IIdentContext)
			i++
		}
	}

	return tst
}

func (s *FullIdentContext) Ident(i int) IIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *FullIdentContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(kdsParserDOT)
}

func (s *FullIdentContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(kdsParserDOT, i)
}

func (s *FullIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FullIdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterFullIdent(s)
	}
}

func (s *FullIdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitFullIdent(s)
	}
}

func (p *kdsParser) FullIdent() (localctx IFullIdentContext) {
	localctx = NewFullIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, kdsParserRULE_fullIdent)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(264)
		p.Ident()
	}
	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == kdsParserDOT {
		{
			p.SetState(265)
			p.Match(kdsParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(266)
			p.Ident()
		}

		p.SetState(271)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldNameContext is an interface to support dynamic dispatch.
type IFieldNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext

	// IsFieldNameContext differentiates from other interfaces.
	IsFieldNameContext()
}

type FieldNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNameContext() *FieldNameContext {
	var p = new(FieldNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_fieldName
	return p
}

func InitEmptyFieldNameContext(p *FieldNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_fieldName
}

func (*FieldNameContext) IsFieldNameContext() {}

func NewFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameContext {
	var p = new(FieldNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_fieldName

	return p
}

func (s *FieldNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNameContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *FieldNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterFieldName(s)
	}
}

func (s *FieldNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitFieldName(s)
	}
}

func (p *kdsParser) FieldName() (localctx IFieldNameContext) {
	localctx = NewFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, kdsParserRULE_fieldName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(272)
		p.Ident()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMessageNameContext is an interface to support dynamic dispatch.
type IMessageNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext

	// IsMessageNameContext differentiates from other interfaces.
	IsMessageNameContext()
}

type MessageNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageNameContext() *MessageNameContext {
	var p = new(MessageNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_messageName
	return p
}

func InitEmptyMessageNameContext(p *MessageNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_messageName
}

func (*MessageNameContext) IsMessageNameContext() {}

func NewMessageNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageNameContext {
	var p = new(MessageNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_messageName

	return p
}

func (s *MessageNameContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageNameContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *MessageNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterMessageName(s)
	}
}

func (s *MessageNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitMessageName(s)
	}
}

func (p *kdsParser) MessageName() (localctx IMessageNameContext) {
	localctx = NewMessageNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, kdsParserRULE_messageName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(274)
		p.Ident()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnumNameContext is an interface to support dynamic dispatch.
type IEnumNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext

	// IsEnumNameContext differentiates from other interfaces.
	IsEnumNameContext()
}

type EnumNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumNameContext() *EnumNameContext {
	var p = new(EnumNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_enumName
	return p
}

func InitEmptyEnumNameContext(p *EnumNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_enumName
}

func (*EnumNameContext) IsEnumNameContext() {}

func NewEnumNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumNameContext {
	var p = new(EnumNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_enumName

	return p
}

func (s *EnumNameContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumNameContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *EnumNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterEnumName(s)
	}
}

func (s *EnumNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitEnumName(s)
	}
}

func (p *kdsParser) EnumName() (localctx IEnumNameContext) {
	localctx = NewEnumNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, kdsParserRULE_enumName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(276)
		p.Ident()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMapNameContext is an interface to support dynamic dispatch.
type IMapNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext

	// IsMapNameContext differentiates from other interfaces.
	IsMapNameContext()
}

type MapNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapNameContext() *MapNameContext {
	var p = new(MapNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_mapName
	return p
}

func InitEmptyMapNameContext(p *MapNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_mapName
}

func (*MapNameContext) IsMapNameContext() {}

func NewMapNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapNameContext {
	var p = new(MapNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_mapName

	return p
}

func (s *MapNameContext) GetParser() antlr.Parser { return s.parser }

func (s *MapNameContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *MapNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterMapName(s)
	}
}

func (s *MapNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitMapName(s)
	}
}

func (p *kdsParser) MapName() (localctx IMapNameContext) {
	localctx = NewMapNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, kdsParserRULE_mapName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(278)
		p.Ident()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMessageTypeContext is an interface to support dynamic dispatch.
type IMessageTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MessageName() IMessageNameContext
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	AllIdent() []IIdentContext
	Ident(i int) IIdentContext

	// IsMessageTypeContext differentiates from other interfaces.
	IsMessageTypeContext()
}

type MessageTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageTypeContext() *MessageTypeContext {
	var p = new(MessageTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_messageType
	return p
}

func InitEmptyMessageTypeContext(p *MessageTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_messageType
}

func (*MessageTypeContext) IsMessageTypeContext() {}

func NewMessageTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageTypeContext {
	var p = new(MessageTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_messageType

	return p
}

func (s *MessageTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageTypeContext) MessageName() IMessageNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageNameContext)
}

func (s *MessageTypeContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(kdsParserDOT)
}

func (s *MessageTypeContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(kdsParserDOT, i)
}

func (s *MessageTypeContext) AllIdent() []IIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentContext); ok {
			len++
		}
	}

	tst := make([]IIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentContext); ok {
			tst[i] = t.(IIdentContext)
			i++
		}
	}

	return tst
}

func (s *MessageTypeContext) Ident(i int) IIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *MessageTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterMessageType(s)
	}
}

func (s *MessageTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitMessageType(s)
	}
}

func (p *kdsParser) MessageType() (localctx IMessageTypeContext) {
	localctx = NewMessageTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, kdsParserRULE_messageType)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == kdsParserDOT {
		{
			p.SetState(280)
			p.Match(kdsParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(283)
				p.Ident()
			}
			{
				p.SetState(284)
				p.Match(kdsParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(291)
		p.MessageName()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnumTypeContext is an interface to support dynamic dispatch.
type IEnumTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EnumName() IEnumNameContext
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	AllIdent() []IIdentContext
	Ident(i int) IIdentContext

	// IsEnumTypeContext differentiates from other interfaces.
	IsEnumTypeContext()
}

type EnumTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumTypeContext() *EnumTypeContext {
	var p = new(EnumTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_enumType
	return p
}

func InitEmptyEnumTypeContext(p *EnumTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_enumType
}

func (*EnumTypeContext) IsEnumTypeContext() {}

func NewEnumTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumTypeContext {
	var p = new(EnumTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_enumType

	return p
}

func (s *EnumTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumTypeContext) EnumName() IEnumNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumNameContext)
}

func (s *EnumTypeContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(kdsParserDOT)
}

func (s *EnumTypeContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(kdsParserDOT, i)
}

func (s *EnumTypeContext) AllIdent() []IIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentContext); ok {
			len++
		}
	}

	tst := make([]IIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentContext); ok {
			tst[i] = t.(IIdentContext)
			i++
		}
	}

	return tst
}

func (s *EnumTypeContext) Ident(i int) IIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *EnumTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterEnumType(s)
	}
}

func (s *EnumTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitEnumType(s)
	}
}

func (p *kdsParser) EnumType() (localctx IEnumTypeContext) {
	localctx = NewEnumTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, kdsParserRULE_enumType)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == kdsParserDOT {
		{
			p.SetState(293)
			p.Match(kdsParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(296)
				p.Ident()
			}
			{
				p.SetState(297)
				p.Match(kdsParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(303)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(304)
		p.EnumName()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIntLitContext is an interface to support dynamic dispatch.
type IIntLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT_LIT() antlr.TerminalNode

	// IsIntLitContext differentiates from other interfaces.
	IsIntLitContext()
}

type IntLitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntLitContext() *IntLitContext {
	var p = new(IntLitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_intLit
	return p
}

func InitEmptyIntLitContext(p *IntLitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_intLit
}

func (*IntLitContext) IsIntLitContext() {}

func NewIntLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntLitContext {
	var p = new(IntLitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_intLit

	return p
}

func (s *IntLitContext) GetParser() antlr.Parser { return s.parser }

func (s *IntLitContext) INT_LIT() antlr.TerminalNode {
	return s.GetToken(kdsParserINT_LIT, 0)
}

func (s *IntLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterIntLit(s)
	}
}

func (s *IntLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitIntLit(s)
	}
}

func (p *kdsParser) IntLit() (localctx IIntLitContext) {
	localctx = NewIntLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, kdsParserRULE_intLit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(306)
		p.Match(kdsParserINT_LIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStrLitContext is an interface to support dynamic dispatch.
type IStrLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STR_LIT() antlr.TerminalNode
	PROTO3_LIT_SINGLE() antlr.TerminalNode
	PROTO3_LIT_DOBULE() antlr.TerminalNode

	// IsStrLitContext differentiates from other interfaces.
	IsStrLitContext()
}

type StrLitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStrLitContext() *StrLitContext {
	var p = new(StrLitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_strLit
	return p
}

func InitEmptyStrLitContext(p *StrLitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_strLit
}

func (*StrLitContext) IsStrLitContext() {}

func NewStrLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StrLitContext {
	var p = new(StrLitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_strLit

	return p
}

func (s *StrLitContext) GetParser() antlr.Parser { return s.parser }

func (s *StrLitContext) STR_LIT() antlr.TerminalNode {
	return s.GetToken(kdsParserSTR_LIT, 0)
}

func (s *StrLitContext) PROTO3_LIT_SINGLE() antlr.TerminalNode {
	return s.GetToken(kdsParserPROTO3_LIT_SINGLE, 0)
}

func (s *StrLitContext) PROTO3_LIT_DOBULE() antlr.TerminalNode {
	return s.GetToken(kdsParserPROTO3_LIT_DOBULE, 0)
}

func (s *StrLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StrLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterStrLit(s)
	}
}

func (s *StrLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitStrLit(s)
	}
}

func (p *kdsParser) StrLit() (localctx IStrLitContext) {
	localctx = NewStrLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, kdsParserRULE_strLit)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(308)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8796294348800) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBoolLitContext is an interface to support dynamic dispatch.
type IBoolLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BOOL_LIT() antlr.TerminalNode

	// IsBoolLitContext differentiates from other interfaces.
	IsBoolLitContext()
}

type BoolLitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolLitContext() *BoolLitContext {
	var p = new(BoolLitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_boolLit
	return p
}

func InitEmptyBoolLitContext(p *BoolLitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_boolLit
}

func (*BoolLitContext) IsBoolLitContext() {}

func NewBoolLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolLitContext {
	var p = new(BoolLitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_boolLit

	return p
}

func (s *BoolLitContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolLitContext) BOOL_LIT() antlr.TerminalNode {
	return s.GetToken(kdsParserBOOL_LIT, 0)
}

func (s *BoolLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterBoolLit(s)
	}
}

func (s *BoolLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitBoolLit(s)
	}
}

func (p *kdsParser) BoolLit() (localctx IBoolLitContext) {
	localctx = NewBoolLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, kdsParserRULE_boolLit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(310)
		p.Match(kdsParserBOOL_LIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFloatLitContext is an interface to support dynamic dispatch.
type IFloatLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FLOAT_LIT() antlr.TerminalNode

	// IsFloatLitContext differentiates from other interfaces.
	IsFloatLitContext()
}

type FloatLitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatLitContext() *FloatLitContext {
	var p = new(FloatLitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_floatLit
	return p
}

func InitEmptyFloatLitContext(p *FloatLitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_floatLit
}

func (*FloatLitContext) IsFloatLitContext() {}

func NewFloatLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatLitContext {
	var p = new(FloatLitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_floatLit

	return p
}

func (s *FloatLitContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatLitContext) FLOAT_LIT() antlr.TerminalNode {
	return s.GetToken(kdsParserFLOAT_LIT, 0)
}

func (s *FloatLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterFloatLit(s)
	}
}

func (s *FloatLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitFloatLit(s)
	}
}

func (p *kdsParser) FloatLit() (localctx IFloatLitContext) {
	localctx = NewFloatLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, kdsParserRULE_floatLit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(312)
		p.Match(kdsParserFLOAT_LIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IKeywordsContext is an interface to support dynamic dispatch.
type IKeywordsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PACKAGE() antlr.TerminalNode
	MAP() antlr.TerminalNode
	INT32() antlr.TerminalNode
	INT64() antlr.TerminalNode
	UINT32() antlr.TerminalNode
	UINT64() antlr.TerminalNode
	SINT32() antlr.TerminalNode
	SINT64() antlr.TerminalNode
	FIXED32() antlr.TerminalNode
	FIXED64() antlr.TerminalNode
	SFIXED32() antlr.TerminalNode
	SFIXED64() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	STRING() antlr.TerminalNode
	DOUBLE() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	BYTES() antlr.TerminalNode
	ENUM() antlr.TerminalNode
	ENTITY() antlr.TerminalNode
	COMPONENT() antlr.TerminalNode
	BOOL_LIT() antlr.TerminalNode

	// IsKeywordsContext differentiates from other interfaces.
	IsKeywordsContext()
}

type KeywordsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordsContext() *KeywordsContext {
	var p = new(KeywordsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_keywords
	return p
}

func InitEmptyKeywordsContext(p *KeywordsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = kdsParserRULE_keywords
}

func (*KeywordsContext) IsKeywordsContext() {}

func NewKeywordsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordsContext {
	var p = new(KeywordsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = kdsParserRULE_keywords

	return p
}

func (s *KeywordsContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordsContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(kdsParserPACKAGE, 0)
}

func (s *KeywordsContext) MAP() antlr.TerminalNode {
	return s.GetToken(kdsParserMAP, 0)
}

func (s *KeywordsContext) INT32() antlr.TerminalNode {
	return s.GetToken(kdsParserINT32, 0)
}

func (s *KeywordsContext) INT64() antlr.TerminalNode {
	return s.GetToken(kdsParserINT64, 0)
}

func (s *KeywordsContext) UINT32() antlr.TerminalNode {
	return s.GetToken(kdsParserUINT32, 0)
}

func (s *KeywordsContext) UINT64() antlr.TerminalNode {
	return s.GetToken(kdsParserUINT64, 0)
}

func (s *KeywordsContext) SINT32() antlr.TerminalNode {
	return s.GetToken(kdsParserSINT32, 0)
}

func (s *KeywordsContext) SINT64() antlr.TerminalNode {
	return s.GetToken(kdsParserSINT64, 0)
}

func (s *KeywordsContext) FIXED32() antlr.TerminalNode {
	return s.GetToken(kdsParserFIXED32, 0)
}

func (s *KeywordsContext) FIXED64() antlr.TerminalNode {
	return s.GetToken(kdsParserFIXED64, 0)
}

func (s *KeywordsContext) SFIXED32() antlr.TerminalNode {
	return s.GetToken(kdsParserSFIXED32, 0)
}

func (s *KeywordsContext) SFIXED64() antlr.TerminalNode {
	return s.GetToken(kdsParserSFIXED64, 0)
}

func (s *KeywordsContext) BOOL() antlr.TerminalNode {
	return s.GetToken(kdsParserBOOL, 0)
}

func (s *KeywordsContext) STRING() antlr.TerminalNode {
	return s.GetToken(kdsParserSTRING, 0)
}

func (s *KeywordsContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(kdsParserDOUBLE, 0)
}

func (s *KeywordsContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(kdsParserFLOAT, 0)
}

func (s *KeywordsContext) BYTES() antlr.TerminalNode {
	return s.GetToken(kdsParserBYTES, 0)
}

func (s *KeywordsContext) ENUM() antlr.TerminalNode {
	return s.GetToken(kdsParserENUM, 0)
}

func (s *KeywordsContext) ENTITY() antlr.TerminalNode {
	return s.GetToken(kdsParserENTITY, 0)
}

func (s *KeywordsContext) COMPONENT() antlr.TerminalNode {
	return s.GetToken(kdsParserCOMPONENT, 0)
}

func (s *KeywordsContext) BOOL_LIT() antlr.TerminalNode {
	return s.GetToken(kdsParserBOOL_LIT, 0)
}

func (s *KeywordsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.EnterKeywords(s)
	}
}

func (s *KeywordsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(kdsListener); ok {
		listenerT.ExitKeywords(s)
	}
}

func (p *kdsParser) Keywords() (localctx IKeywordsContext) {
	localctx = NewKeywordsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, kdsParserRULE_keywords)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(314)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17592253153168) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
