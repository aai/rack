// Code generated by go-bindata.
// sources:
// models/templates/app.tmpl
// DO NOT EDIT!

package models

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

var _templatesAppTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x3c\x6b\x73\xdb\x38\x92\xdf\xf3\x2b\x50\xac\xb9\x72\xb2\x25\xcb\x8f\xa9\xdb\xbb\xd3\xdd\x5e\x95\x23\x2b\x89\x76\xed\x58\x27\x39\xd9\xba\x4d\x52\x5b\x34\x05\x59\x1c\x53\x24\x87\x0f\xc7\x1e\x97\xfe\xfb\x76\x03\x20\xf1\xa6\xe8\xe7\xee\xc4\x16\xd9\x68\x34\x1a\xfd\x46\x43\x0f\x0f\x64\x49\x57\x71\x4a\x49\x10\xe6\x79\x40\xb6\xdb\x37\x84\xc0\xc3\x5f\xe0\x13\x19\xfd\x85\x0c\x4f\xe0\x77\xfb\x70\x13\xa6\xf1\x8a\x96\x15\x7b\x73\xde\x7c\xe0\xaf\xe1\x3f\x42\x82\x93\xbf\x2f\x2e\xe9\x26\x4f\xc2\x8a\x7e\xc8\x8a\x4d\x58\x7d\xa5\x45\x19\x67\x69\x40\x46\x24\x38\x3e\x3c\x3a\xdc\x3f\xfc\x2f\xf8\x7f\x30\xe0\xe0\xe3\x2c\x5d\xc6\x15\xbc\x2f\x83\x91\x40\xc1\x66\xaa\x04\x0e\x12\x5c\x85\x49\x98\x46\xb4\xd8\x8f\x24\xa8\x39\xb7\x35\x28\x2f\xb2\x88\x96\xe5\xa3\xc6\x14\xf4\x3a\x2e\xab\xe2\x7e\xd7\xa0\x60\x9a\x56\xb4\x48\xc3\x04\x29\x26\xc1\x87\x74\x34\x9a\xfc\x5e\x87\x09\xae\xe0\x1b\x3e\x99\xd3\x15\xfc\x29\xc1\xc8\x76\x40\x82\xff\xa7\x80\xed\x07\xfc\xd9\x60\x99\x15\xf1\x2d\xcc\xbb\x03\x49\x03\xe5\xc6\xf1\x1e\x58\x73\xb3\xa0\x51\x5d\xc4\xd5\xfd\xc7\x22\xab\x73\x64\xf3\x83\x8a\x0e\x3e\x7f\x7b\x60\xd8\x70\x03\x74\x58\xc4\x19\xfc\xe0\xeb\x12\x48\x83\x59\x58\x84\x1b\x0a\x94\xb3\xa1\xdd\x3b\x92\x23\xec\x23\x76\xc3\x09\xdf\xac\x65\x9c\xd4\x25\x4c\xab\x88\x01\x3c\xbc\xbc\xcf\x29\x27\xbc\x2a\xe2\xf4\x3a\x18\xc8\x57\xa7\x74\x15\xd6\x49\xc5\xde\xea\xcf\xcb\xa8\x88\xf3\xaa\x91\xb9\x40\xbc\x92\x5c\x3b\xa5\x79\x92\xdd\x6f\x68\x5a\x9d\x87\x77\xf1\xa6\xde\x38\xe6\x84\x81\x9f\xeb\xcd\x15\xd0\xe3\x98\x92\x49\xf2\xa1\x6f\x52\x78\x2b\xf0\x92\x9c\x16\x11\x4c\x13\x5e\x53\x92\xad\x88\x60\x03\x2d\x49\x95\x91\x1b\x4a\x73\x52\xd4\x69\x0a\xcb\x22\x3f\xd7\x71\x42\x41\x0f\x91\x2e\x5c\x66\x17\xc9\x71\xfa\x44\x92\x8f\xba\x49\xe6\x78\x5f\x8e\xe4\x49\x7a\x1b\x17\x59\x8a\x34\xbb\x89\xf5\x6f\x69\xc7\x8e\x3a\x37\x54\x55\xc8\x7e\xf3\x68\x08\x2f\xd2\xe4\x9e\x84\x49\x92\xfd\x24\x61\x84\xcb\xc5\xc5\x56\xeb\xb8\x24\x68\x03\x57\x45\xb6\x21\x71\x5a\xc6\x4b\x0a\x0f\x29\xf9\x3a\x1b\x7b\x68\xfe\x9c\xa9\x2f\x4e\x10\x21\x5d\x7e\x0d\x93\x9a\x72\xad\x66\xfa\x3b\x60\x70\xe4\x87\xb5\x88\xbf\xd1\xfb\xd7\xe6\x93\x62\x72\x9e\xc0\xa6\x2f\x25\x25\x8b\xfa\x2a\xa5\x55\x29\x10\x21\x9f\xca\x9c\x46\xf1\xea\x1e\xd9\xb2\xcf\x78\x94\x64\xe1\x92\x34\x26\x82\xd0\x74\x99\x67\x71\x5a\x95\xaf\xc2\xb3\x39\x4d\x68\x58\xba\x16\xf4\xd2\x36\x63\x4e\xf3\xac\x8c\xab\xac\x70\x6d\xd2\xf3\x26\x5b\x64\x35\xa8\x1c\x89\x32\x60\x5e\x21\xa7\xb1\x48\xd0\x6d\xf7\x4b\x53\x71\x09\xa2\x7d\xa6\x6d\x5d\x29\xe6\x23\xd7\x38\x21\x59\x65\x45\xab\x14\x0e\xe2\xb8\x60\x78\xc8\x3a\x03\xc7\xfa\x3f\x10\x19\x80\x57\x1a\x1f\x8f\x46\x1c\x78\x34\x9a\x2e\xff\xf7\x29\xa4\x82\xa8\x91\x92\xcf\xd7\x8f\x2a\xbf\xdc\xbf\x0e\x71\xb9\x50\x8f\x7e\x44\x36\x01\x92\x46\x9d\xa1\x7b\x6f\xe7\x93\xff\xfb\x32\x9d\x4f\x4e\xdf\x91\xb3\x70\x73\xb5\x0c\xc9\x18\xbc\x65\xb6\xb9\xcc\xf2\x38\x22\x9f\xc2\x74\x99\xc0\x8e\x09\x75\x20\x0d\x46\x85\x4c\x30\xef\x67\x34\xbd\xae\xd6\x8c\xc8\x23\xf5\x95\x61\x00\x6c\xfa\xc0\xe0\xb9\x39\x27\x99\x06\x30\xc8\xb1\xa7\x32\x6c\x07\x83\x66\xe3\xf1\xf4\x74\xfe\xe2\x22\x8f\x33\x23\x62\xf7\xf4\x5a\x54\x74\x0e\x6f\x60\x16\x55\xbe\x83\x59\x56\x54\xb3\x22\xab\xb2\x28\x33\x3c\xcf\xba\xaa\x72\x1e\xd7\xa1\x6c\xd1\x94\x16\x0a\x5c\xf0\xe9\xf2\x72\x86\x26\x6d\x9a\x96\x15\x6a\x9a\xeb\x1d\xd3\x75\xea\x83\x58\x04\x92\x3b\x62\xba\xb2\x7b\xbe\xc5\xb3\x27\xd4\x66\xac\xa2\x8e\xf5\x5d\x8e\xbd\xcb\x13\xaf\xfc\x93\x2d\x16\x67\xe6\x54\x49\xc7\xd2\x10\xfc\x79\x53\x91\xad\x73\xbf\xe7\xb4\x64\x56\x59\xdb\x70\x45\xe5\xe6\x59\xe2\x71\xa3\x4c\x27\xa6\x27\xe7\xa3\x11\x83\x51\x56\x02\x93\x43\x70\x55\xc5\x54\xb7\x92\xe8\xf6\xca\xb2\xde\x50\x84\x9f\x65\x49\x1c\xdd\x9f\x66\x51\x6d\xc5\x4d\x5c\x15\x5a\x5b\x81\xb9\xd4\xf1\x3e\xa4\x53\x47\xff\xa1\x4c\xc2\x80\x16\x15\x18\x1f\x31\xfe\x9b\xf6\x8a\x18\xf8\x18\xf8\x64\xb5\xa2\x11\x73\xc6\xcc\xfd\x1a\xd8\x04\xe9\x71\x1a\xc5\x79\x93\xf2\x2c\x68\x71\x1b\x47\x94\x3b\xe8\x84\xd9\xa3\x61\xb8\x09\xff\xc8\xd2\xf0\x67\x39\x8c\xb2\x8d\x96\xa5\xa8\x0b\x8d\x84\x41\x83\x71\x65\x55\x8e\xe4\xc2\xa5\x77\x6f\x7e\xb6\xda\x67\xf5\xad\x86\x19\x12\x16\x30\x6a\x40\xfc\x01\xa4\x6d\xb7\xd9\xdd\x41\xa0\xbf\x45\x86\x72\x96\xeb\xac\x30\x19\xc1\x21\xef\x3f\x43\xf2\xc3\x58\xb1\xdc\x40\x40\x0c\xe9\x60\x08\xce\xd8\x62\x49\xb0\x63\x9f\x18\x4c\x9f\xbd\x62\x80\xda\x7e\x21\x7f\xad\x1d\x51\x38\x17\xfc\x09\x3f\x36\xf2\xc9\x1f\x90\xed\x0e\xee\xa9\x9f\x24\xe4\xd6\xb2\xb4\x8a\x84\x77\x48\x37\xf7\x40\xa3\xd1\x87\x3a\xe5\x54\xf5\x12\xf2\x31\xc4\x37\xb6\x40\x2f\x7e\x7d\x5f\x47\x37\xb4\x92\x69\xf0\x5f\x21\x5c\xe4\x12\xb2\x0f\x2b\x85\x5f\x7c\x5f\xe1\x6f\x99\x15\x33\x32\xe6\x90\xaf\xa3\x25\x87\xc5\xdb\xe2\x06\x88\x45\x40\x6d\x62\xe5\x48\x0b\xee\x2a\x0f\x34\xb4\x6d\xa9\x02\x13\xe3\x03\x2e\xd8\x07\x2b\x56\xc5\x80\xc7\xc3\x3f\xe2\x3c\xe0\x73\x79\x85\x51\x78\x62\x44\x16\xa7\x4b\x7a\x37\xa4\x77\x22\x35\xd1\xc0\xce\xe9\x06\x42\xbc\x45\xfc\x07\x63\xea\xd1\xf1\x7f\xea\xaf\x1b\xeb\xc2\x49\xff\x48\xab\x93\x8a\xcb\x86\x65\x82\x50\x32\x8a\xd4\x52\xb7\x60\x5e\xa7\x55\xcc\x25\x39\x05\xbe\xff\x56\xea\x13\x5c\xc2\xbb\xac\x66\x12\xf6\x2b\x64\x85\x7e\x89\x70\xe7\xfd\x45\x6b\x1d\xc9\xd0\x93\xf2\x47\x90\xfa\xfd\x96\x5d\xf5\x01\x6d\xaa\x03\x2a\x68\xcf\x82\x42\xc9\x0d\x51\x07\xf2\xb6\xa8\xe3\xc3\xee\x1a\xd4\x44\xbe\x81\x07\x69\x59\xf1\x92\x8c\xee\x33\x2e\xea\x2a\xaf\xab\xdd\x75\xac\x4c\xc0\x91\x61\xf7\xe2\x24\x5c\xdf\xc2\x95\x7b\x84\xcc\x1f\xaa\xca\x88\x61\xd0\x4a\x61\xae\xc5\x85\x4d\x68\x41\x0b\x67\xfa\xc6\x37\xf8\x1f\xcc\x0c\x39\x1d\xc3\xab\x94\x0e\x5d\xf5\xb6\xa6\x68\x58\x84\xe9\x35\x25\xbf\xdc\xb0\x9a\xe1\x24\x05\x42\xd1\xc8\x96\xcd\x62\x82\x49\x1a\x5e\x25\x74\x09\x90\x75\x0e\xb6\x03\x21\xb7\x5b\x29\xfe\x9f\x33\x26\xfb\xce\x22\x19\x3e\x59\x80\x26\x47\x5c\x3d\x0e\x55\x65\xd6\xf1\x7d\x68\xb4\x98\xdb\x0b\x54\xf0\xfd\x23\xa6\x37\x42\x75\xe4\xba\xba\x57\xd8\xd4\xb0\x8c\xd5\x51\xdf\xea\x24\x19\x54\x23\x43\x89\x2b\x1a\xe3\x3a\xce\x36\x9b\xf0\x94\x26\xf1\x26\xae\xe8\x12\xe3\x9d\x40\x29\x00\xc9\x3a\xce\xe0\x70\x70\xfc\xef\x7f\x56\xdf\x69\xb9\x02\x2f\x02\x59\xd5\x9b\xa2\x4e\x07\x64\x3c\xfb\x42\xea\x34\xae\xf8\x13\x8a\xfa\x43\x07\x04\x8c\x16\x39\x7f\x8f\x23\xe6\x27\xe7\xca\x9b\x40\xca\x77\x5f\xf6\xb4\x22\xc8\xd6\x1f\x9c\x65\xd7\x7a\xba\xea\x90\xb7\x16\x86\x4b\xd8\x60\xc7\x0c\x8a\x22\xfb\xe6\xd0\xbd\x55\x76\x5d\xb2\x7f\x39\x50\x9f\x29\xa4\x59\xe9\x55\xf8\xf6\x14\xcb\xe3\x95\x1c\x36\xfc\x14\x42\x16\xda\xec\x86\x90\x0d\x43\x7a\x24\xb0\x88\xaf\x4a\xa5\xe6\xac\x88\xd1\x10\x05\x0c\x5e\x4d\xc6\x8b\xcb\xb0\xbc\x39\x45\xe2\xe3\xca\x91\x41\xe6\xb0\xc4\xf2\x82\xb9\x3d\xcd\xb3\x0f\xda\x08\x8e\xf9\x90\x1f\x8e\x5c\x90\x83\x63\x72\x67\xce\xa1\x00\x2b\x01\xce\xd1\xf0\xb0\x5f\x14\x20\x26\xbe\xcc\x6e\x68\xba\xd3\xc5\x79\xdd\x9b\x88\xd2\x3c\x11\x83\x11\x27\x40\x78\x15\xdd\xb0\x11\x4c\xed\x71\xbb\x5a\x1e\x06\x76\xec\xa0\x16\x95\x5a\x44\xcd\x33\x03\xd4\xa8\x71\xb6\xe0\xea\x73\x63\x48\x1b\x95\x08\x50\xfc\x6c\x80\x20\xc7\x7b\x04\xac\x4d\xa8\xaa\x2f\xc8\x0a\x55\xa7\x9b\xf0\x5a\x81\x63\x1f\x5d\x80\x0f\x0f\x28\xb0\x74\xc8\xac\x50\xba\x1c\x9e\x14\x45\x78\xbf\xdd\xda\xe1\xaa\x00\x70\x24\x17\x1c\x4d\x23\xd4\x2c\x00\x1a\x00\xca\x84\x05\xb7\x4c\xc4\x77\xa3\x57\x89\x61\x18\xb6\xdb\xc1\xc3\x03\x4d\x4a\xba\xdd\xc2\xef\x74\xe9\x1d\x03\x0b\x6c\xe6\x82\xe5\x39\x49\x73\x0f\xff\x61\xb3\x02\xe7\x43\x05\x4e\xa9\x4a\x33\x2f\x35\x40\x08\xd9\xcd\x16\x60\xc1\x2d\x5a\x39\xc7\xd0\xad\x95\x15\xb9\x89\x0a\xc6\x79\x2d\x05\x5c\x71\x71\x47\x6e\x17\xd7\xee\xbf\xe5\xe7\x4c\xc4\x3c\xf4\x74\xe2\x3e\x7e\x2e\x6e\x5f\xc9\xbf\x05\x38\x99\xcd\x1a\x49\x44\x53\xe9\x15\x5a\xd4\xc2\x93\xf1\xdf\x04\x2c\x4d\x6f\xc5\x67\x0f\x2c\xa8\xf9\x3f\xe7\x93\x8f\xd3\x8b\xcf\xea\x08\xe5\xa9\x7b\x9c\x12\x9b\xd0\x7b\x10\x54\xbe\x69\x5c\x4c\x95\xa5\x10\xc7\x6e\x33\xf9\x44\xe1\xe0\x63\x82\xc0\x05\x44\xb8\xdd\x46\xec\x22\xa2\x69\x05\x83\xff\xb2\xa5\xc1\x2f\xa4\xd2\x63\xf5\x5e\xc6\xf0\x2c\x4e\x6f\xbe\x86\x45\xe9\x26\xce\xa2\xad\x93\x2a\xdf\xec\xc1\xd9\xc5\xc7\x7f\x7e\x9c\x5f\x7c\x99\xf9\x9c\xba\xab\x9e\x30\xbf\x18\x4f\x16\x0b\xdb\x7a\x99\x59\xac\x25\x62\x5f\xb3\x04\x52\x6e\xdb\x3a\xea\x8c\xa0\xc3\xf3\x0c\x32\x20\x8c\x2b\xc5\x00\x37\x0b\xb8\x97\xa6\xbf\x93\xe1\xa7\x0c\xdc\x79\x70\x70\x1b\x16\x07\x10\x28\x1d\x2c\x33\x48\x4a\x8b\x61\x09\xbf\x7c\x5b\x8b\xa4\xb3\x61\xdb\xed\x08\xfe\x1a\x67\x30\x1f\x44\x11\x85\x53\xd4\x38\x07\xd1\xa8\x78\x90\x79\xd2\xd4\x83\x5b\x4e\xfe\x81\x9d\xfe\x1a\x6e\xed\x00\xad\x1f\xe3\x23\xda\x49\x0f\x61\xae\x4c\x59\x92\xe7\x15\x2f\xdf\x1b\xd2\x9e\x2b\x33\x8a\x3e\x67\x3c\xb8\x23\x26\xa8\x65\x60\x83\xc9\x5d\x55\x84\x48\xe3\xae\x9d\x74\x68\x66\x3b\xf4\x3c\xcc\x3d\xdb\xea\xde\x2f\x1c\xa4\x3a\x4d\x21\xfb\x2e\x76\xa0\xdf\xcc\x4f\x96\x4b\x08\x36\xcb\x06\xbc\xd1\x0e\x97\x6b\x79\x94\xca\x3c\x83\x6f\x4d\x64\xe8\xe6\xda\xd3\xf1\x62\x2d\x5b\xa9\x71\x77\xec\xc8\x10\x41\x7d\xea\x64\x0a\xf1\x08\xa5\xd8\x27\xef\x7e\x47\x83\x53\xc0\xe3\xe1\xfb\xe6\x28\x6a\xbb\xc5\xbd\x73\xda\x12\x46\x3e\x02\xb7\x72\xee\xd9\x22\x8f\xe8\xbf\xca\x36\xe1\x81\x53\x9c\xd0\x6b\xba\x94\x26\x4e\x3e\xb3\x08\xec\x5b\x9e\x13\xbb\xef\xe0\x98\x1e\xfa\xb7\x4d\x35\x3c\x16\x35\xd2\x6b\x57\xac\xa8\x67\x0b\x6f\x34\xf6\xf0\xb0\x10\x12\x18\x65\x3b\xf4\x5a\x9b\x9a\x20\xb5\x9b\xd6\x94\x28\xd9\x64\x9e\x00\xd5\xc5\x7c\x3d\x01\x70\xe4\x0e\x2c\x69\x79\xe3\xe2\xbe\x9e\xf8\x41\xee\x82\xe1\x0d\x2f\x55\xf7\x2b\x51\xca\x16\x94\x56\x3c\x9b\x67\x46\x8c\x2e\x1b\x32\x80\xd9\xab\xf8\xba\x2e\xcc\xb4\x5e\x00\x8a\xbe\x8a\x4f\x34\x4c\xaa\xf5\xfd\x8c\x77\x57\x48\xa9\xb0\x1a\x3b\x6c\x8b\xd4\x74\x93\x74\x8d\x15\xfd\x26\xba\x60\x99\x14\x97\x71\x41\x97\x63\x74\x8c\xce\xf0\xcf\x53\x3d\xe9\x15\xfe\xb5\x62\xe2\xb4\x0e\x01\x1e\x2c\x37\x72\xe1\xb2\x2f\x8e\x50\xb1\x55\xe7\x7e\x69\x8e\x3a\x02\x69\x10\x23\xde\xb2\x14\x42\x12\x76\xf8\x4e\xb7\x13\x0e\x34\x2a\xad\x32\xcf\x94\x6c\xe9\x2f\xe9\x96\xa1\x30\x8e\x34\x8c\x6d\xf6\x57\x7d\x55\xc1\xf7\xa4\xc4\x4e\x4d\xb2\xcb\x03\x5d\xfb\x6b\xe7\xfa\x0a\xc1\x86\x31\x52\xa7\xdb\x55\x1a\x72\x36\xf0\xe9\xe5\xb3\x96\x95\x6a\x6d\xe4\x17\x51\x8e\x61\xe4\x81\xef\xe7\xf4\x0e\x67\xca\x53\x05\xb8\x99\x65\x56\xc0\xb4\x77\x08\x9f\x43\xc6\x55\xad\x48\xd0\xe0\xfe\x37\x98\x56\xc3\x69\x96\x61\x86\xaa\x17\x54\x6a\x2f\xac\xcb\xce\x31\x87\xd3\x51\x8d\xd1\xb4\xac\xe2\xc8\xea\x37\xf0\xb6\xf8\x99\x4b\xdd\x89\x96\xc5\x7b\x56\x3b\xcc\x93\xb6\xc4\x5d\xcd\x74\x6f\x47\xdb\x18\x82\xa9\x45\x6f\xe6\x49\x41\x6b\xc6\x1b\x3b\xf8\x18\x1e\xbe\x4a\x6b\xcf\x53\x28\x64\xe1\xc8\x53\x48\x43\x43\xc9\x4d\x52\x3b\xd9\x3c\x4c\x97\xd9\xa6\x04\x5b\x55\x65\xa1\x9c\xe5\x9d\xe5\xa1\x3b\x17\xf2\xa4\xed\xd7\xab\xb5\xbe\x42\xa6\xd8\xe0\x73\xd3\xee\xed\x96\x8e\x56\xf7\x5a\x1e\x1b\xac\x35\xf8\xd8\x1d\xb9\x18\x63\x65\x01\x5c\xa9\x29\x9b\xa6\x13\xf7\x4d\xb3\xcf\x38\x0e\x98\xf9\x79\xc1\x53\xa7\x1f\xfa\xc1\xff\xab\x88\x73\xf3\xe7\x63\x82\x34\x0f\x76\xad\xfc\x2a\x56\x1d\x18\xd3\xbd\x8c\x84\x9b\x2e\xf0\x15\x08\x57\xc5\x66\x68\xba\x5d\x52\x15\x58\x28\x09\x65\x95\xe8\x05\xe4\xdd\x3c\x3b\x78\x05\x89\x77\x08\x9c\xaf\x71\xef\x99\x9c\x34\xa3\x5d\xec\x5d\xd3\x66\x52\xfa\x3e\x9d\x11\x6f\xc0\xc0\xf4\xd3\x23\x2b\x85\x23\x3d\x4a\xea\xfb\x0d\xa9\x56\x81\x41\x6f\x5a\x9c\xa6\xd7\x22\xa1\x36\x52\x8c\x4e\x9d\x13\x50\x66\xc8\xc8\x22\x4e\xac\x06\xb0\xf3\x75\x3b\x81\x0b\xc6\xf1\xb2\x98\x22\xbf\x83\xc3\x21\xfb\xdf\xc1\xa1\xa3\xe8\xed\xa9\xca\xc8\xd1\x4a\x7b\x80\xe8\x43\xb3\xd3\x48\x5f\x12\x19\x4c\x73\xb5\xe5\x08\xdb\xa6\xac\x5c\xf1\x43\x91\x6d\x94\x88\x55\xd3\x64\x0b\xf8\x32\xf3\x81\xea\x09\xe5\xae\xd0\xd0\xd8\x4f\x47\x6a\xab\xa6\x55\x5f\xf3\x68\xba\x34\x59\x61\x1d\x0f\x0f\xbc\x0a\xe0\x3a\xec\xe4\x42\x9b\x84\x65\x15\x47\x52\xf7\x61\xe7\xf1\xac\x4e\x9a\x02\x29\xc4\x4f\x73\x0d\x5a\x5e\xdb\x43\x3b\xe5\xba\x7d\x5a\x23\x6b\x86\x8b\x68\x4d\x01\x45\x10\xcb\x8b\x1e\x5a\xf8\xcd\xdf\x07\x23\x05\x42\xcf\x67\x65\xcf\x2c\xd7\xba\xe9\x8a\x53\xd9\xf4\xab\xea\xdb\xaf\x9c\xd0\xeb\x6d\xad\xa6\x3c\x5a\x80\x7a\x26\xa2\x29\xa8\x53\x01\x24\xe5\x06\x61\x6d\xa3\xfd\x40\x5d\x93\x5f\x9a\xac\xb3\x35\xef\x92\xa7\x2e\x6c\xf6\x3a\x9d\x6b\xb3\x57\xa4\x8b\x3b\x8a\x4e\x4a\x59\xab\xd2\x69\x01\xa9\x1f\x08\x19\xef\xdf\xe2\x64\x08\x59\x82\x0f\xe8\x72\x06\x6a\x77\xcc\x9f\x0f\x35\x63\x26\xf1\xa8\xbd\x14\x40\xfc\x32\xa1\x4a\x4b\x0d\x0a\x99\xf2\x88\xa7\x82\x2a\x9a\x22\x2b\xcb\x7f\x64\x29\x6d\xa6\x94\xaf\x78\x99\x60\xbc\xa6\xd1\x8d\x59\x9c\x10\x15\x84\xcb\x35\x98\xd0\x75\x96\xb0\xca\xd2\xb1\x2e\x50\x8c\x89\xb7\xac\x4b\x8f\x11\xc1\x87\x34\x4f\x4d\x83\x02\xe9\x61\x71\xed\xee\xbc\xb2\xca\x76\x0a\xba\xc6\xa0\x01\xba\x91\x57\x42\x7d\x8a\xd9\x04\x1a\x02\x15\x7c\xf2\xd5\xf6\xd4\x19\xc3\x6a\x6d\x98\x38\xfb\xd4\xd6\xe0\x3f\x1f\xa9\xec\x80\x06\xfc\x25\x5d\x3b\xb9\x29\xd3\x5d\x65\x4f\x9a\xc6\xd3\x97\xf4\x5b\x9a\x73\xe7\xec\x1c\x3a\x4f\x57\x54\xf7\xa1\xc7\x4b\x46\x3b\x2c\x1b\xdf\xdf\xbf\xe9\xa8\x0d\x65\x64\x09\xaf\x15\xba\x3f\x31\x81\x1b\xc8\x5e\x5c\xec\xb8\x75\x14\x61\xbd\xde\x53\x75\x04\xbd\x5d\xa4\xab\xd7\x57\xe3\x9c\x09\xe0\xe6\x9c\xc4\xc3\x27\x76\x95\x4e\x1e\x99\x2f\x3a\x8e\xb3\x80\x25\x0a\xaf\x1a\x27\xfb\x7a\x7b\x61\x49\x81\xd7\x74\x77\x81\x3e\x97\x0c\xbb\x82\x6e\xf6\xab\xbe\x6c\xfc\xe2\x69\x02\xee\xa9\xc0\xb6\xc2\xde\xdd\x77\x69\xad\xa3\x98\xa9\xf7\x16\x73\x87\xa3\xe1\x71\x36\x5d\xb3\x41\x4d\xbc\xa4\x81\x2b\xaf\x5c\xc7\xe1\x55\x55\xc4\x57\x75\xc5\x17\xec\x39\x0c\x6b\x88\xd9\x45\x06\x03\x6e\x53\x4d\x74\x57\xf6\xe1\xca\xd6\x3a\x01\x31\xf4\x87\x13\xf2\x12\x1a\x64\xf5\x39\x0f\xcc\xcd\xb2\x65\xe5\xd9\xf2\x73\xf6\x7e\x9c\x65\x37\x31\x5d\x40\xac\x7a\x03\x19\x65\x59\xb6\xf1\x03\xae\x4a\xdf\xdd\x70\xc5\xea\xa7\xd8\xd5\xa3\xe1\x70\x94\x95\x61\x85\xbb\xd3\x5e\x5f\x32\x25\xae\xf6\xb6\xd6\x82\x48\xe1\x76\xdd\x0b\x6e\x3b\x7a\xdb\x43\xad\x9d\xb1\xf0\xd6\x1e\x63\x00\x48\x6e\xb5\x1b\xa3\x24\x03\xbb\xd2\x72\x47\x8f\xb0\xd2\x3b\xc7\x1a\x44\x20\x54\x4a\xff\x9a\x5d\x95\x76\x0f\x2c\x46\x51\xa9\x71\x0d\x63\xd7\x25\x0c\x6f\x22\xdc\xf3\x02\x46\x8f\x96\xfe\x8e\xcb\x17\x56\x03\xd7\xae\x8b\x17\x2f\x73\xed\xe2\x11\x97\x2e\x3c\xe7\x90\xaa\x25\xf5\x5f\xb6\xf0\x5a\x59\x3d\xac\xeb\x7f\xcd\x62\xf7\x25\x8b\x9e\x57\x2c\x3a\x2f\xc4\xb8\xbb\x05\x7a\x5c\x8a\x51\x39\x1b\xd0\xa8\x1c\xcd\xeb\x14\xcf\x4f\xdc\xa0\xfa\x85\x0d\x27\x88\x9a\xe0\x7a\x8c\xf6\x49\x91\xb6\x47\x08\x6e\x10\xc2\x69\x89\xd4\x93\xcc\x1d\xb1\xbd\x36\x38\x2c\xd2\x11\x48\xd3\x08\x91\x78\xbc\x01\xfe\xf8\xae\x62\xf8\x47\x04\x8f\x40\x77\x12\x45\x78\x4a\x09\xe1\x50\x37\x46\xb1\xca\x83\x0e\xcc\x6d\x37\xd8\xf8\xec\xcb\xe2\x72\x32\x0f\x3c\xdd\x01\xf8\x63\x5e\xed\x90\x3f\xae\xa7\xf6\x33\x2b\xaa\xf1\x2a\x97\x62\x38\x9d\x35\x14\xb4\x6d\xed\xed\x1a\x9f\x7d\xf3\x5e\xc3\x31\x8b\x1e\xad\xa5\xdc\x5d\xdc\x08\x1a\x64\xca\xa9\xab\xd6\x7e\xb7\x8f\x5e\x43\x69\xdd\xdb\x47\x23\xae\x5a\x01\xeb\x06\xcc\x5a\x3c\x50\x60\x3a\xee\xb7\x34\xa4\x3a\x8f\x39\x3b\xaf\xb5\x28\x19\xe0\xd1\xa1\x96\xb4\x5b\xf7\x8e\x82\x7f\xc4\xf9\x87\x38\x71\xb4\x10\x07\xdf\x53\x3b\xf7\xdd\xab\x21\x83\x02\x1b\x15\x47\xd5\xde\x7f\x9b\x46\xea\x36\x2c\x08\x28\x0c\xf9\x0b\x29\xe8\xef\x75\x5c\xd0\xb7\x7b\xf0\x71\xbf\x5c\xde\xec\xbd\x73\x02\x83\x66\x01\x70\x4a\x7f\xe2\xb0\xe1\x64\xbc\x78\xeb\x86\x13\xc2\x0d\xb0\x7b\x0e\x19\x76\x13\x52\x69\xe7\xb7\x38\xcf\x83\x99\xa2\xdb\xc5\x6d\xbb\xdb\xbc\x05\x76\x37\xa5\xe1\x2b\xa4\xa9\x91\x81\xbd\x11\xd9\x0b\x9c\xcd\x03\x5d\x47\xcb\xb0\xc3\x7b\x03\x67\x47\x9a\x3b\x1f\xe4\xd3\xee\x8d\xf6\xf6\xcc\x95\x5b\x0d\x13\xf4\x2e\xc7\xa0\xb3\x11\x3d\x60\xc3\x4a\x88\xf5\x5b\x7a\x0b\x6e\x60\x40\xc0\x7d\x55\xf4\xae\x7a\x67\xf1\x87\xcd\x82\xac\xe4\x27\xa3\x0e\x0e\x06\xcd\x1f\x10\xde\x16\x15\x5d\xbe\xbf\x87\xe5\xa3\x16\xc0\x6a\x3c\x90\xfa\xb6\x8c\xcc\x6d\xfa\xc6\x88\x1a\x8a\x33\xe0\x1f\x5e\x34\x42\x22\x46\xcd\x1f\x7e\x40\xb4\x9e\xa0\x06\x5e\x80\xec\x96\x16\x45\xbc\xa4\xe5\xc8\xbf\x3c\x8e\x48\xf4\x4a\x5c\xc8\x01\xdf\xba\x06\xe0\xcf\x03\x49\x61\xd7\x47\xda\xa2\x90\xe3\xac\x11\x7a\xf4\x8d\xec\x95\xeb\xbd\x01\xd9\xdb\x8f\xe0\x5f\x0e\x24\xde\xa1\x24\x74\x21\xff\xe1\x7b\xe9\x1c\xb5\xb5\x34\x04\xff\x01\xed\x1b\x16\xdc\x51\xbf\xe5\x5b\xdc\x0a\xc3\x70\x99\xa5\xd4\xd6\x45\xc0\x63\x16\x9e\x64\x68\xd4\x15\xf4\xaa\xca\xb6\x23\x96\x45\x65\x5a\xac\x41\x66\x85\xbe\xcc\xeb\x8e\xb8\x76\x82\x3c\x2b\xc1\xe9\xd6\x7d\x42\x5b\xc5\x90\x43\xb2\x91\x5e\x37\x76\xbb\x8c\xd6\x74\x59\xeb\x37\x94\x17\xe2\xd9\xe4\x2e\xc7\x73\x1a\x11\xdb\x30\xe2\xc4\x1b\xe3\x24\x8f\x57\x10\xad\x34\x93\x99\x6e\xaf\x7d\x97\xfe\xca\x73\xbb\x63\xba\x74\x10\x2c\x8a\x95\x46\xbd\x33\x17\x15\xbf\xef\xcd\xdd\x9d\xef\xf0\xf1\x3b\x1b\x2b\x32\x4a\x18\xfa\x1d\x66\xfa\x1e\x08\x21\x93\x00\xa2\x35\x9f\x01\x28\x31\x81\x2b\x93\x71\x6c\x11\xf7\xbe\x33\x5a\x6c\xe2\xb2\x74\xb9\x69\x62\xfa\x69\x05\xd6\xb5\x6b\x44\xcf\x48\xa2\xf6\xa0\x9f\x47\xfa\xa3\x29\x04\xdc\x37\xd4\x75\xe7\x56\xf3\xd9\xe4\x89\x7c\x57\x92\x0d\x9c\x94\x29\x66\x69\xa4\x17\xaa\xa8\xb0\x68\x96\xa1\xf1\x9e\xbb\x5b\x12\xad\x4c\xfc\xf8\x6c\xd1\xf9\xf5\x5c\xfc\x06\x18\x0f\x3f\x3f\x85\xe5\x64\xac\x7e\x4b\x04\x23\xea\xa2\xd0\x82\xfe\xae\xef\xda\x32\x42\x59\x12\xd4\xe5\x3e\x0d\xcb\x8a\x5d\x14\x54\x4f\x18\x1e\x89\xe3\x27\x7d\x21\x1c\xc7\xcf\xc0\x41\xeb\x7d\xec\x1e\x2c\xc2\xe4\x59\xa4\x00\x9a\x67\x2f\x27\xcc\xf7\x53\x90\x8c\xf5\xb3\x79\x0b\x88\x20\xab\x7a\x69\x44\x8c\xcb\x02\x0f\xda\x02\xef\x75\x41\xd7\xdd\x62\x29\x90\xf8\xc6\xf1\xad\x41\x8e\x2b\x76\x73\x05\xac\x49\x00\xd4\xc3\x4e\x4d\xbc\x05\x80\xf3\xee\x9d\xdb\x09\x3c\xff\xce\x9d\xf2\xed\x47\x56\x47\xa4\xd5\x0f\xf0\xa6\xd1\xeb\xdd\x7c\x33\x2e\x8b\x36\x5c\xe3\x65\xf0\xbe\x8c\xb0\xee\x93\x5a\x29\xe4\x1b\x61\x61\x3a\xb7\xe5\x31\x13\x99\xec\x73\xe0\x1d\x58\x4c\x13\x56\xcf\xcb\x15\xe3\xd2\xb9\xd6\xf4\xe9\xbe\xd9\x6a\x16\xb8\x3c\xdb\xdf\xab\xb8\xe5\xad\x91\x18\xe5\x1b\x59\x4c\x32\x8b\x2c\x46\x69\x49\x7b\x6d\xb6\xba\x76\x57\x57\xf4\xba\x97\x55\xcc\x91\x55\x30\x2b\x03\xc7\xda\x87\xe1\xb3\x9e\x92\x93\xab\xad\x0f\xce\x0a\x53\xe3\xbd\xa4\x1f\xf7\xd5\xc5\xdc\x55\x31\x2d\x58\xd2\x2b\x62\x5a\xaf\xaf\x0d\xe6\xff\x6a\x98\x97\xff\xd6\x17\xef\x5e\xb3\xb7\x94\xf7\x6e\xe0\xd7\xbd\x5d\xb5\xbd\x1b\xbc\x99\xe8\x8a\x7a\x4a\x5c\x9e\x31\xdc\x20\xd0\xa2\x39\x32\x28\xb1\x27\xc6\xd9\x05\xb2\x1b\xdb\xdc\xc4\xf5\xf7\xb8\x5a\xf7\xc0\x15\x1d\xef\x24\x1e\x40\x4e\xc0\x49\x64\x45\xfc\x07\x75\xf6\x35\x59\xa3\x5c\xe7\x22\x4a\x25\xd0\xc9\xd7\x3f\x39\xd0\x18\x4f\x8c\x3e\x72\xa7\x10\xab\x1e\x6c\x87\x3d\x56\xbf\xce\xc2\xfe\x96\x08\xdd\xe6\x2c\x7e\x1d\x8d\xc4\x37\xb6\x08\xa3\x73\x4a\x13\x8a\x72\xd2\x1e\x89\xc0\x0a\x31\x69\xdc\x61\x94\xd8\x17\x2d\x62\x2f\x7e\xc1\x8f\x68\xcd\x9e\x17\xc8\x2d\x8c\xab\x48\x0f\xcd\xe5\xe9\xa0\xbc\x87\x1d\xde\xa0\x8d\x6d\x0f\xa6\xc4\x97\xc4\x10\x3d\x02\x10\xf0\xf8\xb5\x5e\x03\xaf\xab\x50\x9d\x97\x8b\x6d\x92\x6b\xff\x0a\x00\x00\xff\xff\x29\x87\x49\x9f\x2c\x57\x00\x00")

func templatesAppTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesAppTmpl,
		"templates/app.tmpl",
	)
}

func templatesAppTmpl() (*asset, error) {
	bytes, err := templatesAppTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/app.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"templates/app.tmpl": templatesAppTmpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"app.tmpl": &bintree{templatesAppTmpl, map[string]*bintree{}},
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

