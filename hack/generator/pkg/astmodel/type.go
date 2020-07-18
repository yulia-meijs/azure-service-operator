/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/ast"
)

// Type represents something that is a Go type
type Type interface {
	// RequiredImports returns a list of packages required by this type
	RequiredImports() []*PackageReference

	// References returns the names of all types that this type
	// references. For example, an Array of Persons references a
	// Person.
	References() TypeNameSet

	// AsType renders as a Go abstract syntax tree for a type
	// (yes this says ast.Expr but that is what the Go 'ast' package uses for types)
	AsType(codeGenerationContext *CodeGenerationContext) ast.Expr

	// AsDeclarations renders as a Go abstract syntax tree for a declaration
	AsDeclarations(codeGenerationContext *CodeGenerationContext, name *TypeName, description *string) []ast.Decl

	// Equals returns true if the passed type is the same as this one, false otherwise
	Equals(t Type) bool

	// CreateNamedDefinition gives a name to the type and might generate some associated definitions as well (the second result)
	// that also must be included in the output.
	CreateNamedDefinition(name *TypeName, idFactory IdentifierFactory) (TypeDefinition, []TypeDefinition)

	// NameInternalDefinitions creates definitions for nested types where needed (e.g. nested anonymous enums, structs),
	// and returns the new, updated type to use in this type’s place.
	NameInternalDefinitions(nameHint *TypeName, idFactory IdentifierFactory) (Type, []TypeDefinition)
}
