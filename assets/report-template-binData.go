// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package main generated by go-bindata.// sources:
// report-template.html
package assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _reportTemplateHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\x4b\x6f\xe3\x36\x10\xbe\xfb\x57\x4c\xd5\x2d\xe0\x20\x2b\x39\x49\x77\x81\xad\x2c\xf9\xb0\xd9\xa4\x39\xa4\xdd\xa0\xeb\x1e\x8a\xa2\x07\x5a\x1c\xcb\x8c\x29\x52\x20\x47\x7e\xd4\xf0\x7f\x2f\x28\xc9\x2f\x49\x79\x34\x4b\x20\x0a\x48\xce\x7c\xdf\xf0\x9b\xe1\xd0\xd1\x0f\x5f\xbe\x5e\x8f\xff\x7a\xb8\x81\x19\x65\x72\xd4\x8b\xdc\x3f\x90\x4c\xa5\xb1\x87\xca\x73\x0b\xc8\xf8\xa8\x07\x00\x10\x65\x48\x0c\x92\x19\x33\x16\x29\xf6\xfe\x1c\xdf\xfa\x9f\xbc\x7a\x8b\x04\x49\x1c\x8d\xdd\x37\x1a\x54\x93\x6a\xc3\xd2\x5a\x22\xd0\x3a\xc7\xd8\x23\x5c\xd1\x20\xb1\xb6\x76\x72\x23\x30\x5a\x13\x6c\xf6\x73\x37\x26\x2c\x99\xa7\x46\x17\x8a\xfb\x89\x96\xda\x84\xf0\xe3\xd5\xcf\x57\x17\x1f\x2e\x87\x27\x66\xf5\xde\x72\x26\x08\x0f\x3b\xdb\xde\x01\xdb\x16\x49\x82\xd6\x7e\xde\xe3\x5d\x3b\x97\x17\xd9\x38\x33\xf3\xd4\x20\xaa\x6e\xd4\x29\x13\xf2\x2d\x90\x06\xf9\x13\x61\xce\x45\xfe\xc6\x18\xd7\xdd\x88\x39\x4b\xe6\x2c\xc5\x6b\x66\xf8\x3d\x5b\xeb\xa2\xa9\x70\x6a\x04\xf7\x09\xb3\x5c\x32\x42\x07\x59\x64\xca\x86\x70\x39\x35\xc0\x0a\xd2\xe5\x67\xd8\xf6\xa8\x0c\xfd\x94\xe5\x21\x7c\xca\x57\xa7\x16\x5c\xd8\x5c\xb2\x75\x58\x9a\x76\x87\x45\x68\xe9\xad\x31\xbd\x8a\xcc\x8d\xa5\xe0\x34\x0b\xe1\x97\x8f\x3f\x9d\xae\x4f\xb4\xe1\x68\x7c\xc3\xb8\x28\x6c\x08\x1f\x9a\xf1\x67\xcc\xa4\x42\xf9\x13\x4d\xa4\xb3\x10\x3e\x36\xf7\x73\xc6\xb9\x50\x69\xc3\xf3\xf8\x78\x89\x96\x92\xe5\x56\x4c\x24\x36\xce\x96\x14\xc6\xba\xac\xe5\x5a\x28\x42\xf3\xa2\xfb\x1d\x32\xc7\xd5\x44\xe9\x2c\xf8\x93\xd8\x5a\x59\x79\x5e\x8d\x10\x94\x56\x0d\x2c\x77\x49\x7d\x26\x45\xaa\x42\x90\x38\xa5\xd3\x5d\x5d\x90\x14\x0a\xbb\x1c\xa7\x5a\x91\x6f\xc5\xbf\x18\xc2\x65\x4b\xbd\xef\x51\xff\x79\x91\xae\xb5\x22\x54\xcd\x6a\xda\x2b\x72\x01\x97\x2d\x51\x32\xb6\xf2\x67\x28\xd2\x19\x85\x70\xd1\x38\xdf\x02\xcd\x54\xea\x65\x08\x33\xc1\xf9\x71\x0b\x28\xb5\x31\x4c\x59\x41\x42\xab\xf0\x08\x04\x2e\x82\x2b\x0b\xc8\x2c\xfa\xba\xa0\xa7\x4b\xff\x1b\x31\xb2\x5f\x17\x68\x16\x02\x97\xaf\xaf\xfe\xfa\xef\x7f\xdf\xb6\x9c\x59\x8b\x7c\x8c\x96\x6c\x83\xec\x28\x53\x2b\x5f\x32\x93\x62\x67\x67\x7d\xa1\x07\x7e\x17\x76\xa3\x19\xba\x6f\x34\x28\xdf\x8a\x51\x2f\x1a\x54\x6f\x4e\x34\xd1\x7c\x0d\x89\x64\xd6\xc6\x9e\x7b\x27\xdc\x73\xc4\xc5\x02\x4a\xbb\xd8\x3b\xa2\x2a\x89\xbc\xd1\xaf\x1a\x5c\x4c\xf0\x07\xe6\xda\x50\x34\xe0\x62\x31\xea\x45\xf9\xd3\x0e\xa5\xf5\x17\x46\x18\xc2\x66\x13\xb8\x99\x9b\x6c\xb7\xd1\x20\xaf\xb9\x6a\xfa\x56\xfe\x76\x0f\xdf\x1e\xbc\xae\x60\xd2\xb9\x2b\x2a\x6f\xe7\x78\x94\x05\x6f\xf4\x50\x4e\xc0\x81\xd9\x92\xf1\xe1\xb0\x5b\x93\xbe\x0a\xf4\x48\x7e\x6f\x74\x5b\x4e\x8e\x40\x6f\x0f\xbb\x9d\xa0\x1d\x29\x82\x06\xd1\x68\xac\x89\xc9\x12\x13\x48\x64\xb5\x3c\x6e\xcd\xc1\x8e\x45\xb6\xd7\xa8\xd2\x78\xb3\x31\x4c\xa5\x08\xef\x84\xe2\xb8\x7a\x0f\xef\x50\x62\xe6\x6e\x65\x18\x43\x70\x37\xfe\xed\xfe\xa6\x9a\xdb\x6d\x95\xeb\xcd\x66\x67\xb1\xdd\xf6\x36\x1b\x54\x7c\xbb\xed\x45\x03\x97\xf0\x51\xaf\x17\xd9\xc4\x88\x9c\xaa\xb8\x07\x03\x78\xb4\x50\xad\x00\x69\x48\x0c\x32\x42\x60\x0a\xea\x5e\xc0\x26\x12\x4b\xcb\x05\x33\xe5\x1a\xc4\xc0\x75\x52\x38\xf4\x20\x45\xda\x51\x7f\x5e\x5f\x3b\xf5\x7e\x67\x19\xf6\xbd\xa3\x3e\xe2\x9d\x55\x95\x38\xd5\x06\xfa\x12\x09\x04\xc4\x70\x31\x04\x01\x51\x09\x17\x48\x54\x29\xcd\x86\x20\xce\xcf\xcf\x8e\xaa\x7d\x47\xd7\xe8\xd9\x31\x14\x8a\xe3\x54\x28\xe4\x87\x0b\xb3\xc7\x7e\xac\xb0\x1f\x6b\xec\xbf\xc5\x3f\x41\x32\x13\x92\x1b\x54\x7b\x9e\xc7\x53\x1e\x37\x9c\xeb\x4e\xd2\xb8\xed\x29\x08\xb3\xfe\xe3\xd9\x89\x8b\x98\x42\xbf\x76\x09\x92\xdd\xc1\x03\xa1\x12\x59\x70\xb4\x27\x0a\xd4\xa1\x7b\x67\x4d\xda\xfa\xc2\xb6\x8f\x58\x03\xb7\x8c\x27\x06\xd9\xfc\x64\x75\xdb\xd5\x3f\xda\x98\x01\xe3\xfc\x66\x81\x8a\xee\x85\x25\x54\x68\xfa\x5e\x22\x45\x32\xf7\xde\xc3\xb4\x50\x89\x6b\xb9\xd0\x6f\x86\x57\x65\xa0\x6a\xff\x31\xd0\x4c\xd8\x40\xe1\x6a\x97\xf0\x6f\x62\x22\x85\x4a\x87\x2d\x55\x6a\x97\xa0\xbc\x0e\x41\xc6\x56\x77\x65\x1b\xef\x3e\x7c\xa7\x29\xc4\xa0\x0a\x29\x4f\xa1\xb7\x80\xd2\x36\x5f\xfe\xe7\x41\x96\x42\x71\xbd\x0c\x84\x52\x68\xea\xc5\x73\xf0\xf2\x95\x37\x7c\x4a\xc3\xba\x56\xdd\x65\xd9\x5d\x92\x68\x50\xfe\x74\xff\x2f\x00\x00\xff\xff\x6d\x48\xe3\xcd\xca\x0b\x00\x00")

func reportTemplateHtmlBytes() ([]byte, error) {
	return bindataRead(
		_reportTemplateHtml,
		"report-template.html",
	)
}

func reportTemplateHtml() (*asset, error) {
	bytes, err := reportTemplateHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "report-template.html", size: 3018, mode: os.FileMode(420), modTime: time.Unix(1609034943, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"report-template.html": reportTemplateHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("nonexistent") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"report-template.html": &bintree{reportTemplateHtml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
