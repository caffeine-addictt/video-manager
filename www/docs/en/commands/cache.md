<!-- markdownlint-disable MD013 -->

# Cache commands

The following commands are used to manage caches.

<!-- ## Table of Contents -->

<!--toc:start-->
<!-- - [Cache commands](#cache-commands) -->
<!--   - [Table of Contents](#table-of-contents) -->
<!--   - [Commands](#commands) -->
<!--     - [list|ls <pattern?>](#listls-pattern) -->
<!--     - [remove|rm <pattern>](#removerm-pattern) -->
<!--     - [clear|clr|wipe](#clearclrwipe) -->
<!--   - [Options](#options) -->
<!--   - [See also](#see-also) -->
<!--toc:end-->

## Commands

### list|ls <pattern?>

List all cache entries or those matching the given pattern.

Each cache entry is stripped of its protocol (`https://`, `http://`, `www.`) before matching.
For example, `https://video-site.com` will be stripped to `video-site.com` and `http://www.other-site.com` will be stripped to `other-site.com`.

Patterns are matched in this order:

- Check if line starts with the pattern
- RegExp match against the line

```sh
video-manager cache list
video-manager cache ls myVideoTitle
```

### remove|rm <pattern>

Remove all cache entries matching the given pattern.

<div class="admonition NOTE" markdown>
<p class="admonition-title">Note</p>
The protocol (`https://`, `http://`, `www.`) is not removed before matching.
</div>

Patterns are matched in this order:

- Check if line starts with the pattern
- RegExp match against the line

```sh
video-manager cache remove https
```

### clear|clr|wipe

Clear the entire cache.

```sh
video-manager cache clear
```

## Options

Only [persistent options](./index.md#persistent-options) are supported.

## See also

- [get command](./get.md)
