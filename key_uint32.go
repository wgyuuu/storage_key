package storage_key

import (
	"strconv"
	"strings"
)

type Uint32 uint32

func (this Uint32) ToString() string {
	return strconv.Itoa(int(this))
}

func (this Uint32) ToStringList() []string {
	return []string{this.ToString()}
}

type Uint32List []Uint32

func (this Uint32List) ToString() string {
	return strings.Join(this.ToStringList(), SEPARATE)
}

func (this Uint32List) ToStringList() (strList []string) {
	for _, obj := range this {
		strList = append(strList, obj.ToString())
	}
	return
}