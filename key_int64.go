package storage_key

import (
	"strconv"
	"strings"
)

type Int64 int64

func (this Int64) ToString() string {
	return strconv.Itoa(int(this))
}

func (this Int64) ToStringList() []string {
	return []string{this.ToString()}
}

type Int64List []Int64

func (this Int64List) ToString() string {
	return strings.Join(this.ToStringList(), SEPARATE)
}

func (this Int64List) ToStringList() (strList []string) {
	for _, obj := range this {
		strList = append(strList, obj.ToString())
	}
	return
}
