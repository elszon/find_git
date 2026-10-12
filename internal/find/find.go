package find

import (
	"os"
	"path/filepath"
)

// GitRepos traverses the directory tree starting from root
// and returns a list of directories that contain .git folders
func GitRepos(root string) ([]string, error) {
	var gitRepos []string

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return filepath.SkipDir
		}

		// Skip symlinks to avoid infinite loops
		if d.IsDir() && isSymlink(d) {
			return filepath.SkipDir
		}

		// Check if this directory contains a .git folder
		if d.IsDir() {
			gitPath := filepath.Join(path, ".git")
			// Use Lstat instead of Stat (faster, doesn't follow symlinks)
			if stat, statErr := os.Lstat(gitPath); statErr == nil {
				// only the git repo has `.git` folder, for git worktree, the `.git` is a file
				if stat.IsDir() {
					relPath, _ := filepath.Rel(root, path)
					gitRepos = append(gitRepos, relPath)
				}
				return filepath.SkipDir
			}
		}

		return nil
	})

	return gitRepos, err
}

func isSymlink(d os.DirEntry) bool {
	return d.Type()&os.ModeSymlink != 0
}
