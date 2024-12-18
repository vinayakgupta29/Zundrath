package main

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/google/uuid"
)

var KeyArr []Key = make([]Key, 0)

func CreateKey() (KeyMetaData, error) {

	keyID := uuid.New().String()
	keyMetaData := KeyMetaData{
		TimeStamp: time.Now().UnixNano(),
		KeyId:     keyID,
		IsEnabled: true,
	}
	aesKey := GenerateAesKey()
	key := Key{
		KeyMetaData: keyMetaData,
		Key:         aesKey,
	}
	_, err := SaveKey(key)

	if err != nil {
		return KeyMetaData{}, err
	}
	return keyMetaData, nil
}

func SaveKey(key Key) (bool, error) {
	keyJson, err := json.Marshal(key)
	if err != nil {
		return false, err
	}
	fmt.Println(string(keyJson))
	KeyArr = append(KeyArr, key)
	return true, nil
}
func DeleteKey(keyMetaData KeyMetaData) (bool, error) {
	found := false
	for i, key := range KeyArr {

		if key.KeyMetaData.KeyId == keyMetaData.KeyId {
			found = true
			continue
		}
		KeyArr = append(KeyArr[:i], KeyArr[i+1:]...)
	}
	if !found {
		return found, errors.New("Key not found")
	}
	return found, nil
}
func GetKey(keyId string) (Key, error) {
	for _, key := range KeyArr {
		if key.KeyMetaData.KeyId == keyId {
			return key, nil
		}
	}
	return Key{}, errors.New("Key not found")
}

func GenerateAesKey() []byte {
	key := make([]byte, 32)
	io.ReadFull(rand.Reader, key)
	return key
}
