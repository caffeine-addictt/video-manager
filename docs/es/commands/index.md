<!-- markdownlint-disable MD033 -->

# Documentación de los comandos

Esta es la lista completa de los comandos y las opciones persistentes suportadas por Video Manager.

## Opciones persistentes

### -w, --dir <path>

Utiliza el directorio de trabajo especificado. Por defecto, es ~/Videos.


### -h, --help

Imprime el mensaje de ayuda.

### -v, --verbose

Imprime salida detallada.

### -d, --debug

Imprime salida de depuración.


### -c, --config <path>

Utiliza el archivo de configuración especificado.

Si no se especifica, se buscará el archivo `.video-manager` en este orden:

- El directorio de trabajo actual (`./.video-manager`)
- El directorio principal del usuario (`~/.video-manager`)

### -C, --cache <path>

Utiliza el archivo de caché especificado.

Si no se especifica, se utilizará el archivo `~/.video-manager_history`, o se definirá de otra manera en el [archivo de configuración](../configuration.md).

## Commands

- [cache clear|list|remove](./cache.md)
- [config init|where](./config.md)
- [delete](./delete.md)
- [get](./get.md)
- [list](./list.md)
- [version](./version.md)
