package targz

import (
	"compress/gzip"
	"io"
	"os"
)

func Gzip(srcPath, dstPath string) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	gw := gzip.NewWriter(dstFile)
	defer gw.Close()

	if _, err = io.Copy(gw, srcFile); err != nil {
		return err
	}
	return nil
}

func Gunzip(srcPath, dstPath string) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		return err
	}
	defer gr.Close()

	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	if _, err = io.Copy(dstFile, gr); err != nil {
		return err
	}
	return nil
}

