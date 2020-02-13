package filter

import (
	"strings"

	"github.com/google/go-github/v29/github"
)

// RepoFilter filter function based on a repository
type RepoFilter func(github.Repository) bool

var (
	// ByName filters a repository by name
	ByName = func(name string) RepoFilter {
		return func(r github.Repository) bool {
			return *r.Name == name
		}
	}
	// ByLanguage filters a repository by the main programming language (case insensitive)
	ByLanguage = func(language string) RepoFilter {
		return func(r github.Repository) bool {
			return strings.ToLower(r.GetLanguage()) == strings.ToLower(language)
		}
	}
)

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
