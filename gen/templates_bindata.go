// Code generated by go-bindata.
// sources:
// ../templates/objc/ResourcesManager.h.tpl
// ../templates/objc/ResourcesManager.m.tpl
// ../templates/objc/model/model.h.tpl
// ../templates/objc/model/model.m.tpl
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

var _TemplatesObjcResourcesmanagerHTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x0a\x4a\x2d\xce\x2f\x2d\x4a\x4e\xd5\xcb\x00\x04\x00\x00\xff\xff\x2a\xce\xa2\xe0\x0a\x00\x00\x00")

func TemplatesObjcResourcesmanagerHTplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesObjcResourcesmanagerHTpl,
		"../templates/objc/ResourcesManager.h.tpl",
	)
}

func TemplatesObjcResourcesmanagerHTpl() (*asset, error) {
	bytes, err := TemplatesObjcResourcesmanagerHTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/objc/ResourcesManager.h.tpl", size: 10, mode: os.FileMode(504), modTime: time.Unix(1448202337, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesObjcResourcesmanagerMTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x0a\x4a\x2d\xce\x2f\x2d\x4a\x4e\x55\xc8\x4d\xcc\x4b\x4c\x4f\x2d\x52\xc8\x05\x04\x00\x00\xff\xff\x90\x25\x90\x6b\x12\x00\x00\x00")

func TemplatesObjcResourcesmanagerMTplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesObjcResourcesmanagerMTpl,
		"../templates/objc/ResourcesManager.m.tpl",
	)
}

func TemplatesObjcResourcesmanagerMTpl() (*asset, error) {
	bytes, err := TemplatesObjcResourcesmanagerMTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/objc/ResourcesManager.m.tpl", size: 18, mode: os.FileMode(504), modTime: time.Unix(1448202337, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesObjcModelModelHTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\xc8\xcc\x2b\x49\x2d\x4a\x4b\x4c\x4e\x55\xa8\xae\xd6\x73\x2c\xc8\xf4\x4b\xcc\x4d\xad\xad\x55\xb0\x52\xf0\x0b\xf6\x4f\xca\x4a\x4d\x2e\xe1\xe2\x72\x48\xcd\x4b\x01\x04\x00\x00\xff\xff\x7f\x44\x0c\x78\x28\x00\x00\x00")

func TemplatesObjcModelModelHTplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesObjcModelModelHTpl,
		"../templates/objc/model/model.h.tpl",
	)
}

func TemplatesObjcModelModelHTpl() (*asset, error) {
	bytes, err := TemplatesObjcModelModelHTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/objc/model/model.h.tpl", size: 40, mode: os.FileMode(504), modTime: time.Unix(1448130277, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesObjcModelModelMTpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\xc8\xcc\x2d\xc8\x49\xcd\x4d\xcd\x2b\x49\x2c\xc9\xcc\xcf\xe3\x72\x48\xcd\x4b\x01\x04\x00\x00\xff\xff\xd3\xe6\x58\x8a\x14\x00\x00\x00")

func TemplatesObjcModelModelMTplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesObjcModelModelMTpl,
		"../templates/objc/model/model.m.tpl",
	)
}

func TemplatesObjcModelModelMTpl() (*asset, error) {
	bytes, err := TemplatesObjcModelModelMTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/objc/model/model.m.tpl", size: 20, mode: os.FileMode(504), modTime: time.Unix(1448129855, 0)}
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
	"../templates/objc/ResourcesManager.h.tpl": TemplatesObjcResourcesmanagerHTpl,
	"../templates/objc/ResourcesManager.m.tpl": TemplatesObjcResourcesmanagerMTpl,
	"../templates/objc/model/model.h.tpl":      TemplatesObjcModelModelHTpl,
	"../templates/objc/model/model.m.tpl":      TemplatesObjcModelModelMTpl,
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
	"..": &bintree{nil, map[string]*bintree{
		"templates": &bintree{nil, map[string]*bintree{
			"objc": &bintree{nil, map[string]*bintree{
				"ResourcesManager.h.tpl": &bintree{TemplatesObjcResourcesmanagerHTpl, map[string]*bintree{}},
				"ResourcesManager.m.tpl": &bintree{TemplatesObjcResourcesmanagerMTpl, map[string]*bintree{}},
				"model": &bintree{nil, map[string]*bintree{
					"model.h.tpl": &bintree{TemplatesObjcModelModelHTpl, map[string]*bintree{}},
					"model.m.tpl": &bintree{TemplatesObjcModelModelMTpl, map[string]*bintree{}},
				}},
			}},
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
