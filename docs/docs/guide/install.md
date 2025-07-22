# Installation

<br>

**Make sure to run `Toney init` after installation**

## Linux

### Arch Linux (AUR)

```
yay -S toney
```
> _maintained by [NucleoFusion](https://github.com/NucleoFusion)_

### Debian / Ubuntu (.deb) 

Install the .deb file [_here_](https://github.com/SourcewareLab/Toney/releases/tag/v2.0.0).

```
sudo apt install ./path/to/debfile
```
> _maintained by [NucleoFusion](https://github.com/NucleoFusion)_

### Fedora / RHEL (.dnf) 

Install the .dnf file [_here_](https://github.com/SourcewareLab/Toney/releases/tag/v2.0.0).

```
sudo dnf install ./path/to/debfile
```
> _maintained by [NucleoFusion](https://github.com/NucleoFusion)_

## Windows

Install the .zip file [_here_](https://github.com/SourcewareLab/Toney/releases/tag/v2.0.0).


## MacOS 

Install the .tar.gz file [_here_](https://github.com/SourcewareLab/Toney/releases/tag/v2.0.0).


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
