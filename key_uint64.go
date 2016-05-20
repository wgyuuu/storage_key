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

func (this Uint64List) ToStringList() []string {
	strList := make([]string, len(this))
	for idx, u64Obj := range this {
		strList[idx] = u64Obj.ToString()
	}
	return strList
}
