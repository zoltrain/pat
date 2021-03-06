// Code generated by go-bindata. DO NOT EDIT.
// sources:
// schema/schema.json
// schema/schema.yaml
package pkg

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

var _schemaSchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x55\x4d\x6f\xd4\x30\x10\xbd\xe7\x57\x58\x6e\x0f\x20\x76\x09\x07\x4e\x7b\xab\x04\x3d\x21\x51\x21\x6e\xb4\x48\x26\x99\xec\xba\x8a\xed\x30\x9e\x48\x5d\x91\xfe\x77\xe4\x7c\xad\x1d\x3b\x55\x77\x85\x38\x66\x3c\x33\xef\xcd\x7b\x63\xe7\x4f\xc6\x18\xbf\xb6\xc5\x01\x94\xe0\x3b\xc6\x0f\x44\xcd\x2e\xcf\x1f\xad\xd1\xdb\x21\xfa\xde\xe0\x3e\x2f\x51\x54\xb4\xfd\xf0\x31\x1f\x62\x57\x7c\xe3\xea\x44\x59\x4a\x92\x46\x8b\xfa\x0e\x4d\x03\x48\x12\x2c\xdf\xb1\x4a\xd4\x16\xfa\x84\x12\x2a\xa9\xfb\x14\x17\x77\x58\x8c\xf1\x1b\x6b\x5d\xaa\xd1\x73\xe8\x15\xad\xfa\xa4\xc6\x3f\x9a\x4a\x5d\x31\x05\xdf\x6e\x22\x84\xca\x8d\x73\x95\x7b\x14\xf2\x4f\x2d\x8a\x1e\x78\x4e\x7d\xde\x9c\xba\x14\x46\x29\xd0\x51\x2b\x3a\x36\xe0\x5a\x59\x42\xa9\xf7\xe9\x52\x78\x6a\xa0\x20\x28\x97\xb5\x92\x40\xd9\x45\x70\x95\xdd\xe7\xbe\xc9\x82\x60\x80\xe3\xb1\x11\x88\xe2\xe8\x91\xc9\x16\xc9\x1c\xe1\x77\x2b\xb1\xa7\xf4\x23\x50\x2a\x45\x7b\x0c\x3d\xcc\xd5\x13\x8c\xf9\xf5\x08\x05\x0d\xe7\x63\x6f\x3e\xab\xe8\xd9\xd7\x08\x22\x40\x17\xe2\x3f\xef\xef\xcb\x77\x6f\xb4\xed\x5a\xdb\x29\xdb\xd9\x4e\x75\x87\xb7\xd7\x3c\x6a\xed\xeb\x39\xb5\xf6\x25\xf0\xba\xbf\xc4\xe6\x56\x3e\x51\x8b\x10\x90\x59\x59\x93\x32\x26\x7e\xf1\xb2\x28\x20\x94\xc5\xd2\xdb\x35\xc3\xd7\x76\xe8\x1f\xbb\x3b\x0f\x98\x20\x7a\x96\xc5\xdf\xda\x3a\x10\x8f\x1b\x0d\x5f\xab\x00\x2c\x18\x7b\x45\xf1\xfe\xac\x42\xa3\x6e\x65\x0d\xd1\xc9\x8b\xba\x9c\x86\x8e\x55\x4a\x0e\x1f\x62\x79\xe1\x87\x94\x7f\x67\xb1\xff\x22\x09\x50\xd4\xff\x6b\x80\x09\x2e\x3d\xc3\xeb\x7c\xcc\x46\xc8\xd4\x6c\x5c\x4c\x0f\x70\x60\x71\xbc\xba\x2b\x17\xe3\xf4\x7e\x47\x6b\x99\xd8\xdf\x69\xa5\xaa\xe1\x9e\x5e\x06\x39\x5d\xf2\x73\x00\xb5\x50\x90\x7a\x44\x52\xef\x0e\x2e\xf7\x3d\x4d\x63\xb8\x16\x0b\x85\x23\x33\x07\xe4\xa0\xf3\x26\x12\x3e\x1b\x2d\xe4\x24\xa9\xbf\x1b\xee\x09\x65\x5b\x76\x87\x46\x01\x1d\xa0\xb5\xec\xa6\x06\x24\xf6\x1d\x2c\x49\xbd\x67\x2d\xc9\x5a\xd2\x71\xf8\xf1\x2e\x3d\x7f\xce\xfe\x06\x00\x00\xff\xff\x11\x4f\x07\x19\xc8\x07\x00\x00")

func schemaSchemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_schemaSchemaJson,
		"schema/schema.json",
	)
}

