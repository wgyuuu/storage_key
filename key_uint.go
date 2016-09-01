package storage_key

import (
	"strconv"
	"strings"
)

type Uint uint

func (this Uint) ToString() string {
	return strconv.Itoa(int(this))
}

func (this Uint) ToStringList() []string {
	return []string{this.ToString()}
}

type UintList []Uint

func (this UintList) ToString() string {
	return strings.Join(this.ToStringList(), SEPARATE)
}

func (this UintList) ToStringList() (strList []string) {
	for _, obj := range this {
		strList = append(strList, obj.ToString())
	}
	return
}