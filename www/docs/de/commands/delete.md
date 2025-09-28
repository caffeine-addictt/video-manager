# Der Befehl "delete"

Dieser Befehl löscht Dateien aus dem konfigurierten Verzeichnis.

<!-- ## Inhaltsverzeichnis -->

<!--toc:start-->
<!-- - [Der Befehl "delete"](#der-befehl-delete) -->
<!--   - [Inhaltsverzeichnis](#inhaltsverzeichnis) -->
<!--   - [Grundlegende Verwendung](#grundlegende-verwendung) -->
<!--   - [Optionen](#optionen) -->
<!--     - [-y, --yes](#y-yes) -->
<!--     - [Vererbt](#vererbt) -->
<!--   - [Siehe auch](#siehe-auch) -->
<!--toc:end-->

## Grundlegende Verwendung

```sh
# Spezifische Datei
video-manager delete <Dateiname>

# RegExp (alle mp4 löschen)
video-manager delete -r '[a-z]+\.mp4'
```

## Optionen

### -y, --yes

Bestätigungsaufforderung überspringen.

### Vererbt

Alle unterstützten vererbten Optionen finden Sie unter [persistente Optionen](./index.md#persistente-optionen).

## Siehe auch

- [list](./list.md)
