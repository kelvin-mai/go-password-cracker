package main

import (
	"testing"

	"github.com/kelvin-mai/go-password-cracker/pass"
)

func TestHash_superman(t *testing.T) {
	actual := pass.CrackSha1Hash("18c28604dd31094a8d69dae60f1bcd347f1afc5a")
	if actual == "superman" {
		t.Logf("success: expected %v, got %v", "superman", actual)
	} else {
		t.Errorf("failed: expected %v, got %v", "superman", actual)
	}
}

func TestHash_q1w2e3r4t5(t *testing.T) {
	actual := pass.CrackSha1Hash("5d70c3d101efd9cc0a69f4df2ddf33b21e641f6a")
	if actual == "q1w2e3r4t5" {
		t.Logf("success: expected %v, got %v", "q1w2e3r4t5", actual)
	} else {
		t.Errorf("failed: expected %v, got %v", "q1w2e3r4t5", actual)
	}
}

func TestHash_bubbles1(t *testing.T) {
	actual := pass.CrackSha1Hash("b80abc2feeb1e37c66477b0824ac046f9e2e84a0")
	if actual == "bubbles1" {
		t.Logf("success: expected %v, got %v", "bubbles1", actual)
	} else {
		t.Errorf("failed: expected %v, got %v", "bubbles1", actual)
	}
}

func TestHash_01071988(t *testing.T) {
	actual := pass.CrackSha1Hash("80540a46a2c1a0eae58d9868f01c32bdcec9a010")
	if actual == "01071988" {
		t.Logf("success: expected %v, got %v", "01071988", actual)
	} else {
		t.Errorf("failed: expected %v, got %v", "01071988", actual)
	}
}

func TestHash_NOT_FOUND(t *testing.T) {
	actual := pass.CrackSha1Hash("03810a46a2c1a0eae58d9332f01c32bdcec9a01a")
	if actual == "PASSWORD NOT IN DATABASE" {
		t.Logf("success: expected %v, got %v", "PASSWORD NOT IN DATABASE", actual)
	} else {
		t.Errorf("failed: expected %v, got %v", "PASSWORD NOT IN DATABASE", actual)
	}
}
