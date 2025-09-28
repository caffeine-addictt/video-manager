# Der Befehl "get"

Dieser Befehl lädt Dateien von URL(s) in Ihren konfigurierten Download-Ort herunter.

<!-- ## Inhaltsverzeichnis -->

<!--toc:start-->
<!-- - [Grundlegende Verwendung](#grundlegende-verwendung) -->
<!-- - [Dateiname](#dateiname) -->
<!-- - [Optionen](#optionen) -->
<!--   - [-f, --file](#f-file) -->
<!--   - [-s, --strategy [synchronous|concurrent]](#s-strategy-synchronousconcurrent) -->
<!--   - [-m, --max-concurrency [0 = unbegrenzt]](#m-max-concurrency-0-unbegrenzt) -->
<!--   - [Vererbt](#vererbt) -->
<!-- - [Siehe auch](#siehe-auch) -->
<!--toc:end-->

## Grundlegende Verwendung

```sh
# Um ein Video herunterzuladen
video-manager get https://video-site.com/video.mp4

# Oder mehrere Videos
video-manager get https://video-site.com/video.mp4 https://my.other/video.mp4
```

## Dateiname

Dateinamen werden aus dem letzten Teil der URL generiert.

Beispielweise wird bei der URL `https://video-site.com/video.mp4` der Dateiname `video.mp4` sein.

<div class="admonition IMPORTANT" markdown>
<p class="admonition-title">Important</p>
Das bedeutet auch, dass der Dateiname `video` ist, wenn die URL `https://video-site.com/video` lautet!
</div>


## Optionen

### -f, --file <Pfad>

Sie können eine Datei mit einer Liste von URLs zum Herunterladen übergeben.

Dies kann allein oder mit dem Angeben von URLs aus Argumenten verwendet werden, was bedeutet, dass die folgenden Befehle gültig sind:

- `video-manager get -f urls.txt`
- `video-manager get https://video-site.com/video.mp4 -f urls.txt`

Zum Beispiel,

```text
# urls.txt
# Muss nicht eine `.txt` sein
# Kann nur 1 URL pro Zeile haben

https://video-site.com/video.mp4
https://my.other/video.mp4
```

```sh
video-manager get -f urls.txt
```

### -s, --strategy <synchronous|concurrent>

Sie können die Download-Strategie angeben. Standardmäßig ist sie `concurrent`.

- `synchronous`: Lädt alle Videos sequenziell herunter
- `concurrent`: Lädt Videos gleichzeitig herunter

```sh
video-manager get -s synchronous https://video-site.com/video.mp4 https://my.other/video.mp4
```

### -m, --max-concurrency [0 = unbegrenzt] <Ganzzahl>

Sie können die maximale Anzahl gleichzeitiger Downloads angeben. Standardmäßig ist es `10`, und bei `0` ist es unbegrenzt.

<div class="admonition NOTE" markdown>
<p class="admonition-title">Note</p>
Nur verfügbar, wenn die `concurrent`-Strategie verwendet wird.
</div>

```sh
video-manager get -s concurrent -m 10 https://video-site.com/video.mp4 https://my.other/video.mp4
```

### Vererbt

Alle unterstützten vererbten Optionen finden Sie unter [persistente Optionen](./index.md#persistente-optionen).

## Siehe auch

Vollständige Liste der [Befehle](./index.md).
