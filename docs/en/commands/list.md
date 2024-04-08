<!-- markdownlint-disable MD013 -->

# The list comamnd

This command lists all videos in your configured download location.

## Table of contents

<!--toc:start-->
- [Table of contents](#table-of-contents)
- [Basic usage](#basic-usage)
- [Options](#options)
  - [-n, --count [0 = unlimited]](#n-count-0-unlimited)
  - [-a, --allow (defaults to [])](#a-allow-defaults-to)
  - [-e, --exclude (defaults to [])](#e-exclude-defaults-to)
  - [Inherited](#inherited)
- [See also](#see-also)
<!--toc:end-->

## Basic usage

```sh
video-manager list
```

## Options

### -n, --count [0 = unlimited]

You can specify the maximum number of videos to list. It is `00` by default, and is unlimited if set to `0`.

### -a, --allow (defaults to [])

You can specify the list of allowed extensions, omitting the `.`, and is `[]` by default.

The following is valid:

```sh
video-manager list -a mp4,mkv
video-manager list -a mp4 -a mkv
```

> [!NOTE]
> This cannot be used with `--exclude`.

### -e, --exclude (defaults to [])

You can specify the list of disallowed extensions, omitting the `.`, and is `[]` by default.

The following is valid:

```sh
video-manager list -e mp4,mkv
video-manager list -e mp4 -e mkv
```

> [!NOTE]
> This cannot be used with `--allow`.

### Inherited

See [persistent options](./index.md#persistent-options) for all the supported inherited options.

## See also

Full list of [commands](./index.md).
