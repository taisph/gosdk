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
	"os/exec"
	"path/filepath"

	semver "github.com/hashicorp/go-version"

	"github.com/taisph/gosdk/internal/golangdlversion"
)

func set(version string) error {
	if version[0:2] != "go" {
		return fmt.Errorf("%s: expected prefix: go", version)
	}
	_, err := semver.NewSemver(version[2:])
	if err != nil {
		return err
	}

	root, err := golangdlversion.Goroot(version)
	if err != nil {
		return fmt.Errorf("get go sdk root: %s", err)
	}
	if _, err := os.Stat(filepath.Join(root, golangdlversion.UnpackedOkay)); err != nil {
		if err := download(version); err != nil {
			return fmt.Errorf("%s: not downloaded and failed to download: %s. Run 'go get golang.org/dl/%s' followed by '%s download' to install to %v", version, err, version, version, root)
		}
	}

	home, err := golangdlversion.Homedir()
	if err != nil {
		return err
	}

	homebin := filepath.Join(home, ".local/bin")
	if err := os.MkdirAll(homebin, os.ModePerm); err != nil {
		return err
	}

	dst := filepath.Join(homebin, "go")
	if err := removeSymlink(dst); err != nil {
		return err
	}
	if err := os.Symlink(filepath.Join(root, "bin/go"), dst); err != nil {
		return err
	}

	return nil
}

func download(version string) error {
	if err := run(exec.Command("go", "get", "golang.org/dl/" + version)); err != nil {
		return err
	}

	if err := run(exec.Command(version, "download")); err != nil {
		return err
	}

	return nil
}

func run(cmd *exec.Cmd) error {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
