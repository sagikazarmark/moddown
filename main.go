package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// Module is a downloaded module.
type Module struct {
	Path     string // module path
	Version  string // module version
	Error    string // error loading module
	Info     string // absolute path to cached .info file
	GoMod    string // absolute path to cached .mod file
	Zip      string // absolute path to cached .zip file
	Dir      string // absolute path to cached source root directory
	Sum      string // checksum for path, version (as in go.sum)
	GoModSum string // checksum for go.mod (as in go.sum)
}

// nolint: gochecknoglobals
var (
	dest    = flag.String("dest", "", "destination directory")
	sum     = flag.String("sum", "", "hash of module contents")
	file    = flag.String("f", "", "module download file")
	verbose = flag.Bool("v", false, "enable verbose mode")
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("moddown: ")

	flag.Parse()

	if *file == "" && flag.NArg() != 1 {
		log.Fatal("expected module info or a module to download")
	}

	var moduleContent []byte

	// nolint: gocritic // if-else chain is more expressive
	if *file == "" {
		var err error

		moduleContent, err = downloadModule(flag.Arg(0), *verbose)
		if err != nil {
			log.Fatal(err)
		}
	} else if *file == "-" {
		var err error

		moduleContent, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		var err error

		moduleContent, err = ioutil.ReadFile(*file)
		if err != nil {
			log.Fatal(err)
		}
	}

	var module Module

	// Parse the JSON output
	if err := json.Unmarshal(moduleContent, &module); err != nil {
		log.Fatalf("parsing module: %v", err)
	}

	if module.Error != "" {
		log.Fatal(module.Error)
	}

	if *sum != "" && module.Sum != *sum {
		log.Fatalf("downloaded module with sum %s; expected sum %s", module.Sum, *sum)
	}

	if dest != nil {
		err := copyTree(*dest, module.Dir)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func downloadModule(mod string, verbose bool) ([]byte, error) {
	var buf, bufErr bytes.Buffer

	// nolint: gosec
	cmd := exec.Command(locateGoBinary(), "mod", "download", "-x", "-modcacherw", "-json", mod)
	cmd.Stdout = &buf
	cmd.Stderr = io.MultiWriter(os.Stderr, &bufErr)

	if verbose {
		log.Printf("executing %s with args %q", cmd.Path, strings.Join(cmd.Args, " "))
	}

	err := cmd.Run()
	if err != nil { // Check if the process exited unexpectedly
		if _, ok := err.(*exec.ExitError); !ok {
			return nil, fmt.Errorf("%s %s: %v", cmd.Path, strings.Join(cmd.Args, " "), err)
		}
	}

	return buf.Bytes(), nil
}

func locateGoBinary() string {
	if goCmd, ok := os.LookupEnv("TOOLS_GO"); ok {
		return goCmd
	}

	goCmd := "go"

	if runtime.GOOS == "windows" {
		goCmd += ".exe"
	}

	// If GOROOT is set, we'll use that one; otherwise, we'll use PATH.
	if goroot, ok := os.LookupEnv("GOROOT"); ok {
		goCmd = filepath.Join(goroot, "bin", goCmd)
	}

	return goCmd
}

// copyTree is based on fetch_repo source.
func copyTree(destRoot, srcRoot string) error {
	return filepath.Walk(srcRoot, func(src string, info os.FileInfo, e error) (err error) {
		if e != nil {
			return e
		}
		rel, err := filepath.Rel(srcRoot, src)
		if err != nil {
			return err
		}
		if rel == "." {
			return nil
		}
		dest := filepath.Join(destRoot, rel)

		if info.IsDir() {
			return os.Mkdir(dest, 0777)
		} else { // nolint: golint
			r, err := os.Open(src)
			if err != nil {
				return err
			}
			defer r.Close()
			w, err := os.Create(dest)
			if err != nil {
				return err
			}
			defer func() {
				if cerr := w.Close(); err == nil && cerr != nil {
					err = cerr
				}
			}()
			_, err = io.Copy(w, r)

			return err
		}
	})
}
