# tplate [![Build Status](https://travis-ci.org/paulvollmer/tplate.svg?branch=master)](https://travis-ci.org/paulvollmer/tplate)

tplate is a commandline templating tool. The main goal of **tplate** is simplicity and customisation. If you find yourself typing the same things over and over again when starting a new project, you will save a lot of time using **tplate**. 

## Usage

### Simple Template

First create a new template-file. Let’s create a template for the [go](https://golang.org/) programming language:  

Create a new directory where you store your templates:

    mkdir ~/templates
    
Then create a new file `go.tplate` and put some content in it:    

```go
package main

func main() {
    
}
```

Now you can run

    tplate ~/templates/go > my_new_go_project.go

and tplate will create a new file `my_new_go_project.go` with the contents of your template.
    
_Note: you only have to specify the filename-prefix, in this case `go`, **tplate** will then look for a file called `go.tplate`_


### Template with variables
    
You can also pass variables to the template by using a `key=value` syntax, **tplate** will then replace every occurence of `{{.foo}}` with `bar`. If your template `hello.tplate` looks like this:

```
{{.foo}} {{.bar}}
```

and you run 

    tplate hello foo="Hello" bar="Template"

You will see:  

    Hello Template

    
### Template path / environment variable

You can either specify the full path of your template, e.g. `~/templates/mytemplate.tplate` or just specify the template name and **tplate** tries to find it in your `TPLATE_PATH`-directory. To set the template-path, edit the file `~/.bash_profile` and add the following:

    export TPLATE_PATH="/Users/yourname/templates/"

*Note: After editing the file `~/.bash_profile` restart Terminal or run `source ~/.bash_profile`.*

After you set up your `TPLATE_PATH`, you can simply use the template name like this:

    tplate mytemplate > mynewproject.xyz
    
… and **tplate** will look for the template `/Users/yourname/templates/mytemplate`.


## How it works

The template engine is the core golang text/template.You can find more information about it here: [text/template package documentaion](https://golang.org/pkg/text/template).


## Installation

You can download the latest release from the [GitHub Release Page](http://github.com/paulvollmer/tplate/releases). Put the contents of the zip-file somewhere in you `$PATH`, e.g. `/usr/local/bin`.

You can also install it from source, simply run:

    go get github.com/paulvollmer/tplate
    cd $GOPATH/src/github.com/paulvollmer/tplate
    make install


## License

Licensed under [MIT-License](LICENSE).