func schemaSchemaJson() (*asset, error) {
	bytes, err := schemaSchemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/schema.json", size: 1992, mode: os.FileMode(436), modTime: time.Unix(1530823201, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaSchemaYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x93\x4f\x8b\xdb\x30\x10\xc5\xef\xfa\x14\x83\x37\x87\x96\xe2\xba\x87\x9e\x7c\x5b\xd8\xee\xa9\xd0\x65\xe9\x71\x29\xa8\xf6\x38\xd6\xa2\x3f\xee\x68\x0c\x1b\xf0\x87\x2f\x92\xff\xc7\x4e\x48\xbb\xba\x44\x19\x4d\x7e\xf3\xde\x93\x72\xf0\x45\x8d\x46\xe6\x50\x33\x37\x79\x96\xbd\x7a\x67\xd3\xbe\xf6\xd9\xd1\x31\x2b\x49\x56\x9c\x7e\xf9\x9a\xf5\xb5\x3b\xc1\x8a\x35\xe6\xd0\x48\x86\x14\x9e\xc8\x19\xe4\x1a\x5b\x0f\xf7\x1a\x89\xe1\x27\x7a\x56\xf6\x08\x2d\x2b\xad\xf8\x24\xf8\xd4\x60\x0e\xee\xf7\x2b\x16\x2c\x1a\x72\x0d\x12\x2b\xf4\xb9\x00\xb0\xd2\x60\xf8\x04\xe8\x9b\x3c\x93\xb2\x47\x01\x40\xad\xee\x3b\x00\x0e\x84\x55\x0e\xc9\x5d\x56\x62\xa5\xac\x62\xe5\xac\xcf\x9e\xc3\x79\x22\x00\x2a\xf5\xc6\x2d\x8d\xbd\x3d\x45\x12\xc9\x53\xfc\xae\x18\xcd\x70\x74\x01\xf4\xd8\xff\x3e\xa0\xa4\xf7\x41\x99\xb3\xff\x0d\xbb\x1f\x09\x89\x20\xfc\xd3\x2a\xc2\x32\xf4\xa7\xd1\x67\xdc\x44\x5f\x71\x37\x4f\x13\xb2\x2c\x23\x40\xea\xa7\x39\x1d\xa8\xa4\xf6\x28\x16\xf8\x80\x7a\x9e\x73\x59\xc5\x1a\x0a\xce\xe2\x8f\x6a\xd4\x97\xc2\x3a\xe9\x71\x55\xe4\xcc\xa3\xd2\xb8\xac\x6d\xd2\xef\xd7\xd2\xc2\xb8\xd2\x09\x70\xc3\x9c\xef\x8a\x91\xa4\x7e\xd7\xa8\x81\x21\x00\x86\x9b\xba\x60\x7e\xab\xa2\x6c\x49\x86\xdc\x66\xe6\xee\x9d\x3d\x0c\x6d\xc9\xd0\x66\x90\x49\x15\x0b\x2f\xe7\xcf\x60\xf3\x14\x76\x5d\x9d\x3b\x4a\x27\x3d\x53\x61\x98\x24\x00\xa6\x77\x73\xb3\x39\xc9\xff\x68\x0b\xdf\x1a\x2c\x78\x99\xf0\x4d\xbe\x76\xd9\xdf\x22\x6b\x85\x2f\x9c\x31\x68\xf9\x9c\x7e\x35\x10\xc9\xd3\x76\x54\x17\x0b\xd7\xfe\x0e\x00\x0f\xab\x6b\xdd\x8c\x69\x24\x33\x92\xcd\x21\xf9\xf5\xf2\x52\x7e\xfa\x60\x7d\xd7\xfa\xce\xf8\xce\x77\xa6\xab\x3f\x1e\x82\xe0\x85\xfe\x9d\xc0\xff\x06\x00\x00\xff\xff\x01\xfe\xe5\xd6\x0e\x05\x00\x00")

func schemaSchemaYamlBytes() ([]byte, error) {
	return bindataRead(
		_schemaSchemaYaml,
		"schema/schema.yaml",
	)
}

func schemaSchemaYaml() (*asset, error) {
	bytes, err := schemaSchemaYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/schema.yaml", size: 1294, mode: os.FileMode(436), modTime: time.Unix(1530813958, 0)}
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
	"schema/schema.json": schemaSchemaJson,
	"schema/schema.yaml": schemaSchemaYaml,
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
	"schema": &bintree{nil, map[string]*bintree{
		"schema.json": &bintree{schemaSchemaJson, map[string]*bintree{}},
		"schema.yaml": &bintree{schemaSchemaYaml, map[string]*bintree{}},
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

