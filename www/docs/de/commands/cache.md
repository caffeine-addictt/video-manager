<!-- markdownlint-disable MD013 -->

# Cache-Befehle

Die folgenden Befehle werden verwendet, um Caches zu verwalten.

<!-- ## Inhaltsverzeichnis -->

<!--toc:start-->
<!-- - [Cache-Befehle](#cache-befehle) -->
<!--   - [Inhaltsverzeichnis](#inhaltsverzeichnis) -->
<!--   - [Befehle](#befehle) -->
<!--     - [list|ls <Muster?>](#listls-muster) -->
<!--     - [remove|rm <Muster>](#removerm-muster) -->
<!--     - [clear|clr|wipe](#clearclrwipe) -->
<!--   - [Optionen](#optionen) -->
<!--   - [Siehe auch](#siehe-auch) -->
<!--toc:end-->

## Befehle

### list|ls <Muster?>

Listet alle Cache-Einträge oder solche, die mit dem angegebenen Muster übereinstimmen.

Jeder Cache-Eintrag wird von seinem Protokoll (`https://`, `http://`, `www.`) befreit, bevor er abgeglichen wird.
Beispielsweise wird `https://video-site.com` auf `video-site.com` gekürzt und `http://www.other-site.com` auf `other-site.com`.

Muster werden in dieser Reihenfolge abgeglichen:

- Überprüfen, ob die Zeile mit dem Muster beginnt
- RegExp-Abgleich gegen die Zeile

```sh
video-manager cache list
video-manager cache ls myVideoTitle
```

### remove|rm <Muster>

Entfernt alle Cache-Einträge, die mit dem angegebenen Muster übereinstimmen.

<div class="admonition NOTE" markdown>
<p class="admonition-title">Note</p>
Das Protokoll (`https://`, `http://`, `www.`) wird vor dem Abgleichen nicht entfernt.
</div>

Muster werden in dieser Reihenfolge abgeglichen:

- Überprüfen, ob die Zeile mit dem Muster beginnt
- RegExp-Abgleich gegen die Zeile

```sh
video-manager cache remove https
```

### clear|clr|wipe

Löscht den gesamten Cache.

```sh
video-manager cache clear
```

## Optionen

Es werden nur [persistente Optionen](./index.md#persistente-optionen) unterstützt.

## Siehe auch

- [Befehl "get"](./get.md)
