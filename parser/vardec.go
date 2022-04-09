package parser

import (
	"reflect"
)

type Vardec struct {
	Name string
}

func NewVardec(name string) *Vardec {
	return &Vardec{Name: name}
}

func (VD *Vardec) Equals(other interface{}) bool {
	castOther := reflect.ValueOf(other)
	otherVardec, ok := reflect.Indirect(castOther).Interface().(Vardec)
	return ok && reflect.TypeOf(other) == reflect.TypeOf(VD) && VD.Name == otherVardec.Name
}

// func (VD *Vardec) String() string {
// 	return fmt.Sprintf("Vardec(%s)", VD.Name)
// }
