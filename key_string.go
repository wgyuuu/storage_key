package storage_key

import "strings"

type String string

func (this String) ToString() string {
	return string(this)
}

func (this String) ToStringList() []string {
	return []string{this.ToString()}
}

type StringList []String

func (this StringList) ToString() string {
	return strings.Join(this.ToStringList(), SEPARATE)
}

func (this StringList) ToStringList() []string {
	strList := make([]string, len(this))
	for idx, strObj := range this {
		strList[idx] = strObj.ToString()
	}
	return strList
}
