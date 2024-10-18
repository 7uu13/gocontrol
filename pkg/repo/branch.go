package repo

type Branch struct {
	Name       string `json:"name"`
	CommitHash string `json:"commit_hash"`
}

var CurrentBranch string = "main"
var Branches = make(map[string]Branch)