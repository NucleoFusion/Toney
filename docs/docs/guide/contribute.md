# Contribution

Toney is built with Go using [bubbletea](https://github.com/charmbracelet/bubbletea), [lipgloss](https://github.com/charmbracelet/lipgloss) and other supporting libraries.

## Project Structure

```
├── cmd          // CLI 
└── internal
    ├── colors   // Styles Config
    ├── config   // Config 
    ├── enums    // All enums
    ├── fileTree // Filtree - manages tree creation for directory
    ├── keymap   // Keymap 
    ├── messages // Messages - has all bubbletea message declarations
    ├── models   // Models - Has all bubbletea models
    └── styles   // Styles - has all lipgloss styles declarations
```

## How to Contribute

1. **Fork the Repository from Github**

2. **Clone the Repository**
```
git clone https://github.com/<your-username>/Toney
cd Toney
```

3. **Install dependencies**
```
go get
```

4. **Build and Run**
```
go build
./Toney
```

## What can you work on?

Keep a look at the issues section, there you would find Bugs, Feature Request etc.

**Please** drop a message in the issue and get it approved before you start working, we can discuss potential ways to structure a solution.

## License

By contributing, you agree to license your work under the MIT License used by this project.
