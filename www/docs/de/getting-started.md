<!-- markdownlint-disable MD033 -->

# Erste Schritte

Diese Anleitung bietet einen Überblick über die verschiedenen Funktionen und Optionen.

<!-- ## Inhaltsverzeichnis -->

<!--toc:start-->
<!-- - [Erste Schritte](#erste-schritte) -->
<!--   - [Inhaltsverzeichnis](#inhaltsverzeichnis) -->
<!--   - [Übersicht](#übersicht) -->
<!--   - [Beispielverwendung](#beispielverwendung) -->
<!--   - [Siehe auch](#siehe-auch) -->
<!--toc:end-->

## Übersicht

Dieses Tool wurde entwickelt, um Ihnen beim effizienten Herunterladen und Verwalten von Videos zu helfen.

* Derzeit stehen _3_ Befehle zur Verfügung. Siehe die vollständige Liste [hier](./commands/index.md).
* Dieses Tool ist hoch konfigurierbar mit Konfigurationsdateien. Weitere Informationen finden Sie [hier](./configuration.md).

## Beispielverwendung

Hier sind einige Beispiele zur Verwendung

```sh
#### Herunterladen eines Videos ####
video-manager get https://video-site.com/video.mp4

# Oder mehrere Videos gleichzeitig
video-manager get -s concurrent https://video-site.com/video1.mp4 https://video-site.com/video2.mp4

# Oder mehrere Videos aus einer Datei
video-manager get -f videos.txt


#### Liste heruntergeladener Videos ####
video-manager list

# Lust nicht auf mp4's? Schließen Sie sie aus!
video-manager list -e mp4
```

## Siehe auch

* [Befehle](./commands/index.md)
* [Konfiguration](./configuration.md)
