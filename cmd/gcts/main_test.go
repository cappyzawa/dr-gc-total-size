package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestBlobFind(t *testing.T) {
	cases := []struct {
		name     string
		sha256   []byte
		expect   int64
		existErr bool
	}{
		{
			name:     "with exist file",
			sha256:   []byte("1111"),
			expect:   456,
			existErr: false,
		},
		{
			name:     "with non exist file",
			sha256:   []byte("invalid file"),
			expect:   0,
			existErr: true,
		},
	}

	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			blob := &Blob{
				Dir:    "../../testdata/docker/registry/v2/blobs/sha256",
				SHA256: c.sha256,
			}
			actual, err := blob.Find()
			if !c.existErr && err != nil {
				t.Errorf("error should not be occurred, but actual is %v", err)
			}
			if c.existErr && err == nil {
				t.Error("error should be occurred, but it does not be occurred")
			}
			if actual != c.expect {
				t.Errorf("size should be %d, but actual is %d", c.expect, actual)
			}
		})
	}
}

func TestCLIRun(t *testing.T) {
	cases := []struct {
		name      string
		inStream  func(reader io.Reader)
		expect    int
		expectOut string
	}{
		{
			name: "stdin contains deletion blob",
			inStream: func(reader io.Reader) {
				reader.(*bytes.Buffer).Write([]byte("blob eligible for deletion: sha256:1111"))
			},
			expect:    0,
			expectOut: "456 bytes will be reduced",
		},
		{
			name: "stdin does not contain deletion blob",
			inStream: func(reader io.Reader) {
				reader.(*bytes.Buffer).Write([]byte("myregistry/same: marking manifest sha256:1c61afa"))
			},
			expect:    0,
			expectOut: "0 bytes will be reduced",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			outS := new(bytes.Buffer)
			errS := new(bytes.Buffer)
			inS := new(bytes.Buffer)
			cli := &CLI{
				Out: outS,
				Err: errS,
				In:  inS,
			}
			c.inStream(inS)
			actual := cli.Run([]string{"test", "-d", "../../testdata"})
			if actual != c.expect {
				t.Errorf("exitCode should be %d, actual is %d", c.expect, actual)
			}
			if !strings.Contains(outS.String(), c.expectOut) {
				t.Errorf("output should contain %s, but actual is %s", c.expectOut, outS.String())
			}
		})
	}
}
