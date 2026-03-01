# AI Commit Tracker

AI Commit Tracker is a simple CLI tool that analyzes a Git repository and counts commits made by a specific author (for example, an AI assistant).

## Features

- Analyzes any public Git repository.
- Filters commits by author name.
- Prints the repository name and the total matching commit count.

## Requirements

- Go (1.20+ recommended)
- `git` command available in your environment

## Build

```bash
go build -o ai-commit-tracker
```

## Usage

```bash
./ai-commit-tracker --repo <repository-url> [--author=<author-name>]
```

Options:

- `--repo <repository-url>`: Required. The target Git repository URL.
- `--author=<author-name>`: Optional. Author string to match. Default: `claude`.
- `--version`: Show application name and version.

## Examples

```bash
# Search with default author: "claude"
./ai-commit-tracker --repo https://github.com/ghostty-org/ghostty

# Search with a custom author
./ai-commit-tracker --repo https://github.com/some/repository --author="copilot"
```

## Output Example

```text
Repository: ghostty
AI commit count (author: claude): 5
```

Note: The current implementation prints Japanese labels in output:

```text
リポジトリ: <repo>
AIコミット数 (author: <author>): <count>
```
