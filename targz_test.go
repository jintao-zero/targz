package main

import "testing"

func TestTargz(t *testing.T) {
	var srcPath = "test"
	var dstPath = "test.tar.gz"
	if err := Targz(srcPath, dstPath); err != nil {
		t.Error(err)
	}
}

func TestUnTargz(t *testing.T) {

}
