package storage_key

import (
	"strconv"
	"strings"
)

type Int16 int16

func (this Int16) ToString() string {
	return strconv.Itoa(int(this))
}

func (this Int16) ToStringList() []string {
	return []string{this.ToString()}
}

type Int16List []Int16

func (this Int16List) ToString() string {
	return strings.Join(this.ToStringList(), SEPARATE)
}

func (this Int16List) ToStringList() (strList []string) {
	for _, obj := range this {
		strList = append(strList, obj.ToString())
	}
	return
}
