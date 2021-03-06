// Code generated by go-bindata.
// sources:
// compiled/ScriptEngineEvent.abi
// compiled/ScriptEngineEvent.bin
// DO NOT EDIT!

package gen

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _compiledScriptengineeventAbi = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\xcf\xb1\x0a\xc2\x40\x10\x04\xd0\x7f\x99\xfa\x2a\x41\x8b\x2b\xfd\x07\xab\x90\xe2\xe4\x56\x08\x9a\xd9\x70\xb7\xab\x86\xe0\xbf\x8b\x12\x8c\x45\xb0\xb7\x1c\x98\x37\x30\xcd\x84\x44\xe5\xd8\xab\x57\xc4\x53\xba\x54\x09\xe8\x38\xb8\x55\xc4\x66\x42\xc7\x2c\x77\xc9\x88\x56\x5c\x02\x98\x7a\x41\x84\xde\x28\x05\x01\x36\x0e\xaf\x98\x72\x2e\x52\x2b\x1e\xe1\x0b\xcc\x5b\xb3\x48\xbd\x3a\x6d\x21\xde\xd1\x36\xdb\xdd\x2f\x62\x7a\x16\xae\x88\xf6\xd3\xd8\xab\x33\x2f\x0d\xb9\x0a\xed\xbd\xf8\xbf\x97\x0e\x3c\xae\x9e\x6a\x9f\x01\x00\x00\xff\xff\x45\x51\x86\xa3\xa9\x01\x00\x00")

func compiledScriptengineeventAbiBytes() ([]byte, error) {
	return bindataRead(
		_compiledScriptengineeventAbi,
		"compiled/ScriptEngineEvent.abi",
	)
}

func compiledScriptengineeventAbi() (*asset, error) {
	bytes, err := compiledScriptengineeventAbiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "compiled/ScriptEngineEvent.abi", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _compiledScriptengineeventBin = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xcb\xb1\x91\x45\x21\x0c\x43\xd1\x96\x84\x6d\x19\xb9\x1c\xe0\x99\xfe\x4b\xd8\x21\xd8\xf9\x89\x02\xcd\xb9\x09\x21\x11\xa0\x79\x08\x83\x09\x5c\xce\x04\x20\xdc\x8f\x9b\x48\x38\x9f\x19\xdf\x7b\xbd\xde\x5e\x07\x7e\xe5\xbf\x06\xd6\x48\xa6\xcd\x35\xd7\x34\x07\x65\x70\x3f\xe9\xa6\xa1\x3b\x37\xbb\x07\x3b\x88\xed\xf9\x69\x74\xa7\x19\xf2\x88\xd1\xa9\x75\xc8\x8a\xa9\x5b\x7b\x77\x58\xa9\x7c\x85\x0e\x60\xf5\x17\x00\x00\xff\xff\x47\xae\x54\x0e\xa4\x00\x00\x00")

func compiledScriptengineeventBinBytes() ([]byte, error) {
	return bindataRead(
		_compiledScriptengineeventBin,
		"compiled/ScriptEngineEvent.bin",
	)
}

func compiledScriptengineeventBin() (*asset, error) {
	bytes, err := compiledScriptengineeventBinBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "compiled/ScriptEngineEvent.bin", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"compiled/ScriptEngineEvent.abi": compiledScriptengineeventAbi,
	"compiled/ScriptEngineEvent.bin": compiledScriptengineeventBin,
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
	"compiled": &bintree{nil, map[string]*bintree{
		"ScriptEngineEvent.abi": &bintree{compiledScriptengineeventAbi, map[string]*bintree{}},
		"ScriptEngineEvent.bin": &bintree{compiledScriptengineeventBin, map[string]*bintree{}},
	}},
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

