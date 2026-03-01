# AI Commit Tracker

AI Commit Trackerは、指定されたGitリポジトリ内で、特定の作者（AIアシスタントなど）によって作成されたコミット数を集計するためのシンプルなCLIツールです。

## 機能

- 公開されている任意のGitリポジトリを分析します。
- 特定の作者名（文字列で指定）が含まれるコミットを検索します。
- リポジトリ名と、該当するコミットの総数を表示します。

## インストール

Go言語の環境がセットアップされていることを確認してください。

```bash
# このリポジトリをクローンします（もしあれば）
git clone https://github.com/your-username/ai-commit-tracker.git
cd ai-commit-tracker
```

## ビルド

```bash
go build -o ai-commit-tracker
```

## 使い方

基本的なコマンド構造は以下の通りです。

```bash
./ai-commit-tracker --repo <リポジトリのURL> [--author=<作者名>]
```

- `--repo <リポジトリのURL>`: (必須) 分析したいGitリポジトリのURL。
- `--author=<作者名>`: (任意) 検索したい作者名。デフォルトは `claude` です。
- `--version`: バージョン情報を表示します。

### 例

```bash
# デフォルトの作者 "claude" で検索
./ai-commit-tracker --repo https://github.com/anthropics/claude-code

# "copilot" という名前の作者で検索
./ai-commit-tracker --repo https://github.com/some/repository --author="copilot"
```

### 出力例

```
リポジトリ: claude-code
AIコミット数 (author: claude): 5
```
