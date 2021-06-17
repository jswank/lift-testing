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

func fail_G304p2() {
	repoFile := "/safe/path/../../private/path"
	byContext, err := ioutil.ReadFile(repoFile)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", string(byContext))
}

func fail_G304p3(fname string) {
	byContext, err := ioutil.ReadFile(fname)
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
