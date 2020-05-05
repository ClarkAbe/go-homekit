// Code generated by running "go generate" in ccms/x/text. DO NOT EDIT.

package language

// This file contains code common to the maketables.go and the package code.

// AliasType is the type of an alias in AliasMap.
type AliasType int8

const (
	Deprecated AliasType = iota
	Macro
	Legacy

	AliasTypeUnknown AliasType = -1
)