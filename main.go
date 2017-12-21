package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
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
	flagInit := flag.Bool("init", false, "initialize the template directory")
	flagList := flag.Bool("list", false, "list all templates")
	flagHelp := flag.Bool("help", false, "print the help and exit")
	flagOutpath := flag.String("o", "", "output path")
	flagVersion := flag.Bool("version", false, "print the version and exit")
	flag.Parse()

	if *flagEnv {
		actionEnv()
	} else if *flagInit {
		actionInit()
	} else if *flagList {
		actionList()
	} else if *flagHelp {
		actionHelp()
	} else if *flagVersion {
		actionVersion()
	} else {
		processAction()
	}
}

func actionEnv() {
	fmt.Printf("TPLATE_PATH         = %s\n", EnvVarPath)
	fmt.Printf("TPLATE_AUTHOR       = %s\n", EnvVarAuthor)
	fmt.Printf("TPLATE_AUTHOR_EMAIL = %s\n", EnvVarAuthorEmail)
}

func actionInit() {
	if EnvVarPath == "" {
		fmt.Println("No TPLATE_PATH defined")
		os.Exit(127)
	}
	// TODO: clone a tplate collection repository
	_, err := ioutil.ReadDir(EnvVarPath)
	if err != nil {
		err = os.Mkdir(EnvVarPath, 0644)
		if err != nil {
			fmt.Println("ERROR", err)
			os.Exit(127)
		}
		fmt.Println("created 'tplate' directory")
		fmt.Println("now you can start creating a template file you can call with the cli")
		fmt.Println("")
		fmt.Println("  echo 'hello world' > $TPLATE_PATH/example.tplate")
		fmt.Println("  tplate example")
		fmt.Println("")
		os.Exit(128)
	}
}

func actionList() {
	listPath := EnvVarPath
	if totalArgs > 2 {
		fmt.Println(listPath)
		listPath = filepath.Join(listPath, os.Args[2])
	}
	fmt.Printf("list templates of directory %q\n\n", listPath)
	ListFiles(listPath, "")
	fmt.Println("")
}

func ListFiles(src, prefix string) {
	tmp, err := ioutil.ReadDir(src)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, file := range tmp {
		name := file.Name()
		if file.IsDir() {
			// ignore some dirs
			if name != ".git" {
				ListFiles(src+string(filepath.Separator)+file.Name(), file.Name()+string(filepath.Separator))
			}
		} else {
			ext := path.Ext(name)
			if ext == ".tplate" {
				result := strings.Split(file.Name(), ext)
				fmt.Println(prefix + result[0])
			}
		}
	}
}

func actionHelp() {
	flag.Usage()
}

func actionVersion() {
	fmt.Printf("Version    : %s\n", version)
	fmt.Printf("Revision   : %s\n", gitRev)
	fmt.Printf("Build-Date : %s\n", buildDate)
}

func processAction() {
	// parse the commandline args and read file.
	if totalArgs == 1 {
		fmt.Println("missing template source!")
		fmt.Println("check out the 'tplate -help' for more information")
		os.Exit(1)
	} else {
		tmpFilepath := filepath.Join(EnvVarPath, os.Args[1]) + ".tplate"
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
