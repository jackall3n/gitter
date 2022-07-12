# 💡 gitter (better name pending)

## Installation

```shell
go install github.com/jackall3n/gitter
```

## Example
![Example](example-1.png)

## Usage
```shell
gitter checkout [ticket] [description]
```

#### With alias
```shell
# ~/.zshrc
alias co="gitter checkout"
```

```shell
co [ticket]
```

## Configuration

Add a `.gitter.yaml` configuration to your `$HOME` directory, or your project root
```shell
# Global config
touch ~/.gitter.yaml

# Project specific
touch ~/development/special-project/.gitter.yaml
```

#### Example

```yaml
# ~/.gitter.yaml
prefix: ja
board: ABC
```

