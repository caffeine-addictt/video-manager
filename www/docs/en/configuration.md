<!-- markdownlint-disable MD033 MD013 -->

# Configuration

<!-- ## Table of Contents -->

<!--toc:start-->
<!-- * [Table of Contents](#table-of-contents) -->
<!-- * [Overview](#overview) -->
<!-- * [Default configuration](#default-configuration) -->
<!-- * [See Also](#see-also) -->
<!--toc:end-->

## Overview

Configuration files are searched in this order:

* Working directory `(./)`
* Home directory `($HOME)`

You can also pass the `-w` or `--dir` flag to specify a path to a configuration file.
If `-w` is not provided, configuration files should be named `.video-manager` and follow [yaml](https://yaml.org/) syntax.

## Default configuration

```yaml
# This is where all video files will be stored
dir: ~/.Videos

# This is where download urls are cached
cache: ~/.video-manager_history

# The preferred extensions for files
# The first one that matches is used
preferred_extensions:
  - .mp4
  - .avi
  - .webm
```

## See Also

* [Commands](./index.md)
