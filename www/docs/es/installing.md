<!-- markdownlint-disable MD033 -->

# Instalación
Esta guia te ayudará con la instalación del proyecto

## Tabla de contenidos

<!--toc:start-->
<!-- - [Instalación](#instalación) -->
<!--   - [Tabla de contenidos](#tabla-de-contenidos) -->
<!--   - [Nativamente con *Go*](#nativamente-con-go) -->
<!--   - [Con Homebrew](#con-homebrew) -->
<!--   - [Todos los sistemas operativos](#todos-los-sistemas-operativos) -->
<!--     - [GNU/Linux o MacOS](#gnulinux-o-macos) -->
<!--     - [Windows](#windows) -->
<!--toc:end-->

## Nativamente con *Go*

```sh
# Instala el archivo binario en $GOPATH/bin
go install github.com/caffeine-addictt/video-manager@latest

# Agrega la ruta del programa
echo "alias video-manager=$(go env | grep GOPATH)/bin/video-manager" >> ~/.bashrc

# Carga de nuevo la configuración inicial
source ~/.bashrc
```

## Con Homebrew

```sh
brew tap caffeine-addictt/tap
bew install caffeine-addictt/tap/video-manager
```

## Todos los sistemas operativos

Descarga los tarballs para tu sistema operativo, disponible en [releases](https://github.com/caffeine-addictt/video-manager/releases).

### GNU/Linux o MacOS

```sh
tar -xf video-manager_* video-manager && mv video-manager /usr/local/bin/video-manager
```

### Windows

```sh
# Descomprime el tarball y agraga la ruta
unzip video-manager_*.zip -d video-manager && setx path "%path%;%cd%\video-manager\"
```

Mira [¿Cómo empezar?](./getting-started.md).
