// Code generated from kds.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // kds

import "github.com/antlr4-go/antlr/v4"

// BasekdsListener is a complete listener for a parse tree produced by kdsParser.
type BasekdsListener struct{}

var _ kdsListener = &BasekdsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasekdsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasekdsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasekdsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasekdsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterKds is called when production kds is entered.
func (s *BasekdsListener) EnterKds(ctx *KdsContext) {}

// ExitKds is called when production kds is exited.
func (s *BasekdsListener) ExitKds(ctx *KdsContext) {}

// EnterImportStatement is called when production importStatement is entered.
func (s *BasekdsListener) EnterImportStatement(ctx *ImportStatementContext) {}

// ExitImportStatement is called when production importStatement is exited.
func (s *BasekdsListener) ExitImportStatement(ctx *ImportStatementContext) {}

// EnterImportElement is called when production importElement is entered.
func (s *BasekdsListener) EnterImportElement(ctx *ImportElementContext) {}

// ExitImportElement is called when production importElement is exited.
func (s *BasekdsListener) ExitImportElement(ctx *ImportElementContext) {}

// EnterPackageStatement is called when production packageStatement is entered.
func (s *BasekdsListener) EnterPackageStatement(ctx *PackageStatementContext) {}

// ExitPackageStatement is called when production packageStatement is exited.
func (s *BasekdsListener) ExitPackageStatement(ctx *PackageStatementContext) {}

// EnterOptionStatement is called when production optionStatement is entered.
func (s *BasekdsListener) EnterOptionStatement(ctx *OptionStatementContext) {}

// ExitOptionStatement is called when production optionStatement is exited.
func (s *BasekdsListener) ExitOptionStatement(ctx *OptionStatementContext) {}

// EnterOptionName is called when production optionName is entered.
func (s *BasekdsListener) EnterOptionName(ctx *OptionNameContext) {}

// ExitOptionName is called when production optionName is exited.
func (s *BasekdsListener) ExitOptionName(ctx *OptionNameContext) {}

// EnterField is called when production field is entered.
func (s *BasekdsListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BasekdsListener) ExitField(ctx *FieldContext) {}

// EnterFieldLabel is called when production fieldLabel is entered.
func (s *BasekdsListener) EnterFieldLabel(ctx *FieldLabelContext) {}

// ExitFieldLabel is called when production fieldLabel is exited.
func (s *BasekdsListener) ExitFieldLabel(ctx *FieldLabelContext) {}

// EnterFieldNumber is called when production fieldNumber is entered.
func (s *BasekdsListener) EnterFieldNumber(ctx *FieldNumberContext) {}

// ExitFieldNumber is called when production fieldNumber is exited.
func (s *BasekdsListener) ExitFieldNumber(ctx *FieldNumberContext) {}

// EnterMapField is called when production mapField is entered.
func (s *BasekdsListener) EnterMapField(ctx *MapFieldContext) {}

// ExitMapField is called when production mapField is exited.
func (s *BasekdsListener) ExitMapField(ctx *MapFieldContext) {}

// EnterKeyType is called when production keyType is entered.
func (s *BasekdsListener) EnterKeyType(ctx *KeyTypeContext) {}

// ExitKeyType is called when production keyType is exited.
func (s *BasekdsListener) ExitKeyType(ctx *KeyTypeContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BasekdsListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BasekdsListener) ExitType_(ctx *Type_Context) {}

// EnterTopLevelDef is called when production topLevelDef is entered.
func (s *BasekdsListener) EnterTopLevelDef(ctx *TopLevelDefContext) {}

// ExitTopLevelDef is called when production topLevelDef is exited.
func (s *BasekdsListener) ExitTopLevelDef(ctx *TopLevelDefContext) {}

// EnterEnumDef is called when production enumDef is entered.
func (s *BasekdsListener) EnterEnumDef(ctx *EnumDefContext) {}

