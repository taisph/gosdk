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
	"os"
)

func removeSymlink(name string) error {
	fi, err := os.Lstat(name)
	if err != nil && os.IsNotExist(err) {
		return nil
	}

	if fi.Mode()&os.ModeSymlink == 0 {
		return fmt.Errorf("%s: not a symlink. Inspect and remove it manually to continue", name)
	}

	if err := os.Remove(name); err != nil {
		return err
	}

	return nil
}
