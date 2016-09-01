package storage_key

import (
	"strconv"
	"strings"
)

type Uint16 uint16

func (this Uint16) ToString() string {
	return strconv.Itoa(int(this))
}

func (this Uint16) ToStringList() []string {
	return []string{this.ToString()}
}

type Uint16List []Uint16

func (this Uint16List) ToString() string {
	return strings.Join(this.ToStringList(), SEPARATE)
}

func (this Uint16List) ToStringList() (strList []string) {
	for _, obj := range this {
		strList = append(strList, obj.ToString())
	}
	return
}