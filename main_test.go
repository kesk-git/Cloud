package main

import (
	"bytes"
	"errors"
	"testing"
)

// failWriter rejects every write with a fixed error.
type failWriter struct{ err error }

func (w failWriter) Write([]byte) (int, error) { return 0, w.err }

func TestRunReturnsWriterError(t *testing.T) {
	sentinel := errors.New("sentinel write failure")
	for _, args := range [][]string{{"--smoke"}, {"--smoke", "--name", "Scout"}} {
		err := run(failWriter{err: sentinel}, args)
		if !errors.Is(err, sentinel) {
			t.Fatalf("%v: got %v, want writer error %v", args, err, sentinel)
		}
	}
}

func TestSmoke(t *testing.T) {
	var out bytes.Buffer
	if err := run(&out, []string{"--smoke"}); err != nil {
		t.Fatal(err)
	}
	if got := out.String(); got != "cloud-stub: ok\n" {
		t.Fatalf("unexpected smoke output: %q", got)
	}
}

func TestSmokeWithName(t *testing.T) {
	cases := []struct {
		name string
		want string
	}{
		{"mac", "cloud-stub: ok (mac)\n"},
		{"fleet worker", "cloud-stub: ok (fleet worker)\n"},
		{"  padded  ", "cloud-stub: ok (padded)\n"},
		{"\ttabbed\n", "cloud-stub: ok (tabbed)\n"},
	}
	for _, tc := range cases {
		var out bytes.Buffer
		if err := run(&out, []string{"--smoke", "--name", tc.name}); err != nil {
			t.Fatalf("--name %q: %v", tc.name, err)
		}
		if got := out.String(); got != tc.want {
			t.Fatalf("--name %q: got %q, want %q", tc.name, got, tc.want)
		}
	}
}

func TestSmokeRejectsEmptyName(t *testing.T) {
	for _, name := range []string{"", " ", "   ", "\t", " \n\t "} {
		var out bytes.Buffer
		err := run(&out, []string{"--smoke", "--name", name})
		if err == nil {
			t.Fatalf("expected error for --name %q", name)
		}
		if out.Len() != 0 {
			t.Fatalf("expected no output for --name %q, got %q", name, out.String())
		}
	}
}

func TestInvalidArguments(t *testing.T) {
	cases := [][]string{
		nil,
		{"--unknown"},
		{"--smoke", "extra"},
		{"--smoke", "--name"},
		{"--smoke", "--name", "a", "b"},
		{"--smoke", "--other", "a"},
		{"--name", "a", "--smoke"},
		{"--name", "a"},
	}
	for _, args := range cases {
		var out bytes.Buffer
		if err := run(&out, args); err == nil {
			t.Fatalf("expected error for %v", args)
		}
	}
}
