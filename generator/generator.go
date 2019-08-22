package generator

import (
	"github.com/trntv/bricks/sandbox"
	"os/exec"
)

func GenModules(dir string, name string) {
	cmd := exec.Command("go", "mod", "init", name)
	cmd.Dir = dir
	sandbox.RunWithEnvs(cmd)
}

func InstallModules(dir string) {
	cmd := exec.Command("go", "mod", "vendor")
	cmd.Dir = dir
	sandbox.RunWithEnvs(cmd)
}
