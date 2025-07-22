# General Config

Denoted by the `[general]` tag in TOML.

## Notes Directory

can change the directory where all the notes are stored using the `notes_dir` key.

default value is:
```toml
notes_dir = ".toney"
```

## Editor 

can change the editor for editing notes using the `editor` key, it takes an array. **Make sure to write the command for the editor, not the name**.

For users with _helix_, use `editor=["alacritty", "-e" , "bash", "-c", "hx"]` or change it according to your term and shell. Helix requires a TTY to launch hence this is required.

default value is:
```toml
editor = ["nvim"]
```
