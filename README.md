# 🧴 lotion

*A CLI tool for smooth note-taking, inspired by Notion, but with a little more... lotion.*

Manage your personal notes from the terminal and sync them with GitHub (or any git remote). Lotion keeps your notes silky and organized, so you can focus on what matters.

---

## Features

- **Create** new notes in organized notebooks
- **List** all your notes at a glance
- **Search** and preview notes with your favorite tools
- **Sync** your notes with a remote git repository (like GitHub)

---

## Installation

```bash
# Clone the repo
git clone https://github.com/yourusername/lotion.git
cd lotion
go build main.go

# go ahead and add it to your path
```

## Usage

*   **Create a new note:**
    ```bash
    lotion new -notebook <notebook_name> <note_name>
    ```
    *(Creates a note named `<note_name>` inside `<notebook_name>`)*

*   **List all notes:**
    ```bash
    lotion list
    ```

*   **Set up remote sync (run once):**
    ```bash
    lotion sync -remote <git_repository_url>
    ```

*   **Sync notes with remote:**
    ```bash
    lotion sync
    ```

*   **Search & Open (Recommended Alias):**
    ```bash
    # Add to your .bashrc or .zshrc
    alias sn="lotion list | fzf --preview='cat {}' | xargs -o nvim"
    # Then run:
    sn
    ```
    *(Requires `fzf` and `nvim` or your preferred editor)*
