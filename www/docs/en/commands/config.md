<!-- markdownlint-disable MD013 -->

# Config commands

The following commands are used to manage the configuration file.

<!-- ## Table of Contents -->

<!--toc:start-->
<!-- - [Config commands](#config-commands) -->
<!--   - [Table of Contents](#table-of-contents) -->
<!--   - [Commands](#commands) -->
<!--     - [init|create|reset|new <path?>](#initcreateresetnew-path) -->
<!--       - [-y, --yes](#y-yes) -->
<!--     - [where](#where) -->
<!--   - [Options](#options) -->
<!--   - [See also](#see-also) -->
<!--toc:end-->

## Commands

### init|create|reset|new <path?>

This command creates a new configuration file in the given path or the current working directory.

```sh
video-manager init
video-manager create ~/.video-manager
video-manager reset ./.video-manager
```

#### -y, --yes

Skips the confirmation prompt.

### where

This command prints the location of the currently loaded configuration file.

```sh
video-manager where
```

## Options

Only [persistent options](./index.md#persistent-options) are supported.

## See also

- [get command](./get.md)
- [list command](./list.md)
