package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/google/uuid"
	"golang.org/x/crypto/pbkdf2"

	"time"
)

var KeyArr []Key = make([]Key, 0)

func CreateKey() (KeyMetaData, error) {

	keyID := uuid.New().String()
	iv := make([]byte, 16)
	rand.Read(iv)
	keyMetaData := KeyMetaData{
		TimeStamp: time.Now().UnixNano(),
		KeyId:     keyID,
		IsEnabled: true,
	}
	aesKey := GenerateAesKey()
	key := Key{
		KeyMetaData: keyMetaData,
		Key:         aesKey,
		IV:          iv,
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
	path := CONFIG["KMS_STORE"] + "/" + key.KeyMetaData.KeyId + ".key"
	var mkey = Mk.MasterKey
	var plainText = []byte(keyJson)
	var cipherText, er = EncryptAESGCM(mkey, plainText)
	if er != nil {
		return false, er
	}
	err = os.WriteFile(path, cipherText, 0644)
	if err != nil {
		return false, err
	}
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
func (k Key) GetKey(keyId string) (Key, error) {
	path := CONFIG["KMS_STORE"] + "/" + keyId + ".key"
	fmt.Println(path)
	f, er := os.ReadFile(path)
	if er != nil {
		return Key{}, er
	}
	var mkey = Mk.MasterKey
	plainText, err := DecryptAESGCM(f, mkey)
	fmt.Println("err")
	fmt.Println(err)
	if err != nil {
		return Key{}, err
	}
	js := json.Unmarshal(plainText, &k)
	fmt.Println(js)
	return k, nil
}

func GenerateAesKey() []byte {
	key := make([]byte, 32)
	io.ReadFull(rand.Reader, key)
	return key
}

func (mkp MasterKey) GetMasterKey() []byte {
	passPhrase := "19@obgtrtdznimtblmaxlhexwxoxehixkhymablikhcxvmtgwtgrhgxxelxvetbfbgzlhblerbgz"
	if len(passPhrase) < 10 {
		Exit(fmt.Sprintf("The pass phrase must be at least 10 characters long is only %v characters", len(passPhrase)), 2)
	}
	return pbkdf2.Key([]byte(passPhrase), []byte{}, 4096, 32, sha256.New)
}
func Exit(messages string, errorCode int) {
	// Exit code and messages based on Nagios plugin return codes (https://nagios-plugins.org/doc/guidelines.html#AEN78)
	var prefix = map[int]string{0: "OK", 1: "Warning", 2: "Critical", 3: "Unknown"}

	// Catch all unknown errorCode and convert them to Unknown
	if errorCode < 0 || errorCode > 3 {
		errorCode = 3
	}

	log.Printf("%s %s\n", prefix[errorCode], messages)
	os.Exit(errorCode)
}
