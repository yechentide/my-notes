# Git

## 目次

- [Git](#git)
    - [目次](#目次)
    - [初期設定](#初期設定)
    - [基本的なサブコマンド](#基本的なサブコマンド)
    - [変更と戻す系](#変更と戻す系)
    - [タグ関連](#タグ関連)

## 初期設定

必ずやっておく設定
```shell
git config --global user.name "開発者の名前"
git config --global user.email "開発者のメールアドレス"
```
他にも設定しておくと便利な項目がある
```shell
git config --global core.excludesfile ~/.gitignore_global   # globalなignoreファイルへの参照
git config --global init.defaultBranch main                 # デフォルトのブランチ名を変更
git config --global core.autocrlf false                     # Windows/macOS/Linuxの改行コードの自動変換
git config --global core.ignorecase false                   # ファイル名の大文字と小文字を区別する
git config --global color.ui true                           # 文字に色をつける
git config --global core.editor vim                         # デフォルトのエディターを変更
git config --global core.quotepath false                    # 日本語ファイル名をエンコードしない
git config --global fetch.prune true                        # 削除されたリモートブランチをローカルに自動的に適用する
echo .DS_Store >> ~/.gitignore_global                       # Mac専用設定
git config --global credential helper osxkeychain           # Mac専用設定: 毎回ユーザ名とトークンを入力しなくて済む。Windowsも似たような設定項目がある

### 設定の一覧
git config --global -l

### 設定の削除
# git config --global --unset user.name
```

## 基本的なサブコマンド

- `git init`: これでローカルリポジトリを作れる
    - `--initial-branch=branchName`でデフォルトブランチ名を指定できる
- `git clone`: リモートリポジトリをローカルにクローン
    - `git clone remoteRepoURL localPath`でクローン先の場所を指定できる
    - `git clone -l localRepoPath`でローカルリポジトリをクローンできる
- `git mv`: このコマンドでファイルの移動や名前変更ができる
    - GUIでの移動は、元のファイルを削除して新規ファイル作成という形になる
    - このサブコマンドを使えば、renamedという状態になるためおすすめ
    - `git mv --force file target`で上書きできる
- `git rm`: ファイルを削除する
    - GUIで削除する場合は、自分で`add`をする必要があるが、
    - このサブコマンドは削除と同時にやってくれる
    - `git rm -r directory`で、ディレクトリを削除
    - `git rm --cached file`で、ファイルをリポジトリのindexから外す(ファイルを削除しない)

---

- `git status`: リポジトリの状態を調べる
    - `-s`で、省略形の出力になる
    - `-sb`は、ブランチ情報も表示してくれる
- `git add`: 変更されたファイルをindexに追加
    - `git add PATH`
    - `-A`で、全てのファイルを追加(新規ファイルも含めて)
    - `-u`で、管理されているファイルのみ追加
    - `-f`で、無視するファイルも強制的に追加
- `git commit`: indexに追加された変更点をリポジトリに保存
    - `-m "メッセージ"`でコミットメッセージを書く
    - `-a`オプションをつけると、`git add -A`も一緒にやってくれる
    - `git commit --amend`は、新しい変更点を前のコミットに追加し、そのコミットメッセージを編集できる
- `git diff`: ファイルの変更箇所を確認できる
    - オプションなしの場合、unstaged, uncommitted changesを表示
    - `git diff HEAD`は、前回のコミットからの変更点を表示
    - `git diff --staged`で、indexに追加されて、まだコミットされていないファイルの変更点を表示
    - `git diff fileName`で、指定のファイルの変更点を表示
- `git log`: 履歴を見る
    - オプションなしの場合、全履歴が表示される
    - ファイルを指定して、そのファイルと関係ある履歴だけ見れる
    - `--oneline`で、一行の簡易履歴が表示される
    - `--graph`で、グラフで分岐なども表示される
    - `-i --grep`で、メッセージ検索
    - `-n`で、最後のn個の履歴を見れる
- `git reflog`: gitコマンドの実行履歴を見る
    - ブランチ名を指定できる
    - `-n`で、最後のn個だけ見る

---

- `git branch`: ブランチの一覧を確認できる
    - オプションなしの場合、ローカルブランチの一覧が表示される
    - `--all`または`-a`は、ローカルブランチ＆リモートブランチを表示できる
    - `-m`で、ブランチ名を変える
    - `-d`で、ブランチを削除
- `git checkout`: 機能はたくさんある...
    - オプションなしの場合、ブランチ切り替えになる
    - `-b`で、新規ブランチ作成して切り替える、分岐元も指定できる
    - `git checkout -`で、前のブランチに戻る
    - `git checkout -- fileName`で、`add`前の変更を取り消す
- `git switch`: ブランチを切り替える。git 2.23から使える。
    - オプションなしの場合は、ブランチ切り替えになる
    - `--create`または`-c`で、新規ブランチ作成して切り替える
    - `git switch -`で、前のブランチに戻る
- `git merge`: ブランチを統合
    - `git merge targetBranch`は、targetBranchを今のブランチに統合
    - `-e`をつけると、マージメッセージを編集できる
    - `git merge --abort`で、**コンフリクト時**にマージをやめる

---

- `git remote`: リモートリポジトリの追加＆削除
    - `-v`で、リモート一覧
    - `git remote show remoteName`で、詳しい情報を見る
    - `git remote add remoteName URL`で、新しいリモートを追加
    - `git remote remove remoteName`で、リモートを削除
    - `git remote rename oldName newName`で、リモート名を変更
- `git push`: ローカルの変更をリモートに反映する
    - `-u`で、upstream repositoryを設定できる
    - upstream repositoryを設定指定れば、`git push`だけでよい
- `git fetch`: リモート参照ブランチの中身を最新にする
    - リモート名で指定可能、(upstream repository設定指定れば)省略も可能
    - `git fetch --prune`で余計なリモート参照ブランチを削除
- `git pull`: ローカルの内容をリモートと同じにする

## 変更と戻す系

- reflog
- reset
- rebase
- revert

## タグ関連

[Gitのtagを理解する](https://qiita.com/k-penguin-sato/items/c62b47dd79f144c68dad)
