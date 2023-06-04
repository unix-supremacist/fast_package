package main

import (
	"encoding/json"
	"fmt"
	"flag"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

type Package struct {
	Name string `json:"name"`
	VER string `json:"version"`
	CMD string `json:"build-cmd"`
	BIN []struct {
		LOC string `json:"location"`
		DEST string `json:"destination"`
	} `json:"build_binaries"`
	PKGS []struct {
		Name string `json:"name"`
		VER string `json:"version"`
	} `json:"required_pkgs"`
}

func main() {
	path := flag.String("p", "", "testhelp") //default flags should not exist? weird shit in golang.
	flag.Parse()
	installPkg(*path)
}

func parsePkg(path string) Package {
	var pkg Package
	jso, err := ioutil.ReadFile(filepath.Join(path, "package.json"))
	eror(err)
	eror(json.Unmarshal([]byte(jso), &pkg))
	return pkg
}

func installPkg(path string){
	pkg := parsePkg(path)
	fmt.Println(pkg)
	bld := exec.Command("sh", "-c", pkg.CMD)
	bld.Dir = path
	eror(bld.Run())
	linkPkg(path, pkg)
}

func linkPkg(path string, pkg Package){
	for _, bin := range pkg.BIN {
		eror(os.Symlink(filepath.Join(path, bin.LOC), bin.DEST))
	}
}

func eror(err error) {
	if err != nil {
		fmt.Println(err)
	}
}