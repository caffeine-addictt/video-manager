<!-- markdownlint-disable MD033 -->

# Installieren

Diese Anleitung hilft Ihnen bei der Installation des Projekts.

<!-- ## Inhaltsverzeichnis -->

<!--toc:start-->
<!-- - [Installieren](#installieren) -->
<!--   - [Inhaltsverzeichnis](#inhaltsverzeichnis) -->
<!--   - [Nativer Weg mit Go](#nativer-weg-mit-go) -->
<!--   - [Mit Homebrew](#mit-homebrew) -->
<!--   - [Alle Betriebssysteme](#alle-betriebssysteme) -->
<!--     - [GNU/Linux oder MacOS](#gnulinux-oder-macos) -->
<!--     - [WindowsOS](#windowsos) -->
<!--toc:end-->

## Nativer Weg mit Go

```sh
# Installiert das Binärpaket nach $GOPATH/bin
go install github.com/caffeine-addictt/video-manager@latest

# Zum Pfad hinzufügen
echo "alias video-manager=$(go env | grep GOPATH)/bin/video-manager" >> ~/.bashrc

# Quelle
source ~/.bashrc
```

## Mit Homebrew

```sh
brew tap caffeine-addictt/tap
brew install caffeine-addictt/tap/video-manager
```

## Alle Betriebssysteme

Laden Sie die gültigen Tarballs für Ihr Betriebssystem in unseren [Veröffentlichungen](https://github.com/caffeine-addictt/video-manager/releases) herunter.

### GNU/Linux oder MacOS

```sh
tar -xf video-manager_* video-manager && mv video-manager /usr/local/bin/video-manager
```

### Windows

```sh
# Entpacken des Tarballs
unzip video-manager_*.zip -d video-manager && setx path "%path%;%cd%\video-manager\"
```

Siehe [Erste Schritte](./getting-started.md) als Nächstes.
