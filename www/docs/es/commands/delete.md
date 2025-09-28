# Comando `delete`

Este comando borra archivos de la biblioteca o directorio configurado.

## Tabla de contenidos

<!--toc:start-->
<!-- - [Comando `delete`](#comando-delete) -->
<!--   - [Tabla de contenidos](#tabla-de-contenidos) -->
<!--   - [Uso básico](#uso-básico) -->
<!--   - [Opciones](#opciones) -->
<!--     - [-y, --yes](#-y---yes) -->
<!--     - [Heredado](#heredado) -->
<!--   - [Mira también](#mira-también) -->
<!--toc:end-->

## Uso básico

```sh
# Especificar archivo a borrar
video-manager delete <filename>

# RegExp (delete all mp4)
video-manager delete -r '[a-z]+\.mp4'
```

## Opciones

### -y, --yes

Saltar solicitud de confirmación.

### Heredado

Mira las [opciones persistentes](./index.md#opciones-persistentes) para todos opciones heredadas soportadas.

## Mira también

- [list](./list.md)
