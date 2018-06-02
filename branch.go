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

	rawStr := string(out)
	str := trimBranchStr(rawStr)
	localGitBranchList := strings.Split(str, "\n")

	return localGitBranchList
}

// ReadLocalToListWithoutCurrentActive returns
func ReadLocalToListWithoutCurrentActive() []string {
	branchList := ReadLocalToList()

	// Find and remove current active branch from list
	for i, v := range branchList {
		if strings.Contains(v, "*") {
			branchList = append(branchList[:i], branchList[i+1:]...)
			break
		}
	}

	return branchList
}

// trimBranchStr accepts raw output from `git branch` command
// this function will remove any leading or trailing spaces
func trimBranchStr(rawBranchList string) string {
	trimmedSpaceStr := strings.TrimSpace(rawBranchList)
	str := strings.TrimRight(trimmedSpaceStr, "\n")

	return str
}
