package storage_key

import "strings"

const SEPARATE string = "|"

type Key interface {
	ToString() string
	ToStringList() []string
}

type KeyList []Key

func (this KeyList) ToString() string {
	return strings.Join(this.ToStringList(), SEPARATE)
}

func (this KeyList) ToStringList() []string {
	strList := make([]string, len(this))
	for idx, keyObj := range this {
		strList[idx] = keyObj.ToString()
	}
    return strList
}
