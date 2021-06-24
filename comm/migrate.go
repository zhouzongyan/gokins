package comm

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var __000001_pipeline_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8e\x3d\x0e\xc2\x30\x0c\x46\x77\x4e\x91\x7b\x74\x02\x51\xa4\x4a\x48\x20\xda\x81\x2d\x3f\xd4\x2a\xa1\x71\x1c\xd9\x29\xe7\x47\x62\xaf\x33\xbf\xe7\xcf\xef\xfc\xb8\xdd\xcd\x74\x3c\x5d\x7b\x33\x5c\x4c\xff\x1c\xc6\x69\x34\xae\xda\xb0\xc5\x34\xbb\xee\xb0\xc7\x5f\x38\xdb\x14\x33\x28\xca\x9b\x68\x55\xf0\x87\x82\x42\x11\x44\xfc\xa2\xcd\x97\x58\xa0\x51\x20\x55\x9f\xf8\x7a\x8e\x3e\x24\xf5\x8b\x67\x8f\x0a\x67\xd2\xcf\x81\x51\x24\x52\x56\x9c\x4d\x80\x1b\xd8\xa2\x2c\x2d\x85\xa1\x50\xd3\xd1\x6b\xff\x4e\xa5\x15\xb2\xeb\x7e\x01\x00\x00\xff\xff\x12\x21\xfd\x36\x19\x02\x00\x00")

func _000001_pipeline_down_sql() ([]byte, error) {
	return bindata_read(
		__000001_pipeline_down_sql,
		"000001_pipeline.down.sql",
	)
}

