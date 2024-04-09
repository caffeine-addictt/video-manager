<!-- markdownlint-disable MD033 -->

# Command Documentation

Here is the full list of commands and persistent options that Video Manager supports.

## Persistent options

### -w, --dir [path]

Uses the specified working directory. By default, it is `~/Videos`.

### -h, --help

Prints the help message

### -v, --verbose

Prints verbose output

### -d, --debug

Prints debug output

### -c, --config [path]

Uses the specified config file.

If not specified, the `.video-manager` file is looked for in this order:

- The current working directory (`./.video-manager`)
- The user's home directory (`~/.video-manager`)

## Commands

- [delete](./delete.md)
- [get](./get.md)
- [list](./list.md)
- [version](./version.md)
