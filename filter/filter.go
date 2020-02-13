package filter

import (
	"github.com/google/go-github/v29/github"
)

// RepoFilter filter function based on a repository
type RepoFilter func(github.Repository) bool

// Repositories filters a slice of Repositories using the given Filter
func Repositories(repos []*github.Repository, filter RepoFilter) []github.Repository {
	filtered := make([]github.Repository, 0)
	for _, r := range repos {
		if r != nil {
			repo := *r
			if filter(repo) {
				filtered = append(filtered, repo)
			}
		}
	}
	return filtered
}
