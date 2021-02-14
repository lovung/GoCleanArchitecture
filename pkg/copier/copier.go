package copier

import "github.com/jinzhu/copier"

// MustCopy panic if copy the wrong structure
func MustCopy(toValue, fromValue interface{}) {
	err := copier.Copy(toValue, fromValue)
	if err != nil {
		panic(err)
	}
}
