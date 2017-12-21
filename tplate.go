package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"strings"
	"text/template"
)

// Process is the main tplate file processor
func Process(filename string, vars []string) ([]byte, error) {
	// read the template file
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return []byte(""), err
	}

	// store the template data as slice.
	tmpVars := make(map[string]string)
	// get key/values and set it to the template exec
	if len(vars) > 0 {
		tmpVars, err = parseData(vars)
		if err != nil {
			return []byte(""), err
		}
	}

	resultSource, err := processTemplate(dat, tmpVars)
	return resultSource, err
}

func parseData(vars []string) (map[string]string, error) {
	result := make(map[string]string)

	for i := 0; i < len(vars); i++ {
		splitted := strings.Split(vars[i], "=")
		if len(splitted) <= 1 {
			return result, errors.New("template data " + vars[i] + " not correct")
		}
		if result[splitted[0]] == "" {
			result[splitted[0]] = strings.Join(splitted[1:], "=")
		} else {
			result[splitted[0]] += "\n" + strings.Join(splitted[1:], "=")
		}
	}
	return result, nil
}

// source store the template source as byte array.
func processTemplate(source []byte, vars map[string]string) ([]byte, error) {
	tmpl, err := template.New("tpl").Parse(string(source))
	if err != nil {
		return []byte(""), err
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, vars)
	if err != nil {
		return []byte(""), err
	}
	return buf.Bytes(), nil
}