// ExitEnumDef is called when production enumDef is exited.
func (s *BasekdsListener) ExitEnumDef(ctx *EnumDefContext) {}

// EnterEnumBody is called when production enumBody is entered.
func (s *BasekdsListener) EnterEnumBody(ctx *EnumBodyContext) {}

// ExitEnumBody is called when production enumBody is exited.
func (s *BasekdsListener) ExitEnumBody(ctx *EnumBodyContext) {}

// EnterEnumElement is called when production enumElement is entered.
func (s *BasekdsListener) EnterEnumElement(ctx *EnumElementContext) {}

// ExitEnumElement is called when production enumElement is exited.
func (s *BasekdsListener) ExitEnumElement(ctx *EnumElementContext) {}

// EnterEnumField is called when production enumField is entered.
func (s *BasekdsListener) EnterEnumField(ctx *EnumFieldContext) {}

// ExitEnumField is called when production enumField is exited.
func (s *BasekdsListener) ExitEnumField(ctx *EnumFieldContext) {}

// EnterEntityDef is called when production entityDef is entered.
func (s *BasekdsListener) EnterEntityDef(ctx *EntityDefContext) {}

// ExitEntityDef is called when production entityDef is exited.
func (s *BasekdsListener) ExitEntityDef(ctx *EntityDefContext) {}

// EnterEntityName is called when production entityName is entered.
func (s *BasekdsListener) EnterEntityName(ctx *EntityNameContext) {}

// ExitEntityName is called when production entityName is exited.
func (s *BasekdsListener) ExitEntityName(ctx *EntityNameContext) {}

// EnterEntityBody is called when production entityBody is entered.
func (s *BasekdsListener) EnterEntityBody(ctx *EntityBodyContext) {}

// ExitEntityBody is called when production entityBody is exited.
func (s *BasekdsListener) ExitEntityBody(ctx *EntityBodyContext) {}

// EnterEntityElement is called when production entityElement is entered.
func (s *BasekdsListener) EnterEntityElement(ctx *EntityElementContext) {}

// ExitEntityElement is called when production entityElement is exited.
func (s *BasekdsListener) ExitEntityElement(ctx *EntityElementContext) {}

// EnterComponentDef is called when production componentDef is entered.
func (s *BasekdsListener) EnterComponentDef(ctx *ComponentDefContext) {}

// ExitComponentDef is called when production componentDef is exited.
func (s *BasekdsListener) ExitComponentDef(ctx *ComponentDefContext) {}

// EnterComponentName is called when production componentName is entered.
func (s *BasekdsListener) EnterComponentName(ctx *ComponentNameContext) {}

// ExitComponentName is called when production componentName is exited.
func (s *BasekdsListener) ExitComponentName(ctx *ComponentNameContext) {}

// EnterComponentBody is called when production componentBody is entered.
func (s *BasekdsListener) EnterComponentBody(ctx *ComponentBodyContext) {}

// ExitComponentBody is called when production componentBody is exited.
func (s *BasekdsListener) ExitComponentBody(ctx *ComponentBodyContext) {}

// EnterComponentElement is called when production componentElement is entered.
func (s *BasekdsListener) EnterComponentElement(ctx *ComponentElementContext) {}

// ExitComponentElement is called when production componentElement is exited.
func (s *BasekdsListener) ExitComponentElement(ctx *ComponentElementContext) {}

// EnterConstant is called when production constant is entered.
func (s *BasekdsListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BasekdsListener) ExitConstant(ctx *ConstantContext) {}

// EnterBlockLit is called when production blockLit is entered.
func (s *BasekdsListener) EnterBlockLit(ctx *BlockLitContext) {}

// ExitBlockLit is called when production blockLit is exited.
func (s *BasekdsListener) ExitBlockLit(ctx *BlockLitContext) {}

// EnterEmptyStatement_ is called when production emptyStatement_ is entered.
func (s *BasekdsListener) EnterEmptyStatement_(ctx *EmptyStatement_Context) {}

