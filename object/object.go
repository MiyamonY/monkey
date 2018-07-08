// Package object provides
//
// File:  object.go
// Author: ymiyamoto
//
// Created on Sun Jul  8 15:50:08 2018
//
package object

import "fmt"

type ObjectType string

const (
	NULL_OBJ    = "NULL"
	INTEGER_OBJ = "INTEGER"
	BOOLEAN     = "BOOLEAN"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType {
	return BOOLEAN
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

type Null struct {
}

func (n *Null) Type() ObjectType {
	return NULL_OBJ
}

func (n *Null) Inspect() string {
	return "null"
}
