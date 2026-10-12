package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/elszon/find_git/internal/find"
)

func main() {
	root, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting current working directory: %v\n", err)
		os.Exit(1)
	}

	absRoot, err := filepath.Abs(root)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error resolving path %s: %v\n", root, err)
		os.Exit(1)
	}

	gitRepos, err := find.GitRepos(absRoot)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error traversing directories: %v\n", err)
		os.Exit(1)
	}

	for _, repo := range gitRepos {
		// Make the path relative to the search root for cleaner output
		relPath, relErr := filepath.Rel(absRoot, repo)
		if relErr != nil {
			relPath = repo
		}
		fmt.Println(relPath)
	}
}
