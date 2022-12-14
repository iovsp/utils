// Copyright 2021 utils. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package utils

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Resource struct {
	root string
}

var Public = Resource{
	root: filepath.Dir(os.Args[0]),
}

func (o *Resource) Root(dir string) string {
	if filepath.IsAbs(dir) {
		o.root = dir
	} else {
		o.root = o.Abs(dir)
	}
	return o.root
}

func (o *Resource) Abs(dirs ...string) string {
	public := filepath.Join(dirs...)
	return filepath.Join(o.root, public)
}

// JSONFile load json file
func JSONFile(filename string, obj interface{}) error {
	jfile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer jfile.Close()
	var jsString string
	iReader := bufio.NewReader(jfile)
	for {
		tString, err := iReader.ReadString('\n')
		if err == io.EOF {
			break
		}
		jsString = jsString + tString
	}
	return json.Unmarshal([]byte(jsString), obj)
}

// YamlFile load yaml file
func YamlFile(filename string, obj interface{}) error {
	yfile, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(yfile, obj)
}
