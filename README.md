# Toney

**Toney** is a fast, lightweight, terminal-based note-taking app for the modern developer. Built with [Bubbletea](https://github.com/charmbracelet/bubbletea), Toney brings a sleek TUI interface with markdown rendering, file navigation, and native Neovim editing – all in your terminal.



https://github.com/user-attachments/assets/bf2411e0-2a7e-4669-a12a-6ee1bb63b75b



---

## ✨ Features

- ⚡ **Fast** – Minimal memory usage and snappy performance.
- 📝 **Markdown Renderer** – Styled previews via [`glamour`](https://github.com/charmbracelet/glamour).
- 🧠 **Neovim Integration** – Edit your notes using your favorite editor (`nvim`).
- 📂 **File Management** – Easily navigate, open, and manage markdown files.
- 🧩 **Component Architecture** – Modular codebase using Bubbletea’s `Model` system.
- 🎨 **TUI Styling** – Clean, responsive interface using `lipgloss`.

---

## 🚀 Installation

You can install **Toney** directly using `go install`:

```
go install github.com/SourcewareLab/Toney@latest
```

This will download, build, and install the `Toney` binary into your `$GOBIN` (typically `$HOME/go/bin`).

Run this command to ensure Toney is setup perfectly.

```
  Toney init
```

### 🧪 Requirements

- Go 1.16 or later
- Git (to fetch the module)

Make sure your `GOBIN` is in your system's `PATH`:

```
export PATH=$PATH:$(go env GOBIN)
```

### ✅ Verify Installation

Once installed, you can run:

```
Toney
```

---

## 🔑 Keybinds

| Key Combination | Action                    |
|-----------------|---------------------------|
| **F** / **Shift + F** | Focus on File Tree         |
| **V** / **Shift + V** | Focus on File Viewer       |

### 📁 File Tree Focus Shortcuts

Once the File Tree is focused (`F` or `Shift + F`):

| Key      | Action         |
|----------|----------------|
| **c**    | Create a file/folder |
| **d**    | Delete selected     |
| **r**    | Rename selected     |
| **m**    | Move selected       |
| **Enter**| Edit selected file  |


---

## 🗺 Roadmap

### v2.0.0 Goals

- [X] Daily Tasks 
- [ ] Diary
- [X] Config File
  - [X] Custom Styles
  - [X] Custom Editor
  - [X] Custom Notes Directory
  - [X] Custom Keybinds
- [ ] Search In Notes
- [ ] Search for Notes  
- [X] Keybind Helper (using Bubbles)

### Short Term Goals

- [ ] Overlay support
- [ ] Viewer style improvements
- [ ] Error popups
- [X] Separate package for messages
- [ ] Keybind refactor
- [ ] Config file support (`~/.config/toney/config.yaml`)
- [ ] Custom markdown renderer
- [ ] Custom components:  
  - [ ] [ ] Task Lists  
  - [ ] `code` blocks  
  - [ ] Tables  
- [ ] File Import/Export
- [ ] Configurable external editor support

### Long Term Goals

- [ ] Cross-platform **mobile app**
- [ ] **Server sync** with configuration & cloud storage

---

## 🛠️ Project Structure

```
toney/
├── cmd/              # Entry point (main.go)
├── internal/
│   ├── models/       # All UI models (Home, Viewer, Popups, etc.)
│   ├── enums/        # Typed enums (pages, popup types)
│   ├── messages/     # Message types for tea.Msg (will be modularized)
│   └── utils/        # Shared utility functions
```

---

## 🤝 Contributing

We welcome contributions! Toney follows Go and Bubbletea conventions.

### 🧾 Guidelines

- Follow idiomatic Go formatting (`go fmt`, `go vet`, `golint`).
- Use `Init`, `Update`, and `View` separation for all models.
- Keep component responsibilities well-isolated.
- All exported functions/types should be documented with Go-style comments.
- Prefer `tea.Msg` messages over direct cross-component function calls.

### ✅ How to contribute

1. Fork the repo and create a feature branch:
   ```bash
   git checkout -b feature/my-feature
   ```

2. Write your code and make sure it builds:
   ```bash
   go build ./...
   ```

3. Format your code:
   ```bash
   go fmt ./...
   ```

4. Commit and push:
   ```bash
   git commit -m "feat: add my awesome feature"
   git push origin feature/my-feature
   ```

5. Submit a Pull Request 🎉

Please open an issue or discussion for large changes before starting them.

---

## 📄 License

MIT License. See [LICENSE](./LICENSE).

---

## 💡 Inspiration

Toney is inspired by:
- [Glow](https://github.com/charmbracelet/glow) – for markdown rendering  
- [Lazygit](https://github.com/jesseduffield/lazygit) – for terminal UI polish  
- [Charm ecosystem](https://github.com/charmbracelet) – for all things delightful in the terminal

---

> Made with 💀 by [Nucleo](https://github.com/NucleoFusion) & [SourcewareLab](https://discord.gg/X69MUr2DKm)

