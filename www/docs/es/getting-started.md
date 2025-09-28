<!-- markdownlint-disable MD033 -->

# ¿Cómo empezar?

Esta guia provee un resumen de las distintas características y opciones disponibles

## Tabla de contenidos

<!--toc:start-->
<!-- - [¿Cómo empezar?](#cómo-empezar) -->
<!--   - [Tabla de contenidos](#tabla-de-contenidos) -->
<!--   - [Resumen](#resumen) -->
<!--   - [Ejemplos de uso](#ejemplos-de-uso) -->
<!--   - [Mira también](#mira-también) -->
<!--toc:end-->

## Resumen

Esta herramiente está diseñada para ayudarte a descargar y administrar tu biblioteca de videos de forma eficiente.

* Actualmente tiene _3_ comandos disponibles. Mira la lista completa [aquí](./commands/index.md).
* Esta herramienta es altamente configurable a través de archivos de configuración. Mira [aquí](./configuration.md) para más información.

## Ejemplos de uso

Aquí algunos ejemplos de uso

```sh
#### Descargar un video ####
video-manager get https://video-site.com/video.mp4

# Descargar múltiples videos de forma concurrente
video-manager get -s concurrent https://video-site.com/video1.mp4 https://video-site.com/video2.mp4

# Múltiples videos de un archivo
video-manager get -f videos.txt


#### Listar los videos descargados ####
video-manager list

# No te gusta el formato mp4 últimamente? ¡Exclúyelo!
video-manager list -e mp4
```

## Mira también

* [Commands](./commands/index.md)
* [Configuration](./configuration.md)
