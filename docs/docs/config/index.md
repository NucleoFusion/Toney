# Configuration

Toney supports a flexible configuration file for customizing themes, keybinds etc. We use a TOML format and the default config file is at: 

- Linux:
```
~/.config/toney/config.toml
```

- Windows:
```
C:\Users\<username>\AppData\Roaming\toney\config.yml
```

- Mac:
```
~/Library/Application Support/toney/config.yml
```

## File Format

Toney uses TOML for its configuration file. Making it readable and easy to edit.

An example minimal file with vim motions (hjkl motions, and 'a' for create etc)
```toml
[keybinds.global]
up = "k"
down = "j"
left = "h"
right = "l"

[keybinds.home]
create = "a"
scroll_up = "k"
scroll_down = "k"

[keybinds.daily]
create = "a"
```

## Key Sections

the config file has 3 main sections:

- **General**: for general stuff such as editor, notes_dir etc
- **Styles**: to customize styles and colors
- **Keybinds**: customize keybinds and styles
