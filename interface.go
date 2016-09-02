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

func NewKeyForString(data string) Key {
	index := strings.Index(data, SEPARATE)
	if index == -1 {
		return String(data)
	}

	strList := strings.Split(data, SEPARATE)
	return KeyListForStrList(strList)
}

func KeyListForStrList(strList []string) KeyList {
	keyList := make([]Key, len(strList))
	for k, str := range strList {
		keyList[k] = String(str)
	}
	return NewKeyList(keyList...)
}

func NewKeyList(key ...Key) KeyList {
	return KeyList(key)
}

func AppendKey(keyList KeyList, key Key) KeyList {
	for _, k := range keyList {
		if k.ToString() == key.ToString() {
			return keyList
		}
	}
	keyList = append(keyList, key)
	return keyList
}
