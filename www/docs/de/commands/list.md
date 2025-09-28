<!-- markdownlint-disable MD013 -->

# Der Befehl "list"

Dieser Befehl listet alle Videos in Ihrem konfigurierten Download-Ort auf.

<!-- ## Inhaltsverzeichnis -->

<!--toc:start-->
<!-- - [Inhaltsverzeichnis](#inhaltsverzeichnis) -->
<!-- - [Grundlegende Verwendung](#grundlegende-verwendung) -->
<!-- - [Optionen](#optionen) -->
<!--   - [-q, --query (Standardwert "")](#q-query-standardwert) -->
<!--   - [-n, --count [0 = unbegrenzt]](#n-count-0-unbegrenzt) -->
<!--   - [-a, --allow (Standardwert [])](#a-allow-standardwert) -->
<!--   - [-e, --exclude (Standardwert [])](#e-exclude-standardwert) -->
<!--   - [Vererbte](#vererbte) -->
<!-- - [Siehe auch](#siehe-auch) -->
<!--toc:end-->

## Grundlegende Verwendung

```sh
video-manager list
```

## Optionen

### -q, --query (Standardwert "")

Abfrage zum Suchen von Videos.

<div class="admonition NOTE" markdown>
<p class="admonition-title">Note</p>
RegExp wird unterstützt.
</div>

```sh
video-manager list -q "Abfrage" # Überprüft, ob "Abfrage" im Dateinamen enthalten ist
video-manager list -q ".*"      # Passt zu jedem Zeichen
video-manager list -q "*"       # Überprüft, ob "*" im Dateinamen enthalten ist
```

### -n, --count [0 = unbegrenzt]

Sie können die maximale Anzahl der aufzulisten Videos angeben. Standardmäßig ist es `00`, und wenn es auf `0` gesetzt ist, ist es unbegrenzt.

### -a, --allow (Standardwert [])

Sie können die Liste der erlaubten Erweiterungen angeben, wobei das `.` weggelassen wird, und standardmäßig `[]` ist.

Folgendes ist gültig:

```sh
video-manager list -a mp4,mkv
video-manager list -a mp4 -a mkv
```

<div class="admonition NOTE" markdown>
<p class="admonition-title">Note</p>
Dies kann nicht mit `--exclude` verwendet werden.
</div>

### -e, --exclude (Standardwert [])

Sie können die Liste der unerwünschten Erweiterungen angeben, wobei das `.` weggelassen wird, und standardmäßig `[]` ist.

Folgendes ist gültig:

```sh
video-manager list -e mp4,mkv
video-manager list -e mp4 -e mkv
```

<div class="admonition NOTE" markdown>
<p class="admonition-title">Note</p>
Dies kann nicht mit `--allow` verwendet werden.
</div>

### Vererbte

Siehe [persistente Optionen](./index.md#persistente-optionen) für alle unterstützten vererbten Optionen.

## Siehe auch

Vollständige Liste der [Befehle](./index.md).
