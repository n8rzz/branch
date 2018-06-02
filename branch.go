package branch

import (
	"log"
	"os/exec"
	"strings"
)

// ReadLocalToList reads local git branches and returns a simple array of local git branch names
func ReadLocalToList() []string {
	out, err := exec.Command("git", "branch").Output()

	if err != nil {
		log.Fatalf("cmd.Output() failed with %s\n", err)
	}

	str := string(out)
	trimmedStr := strings.TrimRight(str, "\n")
	localGitBranchList := strings.Split(trimmedStr, "\n")

	return localGitBranchList
}
