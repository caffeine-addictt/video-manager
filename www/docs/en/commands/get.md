# The get command

This command downloads files from URL(s) to your configured download location.

<!-- ## Table of contents -->

<!--toc:start-->
<!-- - [Basic usage](#basic-usage) -->
<!-- - [Filename](#filename) -->
<!-- - [Options](#options) -->
<!--   - [-f, --file](#f-file) -->
<!--   - [-s, --strategy [synchronous|concurrent]](#s-strategy-synchronousconcurrent) -->
<!--   - [-m, --max-concurrency [0 = unlimited]](#m-max-concurrency-0-unlimited) -->
<!--   - [Inherited](#inherited) -->
<!-- - [See also](#see-also) -->
<!--toc:end-->

## Basic usage

```sh
# To download a video
video-manager get https://video-site.com/video.mp4

# Or multiple videos
video-manager get https://video-site.com/video.mp4 https://my.other/video.mp4
```

## Filename

Filenames are generated from the last part of the URL.

For example, if the URL is `https://video-site.com/video.mp4`, the filename will be `video.mp4`.

<div class="admonition NOTE" markdown>
<p class="admonition-title">Note</p>
This also means that the filename will be `video` if the URL is `https://video-site.com/video`!
</div>

## Options

### -f, --file <path>

You can pass a file containing a list of URLs to download.

This can be used alone or with specifying URLs from arguments, meaning the following commands are valid:

- `video-manager get -f urls.txt`
- `video-manager get https://video-site.com/video.mp4 -f urls.txt`

For example,

```text
# urls.txt
# Does not have to be a `.txt`
# Can only have 1 URL per line

https://video-site.com/video.mp4
https://my.other/video.mp4
```

```sh
video-manager get -f urls.txt
```

### -s, --strategy <synchronous|concurrent>

You can specify the download strategy. It is `concurrent` by default.

- `synchronous`: Download all videos sequentially
- `concurrent`: Download videos concurrently

```sh
video-manager get -s synchronous https://video-site.com/video.mp4 https://my.other/video.mp4
```

### -m, --max-concurrency [0 = unlimited] <integer>

You can specify the maximum number of concurrent downloads. It is `10` by default, and is unlimited if set to `0`.

<div class="admonition NOTE" markdown>
<p class="admonition-title">Note</p>
Only available when using `concurrent` strategy.
</div>

```sh
video-manager get -s concurrent -m 10 https://video-site.com/video.mp4 https://my.other/video.mp4
```

### Inherited

See [persistent options](./index.md#persistent-options) for all the supported inherited options.

## See also

Full list of [commands](./index.md).
