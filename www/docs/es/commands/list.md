<!-- markdownlint-disable MD013 -->

# The list command

This command lists all videos in your configured download location.

## Tabla de contenidos

<!--toc:start-->
<!-- - [The list command](#the-list-command) -->
<!--   - [Tabla de contenidos](#tabla-de-contenidos) -->
<!--   - [Uso básico](#uso-básico) -->
<!--   - [Opciones](#opciones) -->
<!--     - [-q, --query (defaults to "")](#-q---query-defaults-to-) -->
<!--     - [-n, --count \[0 = unlimited\]](#-n---count-0--unlimited) -->
<!--     - [-a, --allow (defaults to \[\])](#-a---allow-defaults-to-) -->
<!--     - [-e, --exclude (por defecto \[\])](#-e---exclude-por-defecto-) -->
<!--     - [Heredado](#heredado) -->
<!--   - [Mira también](#mira-también) -->
<!--toc:end-->

## Uso básico

```sh
video-manager list
```

## Opciones

### -q, --query (defaults to "")

Consulta para buscar videos.

<div class="admonition NOTE" markdown>
<p class="admonition-title">Nota</p>
Soporta RegExp.
</div>

```sh
video-manager list -q "query" # Comprobará si "consulta" está en el nombre de archivo
video-manager list -q ".*"   # Coincidirá con cada carácter
video-manager list -q "*"     # Comprobará si "*" está en el nombre de archivo
```

### -n, --count [0 = unlimited]

Puedes especificar el número máximo de videos a enumerar. Por defecto, es `00`, y es ilimitado si se establece en `0`.

### -a, --allow (defaults to [])

Puedes especificar la lista de extensiones permitidas, omitiendo el `.`, y por defecto es `[]`.

Lo siguiente es válido:


```sh
video-manager list -a mp4,mkv
video-manager list -a mp4 -a mkv
```

<div class="admonition NOTE" markdown>
<p class="admonition-title">Note</p>
Esto no se puede usar con `--exclude`.
</div>

### -e, --exclude (por defecto [])

Puedes especificar la lista de extensiones no permitidas, omitiendo el `.`, y por defecto es `[]`.

Lo siguiente es válido:

```sh
video-manager list -e mp4,mkv
video-manager list -e mp4 -e mkv
```

<div class="admonition NOTE" markdown>
<p class="admonition-title">Note</p>
Esto no se puede usar con `--allow`.
</div>

### Heredado

Mira las [opciones persistentes](./index.md#persistent-options) para todos opciones heredadas soportadas.


## Mira también

Lista completa de [comandos](./index.md).

