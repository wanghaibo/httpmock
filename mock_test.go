package main

import (
	"encoding/json"
	"testing"
)

func TestMockJson(t *testing.T) {
	var mock Mock
	jsonData := []byte(`{"url":"http://www.baidu.com","body":"test","headers":{"a":"b"}}`)
	err := json.Unmarshal(jsonData, &mock)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(mock.Url)
	}

	mockObj := Mock{Url: "http://www.baidu.com", Body: "test", Headers: map[string]string{"a": "b"}}
	mockStr, jsonerr := json.Marshal(mockObj)
	if jsonerr != nil {
		t.Error(jsonerr)
	} else {
		t.Log(string(mockStr))
	}
}
