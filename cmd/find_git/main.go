package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// findGitRepos traverses the directory tree starting from root
// and returns a list of directories that contain .git folders
func findGitRepos(root string) ([]string, error) {
	var gitRepos []string

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			// Skip directories we can't access
			return filepath.SkipDir
		}

		// Check if this directory contains a .git folder
		if d.IsDir() {
			gitPath := filepath.Join(path, ".git")
			_, err := os.Stat(gitPath)
			if err == nil {
				gitRepos = append(gitRepos, path)
				return filepath.SkipDir
			}
		}

		return nil
	})

	return gitRepos, err
}

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

	gitRepos, err := findGitRepos(absRoot)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error traversing directories: %v\n", err)
		os.Exit(1)
	}

	for _, repo := range gitRepos {
		// Make the path relative to the search root for cleaner output
		relPath, err := filepath.Rel(absRoot, repo)
		if err != nil {
			relPath = repo
		}
		fmt.Println(relPath)
	}
}
