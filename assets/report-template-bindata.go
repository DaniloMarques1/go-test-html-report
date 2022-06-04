// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

// Mode return file modify time
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

var _reportTemplateHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x57\x4b\x93\xdb\x36\x0c\xbe\xef\xaf\x40\xd5\x74\xe2\x9d\x44\xf2\xa3\x69\x27\x23\x4b\x3e\x64\x93\x34\x87\xb4\xc9\x34\xee\xa1\xd3\xf4\x40\x8b\xb0\xcc\x35\x45\x6a\x48\xc8\x8f\xee\xf8\xbf\x77\xa8\x87\x2d\xcb\xda\x64\x93\xd5\x81\x3b\x24\x81\x0f\xc0\x07\x02\x58\x47\x3f\xbc\xfe\x70\x33\xff\xfb\xe3\x1b\x58\x51\x26\x67\x57\x91\xfb\x03\x92\xa9\x34\xf6\x50\x79\xee\x00\x19\x9f\x5d\x01\x00\x44\x19\x12\x83\x64\xc5\x8c\x45\x8a\xbd\xbf\xe6\x6f\xfd\x97\x5e\x7d\x45\x82\x24\xce\xe6\x6e\x8d\x86\xd5\xa6\xba\xb0\xb4\x97\x08\xb4\xcf\x31\xf6\x08\x77\x34\x4c\xac\xad\x95\xdc\x17\x18\xad\x09\xee\x8e\x7b\xf7\x2d\x58\xb2\x4e\x8d\x2e\x14\xf7\x13\x2d\xb5\x09\xe1\xc7\xc9\xcf\x93\xd1\x8b\xf1\xf4\x4c\xac\xbe\xdb\xae\x04\xe1\xe9\xe6\x70\x75\xc2\xb6\x45\x92\xa0\xb5\xaf\x8e\x78\x37\x4e\xe5\xab\xd6\x38\x33\xeb\xd4\x20\xaa\x7e\xd4\x25\x13\xf2\x7b\x20\x0d\xf2\x7b\xdc\x5c\x8b\xfc\x3b\x7d\xdc\xf7\x23\xe6\x2c\x59\xb3\x14\x6f\x98\xe1\xef\xd9\x5e\x17\x5d\x86\x53\x23\xb8\x4f\x98\xe5\x92\x11\x3a\xc8\x22\x53\x36\x84\xf1\xd2\x00\x2b\x48\x9f\x96\xe9\xa5\x5a\x25\xed\xa7\x2c\x0f\xe1\x65\xbe\x3b\x97\xe0\xc2\xe6\x92\xed\xc3\x52\xb4\xdf\x37\x42\x4b\x8f\x72\xec\x41\x16\xdd\xb7\x15\x9c\x56\x21\x8c\x47\xa3\x9f\x3a\x71\xf4\xfa\xbe\xd0\x86\xa3\xf1\x0d\xe3\xa2\xb0\x21\xbc\xe8\xde\x67\xcc\xa4\x42\xf9\x0b\x4d\xa4\xb3\x10\x7e\xe9\xde\xe7\x8c\x73\xa1\xd2\x8e\x66\x3b\xf4\x44\x4b\xc9\x72\x2b\x16\x12\x3b\x71\x27\x85\xb1\x2e\xad\xb9\x16\x8a\xd0\x7c\x55\xfd\x1d\x32\x67\xab\x8b\xd2\x5b\x11\x67\xbe\x5d\x44\x7d\x2f\x49\x15\x1d\x21\x28\xad\x3a\x60\xae\x8c\x7d\x26\x45\xaa\x42\x90\xb8\xa4\xf3\x5b\x5d\x90\x14\x0a\xfb\x14\x97\x5a\x91\x6f\xc5\x7f\x18\xc2\xf8\x82\xbe\xc7\xd0\xff\x65\x96\x42\xb6\x24\xec\x96\x54\xa2\x15\xa1\xa2\x10\x9e\x7e\x1e\x8d\x26\xaf\x9e\xf6\x83\xb1\x84\xc4\x06\xbf\x0c\xe0\x7d\x9e\x4c\xc6\x13\xef\xa1\xde\xdc\x54\x7a\x1d\xb4\x63\x82\x46\x30\xbe\xc8\x51\xc6\x76\xfe\x0a\x45\xba\xa2\x10\x46\x1d\xb6\x37\x68\x96\x52\x6f\x43\x58\x09\xce\xdb\x2d\xab\xcc\x94\x61\xca\x0a\x12\x5a\x85\x2d\x10\x18\x05\x13\x0b\xc8\x2c\xfa\xba\xa0\xfb\xab\xf4\x13\x31\xb2\x1f\x36\x68\x36\x02\xb7\x0f\x2f\xd4\xa6\x58\xbf\xb9\x33\xe4\xcc\x5a\xe4\x73\xb4\x64\x3b\xd6\x5a\x0f\x67\xe7\x4b\x66\x52\xec\x1d\x05\x5f\x69\xda\x8f\xc2\xee\x74\x6f\xb7\x46\xc3\x72\xb8\xcd\xae\xa2\x61\x35\x24\xa3\x85\xe6\x7b\x48\x24\xb3\x36\xf6\xdc\x60\x73\xf3\x93\x8b\x0d\x94\x72\xb1\x77\xe4\x60\x29\x71\x37\x2d\x57\x9f\x0b\x83\x49\x95\xa1\x8a\xc5\x69\xfd\xd2\x43\x18\xff\x9a\xef\xa6\xd0\x24\x7e\x3c\x1a\x6d\x56\xcd\xb8\x6d\x81\xb6\xfc\x2f\xbd\xf7\x66\xbf\x69\x70\x81\xc2\x9f\x98\x6b\x43\xd1\x90\x8b\xcd\x43\xd4\x4a\x9d\xd7\x8c\x30\x84\xbb\xbb\xc0\xed\xdc\xe6\x70\xe8\x02\xd4\xf1\x5d\xbc\x90\xd6\x54\x8f\xf2\xc6\x4c\x5d\xb5\xa4\x73\xf7\x74\xbd\x46\xb9\x95\x6a\x6f\xf6\xb1\xdc\x80\x03\xb4\xa5\xed\x8f\xa7\x5b\x67\x3e\xff\x06\xe0\x56\x9e\xbd\xd9\xdb\x72\xd3\x02\x7e\x7b\xba\xbd\x17\xb8\xe7\x3d\x40\xc7\xd8\x6c\xae\x89\xc9\x12\x17\x48\x64\x35\x61\xee\xcc\x41\xcf\x45\x86\x2d\xf4\x16\x7d\x77\x77\x86\xa9\x14\xe1\x89\x50\x1c\x77\xcf\xe1\x09\x4a\xcc\x5c\x2f\x08\x63\x08\xde\xcd\x7f\x7f\xff\xa6\xda\xdb\xc3\xa1\x96\x6f\x24\x8e\x07\xa8\xf8\xe1\x70\x55\x63\x46\x43\xf7\xe0\x66\x57\x91\x4d\x8c\xc8\xa9\x32\x32\x1c\xc2\xad\x85\xea\x04\x48\x43\x62\x90\x11\x02\x53\x50\xf7\x22\xb6\x90\x58\x4a\x6e\x98\x29\xcf\x20\x06\xae\x93\xc2\xd9\x09\x52\xa4\xc6\x89\x57\xfb\x1b\xc7\xe9\x1f\x2c\xc3\x81\xd7\xea\x63\xde\x75\x55\x08\x4b\x6d\x60\x20\x91\x40\x40\x0c\xa3\x29\x08\x88\x4a\xb8\x40\xa2\x4a\x69\x35\x05\xf1\xec\xd9\x75\xab\xd8\x1a\x73\x9d\x11\x16\x43\xa1\x38\x2e\x85\x42\x7e\xaa\xd7\x23\xf6\x6d\x85\x7d\x5b\x63\xff\x23\xfe\x0d\x92\x95\x90\xdc\xa0\x3a\xda\xb9\x3d\xb7\xe3\x3e\xa7\xda\x90\x1b\x5f\x6a\x0a\xc2\x6c\x70\x7b\x7d\xa6\x22\x96\x30\xa8\x55\x82\xa4\x09\x3c\x10\x2a\x91\x05\x47\x7b\xc6\x40\xed\xba\x77\xdd\x35\x5b\xf7\x8b\xcb\x10\x6b\xe0\x0b\xe1\x85\x41\xb6\x3e\x3b\x3d\xf4\xb5\xaf\x4b\xcc\x80\x71\xfe\x66\x83\x8a\xde\x0b\x4b\xa8\xd0\x0c\xbc\x44\x8a\x64\xed\x3d\x87\x65\xa1\xca\x86\x02\x83\xae\x7b\xb4\x12\xb6\x8a\xcd\x69\x05\xa4\xd3\x54\xe2\xc0\xab\x86\x9c\x77\x4e\x47\x95\xad\x6a\x54\xc5\x95\xa6\xc2\x5d\xf3\x38\x3e\x89\x85\x14\x2a\x9d\x5e\x30\x58\xab\x04\x65\x31\x05\x19\xdb\xbd\x2b\xdb\x57\x3f\x51\xbd\xa2\x10\x83\x2a\xa4\x3c\x87\x3e\x00\x4a\xdb\xfd\xa7\xe9\xcb\x20\x5b\xa1\xb8\xde\x06\x42\x29\x34\xf5\xe1\x33\xf0\xf2\x9d\x37\xbd\x8f\xef\xfa\x5d\xbb\xf2\x6a\x0a\x2a\x1a\x96\x3f\x8b\xfe\x0f\x00\x00\xff\xff\x71\x13\x5a\x22\x26\x0d\x00\x00")

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

	info := bindataFileInfo{name: "report-template.html", size: 3366, mode: os.FileMode(420), modTime: time.Unix(1653785133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
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
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
