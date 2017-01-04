package main

import "testing"

func TestGzip(t *testing.T) {
	var srcPath = "aa.tar"
	var dstPath = "aa.tar.gz"
	if err := Gzip(srcPath, dstPath); err != nil {
		t.Error(err)
	} else {
		t.Log("gzip suc")
	}
}

func TestGunzip(t *testing.T) {
	var srcPath = "aa.tar.gz"
	var dstPath = "ab.tar"
	if err := Gunzip(srcPath, dstPath); err != nil {
		t.Error(err)
	} else {
		t.Log("gunzip suc")
	}
}
