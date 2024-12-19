package main

var encryptedKeyLength = 128

type MasterKey struct {
	MasterKey []byte
}

type KeyMetaData struct {
	TimeStamp int64  `json:"creationTime"`
	KeyId     string `json:"keyId"`
	IsEnabled bool   `json:"isEnabled"`
}
type Key struct {
	KeyMetaData KeyMetaData `json:"keyMetaData"`
	Key         []byte      `json:"key"`
	IV          []byte      `json:"iv"`
}

type KeyProvider interface {
	CreateKey(keyMetaData string) (KeyMetaData, error)
	DeleteKey(keyMetaData KeyMetaData) (bool, error)
	SaveKey(key Key) (bool, error)
}
