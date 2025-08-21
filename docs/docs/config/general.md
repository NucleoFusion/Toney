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

## Start Script 

runs a script/bash command on startup using the `start_script` key, it takes an array. The internal working, adds `bash -c` so you may omit it. 

default value is:
```toml
start_script = []
```

## Stop Script 

runs a script/bash command on quit/exit using the `stop_script` key, it takes an array. The internal working, adds `bash -c` so you may omit it. 

default value is:
```toml
stop_script = []
```

## Script 

runs a script/bash command on keybind using the `script` key, it takes an array. The internal working, adds `bash -c` so you may omit it. 

default value is:
```toml
script = []
```


## Todo 

Configure synchronization with TODO: Comments in a project. Use the `[general.todo]` tag in TOML.

### Enable

can enable/disable todo integration using the `enable` key.

default value is:
```toml
enable = false 
```

### Todo Projects 

can add project directories using the `[[general.todo.todo_projects]]` key which will be searched for TODO comments.

example value is:
```toml
[[general.todo.todo_projects]]
name = "Toney"
path = "/home/nucleofusion/Programming/projects/toney"
exclude_dir = []
```
