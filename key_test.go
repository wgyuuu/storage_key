package storage_key

import "testing"

func TestKeyList(t *testing.T) {
	keyList := KeyList([]Key{Int(1), Int32(32), String("aa"), Uint64(64)})
	keyList = append(keyList, Int64(664))
	t.Log(keyList.ToString())
	t.Log(keyList.ToStringList())
}

func TestKeyListNew(t *testing.T) {
	keyList := NewKeyList(Int(1), Int(2), String("aa"))
	t.Log(keyList.ToStringList())
}

func TestToKeyList(t *testing.T) {
	keyList := NewKeyList(Int(1), Int(2), String("aa"))
	printStrList(t, keyList)
}

func TestKeyListForString(t *testing.T) {
	keyList := NewKeyList(Int(1), Int(2), String("aa"))
	t.Log(NewKeyForString(keyList.ToString()))
}

func printStrList(t *testing.T, key Key) {
	t.Log(key.ToStringList())
}
