package main

import (
	"crypto/aes"
	"crypto/cipher"
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
func (k Key) GetKey(keyId string) (Key, error) {
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

func EncryptAESGCM(key []byte, plainText []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}
	ciphertext := gcm.Seal(nil, nonce, plainText, nil)

	return append(nonce, ciphertext...), nil
}

func DecryptAESGCMfunc(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("data to decrypt is too small")
	}

	plaintext, err := gcm.Open(nil, ciphertext[:nonceSize], ciphertext[nonceSize:], nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func (mkp MasterKey) GetMasterKey() []byte {
	passPhrase := "19@obgtrtdznimtblmaxlhexwxoxehixkhymablikhcxvmtgwtgrhgxxelxvetbfbgzlhblerbgz"
	if len(passPhrase) < 10 {
		Exit(fmt.Sprintf("The pass phrase must be at least 10 characters long is only %v characters", len(passPhrase)), 2)
	}
	return pbkdf2.Key(mkp.MasterKey, []byte(passPhrase), 4096, 32, sha256.New)
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
