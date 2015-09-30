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

	totalArgs = len(os.Args) // get the total number of arguments
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
	flagHelp := flag.Bool("help", false, "print the help and exit")
	flagVersion := flag.Bool("version", false, "print the version and exit")
	flag.Parse()

	if *flagEnv {
		envAction()
	} else if *flagList {
		listAction()
	} else if *flagHelp {
		helpAction()
	} else if *flagVersion {
		versionAction()
	} else {
		processAction()
	}
}

func envAction() {
	fmt.Printf("TPLATE_PATH         = %s\n", EnvVarPath)
	fmt.Printf("TPLATE_AUTHOR       = %s\n", EnvVarAuthor)
	fmt.Printf("TPLATE_AUTHOR_EMAIL = %s\n", EnvVarAuthorEmail)
}

func listAction() {
	listPath := EnvVarPath
	if totalArgs > 2 {
		listPath = listPath + os.Args[2]
		fmt.Println(listPath)
	}
	fmt.Printf("list templates of directory '%s'\n", listPath)
	fmt.Printf("---------------------------\n")
	ListFiles(listPath, "")
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

func helpAction() {
	flag.Usage()
}

func versionAction() {
	fmt.Printf("Version    : %s\n", version)
	fmt.Printf("Revision   : %s\n", gitRev)
	fmt.Printf("Build-Date : %s\n", buildDate)
}

func processAction() {
	// parse the commandline args and read file.
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
