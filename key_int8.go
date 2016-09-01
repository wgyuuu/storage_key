package storage_key

import (
	"strconv"
	"strings"
)

type Int8 int8

func (this Int8) ToString() string {
	return strconv.Itoa(int(this))
}

func (this Int8) ToStringList() []string {
	return []string{this.ToString()}
}

type Int8List []Int8

func (this Int8List) ToString() string {
	return strings.Join(this.ToStringList(), SEPARATE)
}

func (this Int8List) ToStringList() (strList []string) {
	for _, obj := range this {
		strList = append(strList, obj.ToString())
	}
	return
}
