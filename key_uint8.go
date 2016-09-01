package storage_key

import (
	"strconv"
	"strings"
)

type Uint8 uint8

func (this Uint8) ToString() string {
	return strconv.Itoa(int(this))
}

func (this Uint8) ToStringList() []string {
	return []string{this.ToString()}
}

type Uint8List []Uint8

func (this Uint8List) ToString() string {
	return strings.Join(this.ToStringList(), SEPARATE)
}

func (this Uint8List) ToStringList() (strList []string) {
	for _, obj := range this {
		strList = append(strList, obj.ToString())
	}
	return
}