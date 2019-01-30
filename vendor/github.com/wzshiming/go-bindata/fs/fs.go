package fs

import (
	"bytes"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

// FileInfo implements os.FileInfo interface
type FileInfo struct {
	name      string
	length    int64
	isDir     bool
	timestamp time.Time
}

func (f *FileInfo) Name() string {
	return f.name
}

func (f *FileInfo) Mode() os.FileMode {
	if f.isDir {
		return 0644 | os.ModeDir
	}
	return 0644
}

func (f *FileInfo) ModTime() time.Time {
	return f.timestamp
}

func (f *FileInfo) Size() int64 {
	return f.length
}

func (f *FileInfo) IsDir() bool {
	return f.isDir
}

func (f *FileInfo) Sys() interface{} {
	return nil
}

// AssetFile implements http.File interface
type AssetFile struct {
	*bytes.Reader
	fileInfo *FileInfo
}

func NewAssetFile(name string, content []byte, isDir bool, timestamp time.Time) *AssetFile {
	_, name = filepath.Split(name)
	return &AssetFile{
		Reader: bytes.NewReader(content),
		fileInfo: &FileInfo{
			name:      name,
			length:    int64(len(content)),
			isDir:     isDir,
			timestamp: timestamp,
		},
	}
}

func (f *AssetFile) Close() error {
	return nil
}

func (f *AssetFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *AssetFile) Size() int64 {
	return f.fileInfo.Size()
}

func (f *AssetFile) Stat() (os.FileInfo, error) {
	return f.fileInfo, nil
}

// AssetDirectory implements http.File
type AssetDirectory struct {
	*bytes.Reader
	seek     int
	children []os.FileInfo
	fileInfo *FileInfo
}

func NewAssetDirectory(fs *AssetFS, name string, children []string) *AssetDirectory {
	fileinfos := make([]os.FileInfo, 0, len(children))
	for _, child := range children {
		_, err := fs.AssetDir(filepath.Join(name, child))
		fileinfos = append(fileinfos, &FileInfo{
			name:  child,
			isDir: err == nil,
		})
	}
	_, name = filepath.Split(name)
	return &AssetDirectory{
		children: fileinfos,
		fileInfo: &FileInfo{
			name:  name,
			isDir: true,
		},
	}
}

func (f *AssetDirectory) Close() error {
	return nil
}

func (f *AssetDirectory) Readdir(count int) ([]os.FileInfo, error) {
	if count <= 0 {
		return f.children, nil
	}
	if f.seek+count > len(f.children) {
		count = len(f.children) - f.seek
	}
	buf := f.children[f.seek : f.seek+count]
	f.seek += count
	return buf, nil
}

func (f *AssetDirectory) Size() int64 {
	return 0
}

func (f *AssetDirectory) Stat() (os.FileInfo, error) {
	return f.fileInfo, nil
}

// AssetFS implements http.FileSystem
type AssetFS struct {
	Asset     func(path string) ([]byte, error)
	AssetDir  func(path string) ([]string, error)
	Index     string
	Prefix    string
	Timestamp time.Time
}

func (fs *AssetFS) Open(name string) (http.File, error) {
	name = path.Join(fs.Prefix, name)
	isDir := false
	if name[len(name)-1] == '/' {
		if fs.Index == "" {
			if fs.AssetDir == nil {
				return nil, os.ErrPermission
			}
			children, err := fs.AssetDir(name)
			if err != nil {
				return nil, os.ErrNotExist
			}
			name = strings.TrimPrefix(name, "/")
			return NewAssetDirectory(fs, name, children), nil
		}
		isDir = true
		name += fs.Index
	}
	name = strings.TrimPrefix(name, "/")
	b, err := fs.Asset(name)
	if err != nil {
		return nil, os.ErrNotExist
	}
	return NewAssetFile(name, b, isDir, fs.Timestamp), nil
}
