// Copyright 2019 Tais P. Hansen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	cfg := struct {
		setVersion string
	}{}

	app := kingpin.New(filepath.Base(os.Args[0]), "Manages a symlink to the selected Go SDK version.")
	_ = app.Command("list", "List currently installed Go SDKs").Default()
	appset := app.Command("set", "")
	appset.Arg("version", "Go SDK version, e.g. "+runtime.Version()).Required().StringVar(&cfg.setVersion)

	app.Version("1.0.0-alpha")
	app.HelpFlag.Short('h')
	app.DefaultEnvars()

	log.SetFlags(0)
	log.SetPrefix(app.Name + ": ")

	cmd, err := app.Parse(os.Args[1:])
	if err != nil {
		log.Fatalf("parse commandline: %s", err)
	}

	switch cmd {
	case "list":
		err = list()
	case "set":
		err = set(cfg.setVersion)
	}
	if err != nil {
		log.Fatalf("%s: %s", cmd, err)
	}
}
