# Zing 🚀

**Zing helps developers zing through repetitive commands.**

---

## 🧠 One-liner

Automate your most frequent CLI workflows with human-friendly aliases, templates, and project-aware suggestions — stay in flow and ship faster.

---

## 💡 Why Zing?

- **Kill repetition:** Turn 120-character shell incantations into `zing deploy prod`.
- **Stay consistent:** Share the same commands across your team and CI.
- **Project-aware:** Zing detects tech stacks (Node, Go, Docker, K8s) and suggests relevant tasks.
- **Portable:** Works across macOS, Linux, and Windows (WSL).

---

## ⚙️ Key Features

- **Aliases & templates:** Parameterized commands with variables and prompts.
- **Workflows:** Chain multiple steps with error handling and conditional logic.
- **Context detection:** Autoloads `.zing.yml` in a repo; global fallbacks in `~/.zing`.
- **Dry-run & preview:** See exactly what will execute before it runs.
- **History & re-run:** Fuzzy search your last used tasks.
- **Shell-native:** Works in bash/zsh/fish/PowerShell; autocompletion included.
- **Extensible:** Plugin hooks in any language (exec via stdin/stdout).

---

## 🚀 Quick Start

### Install (Go example)
```bash
go install github.com/yourorg/zing/cmd/zing@latest
```

### check version
```sh
zing version
```

### Add a task interactively
```bash
zing add --tag "deploy" --cmd "docker compose build && docker compose push && kubectl rollout restart deploy/<variable1> -n <variable2>"
```

## Run it (with prompts for variables)
```
zing run deploy variable1=kha variable2=gan
```


## CLI Cheatsheet

```sh
zing add --tag <name> ----cmd "<command>"   # add/append a task interactively
zing run <name> [args]    # run a task or workflow
zing list                 # list tasks with descriptions
zing preview <name>       # render command with variables (no exec)
zing suggest              # generate project-aware task suggestions (Coming Soon)
zing history              # browse commands (Coming Soon)
zing shell                # start an interactive TUI picker
```
