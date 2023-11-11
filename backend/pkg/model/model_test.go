package model

import (
	"encoding/json"
	"testing"
)

func Test_sayHelloResponse(t *testing.T) {
	object := SayHelloResponse{
		Message: "hello",
	}
	got, _ := json.Marshal(object)
	want := "{\"message\":\"hello\"}"
	if string(got) != want {
		t.Errorf("SayHelloResponse(hi) = %v; want %v", string(got), want)
	}
}

func Test_InfoResponse(t *testing.T) {
	object := InfoResponse{
		AuthenticationFlowType: "1",
		Region:                 "2",
		UserPoolId:             "3",
		UserPoolWebClientId:    "4",
	}
	got, _ := json.Marshal(object)
	want := "{\"authenticationFlowType\":\"1\",\"region\":\"2\",\"userPoolId\":\"3\",\"userPoolWebClientId\":\"4\"}"
	if string(got) != want {
		t.Errorf("InfoResponse = %v; want %v", string(got), want)
	}
}

func Test_ApiError(t *testing.T) {
	object := ApiError{
		ErrorMessage: "hello",
	}
	got, _ := json.Marshal(object)
	want := "{\"errorMessage\":\"hello\"}"
	if string(got) != want {
		t.Errorf("InfoResponse = %v; want %v", string(got), want)
	}
}
