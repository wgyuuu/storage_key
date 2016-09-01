package storage_key

import "strings"

type String string

func (this String) ToString() string {
	return string(this)
}

func (this String) ToStringList() []string {
	return []string{this.ToString()}
}

func (this String) ToSum() (value int) {
	for _, byt := range this {
		value += int(byt)
	}
	return
}

type StringList []String

func (this StringList) ToString() string {
	return strings.Join(this.ToStringList(), SEPARATE)
}

func (this StringList) ToStringList() (strList []string) {
	for _, obj := range this {
		strList = append(strList, obj.ToString())
	}
	return
}