var __000001_pipeline_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x5f\x6f\xdb\xd6\x15\x7f\xcf\xa7\xb8\x6f\x94\x01\x0a\x20\x29\x29\xb6\x98\x27\x37\xd1\x06\x61\xb6\x5c\x28\xf2\xd6\x3e\x91\x57\xe2\x95\x4c\x47\xfc\x03\x5e\xca\x70\xde\x1c\x74\x6b\xed\xb5\x6b\x9c\xcd\x69\xea\x75\xde\x92\x62\x40\x3a\x60\xab\x83\x21\xf3\x30\xd9\x9b\xbf\x8c\x25\xd9\x4f\xfd\x0a\x03\x2f\x29\x89\xb2\x64\xf2\xf2\xdf\x80\x01\xcd\x83\x01\x85\xbc\xbf\x73\x79\xcf\xbf\xdf\x39\xe7\xde\xcb\xe7\x41\x3e\xe0\x9f\xf3\xbc\x01\x9b\x5d\x04\xb0\x6d\xf5\x5a\x76\xcf\x42\xa0\x6d\x58\xc0\x96\x9a\x3d\xb5\xab\x84\x2e\x7f\x58\xaf\xac\x36\x2a\xa0\xb1\xfa\xc1\x5a\x05\xc8\xde\x2a\x19\x80\xdc\x3d\x10\xf0\x4f\x56\x15\x19\xec\x40\xab\xb5\x05\xad\xdc\xfd\xe2\x12\xa8\x6d\x34\x40\x6d\x73\x6d\x8d\x0d\x5e\x66\xaa\x26\xea\xaa\x3a\x92\xe6\xd6\x6f\xae\xad\x81\x47\x95\x9f\xac\x6e\xae\x45\x03\xda\x41\x16\x56\x0d\x3d\x31\x20\xb6\xa1\xdd\xc3\x53\x0c\x9e\xe3\x16\x80\x80\x87\x1b\xeb\xeb\x95\x5a\x03\x30\xc3\x3f\xfe\x72\x70\xde\x1f\xfd\xfa\x6c\xb8\xf7\x8c\x09\xc1\x46\x96\x65\x58\x53\xe8\x52\x08\xf4\xcd\xd1\xf1\xf5\xe9\xe9\xd5\xe5\x9b\xe1\xb3\xd3\x50\xe8\x1d\xa4\xdb\xf4\xbb\xbe\xea\x7f\x7e\x75\x7e\x16\x06\x6a\xab\x1a\x92\xb0\x0d\x35\x53\x06\x0a\xb4\x91\xf3\x3b\x17\x72\x1c\x07\x6f\xaf\xdf\x7c\x31\x7c\x75\x76\xf3\xf2\x32\x1c\xde\xee\xa2\xe9\x9e\x85\x52\x29\x18\xfa\xf5\x67\x37\xdf\x7e\x1d\x06\xaa\x21\x8c\x61\x27\x0a\x2c\x51\x20\xdd\x29\x63\x1b\x5a\x36\x52\xa8\x4f\x63\x70\xb1\x37\x78\xfb\xb9\x73\x1a\xaf\xde\x87\x61\xb7\x55\x5d\xc5\x5b\x11\xc0\x47\xe7\xbf\x1b\x9e\xfc\x89\x0e\xbc\x65\x21\x18\x69\xe3\xfb\xdf\x0c\xce\xfb\x74\xd8\x3d\x53\x89\x84\x3d\xfc\xe6\xfd\xf0\xab\x77\x74\xd8\x9e\x57\xd3\x6b\x73\x74\xb0\x3f\xfc\xc3\x5f\x43\x60\x3f\xac\x57\xd7\x57\xeb\x1f\x83\x9f\x55\x3e\x06\x39\x27\x84\x2d\x81\xcd\xc7\xd5\xda\x4f\xc1\x07\x8d\x7a\xa5\x72\x6f\xe9\xc1\xbd\xb9\x70\xd8\xd2\x14\xc9\x89\x33\x5e\x44\x0c\x8c\x7b\x72\xc7\x32\x7a\x26\x5d\x24\x92\x49\x9c\xa5\x7c\x77\xdb\x68\x52\xbe\x79\x3b\x8a\x95\x16\xa9\x84\xbc\xa9\xf7\x34\x19\xa8\xba\x9d\xe3\xf9\xbb\x5e\x69\x19\xba\x4d\xa2\x8b\x8d\x76\x6d\xdf\x7f\x53\xd9\xd4\x78\x3b\x14\x7e\x43\x5e\xa5\x74\x03\xe7\xdd\x18\x6a\xdc\x32\x8c\x27\x61\x49\x2d\x5e\x4e\x93\xed\xa7\x66\x68\xd4\x09\x46\xc0\x3a\x34\xf1\x96\x61\xcb\xa0\x6b\xe8\x9d\x99\xc3\x0e\x58\x74\x4b\xd3\x71\x04\x6b\xb8\x93\x0c\xc0\x39\x56\x29\xf1\x01\x50\xe8\x93\x82\x04\xd5\x51\xcb\xb0\x14\x0c\x8c\x36\x70\xf5\x4d\xb3\x28\x1e\xb3\xda\x36\x9a\x91\x79\xd5\xb6\xd1\x0c\x31\xc0\x58\xf6\x17\x25\x8e\xdc\x89\x81\x6d\xd8\x99\xa5\x64\xa1\xec\xe7\x1f\xcf\x86\xef\xde\x8f\xfa\x97\xaa\x12\x18\x73\x65\x45\xc5\x66\x17\x3e\x95\x74\xa8\x25\x32\x92\x18\x84\x2f\xc6\x6e\x89\x92\x12\x6c\x32\x33\x12\x29\xa3\x5d\xd5\x96\x5a\x86\x82\x64\xd0\x54\x3b\x4e\xdc\x16\x42\x48\xe4\xde\xde\xe0\xb3\xfe\xe8\x75\x18\x6e\x46\xd4\x54\x9e\x55\x78\xd8\x41\x0c\x0e\x7f\x33\xf8\xdb\xab\x60\xc4\xcc\x48\x58\x86\x14\x2c\x3b\x02\x96\x1d\xfd\xca\x84\x7c\x39\x86\xa6\x76\x74\xc3\xf2\x19\x45\x64\xf7\xd2\x7b\x5a\x13\x59\xc1\x1e\x10\xac\x0d\x43\xd3\xa0\xae\xe0\x5b\x94\xe6\xce\xf7\x15\x64\x22\x5d\xc1\x92\x73\x18\xdb\xd8\xd0\xc3\x57\xa8\x1a\x45\x11\x12\x7c\x50\xfa\x8e\x6a\x19\xba\x86\x74\x1b\xd3\x4a\xc5\x86\x65\x4f\x8e\x85\x8f\x7a\x2c\x19\xe4\x5f\x9a\x0c\x19\x3b\xfd\x7a\xa5\x5e\xe4\x14\x3c\x29\x11\x43\x9a\x1b\x5e\x2a\xf6\x5b\x99\x97\x89\xc1\xea\x66\x63\x43\xaa\xd6\x1e\xd6\x2b\x8e\xe5\x07\x57\x1b\x00\xc8\xbb\xf1\xda\x24\x8e\x77\x47\x48\x6f\x83\xe7\x2f\x6e\xf6\x9e\x5d\xef\xfd\xea\x87\x8b\xfd\xc1\xf3\xd3\xd1\x5f\xfa\x3f\x5c\x1c\x84\x94\x42\xb4\x65\x78\x28\xca\xa4\x48\xa0\xe7\xae\x63\xd6\x4c\x53\xa7\x04\xcb\xa6\xae\x44\x02\x61\x54\xbd\x6d\xd0\x86\x04\xb2\xa0\x67\x75\xc3\x52\x66\x18\xc6\x9c\xcf\xb1\xae\xb5\x24\x77\x3d\x5a\xef\x88\xed\x7e\x26\xb4\xa0\x16\xd9\xf9\xc8\x2a\xba\xbe\x62\x52\xc7\xa3\xe2\x1e\x69\xf4\xa8\x42\x30\x14\x68\x43\x6a\xb3\x22\x4d\x37\x9c\xc8\x94\x33\x88\xe3\x74\xba\x8e\x6f\x4a\xc8\xd2\xb0\xc3\x35\xa2\x9b\xd3\x78\x25\x45\x34\x4f\x2d\x9e\x27\x89\xe8\x4e\x09\x03\xad\x99\x66\x6d\x8c\x36\x35\x58\x60\x9b\x71\xec\x9b\xe0\xec\xc0\x6e\x2f\x0d\x9c\x30\xcb\x7d\xb8\x59\xaf\x57\x6a\x0d\xa9\x51\x5d\xaf\x3c\x6e\xac\xae\x7f\x98\xe3\x96\x28\x60\x5d\x52\xb3\xb0\x45\xc5\x73\xe1\xeb\xa9\x22\x6c\x38\x4c\xb5\xf6\xa8\xf2\x11\x90\xab\x8f\x3e\x92\xf0\x53\x3c\xb5\x3b\xc9\xd3\x66\x6e\xac\xd6\x14\x9c\x8d\xda\x1b\xaa\xb5\xc7\x95\x7a\x03\x54\x6b\x8d\x8d\x5b\xce\xf0\xf3\xd5\xb5\xcd\xca\x63\x90\xe3\x59\xc0\x38\x6c\xd7\xd0\x19\xd6\xd5\xa3\x53\x12\xfe\x7e\x74\xf4\x1d\xe3\x7f\xc2\x08\x1c\x5f\xce\xf3\x5c\x5e\x28\x00\x9e\x17\x85\xb2\x58\x10\x18\x16\x70\x4b\x0f\x28\x64\x08\x2c\x60\xba\x46\x47\xd5\x67\x21\x47\xc7\xe7\x83\x7f\xbf\x64\xfc\x0f\xe7\xa4\x14\x39\x86\x05\x3c\x4f\x25\xa6\xc0\x02\xa6\x67\x9a\x10\xe3\x59\x39\x57\x97\xdf\x0f\x8f\xfe\x35\x38\xfd\x94\x14\xb9\xee\x13\x71\xfa\xa2\xc0\x09\x5c\x9e\x2b\xe4\x39\x1e\x08\xbc\x58\x5c\x16\x8b\xce\x97\xf1\x02\x95\xcc\x22\x0b\x18\xa8\x68\xaa\xef\xf4\x06\x87\x5f\x0e\x9e\xbf\x1b\xbd\xfc\xe2\xe6\xe4\x5b\xc6\xf7\x78\xfa\x71\xcb\x80\xe3\x44\x9e\x17\x0b\xce\xc7\xd1\xc9\x29\xb1\x80\xb1\x8c\x2e\xc2\x53\x39\xd7\x6f\x7f\x7b\x7d\xf0\xf7\xd1\xf7\x6f\x46\x87\x9f\x32\xde\x63\xb1\xab\x62\x7b\x2a\x8b\xcf\x73\x05\xc0\x17\xc4\xe2\x8a\x58\x5a\x61\x58\x50\xa0\x92\x75\xdf\x03\xe3\x19\x9f\xd0\xb1\xb8\x8b\xaf\xae\xff\xf3\x62\x22\x0e\x29\xea\x22\x71\x65\x51\xe0\xe9\xf5\xb6\xec\xa1\x09\x0b\xe4\x0d\xf6\x5f\xdf\x1c\xff\x79\x22\x4f\x41\xdd\xc5\xe2\x22\xa8\x6c\x85\x05\x4c\xc7\x82\xba\x3d\x23\x6e\x74\xf4\xdd\x70\xff\x9f\xc3\x2f\xf7\x87\x27\x9f\x38\xbf\x7b\x18\x59\xe2\xe4\xb5\x89\xc0\x22\xe0\xca\x62\x49\x10\x8b\xcb\x8e\x40\xba\xf3\x2c\x7b\x70\x3e\xdd\xb9\xd2\xa6\xba\x23\xd2\xe6\x74\xe7\x0a\xe3\x45\xbe\xc8\xb0\xa0\x48\x25\x8b\xe7\x3c\xb4\xdd\x0e\xe3\x13\x3b\x16\x38\xd1\x1e\x11\x38\xa7\xbd\x22\x10\x38\x51\x58\x16\x39\x72\x9c\x1c\x65\x6c\xba\x23\x65\x7b\x7d\xb8\xe8\x19\xdb\x5b\x48\x93\xb0\x63\x8e\x97\xe7\xb9\x5f\xbc\xa2\x06\xc8\x16\x32\x8d\xa4\x23\xe5\x34\xdb\x9f\x2e\xda\xa4\x07\x9a\xb8\xe5\x4e\xe0\xb6\xb1\xa1\x4b\xf1\xea\xb7\x94\xb8\xe6\x62\x13\x73\x9c\x37\xb2\x79\x39\x8b\xa8\x26\x3c\x09\x69\x60\x4c\x06\x98\x02\x69\x93\x49\x5c\xa0\x2d\x28\xd2\x67\x65\xd9\x95\xac\x54\x0a\xbf\x15\x24\x5d\x7d\x2f\xa4\x3b\x1e\xd1\x19\x9e\x7c\x72\x73\x7c\x38\xe5\x0c\x2c\xa1\x23\xac\x9f\x21\xf8\x59\x49\x81\x17\x0b\x45\x66\x3e\x1a\xcf\x0a\x12\xfc\x99\xdf\x95\xe0\x86\xfb\xc1\x8b\xaf\x27\xac\x80\x25\xe9\x87\xfc\xe5\xc9\x5f\x81\x25\x11\x9b\x9d\x4f\x3c\xe3\x4c\x57\x28\x31\x89\x7c\x86\x8c\x70\x22\x3b\x0d\x59\x95\xfd\x65\x9f\x0c\x47\x36\xe9\x4c\xc0\xb2\xbe\xff\x93\xcd\xf0\x26\xdb\x9b\x45\x34\x29\x2b\xd2\x00\x27\xd5\x74\xf8\xe3\x95\x9c\x45\xd8\xff\x7f\x57\x72\x64\x43\x97\x70\xaf\xd5\x42\x18\xc7\x1f\x0c\xb9\x30\x6d\xa8\x76\x7b\x49\xe6\x4b\x89\x07\x29\xe3\x51\x7a\x32\xd3\x8e\x71\xbd\x65\x07\x5a\xaa\x93\x16\xa8\x9a\x61\x89\xf8\x35\x48\x8d\x63\x83\x05\x2d\xa8\x04\x50\x31\x4e\xcd\xc9\xc9\xff\x13\xca\x98\xb8\x1f\x2d\x3b\x94\x25\xd9\x55\x1d\x5d\x6d\x3d\x49\xb6\x07\xb8\x03\x6d\x18\x9a\x6c\x82\x31\xdc\xe0\x27\x39\xa1\x29\x49\xab\x5b\x26\x54\x2e\x31\x4c\x06\x0d\x73\xc7\xa6\xa2\x52\x59\xd7\x0e\xfd\x54\x76\xc2\x30\x11\xcf\x41\xa5\x55\x28\x17\xcb\x4d\x58\x2a\xc3\x66\x13\x95\xee\x23\xae\xb4\xdc\x16\xb8\xf6\xca\x4a\x01\x91\x1e\x80\x8f\x7f\xba\x7d\x88\xda\xc6\x2f\x72\x4b\xee\x8f\xa5\x07\xf1\x39\xa5\xb3\x2f\x49\xc3\x9d\xc8\xb4\x72\xbc\x90\xba\xd4\x4f\xa1\x35\x2f\x6b\x29\x94\xeb\x51\xa6\xae\xc3\xa3\xb3\xab\xf3\xb3\xab\x7e\x3f\x74\xd2\x9a\xd6\xa0\x12\xc8\x16\x82\x8a\xad\x25\x86\x19\x73\xdd\x85\xbd\xf5\xf0\xd6\xba\xac\xa0\x2e\x22\x9f\x43\x00\x62\xaf\x4f\xec\xbf\x20\x7e\xdc\x97\x2c\x64\x1a\xf4\xb3\xa3\x98\x13\x9f\xc5\x6d\x24\xfa\xf5\x64\xa7\x77\x38\x48\x8c\x09\x08\xed\xd1\x50\xb4\x52\xd2\xd8\x9f\xec\x08\x22\x75\x51\x94\xe9\xbb\x13\xfd\x55\x4d\xb5\xa3\xad\x99\x3d\x8b\xf1\xbe\x29\x0f\xc4\x36\x9e\x20\xba\x41\x63\x7a\x93\x46\x2f\x18\xc5\xbd\xea\x34\x06\x99\xed\x14\xc6\xba\x66\x41\xe8\xb5\x89\xf4\xf0\xfb\xa1\x34\x48\x29\x71\xc7\x5b\x7c\x26\x3e\x4e\x0a\xac\x66\x8c\x44\xaa\x98\xb1\xb1\x4c\x6c\x33\x32\x8e\x85\xda\x16\xc2\x5b\xc9\x81\xd0\xae\xa9\x5a\x08\x4b\xaa\x7e\xb7\x1d\x51\x0c\x52\xa7\x40\x69\xc4\x6b\xff\x07\xa6\x03\x97\x12\xa7\xf4\xd0\xc8\xa9\x47\x8b\x2d\x8e\xab\x46\xbc\x42\x44\x15\x9d\xff\x1b\x00\x00\xff\xff\xfa\x11\xc3\x9c\xfb\x37\x00\x00")

func _000001_pipeline_up_sql() ([]byte, error) {
	return bindata_read(
		__000001_pipeline_up_sql,
		"000001_pipeline.up.sql",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"000001_pipeline.down.sql": _000001_pipeline_down_sql,
	"000001_pipeline.up.sql": _000001_pipeline_up_sql,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"000001_pipeline.down.sql": &_bintree_t{_000001_pipeline_down_sql, map[string]*_bintree_t{
	}},
	"000001_pipeline.up.sql": &_bintree_t{_000001_pipeline_up_sql, map[string]*_bintree_t{
	}},
}}
