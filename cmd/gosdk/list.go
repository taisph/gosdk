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
	"fmt"
	"io/ioutil"
	"sort"

	semver "github.com/hashicorp/go-version"
	"github.com/taisph/gosdk/internal/golangdlversion"
)

func list() error {
	root, err := golangdlversion.Goroot("")
	if err != nil {
		return fmt.Errorf("get go sdk root: %s", err)
	}

	fs, err := ioutil.ReadDir(root)
	if err != nil {
		return fmt.Errorf("%s: read dir: %s", root, err)
	}

	var coll semver.Collection
	for _, f := range fs {
		v, err := semver.NewSemver(f.Name()[2:])
		if err != nil {
			return err
		}
		coll = append(coll, v)
	}
	sort.Sort(sort.Reverse(coll))

	for _, v := range coll {
		fmt.Printf("go%s\n", v.String())
	}

	return nil
}
