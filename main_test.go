package main

import (
	"bytes"
	"testing"
)

func TestSmoke(t *testing.T) {
	var out bytes.Buffer
	if err := run(&out, []string{"--smoke"}); err != nil {
		t.Fatal(err)
	}
	if got := out.String(); got != "cloud-stub: ok\n" {
		t.Fatalf("unexpected smoke output: %q", got)
	}
}

func TestInvalidArguments(t *testing.T) {
	for _, args := range [][]string{nil, {"--unknown"}, {"--smoke", "extra"}} {
		var out bytes.Buffer
		if err := run(&out, args); err == nil {
			t.Fatalf("expected error for %v", args)
		}
	}
}
