package storage_key

import (
	"strconv"
	"strings"
)

type Int int

func (this Int) ToString() string {
	return strconv.Itoa(int(this))
}

func (this Int) ToStringList() []string {
	return []string{this.ToString()}
}

type IntList []Int

func (this IntList) ToString() string {
	return strings.Join(this.ToStringList(), SEPARATE)
}

func (this IntList) ToStringList() (strList []string) {
	for _, obj := range this {
		strList = append(strList, obj.ToString())
	}
	return
}
