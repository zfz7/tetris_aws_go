package main

import (
	"testing"
)

func Test_sayHello(t *testing.T) {
	got := sayHello("hi")
	want := SayHelloResponse{Message: "hi"}
	if got != want {
		t.Errorf("sayHello(hi) = %v; want %v", got, want)
	}
}

func Test_handelInfo(t *testing.T) {
	t.Setenv("REGION", "1")
	t.Setenv("USER_POOL_ID", "2")
	t.Setenv("USER_POOL_WEB_CLIENT_ID", "3")
	got := handelInfo()
	want := InfoResponse{
		Region:                 "1",
		UserPoolId:             "2",
		UserPoolWebClientId:    "3",
		AuthenticationFlowType: "USER_PASSWORD_AUTH",
	}
	if got != want {
		t.Errorf("handelInfo() = %v; want %v", got, want)
	}
}
