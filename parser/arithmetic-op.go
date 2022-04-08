package parser

import "reflect"

type PlusOp struct {
}

func (po *PlusOp) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(po)
}

type MinusOp struct {
}

func (mst *MinusOp) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(mst)
}

type MultiplyOp struct {
}

func (mo *MultiplyOp) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(mo)
}

type DivideOp struct {
}

func (do *DivideOp) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(do)
}

type ModOp struct {
}

func (mo *ModOp) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(mo)
}

type PowOp struct {
}

func (po *PowOp) Equals(other any) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(po)
}
