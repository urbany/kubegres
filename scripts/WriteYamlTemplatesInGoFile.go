/*
Copyright 2021 Reactive Tech Limited.
"Reactive Tech Limited" is a company located in England, United Kingdom.
https://www.reactive-tech.io

Lead Developer: Alex Arica

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/*
 This script is executed every time this project is compiled.
 See main.go where there is a comment triggering this script: "//go:generate go run scripts/WriteYamlTemplatesInGoFile.go"
 It copies each YAML file inside the package "controllers/template/yaml"
 and set them as constants in the file "controllers/template/yaml/Templates.go".
 This mechanism allows developers to work with YAML files when defining templates for resources managed by this operator.
 It is more convenient to work directly with YAML files than GO API classes when developing.
*/

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

const (
	yamlTemplateDir   = "controllers/spec/template/yaml"
	destinationGoFile = yamlTemplateDir + "/Templates.go"
)

func main() {

	fs, _ := ioutil.ReadDir(yamlTemplateDir)
	out, _ := os.Create(destinationGoFile)

	out.Write([]byte("package yaml " +
		"\n\n // This file is auto generated by the script in 'WriteYamlTemplatesInGoFile.go'. " +
		"\n // Any manual modification to this file will be lost during next compilation. " +
		"\n\nconst (\n"))

	fmt.Println("Setting constants in the file: '" + destinationGoFile + "', by copying the YAML contents of the following files:")

	for _, f := range fs {

		if strings.HasSuffix(f.Name(), ".yaml") {

			templateYamlFilePath := yamlTemplateDir + "/" + f.Name()
			fmt.Println("- '" + templateYamlFilePath + "'")

			out.Write([]byte(strings.TrimSuffix(f.Name(), ".yaml") + " = `"))
			f, _ := os.Open(templateYamlFilePath)
			io.Copy(out, f)
			out.Write([]byte("`\n"))
		}
	}
	out.Write([]byte(")\n"))
}