// ExitEmptyStatement_ is called when production emptyStatement_ is exited.
func (s *BasekdsListener) ExitEmptyStatement_(ctx *EmptyStatement_Context) {}

// EnterIdent is called when production ident is entered.
func (s *BasekdsListener) EnterIdent(ctx *IdentContext) {}

// ExitIdent is called when production ident is exited.
func (s *BasekdsListener) ExitIdent(ctx *IdentContext) {}

// EnterFullIdent is called when production fullIdent is entered.
func (s *BasekdsListener) EnterFullIdent(ctx *FullIdentContext) {}

// ExitFullIdent is called when production fullIdent is exited.
func (s *BasekdsListener) ExitFullIdent(ctx *FullIdentContext) {}

// EnterFieldName is called when production fieldName is entered.
func (s *BasekdsListener) EnterFieldName(ctx *FieldNameContext) {}

// ExitFieldName is called when production fieldName is exited.
func (s *BasekdsListener) ExitFieldName(ctx *FieldNameContext) {}

// EnterMessageName is called when production messageName is entered.
func (s *BasekdsListener) EnterMessageName(ctx *MessageNameContext) {}

// ExitMessageName is called when production messageName is exited.
func (s *BasekdsListener) ExitMessageName(ctx *MessageNameContext) {}

// EnterEnumName is called when production enumName is entered.
func (s *BasekdsListener) EnterEnumName(ctx *EnumNameContext) {}

// ExitEnumName is called when production enumName is exited.
func (s *BasekdsListener) ExitEnumName(ctx *EnumNameContext) {}

// EnterMapName is called when production mapName is entered.
func (s *BasekdsListener) EnterMapName(ctx *MapNameContext) {}

// ExitMapName is called when production mapName is exited.
func (s *BasekdsListener) ExitMapName(ctx *MapNameContext) {}

// EnterMessageType is called when production messageType is entered.
func (s *BasekdsListener) EnterMessageType(ctx *MessageTypeContext) {}

// ExitMessageType is called when production messageType is exited.
func (s *BasekdsListener) ExitMessageType(ctx *MessageTypeContext) {}

// EnterEnumType is called when production enumType is entered.
func (s *BasekdsListener) EnterEnumType(ctx *EnumTypeContext) {}

// ExitEnumType is called when production enumType is exited.
func (s *BasekdsListener) ExitEnumType(ctx *EnumTypeContext) {}

// EnterIntLit is called when production intLit is entered.
func (s *BasekdsListener) EnterIntLit(ctx *IntLitContext) {}

// ExitIntLit is called when production intLit is exited.
func (s *BasekdsListener) ExitIntLit(ctx *IntLitContext) {}

// EnterStrLit is called when production strLit is entered.
func (s *BasekdsListener) EnterStrLit(ctx *StrLitContext) {}

// ExitStrLit is called when production strLit is exited.
func (s *BasekdsListener) ExitStrLit(ctx *StrLitContext) {}

// EnterBoolLit is called when production boolLit is entered.
func (s *BasekdsListener) EnterBoolLit(ctx *BoolLitContext) {}

// ExitBoolLit is called when production boolLit is exited.
func (s *BasekdsListener) ExitBoolLit(ctx *BoolLitContext) {}

// EnterFloatLit is called when production floatLit is entered.
func (s *BasekdsListener) EnterFloatLit(ctx *FloatLitContext) {}

// ExitFloatLit is called when production floatLit is exited.
func (s *BasekdsListener) ExitFloatLit(ctx *FloatLitContext) {}

// EnterKeywords is called when production keywords is entered.
func (s *BasekdsListener) EnterKeywords(ctx *KeywordsContext) {}

// ExitKeywords is called when production keywords is exited.
func (s *BasekdsListener) ExitKeywords(ctx *KeywordsContext) {}
