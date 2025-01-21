package tests

import (
	"bytes"
	"encoding/json"
	"infotecs_go/src/settings"
	"infotecs_go/src/wallet"
	"testing"
)

func TestBaseStructBuild(t *testing.T) {
	var schemas wallet.BaseSchemas
	schemas.Detail = "test"
	expected, _ := json.Marshal(schemas)
	result, _ := schemas.BuildJson("test")
	if bytes.Compare(expected, result) == -1 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestErrorStructBuild(t *testing.T) {
	var schemas settings.ErrorSchemas
	schemas.Detail = "testError"
	expected, _ := json.Marshal(schemas)
	result, _ := schemas.BuildJson("testError")
	if bytes.Compare(expected, result) == -1 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, expected)
	}
}
