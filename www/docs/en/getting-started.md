<!-- markdownlint-disable MD033 -->

# Getting started

This guide provides an overview of the various features and options available.

<!-- ## Table of Contents -->

<!--toc:start-->
<!-- - [Getting started](#getting-started) -->
<!--   - [Table of Contents](#table-of-contents) -->
<!--   - [Overview](#overview) -->
<!--   - [Example usage](#example-usage) -->
<!--   - [See also](#see-also) -->
<!--toc:end-->

## Overview

This tool is made to help you download and manage videos efficiently.

* We currently has _3_ commands available. See the full list [here](./commands/index.md).
* This tool is highly configurable with configuration files. See [here](./configuration.md) for more info.

## Example usage

Here are some example usage

```sh
#### Downloading a video ####
video-manager get https://video-site.com/video.mp4

# Or multiple videos concurrently
video-manager get -s concurrent https://video-site.com/video1.mp4 https://video-site.com/video2.mp4

# Or multiple videos from a file
video-manager get -f videos.txt


#### List downloaded videos ####
video-manager list

# Not feeling mp4's lately? Exclude them!
video-manager list -e mp4
```

## See also

* [Commands](./commands/index.md)
* [Configuration](./configuration.md)
