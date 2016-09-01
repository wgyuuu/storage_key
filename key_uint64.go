package storage_key

import (
	"strconv"
	"strings"
)

type Uint64 uint64

func (this Uint64) ToString() string {
	return strconv.Itoa(int(this))
}

func (this Uint64) ToStringList() []string {
	return []string{this.ToString()}
}

type Uint64List []Uint64

func (this Uint64List) ToString() string {
	return strings.Join(this.ToStringList(), SEPARATE)
}

func (this Uint64List) ToStringList() (strList []string) {
	for _, obj := range this {
		strList = append(strList, obj.ToString())
	}
	return
}
