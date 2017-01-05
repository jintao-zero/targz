package targz

import (
	"archive/tar"
	"io"
	"log"
	"os"
	"path/filepath"
)

func Tar(srcPath, dstPath string) error {
	f, err := os.Create(dstPath)
	if err != nil {
		return err
	}

	tw := tar.NewWriter(f)
	if err = tarPath(srcPath, tw); err != nil {
		return err
	}

	// Make sure to check the error on Close.
	if err = tw.Close(); err != nil {
		return err
	}

	return nil
}

func tarPath(path string, tw *tar.Writer) error {

	// if file is file
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}
	if fileInfo.Mode().IsRegular() {
		header, err := tar.FileInfoHeader(fileInfo, "")
		if err != nil {
			return err
		}
		header.Name = path
		if err = tw.WriteHeader(header); err != nil {
			return err
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		if _, err = io.Copy(tw, file); err != nil {
			return err
		}
	}

	if fileInfo.Mode().IsDir() {
		// tar each file and dir in the dir
		var file *os.File
		if file, err = os.Open(path); err != nil {
			return err
		}
		fileInfos, err := file.Readdir(0)
		if err != nil {
			return err
		}
		for _, info := range fileInfos {
			if err = tarPath(filepath.Join(path, info.Name()), tw); err != nil {
				return err
			}
		}
	}
	return nil
}

func UnTar(srcPath, dstPath string) error {
	tarFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	tr := tar.NewReader(tarFile)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		fullPath := filepath.Join(dstPath, hdr.Name)
		os.MkdirAll(filepath.Dir(fullPath), os.ModePerm)
		log.Println("fullPath", fullPath)
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

