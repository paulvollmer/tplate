package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	version   string // the current version
	gitRev    string // the git revision hash (set by the Makefile)
	buildDate string // the compile time

	EnvVarPath        string // the TPLATE_PATH environment var
	EnvVarAuthor      string // the TPLATE_AUTHOR environment var
	EnvVarAuthorEmail string // the TPLATE_AUTHOR_EMAIL environment var
)

func main() {
	// get the environment variables
	EnvVarPath = os.Getenv("TPLATE_PATH")
	EnvVarAuthor = os.Getenv("TPLATE_AUTHOR")
	if EnvVarAuthor == "" {
		EnvVarAuthor = "tplate"
	}
	EnvVarAuthorEmail = os.Getenv("TPLATE_AUTHOR_EMAIL")

	// define and parse commandline flags
	flagEnv := flag.Bool("env", false, "the environment vars")
	flagList := flag.Bool("list", false, "list all templates")
	flagHelp := flag.Bool("help", false, "show the help")
	flagVersion := flag.Bool("version", false, "print the version")
	flag.Parse()

	if *flagEnv {
		PrintEnvVars()
	} else if *flagList {
		fmt.Printf("list directory '%s'\n", EnvVarPath)
		ListFiles(EnvVarPath, "")
	} else if *flagHelp {
		Help()
	} else if *flagVersion {
		fmt.Printf("version: %s rev: %s build-date: %s\n", version, gitRev, buildDate)
	} else {
		// parse the commandline args and read file.
		totalArgs := len(os.Args)
		if totalArgs == 1 {
			fmt.Println("missing template source!")
			fmt.Println("check out the tplate -help for more information")
			os.Exit(1)
		} else {
			tmpFilepath := EnvVarPath + os.Args[1] + ".tplate"
			tmpVars := []string{}
			if totalArgs > 2 {
				tmpVars = os.Args[2:]
			}
			result, err := Process(tmpFilepath, tmpVars)
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
			fmt.Print(string(result))
		}
	}
}

func PrintEnvVars() {
	fmt.Printf("TPLATE_PATH         = %s\n", EnvVarPath)
	fmt.Printf("TPLATE_AUTHOR       = %s\n", EnvVarAuthor)
	fmt.Printf("TPLATE_AUTHOR_EMAIL = %s\n", EnvVarAuthorEmail)
}

func ListFiles(src, prefix string) {
	tmp, err := ioutil.ReadDir(src)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, file := range tmp {
		if file.IsDir() {
			ListFiles(src+string(filepath.Separator)+file.Name(), file.Name()+string(filepath.Separator))
		} else {
			fmt.Println(prefix + file.Name())
		}
	}
}

func Help() {
	flag.Usage()
}
