# Installation

<br>

**Make sure to run `Toney init` after installation**

## Linux

### Arch Linux (AUR)

```
yay -S toney (TBD - not yet works)
```
> _maintained by [NucleoFusion](https://github.com/NucleoFusion)_

### Debian / Ubuntu (.deb) 

Install the .deb file [_here_](https://www.youtube.com/watch?v=xvFZjo5PgG0&list=RDxvFZjo5PgG0&start_radio=1).

```
sudo apt install ./path/to/debfile
```
> _maintained by [NucleoFusion](https://github.com/NucleoFusion)_

### Fedora / RHEL (.dnf) 

Install the .dnf file [_here_](https://www.youtube.com/watch?v=xvFZjo5PgG0&list=RDxvFZjo5PgG0&start_radio=1).

```
sudo dnf install ./path/to/debfile
```
> _maintained by [NucleoFusion](https://github.com/NucleoFusion)_

## Windows

Install the .exe file [_here_](https://www.youtube.com/watch?v=xvFZjo5PgG0&list=RDxvFZjo5PgG0&start_radio=1).

Just run the exe and follow the installer instructions.

## From Source  

### With Git Clone 

```
curl -sSL https://raw.githubusercontent.com/NucleoFusion/toney/main/install.sh | bash
cd Toney
go build
./Toney
```

**Prerequisites**:

- curl
- go

### With Go Install 

```
go install github.com/SourcewareLab/Toney@latest
```
