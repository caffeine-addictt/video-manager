<!-- markdownlint-disable MD013 -->

# Comandos de configuración

Los siguientes comandos son utilizados para administrar los archivos de configuración.

<!-- ## Tabla de contenidos -->

<!--toc:start-->
<!-- - [Comandos de configuración](#comandos-de-configuración) -->
<!--   - [Tabla de contenidos](#tabla-de-contenidos) -->
<!--   - [Comandos](#comandos) -->
<!--     - [init|create|reset|new \<path?\>](#initcreateresetnew-path) -->
<!--       - [-y, --yes](#-y---yes) -->
<!--     - [where](#where) -->
<!--   - [Opciones](#opciones) -->
<!--   - [Mira también](#mira-también) -->
<!--toc:end-->

## Comandos

### init|create|reset|new <path?>

Este comando crea un nuevo archivo de configuración en una ruta establecida o en la ruta predeterminada de trabajo.


```sh
video-manager init
video-manager create ~/.video-manager
video-manager reset ./.video-manager
```

#### -y, --yes

Saltar solicitud de confirmación.

### where

Comando que imprime la ruta actual en la que se encuentra el archivo de configuración cargado.

```sh
video-manager where
```

## Opciones

Solo [opciones persistentes](./index,md/persistent-options)son soportadas.

## Mira también

- [get](./get.md)
- [list](./list.md)
