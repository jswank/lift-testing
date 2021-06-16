package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// See https://securego.io/docs/rules/g304.html

func fail_G304() {
	repoFile := "/safe/path/../../private/path"
	if !strings.HasPrefix(repoFile, "/safe/path/") {
		panic(fmt.Errorf("unsafe input"))
	}
	byContext, err := ioutil.ReadFile(repoFile)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", string(byContext))
}

func pass_G304() {
	repoFile := "/safe/path/../../private/path"
	repoFile = filepath.Clean(repoFile)
	if !strings.HasPrefix(repoFile, "/safe/path/") {
		panic(fmt.Errorf("unsafe input"))
	}
	byContext, err := ioutil.ReadFile(repoFile)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", string(byContext))

}
