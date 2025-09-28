<!-- markdownlint-disable MD033 MD013 -->

# Konfiguration

<!-- ## Inhaltsverzeichnis -->

<!--toc:start-->
<!-- * [Inhaltsverzeichnis](#inhaltsverzeichnis) -->
<!-- * [Übersicht](#übersicht) -->
<!-- * [Standardkonfiguration](#standardkonfiguration) -->
<!-- * [Siehe auch](#siehe-auch) -->
<!--toc:end-->

## Übersicht

Konfigurationsdateien werden in folgender Reihenfolge gesucht:

* Arbeitsverzeichnis `(./)`
* Benutzerverzeichnis `($HOME)`

Sie können auch das `-w` oder `--dir` Flag übergeben, um einen Pfad zu einer Konfigurationsdatei anzugeben.
Wenn `-w` nicht bereitgestellt wird, sollten Konfigurationsdateien den Namen `.video-manager` haben und der [YAML](https://yaml.org/) Syntax folgen.

## Standardkonfiguration

```yaml
# Hier werden alle Videodateien gespeichert
dir: ~/.Videos

# Hier werden die Download-URLs zwischengespeichert
cache: ~/.video-manager_history

#
preferred_extensions:
  - .mp4
  - .avi
  - .webm
```

## Siehe auch

* [Befehle](./commands/index.md)
