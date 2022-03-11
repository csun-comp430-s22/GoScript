package token

import "reflect"

type VarToken struct {
}

func (vrt *VarToken) Equals(other interface{}) bool {
	return reflect.TypeOf(other) == reflect.TypeOf(vrt)
}

func (vrt *VarToken) String() string {
	return "var"
}
