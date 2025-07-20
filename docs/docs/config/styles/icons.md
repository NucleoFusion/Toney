# Icons 

Denoted by the `[styles.icons]` tag in TOML. Uses nerd fonts, so may require nerd fonts to be downloaded.

## Folder Icon 

can change the folder icon shown in the filetree using the `folder_icon` key.

default value is:
```toml
folder_icon = "󰷏"
```

## File Icon 

can change the file icon shown in the filetree using the `file_icon` key.

default value is:
```toml
file_icon = ""
```

## Task Icons 

used to change the Icons shown in daily tasks for each status. uses the `[styles.icons.task_icons]`

### Completed Task

default value is:
```toml
completed_icon = "✓"
```

### Pending Task 

default value is:
```toml
pending_icon = "~"
```
### Started Task 

default value is:
```toml
started_icon = "○"
```

### Abandoned Task 

default value is:
```toml
abandoned_icon = "×"
```

