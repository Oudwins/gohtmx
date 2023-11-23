package filestorage

import (
	"bufio"
	"bytes"
	"io"
	"mime/multipart"
	"os"
	"path"
)

type FileStorage interface {
	SaveFile(f *multipart.File, name string) error
	// GetFile(url string) (f *multipart.File, e error)
}

type LocalFileStorage struct {
	path string
}

var Api FileStorage

func Init() error {

	localPath := "./tmp/files/"
	// _ = os.Mkdir(localPath, os.ModePerm)
	Api = &LocalFileStorage{
		path: localPath,
	}

	return nil
}

// func (s *LocalFileStorage) GetFile(url string) (f *multipart.File, e error) {
// }

func (s *LocalFileStorage) SaveFile(f *multipart.File, name string) error {

	fpath := path.Join(s.path, name)
	dst, err := os.Create(fpath)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, *f); err != nil {
		return err
	}

	return nil
}

func (s *LocalFileStorage) SaveBuf(f *bytes.Buffer, name string) error {
	fpath := path.Join(s.path, name)
	dst, err := os.Create(fpath)
	if err != nil {
		return err
	}
	defer dst.Close()
	writer := bufio.NewWriter(dst)
	for {

	}
}
