<!-- markdownlint-disable MD033 MD013 -->

# Configuración

## Tabla de contenidos

<!--toc:start-->
<!-- - [Configuración](#configuración) -->
<!--   - [Tabla de contenidos](#tabla-de-contenidos) -->
<!--   - [Resumen](#resumen) -->
<!--   - [Configuración por defecto](#configuración-por-defecto) -->
<!--   - [Mira también](#mira-también) -->
<!--toc:end-->

## Resumen

Los archivos de configuración tienen la siguiente estructura:

* Directorio de trabajo `(./)`
* Directorio inicial `($HOME)`

También puedes especificar en la terminal los argumentos  `-w` o `--dir` para establecer el directorio al archivo de configuración.

Si `-w` no es establecido, los archivos de configuración deberán ser nombrados como `.video-manager` y seguir la sintáxis [yaml](https://yaml.org/).

## Configuración por defecto

```yaml
# Aquí es donde los videos serán guardados
dir: ~/.Videos

# Este es el caché de los archivos descargados
cache: ~/.video-manager_history

# Extensiones preferidas para los videos
# El primer caso que coincida será el utilizado
preferred_extensions:
  - .mp4
  - .avi
  - .webm
```

## Mira también

* [Commands](./commands/index.md)
