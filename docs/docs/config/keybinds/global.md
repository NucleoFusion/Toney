# Global 

Denoted by the `[keybinds.global]` tag in TOML.

## Up 

can change the 'up' key, generally used for menu's etc using the `up` toml key.

default value is:
```toml
up = "up"
```

## Down 

can change the 'down' keybind, generally used for menu's etc using the `down` toml key.

default value is:
```toml
down = "down"
```

## Script 

can change the keybind for executing a command/script using the `script` toml key. The internal working, adds `bash -c` so you may omit it. 

default value is:
```toml
script = []
```
