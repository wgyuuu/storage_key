package storage_key

import (
	"strconv"
	"strings"
)

type Int32 int32

func (this Int32) ToString() string {
	return strconv.Itoa(int(this))
}

func (this Int32) ToStringList() []string {
	return []string{this.ToString()}
}

type Int32List []Int32

func (this Int32List) ToString() string {
	return strings.Join(this.ToStringList(), SEPARATE)
}

func (this Int32List) ToStringList() []string {
	strList := make([]string, len(this))
	for idx, i32Obj := range this {
		strList[idx] = i32Obj.ToString()
	}
	return strList
}
