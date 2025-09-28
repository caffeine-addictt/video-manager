<!-- markdownlint-disable MD013 -->

# Comandos Cache

Los siguientes comandos son utilizados para administrar el caché.

## Tabla de contenidos

<!--toc:start-->
<!-- - [Comandos Cache](#comandos-cache) -->
<!--   - [Tabla de contenidos](#tabla-de-contenidos) -->
<!--   - [Comandos](#comandos) -->
<!--     - [list|ls \<pattern?\>](#listls-pattern) -->
<!--     - [remove|rm ](#removerm-) -->
<!--     - [clear|clr|wipe](#clearclrwipe) -->
<!--   - [Opciones](#opciones) -->
<!--   - [Mira también](#mira-también) -->
<!--toc:end-->

## Comandos

### list|ls <pattern?>

Lista todas las entradas del caché que coincidan con el patrón de búsqueda especificado.

Cada entrada en la caché se despoja de su protocolo (https://, http://, www.) antes de hacer coincidencias.
Por ejemplo, `https://video-site.com` se reformateará a `video-site.com` y `http://www.other-site.com` se reformateará a `other-site.com` para realizar la búsqueda.

Los patrones se emparejan en el siguiente orden:

- Verificar si la linea inicia con el patrón de búsqueda
- Coincidencia de RegExp para la linea de búsqueda

```sh
video-manager cache list
video-manager cache ls myVideoTitle
```

### remove|rm <pattern>

Borra todas las entradas del caché de acuerdo al patrón establecido.

<div class="admonition NOTE" markdown>
<p class="admonition-title">Note</p>
El protocolo (`https://`, `http://`, `www.`) no es removido antes de hacer la búsqueda.
</div>

Los patrones son emparejados en el siguiente orden:

- Comprueba si la línea empieza con el patrón.
- Coincidencia de RegExp contra la línea.

```sh
video-manager cache remove https
```

### clear|clr|wipe

Limpia toda el caché.


```sh
video-manager cache clear
```

## Opciones

Solo [opciones persistentes](./index.md#opciones-persistentes) son soportadas.

## Mira también

- [get](./get.md)
