// Code generated from kds.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // kds

import "github.com/antlr4-go/antlr/v4"

// kdsListener is a complete listener for a parse tree produced by kdsParser.
type kdsListener interface {
	antlr.ParseTreeListener

	// EnterKds is called when entering the kds production.
	EnterKds(c *KdsContext)

	// EnterPackageStatement is called when entering the packageStatement production.
	EnterPackageStatement(c *PackageStatementContext)

	// EnterImportStatement is called when entering the importStatement production.
	EnterImportStatement(c *ImportStatementContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterFieldLabel is called when entering the fieldLabel production.
	EnterFieldLabel(c *FieldLabelContext)

	// EnterFieldNumber is called when entering the fieldNumber production.
	EnterFieldNumber(c *FieldNumberContext)

	// EnterMapField is called when entering the mapField production.
	EnterMapField(c *MapFieldContext)

	// EnterKeyType is called when entering the keyType production.
	EnterKeyType(c *KeyTypeContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterTopLevelDef is called when entering the topLevelDef production.
	EnterTopLevelDef(c *TopLevelDefContext)

	// EnterEnumDef is called when entering the enumDef production.
	EnterEnumDef(c *EnumDefContext)

	// EnterEnumBody is called when entering the enumBody production.
	EnterEnumBody(c *EnumBodyContext)

	// EnterEnumElement is called when entering the enumElement production.
	EnterEnumElement(c *EnumElementContext)

	// EnterEnumField is called when entering the enumField production.
	EnterEnumField(c *EnumFieldContext)

	// EnterEntityDef is called when entering the entityDef production.
	EnterEntityDef(c *EntityDefContext)

	// EnterEntityName is called when entering the entityName production.
	EnterEntityName(c *EntityNameContext)

	// EnterEntityBody is called when entering the entityBody production.
	EnterEntityBody(c *EntityBodyContext)

	// EnterEntityElement is called when entering the entityElement production.
	EnterEntityElement(c *EntityElementContext)

	// EnterComponentDef is called when entering the componentDef production.
	EnterComponentDef(c *ComponentDefContext)

	// EnterComponentName is called when entering the componentName production.
	EnterComponentName(c *ComponentNameContext)

	// EnterComponentBody is called when entering the componentBody production.
	EnterComponentBody(c *ComponentBodyContext)

	// EnterComponentElement is called when entering the componentElement production.
	EnterComponentElement(c *ComponentElementContext)

	// EnterEmptyStatement_ is called when entering the emptyStatement_ production.
	EnterEmptyStatement_(c *EmptyStatement_Context)

	// EnterIdent is called when entering the ident production.
	EnterIdent(c *IdentContext)

	// EnterFullIdent is called when entering the fullIdent production.
	EnterFullIdent(c *FullIdentContext)

	// EnterFieldName is called when entering the fieldName production.
	EnterFieldName(c *FieldNameContext)

	// EnterMessageName is called when entering the messageName production.
	EnterMessageName(c *MessageNameContext)

	// EnterEnumName is called when entering the enumName production.
	EnterEnumName(c *EnumNameContext)

	// EnterMapName is called when entering the mapName production.
	EnterMapName(c *MapNameContext)

	// EnterMessageType is called when entering the messageType production.
	EnterMessageType(c *MessageTypeContext)

	// EnterEnumType is called when entering the enumType production.
	EnterEnumType(c *EnumTypeContext)

	// EnterIntLit is called when entering the intLit production.
	EnterIntLit(c *IntLitContext)

	// EnterKeywords is called when entering the keywords production.
	EnterKeywords(c *KeywordsContext)

	// ExitKds is called when exiting the kds production.
	ExitKds(c *KdsContext)

	// ExitPackageStatement is called when exiting the packageStatement production.
	ExitPackageStatement(c *PackageStatementContext)

	// ExitImportStatement is called when exiting the importStatement production.
	ExitImportStatement(c *ImportStatementContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitFieldLabel is called when exiting the fieldLabel production.
	ExitFieldLabel(c *FieldLabelContext)

	// ExitFieldNumber is called when exiting the fieldNumber production.
	ExitFieldNumber(c *FieldNumberContext)

	// ExitMapField is called when exiting the mapField production.
	ExitMapField(c *MapFieldContext)

	// ExitKeyType is called when exiting the keyType production.
	ExitKeyType(c *KeyTypeContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitTopLevelDef is called when exiting the topLevelDef production.
	ExitTopLevelDef(c *TopLevelDefContext)

	// ExitEnumDef is called when exiting the enumDef production.
	ExitEnumDef(c *EnumDefContext)

	// ExitEnumBody is called when exiting the enumBody production.
	ExitEnumBody(c *EnumBodyContext)

	// ExitEnumElement is called when exiting the enumElement production.
	ExitEnumElement(c *EnumElementContext)

	// ExitEnumField is called when exiting the enumField production.
	ExitEnumField(c *EnumFieldContext)

	// ExitEntityDef is called when exiting the entityDef production.
	ExitEntityDef(c *EntityDefContext)

	// ExitEntityName is called when exiting the entityName production.
	ExitEntityName(c *EntityNameContext)

	// ExitEntityBody is called when exiting the entityBody production.
	ExitEntityBody(c *EntityBodyContext)

	// ExitEntityElement is called when exiting the entityElement production.
	ExitEntityElement(c *EntityElementContext)

	// ExitComponentDef is called when exiting the componentDef production.
	ExitComponentDef(c *ComponentDefContext)

	// ExitComponentName is called when exiting the componentName production.
	ExitComponentName(c *ComponentNameContext)

	// ExitComponentBody is called when exiting the componentBody production.
	ExitComponentBody(c *ComponentBodyContext)

	// ExitComponentElement is called when exiting the componentElement production.
	ExitComponentElement(c *ComponentElementContext)

	// ExitEmptyStatement_ is called when exiting the emptyStatement_ production.
	ExitEmptyStatement_(c *EmptyStatement_Context)

	// ExitIdent is called when exiting the ident production.
	ExitIdent(c *IdentContext)

	// ExitFullIdent is called when exiting the fullIdent production.
	ExitFullIdent(c *FullIdentContext)

	// ExitFieldName is called when exiting the fieldName production.
	ExitFieldName(c *FieldNameContext)

	// ExitMessageName is called when exiting the messageName production.
	ExitMessageName(c *MessageNameContext)

	// ExitEnumName is called when exiting the enumName production.
	ExitEnumName(c *EnumNameContext)

	// ExitMapName is called when exiting the mapName production.
	ExitMapName(c *MapNameContext)

	// ExitMessageType is called when exiting the messageType production.
	ExitMessageType(c *MessageTypeContext)

	// ExitEnumType is called when exiting the enumType production.
	ExitEnumType(c *EnumTypeContext)

	// ExitIntLit is called when exiting the intLit production.
	ExitIntLit(c *IntLitContext)

	// ExitKeywords is called when exiting the keywords production.
	ExitKeywords(c *KeywordsContext)
}
