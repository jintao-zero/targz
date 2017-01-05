package targz

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
	"path/filepath"
)

func Targz(srcPath, dstPath string) error {
	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	gw := gzip.NewWriter(dstFile)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()
	if err = tarPath(srcPath, tw); err != nil {
		dstFile.Close()
		os.Remove(dstPath)
		return err
	}
	return nil
}

func UnTargz(srcPath, dstPath string) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		return err
	}
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		fullPath := filepath.Join(dstPath, hdr.Name)
		os.MkdirAll(filepath.Dir(fullPath), os.ModePerm)
		file, err := os.Create(fullPath)
		if err != nil {
			return err
		}
		if _, err := io.Copy(file, tr); err != nil {
			return err
		}
		file.Close()
	}
	return nil
}
