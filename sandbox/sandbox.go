package sandbox

import "os/exec"

func RunWithEnvs(cmd *exec.Cmd) error {
	cmd.Env = []string{"GOMAXPROCS=1", "GO111MODULE=on", "GOOS=nacl", "CGO_ENABLED=0"}

	return cmd.Run()
}
