<!-- markdownlint-disable MD033 -->

# Befehlsdokumentation

Hier ist die vollständige Liste der Befehle und persistenter Optionen, die von Video Manager unterstützt werden.

## Persistente Optionen

### -w, --dir <Pfad>

Verwendet das angegebene Arbeitsverzeichnis. Standardmäßig ist es `~/Videos`.

### -h, --help

Gibt die Hilfemeldung aus

### -v, --verbose

Gibt ausführliche Ausgaben aus

### -d, --debug

Gibt Debug-Ausgaben aus

### -c, --config <Pfad>

Verwendet die angegebene Konfigurationsdatei.

Wenn nicht angegeben, wird die `.video-manager`-Datei in folgender Reihenfolge gesucht:

- Das aktuelle Arbeitsverzeichnis (`./.video-manager`)
- Das Benutzerverzeichnis (`~/.video-manager`)

### -C, --cache <Pfad>

Verwendet die angegebene Cache-Datei.

Wenn nicht angegeben, wird die `~/.video-manager_history`-Datei verwendet oder anderweitig in der [Konfigurationsdatei](../configuration.md) definiert.

## Befehle

- [cache clear|list|remove](./cache.md)
- [delete](./delete.md)
- [get](./get.md)
- [list](./list.md)
- [version](./version.md)