package tools

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"go_fb_plugin/fb"
	"log"
	"net/http"
)

// Check checks for the error, logs it and sends back the bool value.
func Check(err error, info ...string) bool {
	if err != nil {
		log.Println(info, err.Error())
		return false
	}
	return true
}

// NumToBytes - converts integer to bytes, supports int32 and int64
func NumToBytes(num interface{}) []byte {
	var len int
	switch num.(type) {
	case int64:
		len = 8
		b := make([]byte, len)
		binary.BigEndian.PutUint64(b, uint64(num.(int64)))
		return b
	case int:
		len = 4
		b := make([]byte, len)
		binary.BigEndian.PutUint32(b, uint32(num.(int)))
		return b
	case int32:
		len = 4
		b := make([]byte, len)
		binary.BigEndian.PutUint32(b, uint32(num.(int32)))
		return b
	}
	return []byte{}
}

// BytesToInt - converts byte array to int
func BytesToInt(arr []byte) int {
	return int(binary.BigEndian.Uint32(arr))
}

// BytesToInt64 - converts byte array to int
func BytesToInt64(arr []byte) int64 {
	return int64(binary.BigEndian.Uint64(arr))
}

// SendResponse - sends response back to client
func SendResponse(response interface{}) {
	client := &http.Client{}
	url := fb.FacebookAPI
	var req *http.Request
	var err error
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(&response)
	req, err = http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}
