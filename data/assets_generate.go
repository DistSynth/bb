// +build ignore

package main

import (
	"log"

	"ab.ru/bigbrother/data/admin"
	"ab.ru/bigbrother/data/js"
	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(js.Assets, vfsgen.Options{
		Filename:     "data/js/assets_notdev.go",
		PackageName:  "js",
		BuildTags:    "!dev",
		VariableName: "Assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
	err = vfsgen.Generate(admin.Assets, vfsgen.Options{
		Filename:     "data/admin/assets_notdev.go",
		PackageName:  "admin",
		BuildTags:    "!dev",
		VariableName: "Assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
