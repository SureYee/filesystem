package filesystem

import (
	"io/fs"
)

type FileSys struct {
	adapter Adapter
}

type DirEntry interface {
	IsDir() bool
	Stat()
}

type File interface {
	Read([]byte) (int, error)
	Write([]byte) (int, error)
	Close() error
}

type Adapter interface {
	OpenFile(name string, conf Config) (File, error)
	WriteFile(name string, content []byte, conf Config) error
	ReadFile(name string) ([]byte, error)
	Remove(name string) error
	RemoveAll(name string) error
	List(name string) ([]fs.DirEntry, error)
	FileExists(name string) bool
	DirectoryExists(name string) bool
	Mkdir(name string, conf Config) error
}

func NewFileSystem(adapter Adapter) *FileSys {
	return &FileSys{
		adapter: adapter,
	}
}

// WriteFile
func (fs *FileSys) WriteFile(name string, data []byte, config Config) error {
	return fs.adapter.WriteFile(name, data, config)
}

func (fs *FileSys) ReadFile(name string) ([]byte, error) {
	return fs.adapter.ReadFile(name)
}

func (fs *FileSys) Mkdir(name string, config Config) error {
	return fs.adapter.Mkdir(name, config)
}

func (fs *FileSys) FileExists(name string) bool {
	return fs.adapter.FileExists(name)
}

func (fs *FileSys) DirectoryExists(name string) bool {
	return fs.adapter.DirectoryExists(name)
}

func (fs *FileSys) List(name string) ([]fs.DirEntry, error) {
	return fs.adapter.List(name)
}

func (fs *FileSys) Remove(name string) error {
	return fs.adapter.Remove(name)
}
