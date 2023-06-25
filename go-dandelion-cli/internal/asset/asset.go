// Code generated for package asset by go-bindata DO NOT EDIT. (@generated)
// sources:
// internal/template/boot/boot.tmpl
// internal/template/cmd/apiserver.tmpl
// internal/template/cmd/cobra.tmpl
// internal/template/cmd/rpcserver.tmpl
// internal/template/config/config.tmpl
// internal/template/global/global.tmpl
// internal/template/internal/route.tmpl
// internal/template/internal/rpcapi.tmpl
// internal/template/main.tmpl
package asset

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

var _internalTemplateBootBootTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\x52\x00\xad\xff\x70\x61\x63\x6b\x61\x67\x65\x20\x62\x6f\x6f\x74\x0a\x0a\x66\x75\x6e\x63\x20\x49\x6e\x69\x74\x28\x29\x20\x7b\x0a\x09\x2f\x2f\x20\xe5\xb0\x86\xe9\x9c\x80\xe8\xa6\x81\xe5\x88\x9d\xe5\xa7\x8b\xe5\x8c\x96\xe7\x9a\x84\xe6\x96\xb9\xe6\xb3\x95\xe5\x9c\xa8\xe8\xaf\xa5\xe5\xa4\x84\xe6\xb3\xa8\xe5\x86\x8c\x20\x54\x4f\x44\x4f\x0a\x7d\x0a\x01\x00\x00\xff\xff\xa8\x42\xb5\x5d\x52\x00\x00\x00")

func internalTemplateBootBootTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplateBootBootTmpl,
		"internal/template/boot/boot.tmpl",
	)
}

