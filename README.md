# tplate [![Build Status](https://travis-ci.org/paulvollmer/tplate.svg?branch=master)](https://travis-ci.org/paulvollmer/tplate)


## What is tplate
A commandline template tool. The main goal of `tplate` is simplicity and customisation.  


## Usage
You can run...

    tplate your/template

and the tool print out the selected template.  
You can also pass variables to the template by using a `key=value` syntax like this...

    tplate your/template foo=5 bar=7


## How it works
`tplate` scan the directory set by the `TPLATE_PATH` environment variable
(this is useful to shortcut the template path you need to type).  
The template engine is the core golang text/template and you can find more usage infos at the [docs](https://golang.org/pkg/text/template).


## Installation
You can download the latest release from the [GitHub Release Page](http://github.com/paulvollmer/tplate)
and put it into your `$PATH`. Also you can install it from source. Simple run...

    go get github.com/paulvollmer/tplate
    cd $GOPATH/github.com/paulvollmer/tplate
    make install


## License
Licensed under [MIT-License](LICENSE).
