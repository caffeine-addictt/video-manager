# The delete command

This command deletes files from the configured directory.

<!-- ## Table of Contents -->

<!--toc:start-->
<!-- - [The delete command](#the-delete-command) -->
<!--   - [Table of Contents](#table-of-contents) -->
<!--   - [Basic usage](#basic-usage) -->
<!--   - [Options](#options) -->
<!--     - [-y, --yes](#y-yes) -->
<!--     - [Inherited](#inherited) -->
<!--   - [See also](#see-also) -->
<!--toc:end-->

## Basic usage

```sh
# Specific file
video-manager delete <filename>

# RegExp (delete all mp4)
video-manager delete -r '[a-z]+\.mp4'
```

## Options

### -y, --yes

Skip confirmation prompt.

### Inherited

See [persistent options](./index.md#persistent-options) for all the supported inherited options.

## See also

- [list](./list.md)
