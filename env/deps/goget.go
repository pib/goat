package deps

import (
	"fmt"

	. "github.com/pib/goat/common"
	"github.com/pib/goat/exec"
)

func GoGet(depdir string, dep *Dependency) error {
	fmt.Println("go", "get", dep.Location)
	return exec.PipedCmd("go", "get", dep.Location)
}
