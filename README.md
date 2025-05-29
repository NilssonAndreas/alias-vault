```
  ___   _  _               _   _                _  _   
 / _ \ | |(_)             | | | |              | || |  
/ /_\ \| | _   __ _  ___  | | | |  __ _  _   _ | || |_ 
|  _  || || | / _` |/ __| | | | | / _` || | | || || __|
| | | || || || (_| |\__ \ \ \_/ /| (_| || |_| || || |_ 
\_| |_/|_||_| \__,_||___/  \___/  \__,_| \__,_||_| \__|
```

**AliasVault** — Terminal-based alias manager built in Go. Save, tag, run, and organize your shell commands with ease.

---

## 🚀 Features

- Save complex commands with a short alias and tags
- Run commands directly: `vault run "deploy staging"`
- Interactive TUI: browse, filter, run, add, delete
- Fuzzy search and live filtering
- Simple SQLite backend

---

## 🛠️ Installation

```bash
git clone https://github.com/yourname/alias-vault.git
cd alias-vault
go build -o vault
sudo mv vault /usr/local/bin/
```

Now you can run `vault` from anywhere in your terminal.

---

## 🧪 Usage

### Add a new alias (interactive)
```bash
vault add
```

Or non-interactive:
```bash
vault add "alias" "full command" --tags tag1,tag2
```

### Run a command
```bash
vault run "alias"
```

### Launch TUI
```bash
vault tui
```
- `↑↓` Navigate
- `/` Filter
- `a` Add
- `d` Delete
- `Enter` Run
- `q` Quit

### List aliases
```bash
vault list
```

### Delete alias
```bash
vault delete "alias"
```
