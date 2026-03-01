package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

const version = "v0.0.1"
const clientName = "ai-commit-tracker"

var (
	repo   string
	author string
	info   bool
)

func init() {
	flag.StringVar(&repo, "repo", "", "repository URL to analyze (required)")
	flag.StringVar(&author, "author", "claude", "author name to search")
	flag.BoolVar(&info, "version", false, "show version")
}

func main() {
	flag.Parse()

	if info {
		fmt.Printf("Application Name: %s\n", clientName)
		fmt.Printf("Version: %s\n", version)
		return
	}

	if strings.TrimSpace(repo) == "" {
		log.Fatal("--repo is required")
	}

	repoName := extractRepoName(repo)
	tmpDir, err := os.MkdirTemp("", clientName+"-*")
	if err != nil {
		log.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	clonePath := filepath.Join(tmpDir, repoName)
	if err := cloneRepo(repo, clonePath); err != nil {
		log.Fatalf("failed to clone repository: %v", err)
	}

	count, err := countAuthorCommits(clonePath, author)
	if err != nil {
		log.Fatalf("failed to count commits: %v", err)
	}

	fmt.Printf("リポジトリ: %s\n", repoName)
	fmt.Printf("AIコミット数 (author: %s): %d\n", author, count)
}

func cloneRepo(repoURL, dst string) error {
	cmd := exec.Command("git", "clone", repoURL, dst)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git clone failed: %w: %s", err, strings.TrimSpace(string(output)))
	}
	return nil
}

func countAuthorCommits(repoPath, author string) (int, error) {
	cmd := exec.Command(
		"git",
		"log",
		"--all",
		"--regexp-ignore-case",
		"--author="+author,
		"--oneline",
	)
	cmd.Dir = repoPath
	output, err := cmd.CombinedOutput()
	if err != nil {
		return 0, fmt.Errorf("git log failed: %w: %s", err, strings.TrimSpace(string(output)))
	}

	trimmed := strings.TrimSpace(string(output))
	if trimmed == "" {
		return 0, nil
	}
	return len(strings.Split(trimmed, "\n")), nil
}

func extractRepoName(repoURL string) string {
	trimmed := strings.TrimSuffix(strings.TrimSpace(repoURL), "/")
	base := path.Base(trimmed)
	base = strings.TrimSuffix(base, ".git")
	if base == "." || base == "/" || base == "" {
		return "repository"
	}
	return base
}
