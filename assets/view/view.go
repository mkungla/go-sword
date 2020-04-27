// Code generated for package view by go-bindata DO NOT EDIT. (@generated)
// sources:
// view/layout/default.html
// view/layout/model.html
package view

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

var _viewLayoutDefaultHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x56\x4d\x73\xab\x36\x14\xdd\xe7\x57\xdc\xaa\x0b\x92\x4e\x30\x7e\xe9\x5b\xa9\xc6\x5d\x76\xdb\xe9\x74\xf7\x26\xd3\x91\xd1\x05\x34\x11\x12\x23\x09\x92\xd4\xe3\xff\xfe\x46\x02\xdb\x60\xc0\x76\xb4\xc8\x38\x47\xf7\xf3\xdc\xc3\x85\x8d\x75\x9f\x12\xc1\x66\xba\x46\xbe\x7d\x00\x00\x58\x49\xf6\xa9\x1b\x17\x2b\xd6\xc2\x3e\x20\xfe\xbc\x0b\xee\x4a\x0a\xdf\x5f\xd6\xf5\xc7\x1f\x27\xb4\x62\xa6\x10\x8a\xc2\x1a\x58\xe3\xf4\x25\x1e\x1b\x51\x94\x8e\xc2\xd8\x27\xd3\x52\x1b\x0a\xef\xa5\x70\xd8\xa1\x87\x87\x51\x62\xa9\x0b\x3d\xc8\xbc\xd3\x86\xa3\x89\x0d\xe3\xa2\xb1\x14\x7e\x1f\xc6\xca\xa5\x66\x8e\x82\xc4\xdc\x1d\x43\x6d\x92\xd0\xd2\xf6\x61\xe3\xb0\xaa\x25\x73\xd8\xb5\xb5\xe1\xa2\xdd\x9e\x1c\x37\x22\xee\xb2\x9d\xa1\x1e\x2e\x91\x71\x34\x40\x43\x94\x94\xec\x77\x2c\x7b\x2b\x8c\x6e\x14\xa7\xd1\xaf\xeb\xf5\x3a\x3a\x90\xb1\x4f\xef\x57\xa1\x6a\xa0\xd2\x1c\x53\x52\x6a\x23\xfe\xd7\xca\x31\x49\xc0\x95\x58\x61\x4a\x38\x33\x6f\x04\x58\xe6\x44\x8b\xb1\x62\x1e\xfa\x46\xbe\x98\x24\x24\x2a\x5f\x20\x93\xcc\xda\x94\x0c\xd8\x22\xd0\x07\xea\xb8\xed\xa8\x25\xdb\x7f\x85\x93\xb8\x49\xca\x97\x85\x58\x5c\xb4\x17\xc1\x14\x6b\x97\x12\x27\x23\xfe\xce\x70\xd7\xf9\x05\x8d\xc9\x91\xc7\x09\xbd\x73\xac\xf7\x57\x56\x78\xe2\x4b\xc1\x31\x76\x46\x14\xc5\x70\x0a\x50\x62\xa7\xa5\xe8\xdb\x7a\xdd\x96\xd1\xf3\x88\xb1\x3c\xcf\x97\x19\xeb\x47\xd3\x0f\x42\xfa\x30\xa4\x93\x73\x4a\xbc\x6a\x17\xfc\x82\xef\x51\x42\xd0\xc6\xb9\x36\x29\x11\x0e\x2b\x10\x0a\x7c\xc4\x2b\x7e\x83\xbc\x71\x70\x69\x63\x91\xa7\xc4\x7d\xd6\xa8\xf3\x47\x8f\xac\x6c\xb3\xfb\xcf\xdf\x3f\x41\x9a\x42\xd4\x28\x8e\xb9\x50\xc8\x23\x02\xb4\x93\x47\xb0\xf2\x3f\x6f\xe4\xe9\x73\x89\x4c\x2b\xa0\x3e\x43\xef\xea\x01\xb2\xf5\x93\xf0\xbf\xee\x88\x61\x74\xe3\xd0\xc4\x52\xa8\x37\xa0\x4e\xa7\xa4\x40\xf7\x8f\xc7\xfe\x66\xae\x7c\x3c\x95\xf3\x44\xc0\xb1\x22\x25\xb6\x66\x8a\x6c\xf7\xfb\x70\xe1\xbc\xd2\x0e\x87\x4d\x32\x08\x72\x83\x9e\x64\xc0\xcf\x4d\x26\x6d\xb3\x0b\x43\xbc\xc6\xe3\x2f\x63\x1e\xef\xe8\xf8\x34\x5d\x2b\xb5\x4b\x49\x68\xe2\x0e\x3f\xb8\x8f\xf1\x31\x35\xb7\xab\x49\xc6\xfb\xea\x46\xf6\xa1\xb6\x82\x34\x6d\xb3\xf3\xca\x1c\x91\x72\x12\x93\x6d\x76\xf7\x6a\x09\x6e\x6b\xe1\x18\xed\x52\x0a\x1e\xff\xba\x12\xe0\x8b\x6a\x48\x4e\x72\xb8\xf2\xe0\xde\x20\x73\x7e\x6d\x9d\xc2\x8b\xc9\xe6\x82\xe1\xf6\x3a\xaf\xa5\x9a\x71\x2e\x54\x41\x21\x5a\xc3\xcb\xf7\xfa\x23\xfc\x59\x5e\x45\x3d\x27\xad\xc0\xf7\xed\x89\xa1\xf0\xdf\x6c\x1d\xb3\x6f\xa8\x19\xfc\x12\xeb\x77\xf5\x90\x84\x8d\xcd\x8c\xa8\xfb\xfb\x4a\xf3\x46\xe2\x0a\x3f\x6a\x6d\x9c\x85\x74\xf0\xb2\xe5\xcc\x31\x0a\x79\xa3\x32\x27\xb4\x82\xc7\xa7\xc1\x9d\x3f\x06\x5d\x63\xd4\x05\xe8\x4f\x92\x40\xa6\x55\x2e\x0a\x28\xa4\xde\x31\x19\x56\x24\x48\x61\xdd\xc4\xd4\xdf\x50\xf8\x31\x8d\xe1\x8f\x7f\x74\x28\x44\x42\xdb\xb8\x46\x5d\x4b\x8c\x9e\x67\xed\x82\xce\x28\x44\x7f\xa1\x42\xc3\xdc\x92\x99\x97\x29\x85\xc8\xbf\x97\x65\x34\xb1\x38\xbc\x4e\xbd\x26\x8d\x84\x31\xd9\x89\x5d\x07\x2f\xf6\x51\x33\xff\xb9\x14\xfd\xb6\x50\x97\x41\x2e\x0c\x66\x8e\xc2\x7e\x54\xe2\x74\x4f\x1c\x9e\x67\xd8\x9e\xb4\x36\x9f\xa5\x2f\x22\xb9\x66\xd3\x18\xe9\x4d\x0c\x2a\x8e\xe6\x4f\xef\x91\x26\x9d\x98\x92\x45\xd6\x46\xd0\xb9\xe4\xc3\xf3\xe0\x23\xaf\xaa\x1b\x87\x9c\xc2\x7e\x80\x56\xe8\x4a\xcd\x2d\xbd\xe8\x68\xb8\x5c\x86\xe2\x0b\x3b\x66\xa6\xfb\x96\x99\x9e\x7e\x48\xc1\x95\xc2\xae\x16\x66\x94\x6b\x03\x8f\xde\x5a\xf8\xd5\xd8\x19\xcd\x05\xf4\x47\xe4\x5d\x42\xff\x3e\xee\x2c\x7f\x88\xd7\xd5\x52\x09\xc7\xd3\x3f\x0f\x67\x07\x4f\xe0\xac\xf5\xcc\x64\xa7\x9a\xea\xa2\x45\xd1\x12\xc1\xe7\x6f\xdc\xee\x71\xfe\x19\x00\x00\xff\xff\x19\x51\x03\xb6\xbf\x0b\x00\x00")

