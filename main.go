package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

var (
	version   string // the current version
	gitRev    string // the git revision hash (set by the Makefile)
	buildDate string // the compile time
	repoURL   = "github.com/paulvollmer/tplate"

	envVarPath        string         // the TPLATE_PATH environment var
	envVarAuthor      string         // the TPLATE_AUTHOR environment var
	envVarAuthorEmail string         // the TPLATE_AUTHOR_EMAIL environment var
	envVarEditor      string         // the TPLATE_EDITOR environment var
	totalArgs         = len(os.Args) // get the total number of arguments
	err               error
)

func usage() {
	fmt.Println("USAGE: tplate [flags]")
	fmt.Println("\nFLAGS:")
	flag.PrintDefaults()
	fmt.Println("")
	fmt.Println("GETTING STARTED:")
	fmt.Println("  create a template by running")
	fmt.Println("    $ echo 'hello template' > $TPLATE_PATH/hello.tplate")
	fmt.Println("")
	fmt.Println("  calling a template file is easy to do by writing tplate and the name of the template")
	fmt.Println("    $ tplate hello")
	fmt.Println("")
	fmt.Println("  we need to edit our hello.tplate file to add a variable. open the file and change the content to")
	fmt.Println("    'hello {{.Var1}}'")
	fmt.Println("")
	fmt.Println("  now you can set the Var1 value over the terminal")
	fmt.Println("    $ tplate hello Var1=314")
	fmt.Println("")
	fmt.Println("  write templates to file")
	fmt.Println("    $ tplate hello -o myfile.src")
	fmt.Println("")
	fmt.Println("VERSION:")
	actionVersion()
	fmt.Println("")
	fmt.Println("ISSUE AND SUPPORT:")
	fmt.Printf("  Repository %s\n", repoURL)
	fmt.Printf("  Releases   %s\n", repoURL+"/releases")
	fmt.Printf("  Issues     %s\n", repoURL+"/issues")
	fmt.Println("")
	fmt.Printf("Copyright 2017 (c) Paul Vollmer\n\n")
}

func main() {
	// get the environment variables
	envVarPath = os.Getenv("TPLATE_PATH")
	envVarAuthor = os.Getenv("TPLATE_AUTHOR")
	if envVarAuthor == "" {
		envVarAuthor = "tplate"
	}
	envVarAuthorEmail = os.Getenv("TPLATE_AUTHOR_EMAIL")
	envVarEditor = os.Getenv("TPLATE_EDITOR")
	if envVarEditor == "" {
		envVarEditor = "atom"
	}

	// define and parse commandline flags
	flagEnv := flag.Bool("env", false, "the environment vars")
	flagInit := flag.Bool("init", false, "initialize the template directory")
	flagList := flag.Bool("list", false, "list all templates")
	flagHelp := flag.Bool("help", false, "print the help and exit")
	flagEdit := flag.String("edit", "", "open a template in your editor or create one if not exist")
	flagVersion := flag.Bool("version", false, "print the version and exit")
	flag.Usage = usage
	flag.Parse()

	if *flagEnv {
		actionEnv()
	} else if *flagInit {
		actionInit(envVarPath)
	} else if *flagList {
		actionList()
	} else if *flagHelp {
		actionHelp()
	} else if *flagEdit != "" {
		actionEdit(*flagEdit)
	} else if *flagVersion {
		actionVersion()
	} else {
		actionProcess()
	}
}

func actionEnv() {
	fmt.Printf("TPLATE_PATH         = %s\n", envVarPath)
	fmt.Printf("TPLATE_AUTHOR       = %s\n", envVarAuthor)
	fmt.Printf("TPLATE_AUTHOR_EMAIL = %s\n", envVarAuthorEmail)
}

func actionEdit(src string) {
	content := []byte("")
	fPath := path.Join(envVarPath, src+".tplate")
	// check if file exist

	if _, err = os.Stat(fPath); os.IsNotExist(err) {
		fmt.Printf("create new template %q\n", src)
		err = ioutil.WriteFile(fPath, content, 0644)
		if err != nil {
			fmt.Println("Create new template error:", err)
			os.Exit(127)
		}
	} else {
		fmt.Printf("template %q already exist\n", src)
	}

	// open editor
	cmd := exec.Command(envVarEditor, fPath)
	fmt.Printf("open file in %q editor...\n", envVarEditor)
	err = cmd.Run()
	if err != nil {
		fmt.Println("TERMINAL ERROR", err)
	}
}

func actionInit(src string) {
	if src == "" {
		fmt.Println("No TPLATE_PATH env var defined.")
		fmt.Println("")
		fmt.Println("  export TPLATE_PATH=$HOME/tplate")
		fmt.Println("")
		os.Exit(127)
	}
	// TODO: clone a tplate collection repository
	_, err = ioutil.ReadDir(src)
	if err != nil {
		err = os.Mkdir(src, 0644)
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
	listPath := envVarPath
	if totalArgs > 2 {
		fmt.Println(listPath)
		listPath = filepath.Join(listPath, os.Args[2])
	}
	fmt.Printf("list templates of directory %q\n\n", listPath)
	listFiles(listPath, "")
	fmt.Println("")
}

func listFiles(src, prefix string) {
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
				listFiles(src+string(filepath.Separator)+file.Name(), file.Name()+string(filepath.Separator))
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
	fmt.Printf("  Version    %s\n", version)
	fmt.Printf("  Revision   %s\n", gitRev)
	fmt.Printf("  Build-Date %s\n", buildDate)
}

func actionProcess() {
	flagArgs := flag.Args()
	if len(flagArgs) == 0 {
		fmt.Println("missing template source!")
		fmt.Println("check out the 'tplate -help' for more information")
		os.Exit(1)
	}

	command := flag.NewFlagSet("program", flag.ExitOnError)
	flagOut := command.String("o", "", "output path")
	command.Parse(flag.Args()[1:])

	// parse the commandline args and read file.
	tmpFilepath := filepath.Join(envVarPath, flagArgs[0]) + ".tplate"
	tmpVars := []string{}
	cmdArgs := command.Args()
	if len(cmdArgs) >= 1 {
		tmpVars = cmdArgs[0:]
	}
	// fmt.Println("VARS", tmpVars)
	result, err := Process(tmpFilepath, tmpVars)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	if *flagOut == "" {
		fmt.Print(string(result))
	} else {
		err = ioutil.WriteFile(*flagOut, result, 0644)
		if err != nil {
			return
		}
	}
}
