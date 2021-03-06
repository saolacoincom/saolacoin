// Code generated by go-bindata.
// sources:
// mainnet/delegates.json
// shoal/delegates.json
// DO NOT EDIT!

package preset

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

var _mainnetDelegatesJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\x90\x4b\xaf\x9a\x40\x18\x86\xf7\xfc\x8a\xc9\x6c\x6d\xed\x0c\xd7\x91\xc4\x05\xa3\xad\x56\x41\xf1\x0a\xd2\x34\x06\x66\x46\xc0\x0b\x58\x40\xb0\x9e\x9c\xff\x7e\xc2\x39\xba\xf8\x36\xdf\x93\xf7\x4d\xde\xe7\x8f\x04\xc0\x9b\x04\x00\x00\x30\x0b\x2f\x02\x9a\x00\x0a\x9e\x96\x79\xf6\x1d\x61\xf8\xed\x0b\x84\x9c\x17\xa2\x2c\x5b\x86\xee\x5c\x21\xb8\x17\x12\x16\xb1\x48\x37\x30\x42\x72\xc8\x89\xa2\x30\xa1\x19\x8a\x2c\x23\x7c\x20\x5c\x13\xd1\x81\x89\x57\xf8\x7a\x8b\xf6\x27\xf1\xbf\x0d\x53\x6a\x2d\xbd\x7a\x8d\x44\x30\x40\x9d\xc5\xe0\x61\x50\xb7\x77\xdc\x8e\x48\xba\x73\x23\x8e\xa9\xc0\x3b\x77\x9a\x8e\x7d\xbc\xa3\x79\xbc\xb0\xcb\xe4\x50\xfb\x46\x35\x99\x6e\x0a\xce\xf9\xb4\x76\x6c\xc6\x48\x42\xd5\x99\xb2\x2d\xef\xdc\xf7\x9d\x7a\x72\xd6\xd4\x60\xee\xf4\x4d\xd3\xdc\xae\x48\x5c\x5b\x96\x13\x0f\x83\xa4\x93\x56\x34\x0a\xb4\x9b\xa5\xb1\x39\x8e\x1b\x32\x5c\x51\x67\x93\x14\x1b\x25\x5e\xcf\x3b\x3f\x3d\x36\x8e\x7f\x0d\xd9\xea\xdf\xd9\x9d\x5d\x92\x1a\xdb\x47\xf6\x3b\x27\xeb\xd4\x5e\xb0\x91\xa7\x3b\xde\x44\x7e\x8c\xf0\x8f\x25\x89\xc6\xac\xb1\xfa\xaf\x15\x75\x5e\xa5\x59\xbc\xbf\xe6\x8d\x28\xa0\x09\x30\x42\x4f\x90\x89\xaa\xc9\x8b\xd3\xbe\x75\x04\xcd\xa7\x4a\x00\x60\x7a\x6d\x17\xab\x46\x57\x56\x95\x2e\x21\x5d\x55\x7f\x56\xb5\x4a\xf2\xa2\x82\x26\x20\xba\x81\x3e\x5f\xef\x52\x7b\x7f\xa5\x8f\x00\x00\x00\xff\xff\x9e\x30\x2c\xc9\x8c\x01\x00\x00")

func mainnetDelegatesJsonBytes() ([]byte, error) {
	return bindataRead(
		_mainnetDelegatesJson,
		"mainnet/delegates.json",
	)
}