func viewLayoutDefaultHtmlBytes() ([]byte, error) {
	return bindataRead(
		_viewLayoutDefaultHtml,
		"view/layout/default.html",
	)
}

func viewLayoutDefaultHtml() (*asset, error) {
	bytes, err := viewLayoutDefaultHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "view/layout/default.html", size: 3007, mode: os.FileMode(420), modTime: time.Unix(1587958589, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewLayoutModelHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\x5f\x8b\xe3\x36\x10\x7f\xdf\x4f\x31\x15\x85\x64\x61\xe3\x6c\x17\x5a\x8a\x2f\x0a\x57\x7a\xd0\x97\x3b\xb8\x72\xed\xd3\x72\x04\xad\x3d\x8e\xd5\x93\x25\x23\x29\xc9\xa6\x47\xbe\x7b\x91\xec\x24\xb6\x23\xdb\xb9\x6b\xcb\xad\x1f\x96\x78\x34\xff\x34\x33\xbf\x99\xf1\x2e\x2c\x16\xa5\x60\x16\x97\x37\x00\x00\x8b\x94\x6f\xc1\xd8\xbd\x40\x4a\x0a\xa6\xd7\x5c\xce\xac\x2a\x63\xf8\xf1\xbe\x7c\x26\x15\x8b\x67\xe3\x33\xad\x76\xe7\xf7\x9a\x96\x28\x01\xa6\x64\x92\x92\x1f\x1e\x48\xfb\xb4\xe6\xc8\x94\x2e\x20\x16\xec\x09\xc5\x6c\xc7\x53\x9b\x53\xf2\x70\x7f\x4f\xa0\x40\x9b\xab\x94\x92\x52\x19\x1b\x90\x6c\x48\xcf\xb8\xc5\x02\xbc\x06\x4a\xde\xed\x3f\xfc\xfe\x16\x2c\x7b\x12\xd8\x23\x55\x4b\xb2\x8d\x55\xb3\x44\x15\xa5\x40\x8b\x10\xa7\xcc\x32\x4a\xbc\xdc\x4a\x70\x63\x09\x6c\x67\x85\x4a\x9d\xca\x8a\x28\x59\x81\xa4\x57\xe1\xc5\x53\x0a\x96\x60\xae\x44\x8a\x9a\x92\x5f\x05\x4f\x3e\x41\x8e\x1a\x21\x8a\x22\xb2\x5c\xcc\x3b\xf6\x7b\xee\x37\x6f\x5c\x70\x3c\x04\x83\xd7\x7d\xda\x58\xab\x24\xd8\x7d\x89\x94\x94\x9a\x17\x4c\xef\x09\xbc\x4e\x9c\x67\x8e\x80\x5b\x8e\x3b\xb2\x7c\x5f\xfd\x70\x96\x2b\x89\xeb\x95\x9a\x4d\x92\xa0\x31\x67\xa5\x6b\x94\xa8\x99\x45\x02\x71\xca\x8d\x0b\x62\x4a\x09\x37\xab\xe3\xcb\xea\xc8\xb0\x7a\xb2\x92\x2c\x7f\xab\xdf\xc6\x6c\x8f\x44\xe5\x78\x1c\x2c\x36\x26\x50\xdb\xda\x5f\x02\x26\x57\xbb\x19\x4f\x7a\x0d\xb9\xc2\x05\x23\x94\xa5\x24\x45\x93\xb8\xbc\x39\x52\xd8\xa6\x57\xdd\xa9\xff\xb9\x07\x40\x03\x24\xf3\x0e\x4a\x46\x51\xf3\x70\x1f\x42\x4d\x03\x92\x16\x9f\xed\x8c\x09\xbe\x96\x31\x68\xbe\xce\x07\xb0\x52\x67\xeb\x98\x9e\x9c\xc9\x54\xe0\x07\x14\x98\xd8\x5f\x84\x98\x5a\xbd\xc1\x5b\xb2\xac\x08\xc0\x84\x18\x4d\xc4\xa8\xca\x8c\x09\xe3\x74\xfe\x29\xcd\x55\x5a\x17\xf3\x94\x6f\x03\xe4\x27\x1d\xcc\xa6\x07\x26\xc4\x89\x12\x9b\x42\x9a\x23\x50\x13\x25\x48\x1b\xcf\xee\x37\x01\x8d\x19\x25\x95\x1f\x5c\x49\x02\xaf\x95\x9c\x9d\x5e\x67\x49\xce\xe4\x1a\x8f\x12\x27\xfa\xaa\xa2\xf7\x05\xf5\xd8\x2a\x7d\x95\xcc\x4c\xa2\x5c\x61\x7d\x06\xad\x76\x77\xc0\x65\x8a\xcf\x70\x20\x75\x05\x55\x8a\xab\x9b\x0f\x77\xa6\x16\xaa\xb8\xcc\x14\x01\xc3\xff\x76\x08\x2b\x98\x10\x67\x7c\xa9\x12\xe5\x3b\x95\x32\x31\xf5\xa6\x6e\xc9\xf2\x0d\x5a\xc6\x47\xf3\xd6\xf6\x7c\x3b\xe3\x19\x25\x5a\xed\x22\x0f\xcd\x2c\x03\x4a\x61\xb2\x9f\x0c\xf8\x18\xf2\x73\xc7\xb4\xe4\x72\x3d\xe0\xea\x1b\x9e\x65\x6d\x77\x9d\xb5\xab\x9c\x9d\xb7\x47\xd2\x97\x9c\x3b\xfd\x3e\xf4\x5f\x01\xcd\xc2\x79\x7b\x9e\x03\xdc\xac\x5c\xc7\x58\xa5\x3e\xca\x2b\x4f\x25\x10\x5b\x6e\x1d\x10\x9b\xd4\x95\xa7\x11\xa8\xe7\xd9\xcf\x5d\x10\x2f\x4a\x8d\xcb\x45\xa2\x52\x17\x7e\x87\xe0\x8e\x78\xa2\xa4\x45\x69\x5d\xc3\x71\x4c\xcb\xc5\xdc\x09\xb4\x7c\xf5\xbe\xd5\x13\xba\x42\x4d\x33\x0a\x0b\x93\x68\x5e\x1e\x3b\x9d\xb3\x30\xff\x8b\x6d\x59\x45\xad\x9d\x29\x54\xba\x11\x18\xe1\x73\xa9\xb4\x35\x40\xe1\xf3\xc9\x80\x03\x4c\x0c\xd9\x46\x7a\x0c\xc0\xf4\xb6\x71\xe6\x1e\x8d\x76\xa3\x65\x87\xe8\x9e\xf3\xfc\x8c\xe1\xf1\xe3\x5d\xcf\xb9\x1b\xa5\x31\x4c\x26\x7d\xe7\x89\x12\x31\x3c\x06\x53\x7d\x69\xf2\x24\xba\x2f\x9d\x52\x5f\x5b\x01\xcd\xc7\xc7\xa7\x24\x86\x9f\xee\xfb\x59\xea\x6e\x3a\x49\x50\x5a\xd4\x93\x20\xdf\x21\x2c\x3e\xe0\x9e\xab\x88\x18\x26\x25\xb3\xf9\x80\x7b\x9f\x70\x1f\xc3\x24\xe3\x02\x57\x03\x9c\x5f\x6d\x5d\x95\x6e\xc4\x72\x25\x07\x5c\x70\xed\x2a\x86\x76\xbf\x1a\x8d\xe7\xc3\x7d\x4f\x40\xbf\xdc\xd5\x2a\x8f\xa7\x0e\xfc\x4d\x72\xd9\x5b\xba\x15\x32\x42\xa5\xbd\x63\xdc\x72\xb9\x76\xa9\x33\x61\x8e\x50\xff\x88\xc1\xcf\xc8\x4b\xe6\x50\x3f\x08\x63\xe6\xb2\xf1\x84\xf9\xfa\x36\xaf\xd8\x0d\xfe\x36\xfb\xe1\x26\x10\x9b\x6a\x21\x37\x71\x27\x77\x6b\xb4\x7f\x38\xa5\x6f\x3d\xe8\xfb\x7b\x86\x7b\x04\x5a\xb0\x39\xb3\x40\xc1\xe6\xdc\xdc\x5c\x06\x39\xe7\x26\x62\xcf\x5c\x99\x60\xa6\x22\xf7\x31\x30\x9d\xcc\x59\xc9\xe7\xfe\xbe\xf3\x73\xc3\x99\xdc\x86\x45\x6c\x8e\x72\xaa\xd1\x94\x4a\x1a\x04\xba\x1c\x28\x3d\x9e\xc1\x89\x33\x72\x99\x8e\x7c\x8b\xa6\xd4\x95\x77\xe8\x3e\x6d\xd7\x99\x8d\xce\xee\x00\x85\xb6\x2a\xff\xc7\x9d\xf4\x6a\x39\x00\x0a\x83\x23\x66\xfc\x9a\xd9\xf1\xb2\x30\xeb\xf0\xe5\xa1\x95\xcb\x16\xb5\x2d\xd0\x81\x40\xb5\xc7\xf9\xb4\xbe\xe9\x8c\x02\x7b\x24\x86\xe2\xe1\xd3\x77\x06\x8a\x4b\xf3\x91\xfd\x82\x37\x51\xd2\x28\x81\x91\x5b\x70\xa6\x1d\xc1\x41\xe7\xea\x6f\x95\x7f\x5d\x6b\x2e\xdb\x0d\xbb\x6e\x28\x01\xa5\x14\x08\xe9\x4b\x75\x15\x7b\xf2\x5e\x20\x33\x08\x49\xae\x94\x41\x68\x7e\x70\xde\xbe\x0a\xca\x55\xe3\xf2\xe2\xe8\xf0\x1f\xd4\x7f\x1d\x8c\xc9\x1d\x7c\x6e\xce\xd6\xce\xbd\x0e\xff\x23\x36\xae\x07\x47\xa7\xaa\xa6\x3d\xf8\xe8\x2f\xe5\x93\xa6\xbe\x56\x06\xb4\x6a\xa7\xe1\x34\xc0\x8b\xc2\xd8\xd1\xef\xb1\x3a\x3e\x95\x69\x73\xc0\x44\x02\xe5\xda\xe6\x3e\xfc\xbd\xc1\x0f\x96\xab\x17\xff\x8e\x84\xef\x72\x7d\xa1\xbe\x44\xf0\x74\x11\xdf\x3d\xff\x62\x70\x1d\x53\x34\x82\xae\xbb\x6a\xe6\x5f\x24\xe9\x9b\xa3\xae\xd5\x62\x5b\x6a\x06\x4a\xf9\xc5\x00\xa4\xfb\x79\xdb\x37\x6f\x42\x9b\x52\x5d\x02\x8d\x91\xf2\xe8\x95\x7c\x8c\xfc\x6a\x5d\x73\x5d\xa1\xcf\xef\x53\x23\xda\xdc\xa2\x1e\x56\x15\xda\xf8\x9c\x32\xbd\xc1\xf1\xf1\x7b\xfe\x37\x8a\xb1\xcc\x6e\x4c\x6f\x00\xbe\xd7\x98\x99\xe8\xb4\x30\xd7\xbf\x1a\x82\xaf\x86\x8c\x85\xff\xe1\x31\xd5\x6a\xd7\x6f\xb1\x59\xe5\x40\xe1\xf1\xe3\x05\x5b\xa6\x34\x4c\xb7\x4c\x03\x07\x2e\xa1\x4f\x59\x50\x61\x54\x6e\x4c\xee\xed\x3f\xf2\x46\x84\x03\xed\xe1\x70\xc5\xd2\xaa\x36\xd2\x62\x3a\xd4\x64\xbd\x03\xcd\x2d\x76\x7a\x2e\xcb\x4a\xe7\xe1\x66\x31\xaf\x3e\x9d\x97\xff\x04\x00\x00\xff\xff\x6d\xdb\x3e\x8e\x1d\x17\x00\x00")

func viewLayoutModelHtmlBytes() ([]byte, error) {
	return bindataRead(
		_viewLayoutModelHtml,
		"view/layout/model.html",
	)
}

func viewLayoutModelHtml() (*asset, error) {
	bytes, err := viewLayoutModelHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "view/layout/model.html", size: 5917, mode: os.FileMode(420), modTime: time.Unix(1587958667, 0)}
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
	"view/layout/default.html": viewLayoutDefaultHtml,
	"view/layout/model.html":   viewLayoutModelHtml,
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
	"view": &bintree{nil, map[string]*bintree{
		"layout": &bintree{nil, map[string]*bintree{
			"default.html": &bintree{viewLayoutDefaultHtml, map[string]*bintree{}},
			"model.html":   &bintree{viewLayoutModelHtml, map[string]*bintree{}},
		}},
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