func internalTemplateBootBootTmpl() (*asset, error) {
	bytes, err := internalTemplateBootBootTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/template/boot/boot.tmpl", size: 82, mode: os.FileMode(420), modTime: time.Unix(1686111346, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _internalTemplateCmdApiserverTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x56\x5d\x6f\xdc\x44\x14\x7d\xf6\xfc\x8a\x8b\xa5\x54\x76\xb5\x6b\x0b\xf1\xb6\x22\x0f\x28\xe4\x4b\x42\xe9\x2a\xdb\xf2\xd2\x22\x34\xb1\xc7\xde\x51\xec\x19\x33\xbe\xde\x6c\x15\xad\x14\x5a\x68\xa1\xa4\x02\x41\xda\x87\x50\x24\x2a\xf1\x11\xf1\xd0\x14\x15\xa9\x2a\xa2\xf4\xc7\x50\x27\xbb\xff\x02\x8d\xc7\xbb\xd9\x26\x4d\x49\xf0\x83\x3d\xf6\x9c\x3d\x73\xef\xb9\xe7\xce\x6c\x46\x83\x75\x1a\x33\xa0\x19\x27\x84\xa7\x99\x54\x08\x0e\xb1\xec\x28\x45\x9b\x58\x4a\x16\xc8\x45\x0c\x76\xcc\xb1\x5b\xac\x79\x81\x4c\xfd\x38\xb9\xde\xec\x16\x6b\x7e\x44\x73\xec\x22\x66\xcd\x1a\x64\x13\xcb\xde\xdc\x04\xaf\x6d\x18\x57\x68\xca\x60\x30\xf0\xb9\x40\xa6\x04\x4d\x7c\x0d\x63\x1a\xf4\x1a\xae\x58\x36\x43\x2a\x42\x96\x70\x29\x7c\x9a\x65\x09\x0f\x28\x72\x29\xce\x02\x0f\xa4\x88\x78\x7c\x16\x64\x22\xe3\x98\xa9\x53\x90\x28\x65\xb2\x26\xfb\x3e\xcf\xfe\x03\x90\xa3\xe2\x22\xee\x1f\x43\xe5\x59\xf4\xf6\x3b\x7e\x20\xd7\x14\xd5\x33\x5c\xfa\x5c\xcb\x92\xe8\x17\x99\x9b\xbb\x9f\xf3\x58\xd0\xc4\x26\x2e\x21\x3d\xaa\xb4\xce\x4c\xf4\xa0\xba\x0c\x2b\xb1\x3a\x48\x15\xce\xa5\x21\xcc\xc2\x85\x8a\xcd\x9b\x93\x69\x4a\x45\xb8\x49\x2c\xeb\x4a\xce\x5a\x30\xb9\xec\x9c\xa9\x1e\x53\x76\x83\x58\x56\xa7\x2b\x15\x4e\xe6\xec\x8a\x05\xde\x6b\x2f\xc3\x14\x66\xbe\x4f\xd3\x2c\x19\x33\x54\xb5\xea\x54\xb3\x75\xa9\x6a\x2c\x34\x19\x24\x32\xa0\x89\x21\xe6\x09\x13\x01\xbb\x92\xd3\x98\xb5\x00\x55\xc1\xf4\xd7\xb6\x62\xab\x85\x68\x41\x54\x88\xc0\x09\xd2\x10\x2e\xbe\x12\x6b\x03\xa8\x8a\x73\xb8\xfa\x91\xc9\xca\x05\x1d\xbc\x95\x33\x2c\x32\xc7\x25\x96\x35\xd0\x1c\xab\x85\x98\x3f\x3b\x03\x53\x4a\x2a\xc3\xa3\x18\x16\x4a\x80\x2a\xc4\x84\x6c\xa0\x25\xd5\x54\xc0\x05\x47\xa7\x5a\x70\xac\xa4\xd7\x66\x2a\xe7\x39\x32\x81\x0b\x09\x8d\x73\xc7\xf5\x3a\x15\xe9\x87\x54\xb5\x9d\x0b\x4c\xf4\x1a\x60\x33\xd1\xb3\xf5\x43\xdf\xea\xe4\xc1\x9e\x17\x3d\xdb\x25\x83\x9a\xb9\x0e\x5f\x53\xfb\x3e\x8c\x3e\xbf\x7b\xf8\xfc\x51\xf9\xc5\x0f\xe5\xaf\x5f\x95\xdb\xf7\x89\x65\x6c\xe8\x2d\x0b\x8e\x73\xd5\xd0\x61\xa2\xe7\x56\xd8\xf2\xcf\x9d\xc3\x9d\xbd\x29\xec\x94\xc3\xab\x1f\x38\x06\x37\x7c\xba\x7f\xb8\xf3\xfb\x14\xae\x6a\x99\x0a\xb1\xaa\x47\x35\xec\xe0\xc9\x5e\x79\x6b\xbb\xfc\xe9\x8f\xd1\xcd\xbd\x40\x0a\x64\x7d\x1c\x7d\xf7\xf7\xf0\xe9\xfe\xab\xc4\xab\x2c\xd6\x59\xab\x25\x46\x43\xa6\x16\xb4\xce\x47\xc3\xa3\xb4\xa6\xa6\x03\xec\xc3\xc5\xba\x9b\xbd\x39\xc3\xdc\x80\x90\x22\x85\x94\x66\x57\x4d\x29\x26\x15\x39\xf1\xa9\x56\x66\x78\xfb\xb7\xf2\xd1\xee\xcb\x67\x5f\x9a\x08\x4d\x68\xff\x6c\xdd\x18\xee\xff\x7c\x70\xff\xd9\xc1\x93\x7b\xc3\x9b\xcf\xcb\xc7\xb7\x46\x0f\xb6\x86\xbf\x7c\x7a\xb8\xfb\x59\xf9\xf5\x8d\x83\x7b\x8f\x47\x5b\xbb\xc3\x17\xb7\x55\x16\x0c\x5f\x7c\x3f\x7c\xb8\xfd\xf2\xaf\x1f\x47\x5b\xdf\xc2\xe5\x4b\xef\x5f\x22\x64\x5c\x71\x1d\xc9\x24\xee\xaa\xfc\x47\xb6\xd0\x32\x7f\xb3\x5f\xde\xd9\xd3\x5b\xd2\xc1\x83\xbb\xe5\x9d\x87\xc4\x8a\xa5\xf1\x97\x31\xe0\xb4\x38\x4b\x88\x99\x31\xbf\xb6\x43\x3d\x20\xd6\x40\xdf\x2a\x4d\x05\x36\xe0\x63\x68\xcd\x82\x69\x63\x6f\x95\xd1\x70\x81\x27\xcc\xb1\x3d\x3f\x47\x8a\x3c\xf0\x4f\x34\x90\x87\x7d\xb4\x5d\x62\x45\x29\x7a\x6d\xc5\x05\x26\xc2\x31\x7b\x8e\xb7\xa8\x18\x13\x8e\xd1\xc9\xa9\x17\x70\x5d\x97\xe8\x4e\x3c\x15\x6e\x1b\x76\x9d\x2a\x50\x6c\xd9\xc7\xf1\x91\x63\x37\x01\x3e\xd0\x76\xd5\x4d\xad\x33\x6f\xf9\x7e\x65\xdf\xae\xcc\xb1\x35\x13\xfa\x70\x4d\x5d\x13\x76\x03\x4e\xcd\xbd\x2d\x15\x3a\xaf\x27\x5e\x61\xb8\x21\xd5\x7a\x6b\x4c\x3c\x93\x4f\x33\xf2\xcc\x5b\x64\x58\x2d\xbe\x24\x73\x74\xdc\xf3\x2e\x92\xe8\xf6\xd5\x5f\x78\x04\x75\xf3\x2c\x32\x9c\x17\x3d\xc7\x85\xb7\x66\xc1\xce\x94\x0c\x8b\xa0\x3a\x05\x60\x93\x8c\xb7\xb6\x37\x88\xb5\x41\xf5\xeb\x71\xb5\xce\xab\x18\xcd\xb8\x9f\x1b\x2a\x9f\x8b\x90\xf5\xbd\x2e\xa6\xc9\xb9\x74\x3c\xb3\x96\x6f\x5e\xeb\xff\x2a\x3c\x20\xd5\xe3\x93\x82\xa3\x36\x70\x4a\xd7\x99\x13\x74\xa9\x00\x99\x7b\x9d\xea\x04\x32\x38\x73\x1a\x79\x2b\x12\x79\x74\xdd\xd1\xf0\x86\x86\x2c\xeb\x03\x5b\x15\x19\x1a\xd4\xbb\x4d\x3d\x73\xc2\x1f\x33\x39\x74\xba\x05\x86\x72\x43\x40\x6d\x53\xcf\xf3\xc6\xa1\xd7\x67\xa4\x8e\x7f\xae\x50\x8a\x09\xbc\xcc\x53\xd6\x41\x35\x8e\xb1\x2e\xdd\xb2\x88\xe4\xc4\xe6\xac\xcf\xab\xff\x11\x06\x51\xf7\xbc\xe0\x09\x19\x90\x7f\x03\x00\x00\xff\xff\xfb\x8b\xe4\x1a\xa6\x08\x00\x00")

func internalTemplateCmdApiserverTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplateCmdApiserverTmpl,
		"internal/template/cmd/apiserver.tmpl",
	)
}

func internalTemplateCmdApiserverTmpl() (*asset, error) {
	bytes, err := internalTemplateCmdApiserverTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/template/cmd/apiserver.tmpl", size: 2214, mode: os.FileMode(420), modTime: time.Unix(1686125096, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _internalTemplateCmdCobraTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x91\x41\x8f\xd3\x30\x10\x85\xcf\x9e\x5f\x31\xe4\x80\x12\xd4\x26\xaa\xb8\x05\x7a\x58\xad\x7a\x43\xab\x6a\xab\x3d\x21\x0e\xae\x3d\x75\x2d\x62\x3b\x8c\xed\x65\x21\xca\x7f\x47\x4e\x4b\x05\x2b\x4e\x5c\xe7\x7b\xf3\xde\xd3\xcc\x28\xd5\x57\x69\x08\x95\xd3\x00\xd6\x8d\x81\x13\xd6\x20\x2a\x62\x0e\x1c\x2b\x10\xd5\x34\x61\xbb\xbf\xa8\x1e\xa4\x23\x9c\xe7\x4e\x39\xdd\xc9\xd1\x16\x6a\x6c\x3a\xe7\x63\xab\x82\xeb\xcc\xf0\x63\x7d\xce\xc7\xce\x84\xb5\x96\x5e\xd3\x60\x83\xef\x86\x60\x0c\xf1\x2b\x65\x1c\x4f\x9b\xf7\x9d\x0a\x47\x96\x85\x84\x58\x41\x03\xf0\x2c\x19\x39\x84\x74\xef\x34\x6e\xf1\xed\x82\xdb\xfb\xe0\x9c\xf4\x7a\x02\xf1\x14\xa9\xc7\x7f\xb4\xa9\x56\x20\x0e\xe7\xc0\xe9\x4a\x0f\xc4\xcf\xc4\x7f\x42\x3b\x90\x57\xf4\x14\xa5\xa1\x3e\x71\xa6\x15\x88\x4f\xc1\x9b\x1e\x2b\x99\xd3\x39\xb0\xfd\x49\x45\x77\xc7\x26\xf6\x78\xca\x5e\xd5\xca\x69\x7c\xf7\x57\x81\x15\x4a\x36\x11\x3f\x7f\x89\x89\xad\x37\x0d\x2e\x07\xc2\x09\x84\xb0\x27\x1c\xc8\xd7\x85\x37\xf8\x11\x37\xcb\x50\x30\xa5\xcc\xfe\x22\x8b\xed\x03\x7d\xaf\x2f\xa7\x68\x1f\x49\xd7\x15\xd3\xb7\x6c\x99\x22\xca\x84\x03\xc9\x98\x30\x78\x2a\x11\x55\xd3\x80\x10\x33\xdc\x0c\xbc\x1d\x40\xcc\x2b\x10\x7b\xe2\x68\x63\x22\x9f\xf6\x21\xa6\xc7\xec\x77\xff\x53\xf6\x95\xeb\x0c\x50\x3c\xd0\x7a\x9b\xea\x66\x02\x71\x7d\x40\x7b\xa7\xf5\xd5\xac\x96\xa3\x6d\x0f\x49\x72\x99\x37\xb7\x8d\xdd\x0b\xa9\x9c\x68\x59\xb2\xa7\x92\x80\xfd\xf6\xf7\xff\xda\x1b\xfd\xb0\x90\x37\xdb\x92\x58\xf2\x43\x6c\x77\x2f\x36\xd5\xeb\x4d\x03\x62\x86\x19\x7e\x05\x00\x00\xff\xff\xda\xd0\x84\x07\x81\x02\x00\x00")

func internalTemplateCmdCobraTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplateCmdCobraTmpl,
		"internal/template/cmd/cobra.tmpl",
	)
}

func internalTemplateCmdCobraTmpl() (*asset, error) {
	bytes, err := internalTemplateCmdCobraTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/template/cmd/cobra.tmpl", size: 641, mode: os.FileMode(420), modTime: time.Unix(1686118345, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _internalTemplateCmdRpcserverTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x94\xcf\x8f\x1b\x35\x14\xc7\xcf\xf6\x5f\xf1\x64\x89\xca\x83\xb2\x1e\x21\x6e\x2b\x7a\x40\xd1\xb6\xda\x4b\x15\x25\x94\x0b\x45\xc8\x99\xbc\x99\x58\xf5\x3c\x0f\x9e\x37\x69\xaa\x55\x8e\x48\x08\xe8\xad\xf4\x04\x07\x4e\x70\xe2\xc2\x09\x24\xfe\x9c\x6c\xc5\x7f\x81\x3c\x9e\x2c\x01\x54\x69\x7d\x70\x12\xbf\x8f\xbf\xef\xf9\xfd\x48\x67\xab\xe7\xb6\x41\xb0\x9d\x93\xd2\xb5\x5d\x88\x0c\x5a\x0a\x55\xb7\xac\xa4\x50\x37\x37\x60\x16\x19\x79\x62\x5b\x84\xc3\xa1\x5c\x87\xf0\x2e\x93\x23\xc6\x48\xd6\x97\x3d\xc6\x9d\xab\x30\x61\x8d\xe3\xed\xb0\x36\x55\x68\xcb\xc6\xbf\xbc\xd8\x0e\xeb\xb2\x09\x17\x1b\x4b\x1b\xf4\x2e\x50\x69\xbb\xce\xbb\xca\xb2\x0b\x74\x1f\xbc\x0a\x54\xbb\xe6\x3e\xa4\x0f\x4d\x83\xf1\x1d\x24\x87\xe0\xd7\x61\x5f\xf6\x1c\x1d\x35\xfb\xff\x50\x7d\x57\x7f\xf0\x61\x59\x85\x75\xb4\xc9\xe2\x42\xe9\xc2\xc0\xce\xa7\x1f\xa1\xcf\x7b\xd9\xbb\x86\xac\x57\xb2\x90\x72\x67\x63\x4a\x1a\xd2\x0e\xc6\x95\x55\xa5\x58\xb1\x8d\x3c\x6f\x37\xf0\x10\x1e\x8c\x6a\x66\x1e\xda\xd6\xd2\xe6\x46\x0a\xf1\xb4\xc7\x4b\xb8\x5b\x2a\xa5\x0c\xa3\x9a\x49\x21\x56\xdb\x10\xf9\xce\xa6\x46\x15\x58\x2e\xe6\x70\xc6\x5c\xed\x6d\xdb\xf9\x93\xc2\x58\x8c\xd5\x68\x9d\x6a\x31\xb1\x70\x81\xe0\x43\x65\x7d\x16\x76\x1e\xa9\xc2\xa7\xbd\x6d\xf0\x12\x38\x0e\x98\x4e\x17\x11\x97\x03\x5d\x42\x3d\x50\xa5\xab\x76\x03\xef\xff\x2b\xd6\x19\xd8\xd8\xf4\xf0\xd9\xe7\xf9\x55\x05\xa4\xe0\x45\x8f\x3c\x74\xba\x90\x42\x1c\x92\xc6\x72\xa0\xab\xfb\x2b\x60\x8c\x21\x66\x9d\x88\x3c\x44\x82\x38\xd0\x9d\xd8\x21\xa5\x34\x49\x81\x23\xc7\x7a\x74\x78\xca\xa4\x59\x60\xec\x5d\xcf\x48\xfc\xc8\xdb\xa6\xd7\x85\x59\x8d\xa2\x9f\xda\xb8\xd0\x0f\x90\x76\x33\x50\x48\x3b\x95\x3e\xd2\x36\x3d\x1e\xd4\x15\xed\x54\x21\x0f\x93\xf2\x14\x7e\x92\x2e\x4b\xf8\xeb\xab\x57\x6f\xff\xfc\xf5\xf8\xf5\x8f\xc7\x9f\xbf\x3d\x7e\xf7\x46\x8a\xdc\x65\xe6\x9a\x1c\xcf\xc7\xaf\x1a\x69\x57\x8c\xec\xf1\x8f\xd7\x6f\x5f\xff\x72\xc6\x9e\x35\xf0\x78\x41\x4f\xdc\x89\xb8\xfd\xe1\xd5\xf1\x9b\x9f\x6e\xdf\xfc\x7e\xfb\xdb\xf7\x52\xa4\xe9\x39\x61\xa7\x60\xc6\xc7\xff\x93\x94\xf3\xcb\xb1\xab\xa0\x0d\x1b\xf4\x52\x34\x21\xe7\x37\x17\xe0\xdc\xeb\xb2\xab\x72\xed\x35\xe1\x0b\x3d\xcd\x5e\x3a\xfd\xb8\x73\x45\x21\xc5\x21\x85\x54\x05\x4a\x59\x9b\xc1\x17\x70\xf9\x10\x72\x3f\x9b\x25\xda\xcd\x23\xe7\x51\x2b\x53\xf6\x6c\xd9\x55\xe5\xff\x3a\xc9\xf0\x9e\x55\x21\x45\xdd\xb2\x59\x44\x47\xec\x49\xe7\xd9\x32\x8f\x23\x22\xe9\x5c\x55\x3d\x39\x28\x92\xc7\x2f\x07\xc7\xc9\x4d\x6b\x9f\xa3\xae\xb6\x96\x20\xf4\x66\x35\x0e\x4c\x21\x45\x9e\x1c\xf3\x24\xb0\xab\x5f\xea\xc4\xce\x92\xfd\x3a\xfd\x7b\xc4\xa1\xe3\x42\x8a\x8f\x2e\xd2\xf1\x99\xd3\x5a\xab\xf7\x7a\x58\x6d\x07\xde\x84\x17\x04\x39\x42\x30\xc6\xc0\xb3\xf8\x8c\xd4\x6c\x9a\xb9\xbd\x79\x8c\x3c\x1f\x62\x44\xe2\x4f\x5c\x8b\x2b\x8e\xba\x28\x64\x9a\x91\x29\xe6\x6b\xaa\x83\x56\xd3\x7d\xdc\x3b\x76\xd4\xa8\x4c\x4c\xbd\x48\xce\xcb\x83\xfc\x3b\x00\x00\xff\xff\xc1\x6a\xdc\x2a\x19\x05\x00\x00")

func internalTemplateCmdRpcserverTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplateCmdRpcserverTmpl,
		"internal/template/cmd/rpcserver.tmpl",
	)
}

func internalTemplateCmdRpcserverTmpl() (*asset, error) {
	bytes, err := internalTemplateCmdRpcserverTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/template/cmd/rpcserver.tmpl", size: 1305, mode: os.FileMode(420), modTime: time.Unix(1686135118, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _internalTemplateConfigConfigTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x54\x41\x6f\xd3\x4a\x10\xbe\xfb\x57\x8c\xfc\xce\x2f\xb2\xe3\xbe\x24\xdd\x5b\x5f\x11\x02\xa9\x40\xd5\x46\xe2\x80\x38\x6c\xbd\x93\x64\xc5\xc6\x6b\x76\x37\x09\xc5\xb2\x44\x05\x12\x52\x25\x04\x48\x70\x80\x1b\x12\x2d\xa7\x2a\x1c\x38\x71\xe1\xcf\x40\x9a\xfe\x0b\xb4\x6b\x27\x4e\x8b\xd3\x13\xa7\xcc\x7e\x33\x9e\x6f\xbf\x99\x2f\x9b\x65\xc0\x7b\xd0\xe8\x4a\x29\x74\x63\x47\xf6\xfb\xa8\x20\xcf\x3d\xe1\x22\xe2\x01\xfc\xb3\x83\x63\x14\x10\xc0\xf9\xb7\x2f\xb3\x67\x27\xe7\x1f\x5f\x40\x38\x3f\x3b\x9d\x1d\x9f\x40\xf3\xe2\xe5\xab\xf9\xe9\x91\x85\xa2\x8b\x77\x1f\xe6\xd3\x29\x6c\xcc\xcf\x4e\x7f\xbd\x3d\x86\xff\x66\xaf\xdf\x9c\x7f\xfe\x0e\xad\x9f\x3f\x3e\xcd\x8e\xa6\xd0\x9e\x7f\x7d\x3e\x9f\xbe\xf7\x00\x62\x99\x68\x29\x70\x7f\x20\x27\x04\x8c\x1a\x61\x85\x39\x26\x02\xd0\xf6\x00\x7a\x5c\xe0\x7d\xc5\x0d\x12\x80\x1e\x15\x1a\x4b\x6c\xb5\x66\x38\x12\x86\xdf\xac\x0a\x17\x75\x4b\xbc\x2c\x6e\x7b\x59\x06\x98\x30\x2b\xcc\xbb\x24\xf8\x96\x31\x29\xe4\xf9\xc0\x98\x74\x1f\xd5\xb8\x50\x9c\x4a\x65\x08\x74\x82\x4e\x60\x0f\xa9\x92\x3d\x77\xea\x5c\xee\x52\x35\xd9\x4b\xe3\x6d\xc1\x31\x31\x79\xee\xa9\x45\x6c\x1b\xc5\x2e\xba\x4b\x87\x48\xc0\xcf\x32\x68\x14\x1c\x16\x80\x3c\xf7\x3d\x80\x03\xaa\x71\x97\x9a\x41\x99\xdf\x4a\xd3\x32\xa1\xb0\xcf\xb5\x41\xb5\x2b\x46\x7d\x9e\x10\xf0\xd1\xc4\x6c\x35\x51\xb4\xd2\x96\x06\xe0\x5f\xf0\xc3\x66\xbb\x11\x34\x82\x46\x48\x9a\x51\x7b\xd3\x56\xf6\x28\x17\x7b\x68\xd4\xe1\x1d\xc9\xec\x14\x22\xc7\x27\x68\x12\x63\x89\x34\x9d\x58\x29\xf6\xf9\x53\x24\x10\x7a\x59\x86\x09\xab\x51\x57\x70\x15\xea\xaa\x31\xe9\xa5\x98\x75\xea\xfe\x82\x88\xb5\x03\xa2\x8c\x29\x02\xbe\x5f\xed\xab\xb3\xb9\x59\xed\x2b\x74\xc7\x5a\x41\x37\xfe\xcf\x73\x8f\x1d\x58\x52\x76\xd0\x3d\x4c\xed\xf5\x87\x87\xfa\xb1\xb0\xbd\x86\xf4\xc9\xbd\x14\x93\x6d\x99\x24\x04\x9a\x41\x81\xdc\x66\x02\x0b\x64\xa3\x02\xba\xdc\x0a\x0f\x83\xb2\x66\x87\xf7\x4a\x28\x6a\x39\x4c\x14\xde\xb3\x5f\x68\x21\x27\xdd\x81\x42\x3d\x90\x82\x11\xf0\xc3\x20\x18\xea\x82\xcd\x4e\xa1\x90\x3f\xd2\x68\x05\x29\x29\x8d\xef\x80\x94\x6a\x3d\x91\xca\x7e\xb0\x08\x8b\xc4\x40\x6a\x43\x56\x86\x55\x96\xbb\x29\xf8\x51\x14\xb4\x0a\x80\x51\x43\xed\xf8\x08\xf8\x8b\xd0\x77\x97\xa1\x63\x5c\x4c\xfc\x0f\xd2\x6b\x68\xd7\x11\xd7\x50\xaf\x21\xaf\xf7\x17\x32\xae\xad\xb7\xec\x2f\x71\xe6\x60\x5c\x97\x7b\xa1\x42\x26\xee\xd6\x09\x9a\x89\x54\x8f\x56\xd9\x49\xab\xf4\x88\x36\x54\x99\x2d\xe7\x87\x07\x57\xd3\x0f\xad\x55\x62\xc3\xc7\xcb\x5d\x71\x26\x96\x31\x1d\x39\x67\xf9\xc5\x13\x94\xd8\xfd\xc9\x91\x59\xdd\x90\x42\xca\x6a\xe0\x89\x7d\x70\x6a\x70\x5e\x1a\xe3\x12\x5c\x2b\xbb\xab\x68\x8c\x79\xee\x19\xfb\xeb\x2c\x20\x53\x4c\x1c\xba\x7c\x16\x5d\xee\xba\x3f\xd8\xd5\x85\x90\x56\x27\x0a\x2b\xc2\xdf\x01\x00\x00\xff\xff\x0d\xd9\xda\x2f\xdc\x05\x00\x00")

func internalTemplateConfigConfigTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplateConfigConfigTmpl,
		"internal/template/config/config.tmpl",
	)
}

func internalTemplateConfigConfigTmpl() (*asset, error) {
	bytes, err := internalTemplateConfigConfigTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/template/config/config.tmpl", size: 1500, mode: os.FileMode(420), modTime: time.Unix(1687686456, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _internalTemplateGlobalGlobalTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x48\xcf\xc9\x4f\x4a\xcc\xe1\x02\x04\x00\x00\xff\xff\x32\x37\xa5\xbc\x0f\x00\x00\x00")

func internalTemplateGlobalGlobalTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplateGlobalGlobalTmpl,
		"internal/template/global/global.tmpl",
	)
}

func internalTemplateGlobalGlobalTmpl() (*asset, error) {
	bytes, err := internalTemplateGlobalGlobalTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/template/global/global.tmpl", size: 15, mode: os.FileMode(420), modTime: time.Unix(1686108961, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _internalTemplateInternalRouteTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8f\x31\x6b\x1b\x31\x14\x80\x67\xbd\x5f\xa1\x6a\xba\x33\xb5\xb4\x17\x8c\x97\x96\xda\x50\x30\xd8\x85\xce\xf2\x9d\x2c\x8b\x9e\x25\x21\xeb\x5c\x4a\x31\x74\x68\x3d\xb5\x24\x10\x67\x72\xc6\x18\xbc\x04\x13\x3c\x24\x84\x84\xfc\x1a\xc5\xce\xbf\x08\xe7\x73\x82\x33\xc5\xa3\xc4\xf7\xde\xfb\x3e\xcb\x93\xef\x5c\x0a\xec\x4c\xee\x05\x80\x1a\x58\xe3\x3c\x8e\x00\x11\xa9\x7c\x3f\xef\xd2\xc4\x0c\x98\xcc\x7e\x56\xfb\x79\x97\x49\x53\x4d\xb9\x4e\x45\xa6\x8c\x66\xdc\xda\x4c\x25\xdc\x2b\xa3\xc9\x01\x78\x62\x74\x4f\xc9\x43\xc8\xa1\x70\x23\xe1\x58\xdf\x7b\x4b\x00\x15\x62\x4a\xcb\xce\x0f\x2e\xa5\x70\xf8\xed\xe9\x12\x24\x10\x03\xf4\x72\x9d\xe0\xa6\x56\xbe\x5d\xd4\x45\x31\xfe\x05\xa8\xcb\x87\x62\xfb\x74\xf8\x43\x0d\xef\x45\xd0\x86\xf7\xb6\xb3\xbd\x1d\xc5\xb4\x44\xa2\x18\x90\xea\xe1\xd2\x9d\x7e\x16\xfe\x93\x1e\x45\x31\x7e\x57\xc3\xc4\x3a\x93\xe6\xc9\xb6\xbe\x58\x8b\x18\xc3\x0f\xab\x45\x98\xfc\xdb\x09\x00\xda\x3b\x55\x8c\x46\xe4\xd9\x8d\x55\xc8\x7b\xfc\xba\x8b\x7e\x73\xdc\x36\xb8\x4e\x33\xe1\x62\x40\xa8\x88\xa7\x5f\x8c\x6c\x4a\x6d\x9c\x68\x8b\x61\x9e\xf9\x88\xd0\x4a\xfd\x65\x07\xad\xd4\x49\x8c\x19\xc3\xe1\xfe\x6e\x7d\x3a\xdf\x7d\x87\x93\xff\xe1\x66\x1a\x7e\xdf\x02\x1a\x03\x14\x52\xe1\x68\x19\xce\x16\x9b\xe5\x3c\x9c\xff\x29\x05\xd7\xb3\xeb\xf0\x77\x15\x2e\x8e\x1f\x67\x93\xcd\xd5\x72\x3d\xbd\xc4\x5f\x5b\x1f\x5b\x00\xc8\x09\x9f\x3b\x0d\x63\x78\x0a\x00\x00\xff\xff\xb1\xcf\x6a\xbd\x16\x02\x00\x00")

func internalTemplateInternalRouteTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplateInternalRouteTmpl,
		"internal/template/internal/route.tmpl",
	)
}

func internalTemplateInternalRouteTmpl() (*asset, error) {
	bytes, err := internalTemplateInternalRouteTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/template/internal/route.tmpl", size: 534, mode: os.FileMode(420), modTime: time.Unix(1686108871, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _internalTemplateInternalRpcapiTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x28\x4e\x2d\x2a\xcb\x4c\x4e\xe5\xe2\x2a\xa9\x2c\x48\x55\x08\x2a\x48\x76\x2c\xc8\x54\x28\x2e\x29\x2a\x4d\x2e\x51\xa8\xe6\xe2\xaa\xe5\x02\x04\x00\x00\xff\xff\xbd\xf9\x40\x98\x29\x00\x00\x00")

func internalTemplateInternalRpcapiTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplateInternalRpcapiTmpl,
		"internal/template/internal/rpcapi.tmpl",
	)
}

func internalTemplateInternalRpcapiTmpl() (*asset, error) {
	bytes, err := internalTemplateInternalRpcapiTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/template/internal/rpcapi.tmpl", size: 41, mode: os.FileMode(420), modTime: time.Unix(1686108931, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _internalTemplateMainTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\xc8\x4d\xcc\xcc\xe3\xe2\xca\xcc\x2d\xc8\x2f\x2a\x51\x50\xaa\xae\x56\xd0\x0b\x80\xc8\xf8\x25\xe6\xa6\x2a\xd4\xd6\xea\x27\xe7\xa6\x28\x71\x71\xa5\x95\xe6\x25\x83\xd5\x6a\x68\x2a\x54\x73\x71\x26\xe7\xa6\xe8\xb9\x56\xa4\x26\x97\x96\xa4\x6a\x68\x72\xd5\x72\x01\x02\x00\x00\xff\xff\x12\xa7\xcc\x37\x4e\x00\x00\x00")

func internalTemplateMainTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplateMainTmpl,
		"internal/template/main.tmpl",
	)
}

func internalTemplateMainTmpl() (*asset, error) {
	bytes, err := internalTemplateMainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/template/main.tmpl", size: 78, mode: os.FileMode(420), modTime: time.Unix(1686118271, 0)}
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
	"internal/template/boot/boot.tmpl":       internalTemplateBootBootTmpl,
	"internal/template/cmd/apiserver.tmpl":   internalTemplateCmdApiserverTmpl,
	"internal/template/cmd/cobra.tmpl":       internalTemplateCmdCobraTmpl,
	"internal/template/cmd/rpcserver.tmpl":   internalTemplateCmdRpcserverTmpl,
	"internal/template/config/config.tmpl":   internalTemplateConfigConfigTmpl,
	"internal/template/global/global.tmpl":   internalTemplateGlobalGlobalTmpl,
	"internal/template/internal/route.tmpl":  internalTemplateInternalRouteTmpl,
	"internal/template/internal/rpcapi.tmpl": internalTemplateInternalRpcapiTmpl,
	"internal/template/main.tmpl":            internalTemplateMainTmpl,
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
	"internal": &bintree{nil, map[string]*bintree{
		"template": &bintree{nil, map[string]*bintree{
			"boot": &bintree{nil, map[string]*bintree{
				"boot.tmpl": &bintree{internalTemplateBootBootTmpl, map[string]*bintree{}},
			}},
			"cmd": &bintree{nil, map[string]*bintree{
				"apiserver.tmpl": &bintree{internalTemplateCmdApiserverTmpl, map[string]*bintree{}},
				"cobra.tmpl":     &bintree{internalTemplateCmdCobraTmpl, map[string]*bintree{}},
				"rpcserver.tmpl": &bintree{internalTemplateCmdRpcserverTmpl, map[string]*bintree{}},
			}},
			"config": &bintree{nil, map[string]*bintree{
				"config.tmpl": &bintree{internalTemplateConfigConfigTmpl, map[string]*bintree{}},
			}},
			"global": &bintree{nil, map[string]*bintree{
				"global.tmpl": &bintree{internalTemplateGlobalGlobalTmpl, map[string]*bintree{}},
			}},
			"internal": &bintree{nil, map[string]*bintree{
				"route.tmpl":  &bintree{internalTemplateInternalRouteTmpl, map[string]*bintree{}},
				"rpcapi.tmpl": &bintree{internalTemplateInternalRpcapiTmpl, map[string]*bintree{}},
			}},
			"main.tmpl": &bintree{internalTemplateMainTmpl, map[string]*bintree{}},
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
