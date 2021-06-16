package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func fail_G304() {
	repoFile := "/safe/path/../../private/path"
	if !strings.HasPrefix(repoFile, "/safe/path/") {
		panic(fmt.Errorf("Unsafe input"))
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
		panic(fmt.Errorf("Unsafe input"))
	}
	byContext, err := ioutil.ReadFile(repoFile)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", string(byContext))

}