func mainnetDelegatesJson() (*asset, error) {
	bytes, err := mainnetDelegatesJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mainnet/delegates.json", size: 396, mode: os.FileMode(420), modTime: time.Unix(1632293263, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _shoalDelegatesJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\xd7\x5b\xaf\xaa\x48\x9b\x07\xf0\xfb\xfd\x29\x56\xd6\xad\x3d\x5a\x27\xea\xb0\x93\xbe\x00\xe4\x20\x28\x8a\xa8\xa0\x93\x49\x87\xa3\x8a\x88\x08\x08\xea\xa4\xbf\xfb\xc4\x3d\xdd\x73\x31\xb1\xf3\x26\x6b\xbd\xb7\x54\x8a\xa4\xea\x97\xff\xf3\xd4\xf3\x9f\x3f\x3e\x3e\xfe\xfb\xc7\xc7\xc7\xc7\xc7\x67\x19\x9e\xd3\xcf\x9f\x1f\x9f\xcd\xe1\x12\x16\xff\x01\xe0\xe7\x6f\xff\xfb\x3d\x4c\x92\x3a\x6d\x9a\xd7\x12\xb8\x53\x89\x2b\x34\x09\x29\x43\x58\x30\x91\xaa\x12\xcd\xd2\x0c\x22\x41\x75\x08\x25\xc2\x18\xa1\x40\xe6\x4c\xb0\xbf\x37\x57\xb7\xe8\x8f\x53\xfa\x78\x6d\x56\xc6\xbc\xde\xf5\x79\x72\xac\xc5\xe1\x98\xb4\xd7\xfb\xd8\x96\xea\x28\x32\x6a\xc3\xdb\x09\x65\x07\x96\x82\xe7\x72\x5b\xdd\xed\x64\x6a\x81\x12\x2b\x91\x9e\x4e\x80\xd4\x46\x55\xeb\xb7\x9d\x25\x05\xe7\xd3\x99\x48\xb1\xdd\x57\xc6\xee\x7a\xc7\x22\x04\xb7\xd9\x29\x14\xe0\xf7\x9f\x3f\x7f\x66\xe2\x76\xf0\x8c\x78\x15\x2c\x95\xd3\xfc\x5a\x6d\x6f\xe3\x53\xde\x93\xd3\x39\xee\xce\xb2\x36\x35\xc2\x1b\x9b\x59\x1b\x63\xef\xf6\x8b\x81\x41\xf1\xa0\x7b\xf6\xda\x7a\xe0\xab\x37\x90\x02\x47\x57\xce\x96\x5e\x35\x2b\xb1\x73\xf7\x93\x2e\xac\x9e\xd5\x16\x26\xdd\xd1\xd6\x5d\xf9\xf7\xbf\x4f\xd1\x5d\xda\x63\xb9\xff\xa3\xba\xf4\x69\xfd\xf9\xf3\x03\x02\xf0\xd7\x42\x99\xb6\xfd\xa5\x3e\xfd\xf1\xba\xa3\xcf\x9f\x7f\xdd\xe4\xc7\xc7\xe7\xb1\x7a\x9d\x18\xe2\x21\x42\x7c\x28\xe0\x10\x32\xf4\xd7\xbf\x5e\x77\x72\xa9\xdb\xcf\x9f\x1f\x9c\x32\xf0\xeb\xd3\x9f\x3f\x3e\x3e\xfe\xfc\xed\x1f\x21\xd0\x5b\x08\x4c\x94\xb1\x60\x08\x64\x24\xe1\x78\x1c\xa1\x98\x27\x2c\x49\x39\x4b\x55\xcc\x23\xc6\x31\xc2\x00\xaa\x32\x7d\x0b\x61\x5e\x47\xce\x39\x8e\xd2\xb9\x37\xea\xb5\xeb\xce\x30\xb6\xf3\xfd\x99\x89\x51\x3b\x3d\xed\x1e\x00\x90\x89\x7e\x23\xfb\xdc\x6b\x59\x7e\x1c\x8f\xb2\x69\x6d\x35\x89\xed\x99\xd7\x74\x74\x71\xaf\x2b\xe9\xc1\x78\xdb\x52\x13\x3f\x6b\x00\xe9\xa1\x0b\xda\x74\xe4\xb4\x3d\xfb\x05\x11\x39\x55\x9f\x79\xb9\x15\x35\xa7\x35\x8c\x9f\x4a\x0e\xaa\x11\xda\x92\x06\xe0\xdc\xe4\xe7\xbb\x28\x37\x83\xf0\x10\x48\x12\x87\x92\xe9\x2c\x0d\x80\x07\xa5\x96\xcf\xa7\xa7\x03\xf1\x63\x67\xbf\x4e\xd1\xcc\xb5\x37\xab\x9e\xe3\xf3\xde\x59\x73\xeb\xc6\xc4\x4c\xe7\x7b\xed\xdb\x10\x7c\x08\x31\x1d\x42\x2e\x86\x04\x7f\x19\x02\xbf\x85\x48\x69\xcc\x30\x64\x32\xa2\x50\x03\x44\xc2\x21\x65\x02\xeb\x20\xd6\x10\xa1\x34\x96\x00\x4f\x70\x86\x62\x40\xde\x41\xc8\xdb\x9b\x3b\x48\xe2\xca\xdc\x2d\x58\x76\x54\x6e\x76\x06\x92\xa4\x15\xd7\xab\x16\xdd\xd3\xd3\x48\x1a\xeb\xe3\xfd\x85\x0b\x16\x2f\x89\x79\xd9\x8c\x83\xa2\x36\x5b\xab\x5b\x24\x2a\xdd\xaa\xca\x8c\xac\x63\x4f\xb9\x8a\xc1\x58\x31\x84\x98\x34\x2b\xcd\x2d\x33\x6f\x17\x6a\x2f\x88\x85\xed\x5f\xc9\xc6\xbd\x76\x01\x1d\x28\x03\x6c\xd4\xdd\xbc\xf4\xb4\xe9\x65\xd1\x22\xf9\x60\x67\xce\xe5\x19\x6e\xc9\x71\xf2\xc4\x2d\x1a\x77\x97\xb4\xb5\x93\x31\x03\xf5\x02\xb8\xfd\x8a\x53\xbf\x8c\x6d\xd8\xac\x65\x06\x66\x56\x3e\xb7\x46\x03\x2a\x76\xb2\x3b\x90\xbf\x0f\x21\xa1\x21\x42\x60\x08\x25\x30\x94\xbe\x0e\x41\xde\x42\xa0\x08\x02\x80\x13\x02\xe8\x18\x03\xce\xb9\xa2\xf1\x48\xa5\x44\x4a\x25\x82\xc6\x69\xcc\x04\x05\x2c\x0a\xc7\xd2\x3b\x08\xe5\x7e\x31\xf3\xd9\x6e\xaa\x92\xf0\xd4\x9c\x3a\x78\x9b\xd5\xbc\xb8\x56\xe7\x4c\x9b\xa6\xe7\xad\x82\x17\xb0\x3a\x6a\x17\xeb\x76\xb2\x8a\x20\x28\xf7\xb3\x31\xd1\x47\x84\xe5\xe6\x2a\x6b\xf2\xc5\x55\x76\xd3\x65\x56\xb6\x93\x85\x2f\xb0\xfe\x50\xd6\x6c\xb6\x93\xd1\xe5\x26\xff\x2a\x4d\xbd\x73\x6b\x8f\xe4\xe0\x60\xa3\x4f\x56\x0b\x2b\xf4\x3c\xd7\xe8\xd7\x68\xde\xfa\xea\x7d\xc4\xd7\x89\x99\x6e\xdb\xa9\x26\x97\x92\xef\x1f\xe7\x79\x7e\x50\xec\x3a\xbf\xa3\x38\x8e\x36\x9a\x7b\x29\x8f\x0f\x64\x8e\xac\xa9\xa0\xa7\xed\xe2\x5a\x6d\xfb\x82\x41\xa4\x1c\xe5\x6f\x97\x26\x89\x0c\x91\x04\x87\x08\x80\x21\xe6\x5f\x86\x90\xde\x42\x00\x4a\x10\x96\x09\x87\x8c\x53\x28\xa9\x3a\x41\x99\x60\x0a\x4f\x85\x42\x75\x09\x4b\x50\xc3\x12\xd4\xa9\x84\xde\x41\xa8\xc1\xfc\x32\xbf\x85\x75\xb2\xec\xb7\xb2\x99\x37\x2e\x9a\x2e\x2c\xd6\xdd\x1c\x7f\x1e\x4d\xc7\x42\x82\xea\x99\xfb\xbe\xb1\x2b\x6c\xdd\x6c\x46\x97\x81\x21\x26\xe6\xe6\xde\xf2\x4e\xf6\xb1\x83\x15\x3a\x39\x64\xe6\x12\xb2\x7c\x62\x79\xc5\xad\x21\xdb\xd5\x66\xd5\x1f\xdd\x17\x44\xed\x2b\x63\x32\xd9\x9e\x5c\xdb\x18\x3d\xad\xf5\x83\xd4\xa7\x06\xcf\x76\xb2\x71\x8d\xcc\xd9\xa6\x7a\x5c\x82\x89\x77\x43\x45\xee\x5a\x2b\x6d\xa1\x6f\x92\xb2\x70\xe5\x63\xe6\x9d\x79\xd7\x1a\xb3\x62\xaf\xdf\x8f\x99\x28\x72\x33\x68\xd7\xb9\x65\xfb\x2a\x08\x56\x1b\xc5\x95\xbf\x9f\x88\x17\x04\x19\x42\x38\x44\x10\x7c\x19\x82\xfe\x43\x69\x12\x4a\x82\xa9\x90\x80\x40\x00\x47\x29\x06\x3a\xa7\x58\x4f\x90\x4e\x18\xd1\x28\xcb\xd2\x98\x4b\x02\xc4\x6f\x21\x1e\x07\x11\xd3\x83\xa9\x47\xa3\x2a\x4d\x88\x8c\xb6\x30\xb0\x56\x6e\x1e\xb2\xcc\xaf\xdb\xee\xf9\x38\xd2\x46\x99\xef\x4e\x97\xc0\x89\x12\xb4\x8e\x0b\xbb\x98\xc0\xb5\x54\x68\x3d\x16\xc7\xc5\x4e\xa9\xe4\x7e\x39\x9f\x8c\x84\xe1\x1d\xd0\xc9\x0d\x69\xd0\xb6\xeb\xd5\xaf\x1e\x91\x2f\xd4\x4c\x1b\xd8\x6d\x49\xac\xc5\x4c\x8b\x0d\xe3\xd0\x9c\xba\x49\xe2\x80\x9c\x47\xb9\xe7\xcf\xf2\x68\xba\xc8\x77\x2d\xdb\xf8\x7a\xe4\x21\x58\xfa\x46\x52\xba\xba\xd6\x1f\xd6\xb6\xc3\x71\x9d\xce\xd2\x32\x8b\x02\x79\xaf\x6e\x93\xeb\xcc\xa7\xd3\x07\x0e\x9f\x41\xff\xef\x49\x04\x19\x42\xc0\x87\x08\x89\x2f\x4b\xb0\xf7\x12\x24\xa1\x98\xeb\x21\x53\x25\x0c\x39\x94\x20\xd0\x11\x81\x8c\x52\x3d\x44\x80\x22\x40\x75\x20\x71\x2d\x7a\x5b\x9b\x1c\x5b\xbe\x5c\x5b\x53\x8f\x54\xed\x51\x25\xb6\x5b\x78\x61\x77\x18\x75\x60\xb1\x90\x3b\xac\xfb\xfb\xd9\xce\xe8\xc4\x60\x77\x73\xc4\xd4\x3e\x69\x24\x2d\x7d\xf8\x5c\xaf\x1f\x87\xba\xba\x59\x78\xab\x51\x7e\x98\xad\x27\x03\xb4\xf4\xca\x07\xdd\xf6\x58\x26\x42\x5b\xa8\xfb\x97\xc4\xba\x66\x85\x55\x54\x22\x52\x07\xb7\x99\x7c\x3c\x74\x77\xca\xe5\xc3\xca\xf3\xcb\x75\x79\x08\xa1\x47\x17\xa7\xe9\xac\xa5\xf7\xbd\x47\xcd\xe5\x6a\x16\x61\x25\xa7\x47\xc3\x91\xb7\x46\x03\x6f\x99\x1f\xa4\xd7\xa7\xf0\x6e\xc8\x1e\x08\x1c\x08\xba\x0e\xd2\xc0\x57\x07\xfd\xb7\x23\x41\xc8\x10\x11\x38\xe4\xd2\x50\x82\x5f\x76\xe0\x6f\x1d\x44\x94\x00\x99\x11\x95\xa9\x28\x96\x60\xac\x63\x80\x08\x8d\x39\x53\x38\x93\x54\x9e\x49\x90\x01\x2a\x03\x2e\xde\x39\x58\x69\xbf\xc2\x1d\xd2\xab\xa4\xb5\xef\xed\xc4\x13\x03\x3d\x96\xdc\x49\xa4\x04\x2b\xde\x38\xab\xb6\x62\x41\xbe\x9b\x5c\xca\x6e\xb3\x54\x62\xa3\x10\x9b\x68\x15\x5a\xf9\xf9\xe1\xba\x88\xc8\xcb\x6d\x29\xdf\x12\x7c\x59\x3f\x7c\x7e\xf5\x6e\x03\xe3\x74\x85\x77\x3d\x2d\xf9\xcb\x61\xc2\x58\xf4\x74\xca\x29\x0b\x0e\x92\xd7\xf3\x67\xb7\x01\x79\x19\x64\x17\xed\x98\x58\x55\x57\x8a\x69\x32\x2e\x77\x22\xe2\x67\x0f\x58\xfb\xe0\xb2\x54\xf6\x52\x30\xbf\x6d\xb9\x7f\xbc\xe8\x45\x07\x94\x31\x1e\xa8\x7d\x44\x7b\x5c\xa5\xce\x15\x7b\x9b\xd0\x7c\xdc\xcf\xfb\xef\x27\x02\x0d\x11\x19\xf2\x21\x61\x5f\x56\x10\x6f\x15\x48\xa8\x67\x4a\x84\x24\x28\x65\x50\x62\x82\x47\x0a\x88\x63\x94\x72\x28\xb8\xa6\x8c\x01\x00\x0c\x40\x69\xfc\xfe\xed\x2a\xcf\x60\xe1\x4c\x96\x00\x10\x2f\xbe\xaa\xcd\xd9\xc1\x6a\xdf\x48\xc1\x56\xad\xb2\xa8\x6f\xd2\xda\x5b\x34\xbe\x5d\x99\xfa\xca\xd8\x8d\xd4\x04\x2d\x4c\xea\x8d\x5a\xe9\x90\x45\xca\xaa\x9c\xa8\xa4\xb1\x09\xde\xac\xf7\x89\x2d\xcb\xdd\x88\x8c\xf7\x93\x3e\xbb\xae\xa6\xb3\x97\x42\x7b\xd7\x89\xdf\x2f\x8a\xf1\xac\x5e\x37\xe9\x26\x83\x3a\x0e\xc0\x44\x2e\x23\x37\xbb\xba\x62\x6b\x38\xc8\x55\xcd\x6a\x7e\x6e\xeb\x09\x9b\x02\x70\x48\x8c\xbd\x5f\xc8\x7b\x70\xed\xd9\x02\x14\xee\x34\x7c\x8c\x9b\x4d\x4e\x16\x1e\x59\x7b\xe8\xa9\x84\xd9\xb2\xf6\xc9\xf7\x87\x88\x5f\x69\x00\x43\x8e\x87\x8c\x7e\xd5\xe1\xff\x3a\xcb\xff\x4b\x43\x3a\x86\x02\x12\x1d\xca\x94\x49\x82\x11\x28\x23\x1c\x03\x39\x86\x98\x41\x44\x30\x81\x44\x93\x29\x44\xca\x3b\x87\x89\xeb\x9d\xe1\xa2\x24\xa5\xda\xd5\x22\x69\x45\xe1\xf4\x53\x9c\xab\xb9\xd1\x3f\xc5\x34\x34\x7b\x68\xce\x48\x1c\x2b\x7a\xf9\x44\x9b\x93\x0d\xe6\xcd\x63\xf4\xcc\x96\xc7\x50\x1a\x9c\x97\xb6\xea\x3d\x17\xb6\xb2\xd9\xe9\x91\x9d\x49\x9e\x6b\xd9\x73\xeb\x1c\x2c\x07\xfa\xe9\xe5\xe0\xc5\x83\x78\x7c\x2f\x78\x7c\x49\xb6\x07\x34\x5b\xad\xdb\xb4\x6b\x76\x1e\xdd\xf5\x7e\x30\xd8\xcb\x03\x54\xa2\x67\x5f\xd5\x83\x91\xb5\x22\x65\xb3\xce\x4b\x4f\xdb\x37\x92\xef\x0e\x0e\xa6\xfc\x68\x7c\x52\xa4\x8f\x68\x32\x51\xf7\xe9\xf6\xba\x43\x1b\x18\x6c\xaf\xb6\x9c\x7e\xff\xc5\xf4\x72\x40\xbf\x7a\xc3\x10\x82\x2f\x27\x02\xbe\x1f\xab\xe5\x98\x84\x71\x12\xd2\x44\xa4\x84\x30\x38\x0e\xa5\x08\xea\x98\x43\x29\x66\x29\x15\x12\x52\x13\x4c\x54\x05\x46\xef\x24\xa6\xd2\xac\xc8\x99\x81\x9b\xaa\x76\xb4\xe5\x33\xa9\x37\x56\xd9\x64\x0c\xfb\xb0\x53\x0a\x3f\xd0\xc3\x33\xdb\x4d\x59\xa6\x5a\xbe\x6e\x99\xc9\xd2\x19\x78\xb2\x5c\x2d\x6f\x33\x43\x98\xe3\x8d\x31\xdf\x8c\x84\xef\xdc\xf6\x4d\x82\x3a\xd7\x2f\x96\x30\xdf\x66\x7c\x89\x2f\x2f\x89\xdd\xe9\x50\xa6\x0b\x2d\x20\xfb\xc3\xd9\xd2\x55\x10\x2d\x73\x0b\xf3\xfb\xb8\x7a\x96\x4b\xd8\x3c\x27\x1c\x13\xd9\x08\x94\x8b\x59\x65\xd5\x92\xb8\x87\xb0\x9c\x16\xfb\xcd\xb3\x75\x07\x5a\x72\xdc\x8b\xdd\xed\xe8\x08\xa2\x75\x93\xfd\x61\xf7\x4c\x1c\x9f\x3b\x0b\x00\x67\xdf\x1f\x22\x5e\x12\x98\x0e\x11\xc0\x43\x28\xfe\xd5\x5c\xfd\xe3\xbf\xfe\x27\x00\x00\xff\xff\x61\x0f\xee\x71\xe3\x10\x00\x00")

func shoalDelegatesJsonBytes() ([]byte, error) {
	return bindataRead(
		_shoalDelegatesJson,
		"shoal/delegates.json",
	)
}

func shoalDelegatesJson() (*asset, error) {
	bytes, err := shoalDelegatesJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "shoal/delegates.json", size: 4323, mode: os.FileMode(420), modTime: time.Unix(1632280835, 0)}
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
	"mainnet/delegates.json": mainnetDelegatesJson,
	"shoal/delegates.json": shoalDelegatesJson,
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
	"mainnet": &bintree{nil, map[string]*bintree{
		"delegates.json": &bintree{mainnetDelegatesJson, map[string]*bintree{}},
	}},
	"shoal": &bintree{nil, map[string]*bintree{
		"delegates.json": &bintree{shoalDelegatesJson, map[string]*bintree{}},
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

