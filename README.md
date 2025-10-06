
# Gooey CLI

The Gooey CLI is a management tool for the [Gooey](https://github.com/cookiengineer/gooey) WebASM
framework. Its goal is to simplify the bootstrapping process of Controllers, Views, and Components
built with Gooey.


## Installation

You can use the `go install` command to install the program directly from GitHub:

```bash
# Install directly from GitHub
go install github.com/cookiengineer/gooey-cli/cmds/gooey@master;
```


## Building

Building Requirements are `go`, `cgo`, `gcc`, and `WebKit2GTK` version `4.1`.

(Upstream [webview/webview](https://github.com/webview/webview_go) isn't ready for `6.0` yet as of 2025-Q4)

```bash
# On ArchLinux
sudo pacman -S gcc go webkit2gtk-4.1;

cd /path/to/gooey-cli;
go build ./cmds/gooey;
```


## Usage

The Gooey CLI is currently highly experimental and will break regularly until upstream has reached
a stable `1.0.0` version release.

Execute the program without any parameters to show its help and usage examples:

```bash
# Show Help
gooey;
```


## License

This project is licensed under [MIT](./LICENSE.txt) license.
