package storage_key

import "testing"

func TestKeyList(t *testing.T) {
	keyList := KeyList([]Key{Int(1), Int32(32), String("aa"), Uint64(64)})
    keyList = append(keyList, Int64(664))
	t.Log(keyList.ToString())
	t.Log(keyList.ToStringList())
}
