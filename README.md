# **GIT COMMANDS CHEAT SHEET**

---

## **Setup & Configuration**

| Command                                            | Description                          |
| -------------------------------------------------- | ------------------------------------ |
| `git config --global user.name "Your Name"`        | Set your global username             |
| `git config --global user.email "you@example.com"` | Set your global email                |
| `git config --global core.editor "code --wait"`    | Set default editor (VS Code example) |
| `git config --list`                                | Show all configs                     |
| `git help <command>`                               | Show help for any command            |

---

## **Repository Initialization**

| Command           | Description                          |
| ----------------- | ------------------------------------ |
| `git init`        | Create a new local Git repo          |
| `git clone <url>` | Clone a remote repo to your computer |

---

## **Basic Snapshotting**

| Command                                | Description                           |
| -------------------------------------- | ------------------------------------- |
| `git status`                           | See changed files                     |
| `git add <file>`                       | Stage a file for commit               |
| `git add .`                            | Stage all modified files              |
| `git commit -m "message"`              | Commit staged changes                 |
| `git commit -am "message"`             | Add + commit modified (tracked) files |
| `git diff`                             | Show unstaged changes                 |
| `git diff --staged`                    | Show staged changes                   |
| `git log`                              | View commit history                   |
| `git log --oneline --graph --decorate` | Pretty log view                       |

---

## **Branching & Merging**

| Command                  | Description                      |
| ------------------------ | -------------------------------- |
| `git branch`             | List branches                    |
| `git branch <name>`      | Create a new branch              |
| `git checkout <branch>`  | Switch branch                    |
| `git checkout -b <name>` | Create & switch to new branch    |
| `git merge <branch>`     | Merge branch into current one    |
| `git rebase <branch>`    | rebase the current branch into other one    |
| `git branch -d <branch>` | Delete branch                    |
| `git branch -D <branch>` | Force delete branch              |

---

## **Remote Repositories**

| Command                       | Description                       |
| ----------------------------- | --------------------------------- |
| `git remote -v`               | List remotes                      |
| `git remote add origin <url>` | Add remote connection             |
| `git fetch origin`            | Download new data (no merge)      |
| `git pull origin master`      | Fetch + merge changes from remote |
| `git push origin master`      | Upload commits to remote          |
| `git push -u origin <branch>` | Push & set default upstream       |
| `git push --force`            | Force push (after rebase)         |
| `git remote remove <name>`    | Remove a remote                   |

---

## **Undoing & Fixing**

| Command                        | Description                                    |
| ------------------------------ | ---------------------------------------------- |
| `git restore <file>`           | Undo local unstaged changes                    |
| `git restore --staged <file>`  | Unstage a file                                 |
| `git reset HEAD~1`             | Undo last commit but keep changes              |
| `git reset --hard HEAD~1`      | Delete last commit and changes                 |
| `git checkout <commit> <file>` | Restore a file from an older commit            |

---

## **Viewing & Comparing**

| Command                         | Description                 |
| ------------------------------- | --------------------------- |
| `git show <commit>`             | Show details of a commit    |
| `git blame <file>`              | Show who changed each line  |
| `git shortlog -sn`              | List contributors           |
| `git log --stat`                | Show commits + file changes |
| `git diff <branch1>..<branch2>` | Compare two branches        |

---

## **Clean Up**

| Command        | Description                        |
| -------------- | ---------------------------------- |
| `git clean -n` | Show untracked files to be deleted |
| `git clean -f` | Delete untracked files             |
| `git gc`       | Clean up unnecessary files in repo |
| `git prune`    | Remove unreachable objects         |

---

## **Advanced / Useful Extras**

| Command                                       | Description                                        |
| --------------------------------------------- | -------------------------------------------------- |
| `git reflog`                                  | View all recent branch movements (saves your life) |
| `git cherry-pick <commit>`                    | Apply a specific commit to another branch          |
| `git bisect`                                  | Find which commit introduced a bug                 |
| `git submodule add <url>`                     | Add a submodule                                    |
| `git archive --format=zip HEAD > project.zip` | Export your repo as a ZIP                          |
| `git alias`                                   | Create shortcuts for commands                      |


