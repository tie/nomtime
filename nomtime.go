package nomtime

import (
	"net/http"
	"os"
	"time"
)

type fileSystem struct {
	http.FileSystem
}

func Nomtime(fs http.FileSystem) http.FileSystem {
	return fileSystem{fs}
}

func (fs fileSystem) Open(name string) (http.File, error) {
	f, err := fs.FileSystem.Open(name)
	if err != nil {
		return f, err
	}
	return file{f}, nil
}

type file struct {
	http.File
}

func (f file) Stat() (os.FileInfo, error) {
	fi, err := f.File.Stat()
	if err != nil {
		return fi, err
	}
	return fileInfo{fi}, nil
}

type fileInfo struct {
	os.FileInfo
}

func (fileInfo) ModTime() time.Time {
	return time.Time{}
}
