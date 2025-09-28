# Comando get

Descarga videos de una dirección especificada en la ruta de descarga configurada.

## Tabla de contenidos

<!--toc:start-->
<!-- - [Comando get](#comando-get) -->
<!--   - [Tabla de contenidos](#tabla-de-contenidos) -->
<!--   - [Uso básico](#uso-básico) -->
<!--   - [Nombre de archivo](#nombre-de-archivo) -->
<!--   - [Opciones](#opciones) -->
<!--     - [-f, --file ](#-f---file-) -->
<!--     - [-s, --strategy \<synchronous|concurrent\>](#-s---strategy-synchronousconcurrent) -->
<!--     - [-m, --max-concurrency \[0 = unlimited\] ](#-m---max-concurrency-0--unlimited-) -->
<!--     - [Heredado](#heredado) -->
<!--   - [See also](#see-also) -->
<!--toc:end-->

## Uso básico

```sh
# Descarga de un video
video-manager get https://video-site.com/video.mp4

# O múltiples videos
video-manager get https://video-site.com/video.mp4 https://my.other/video.mp4
```

## Nombre de archivo

Los nombres de archivo son generados con la última parte de la URL.
Por ejemplo, si la URL es `https://video-site.com/video.mp4`, el nombre de archivo será `video.mp4`.

<div class="admonition IMPORTANT" markdown>
<p class="admonition-title">Importante</p>
¡Esto también significa que el nombre de archivo será `video` si la URL es `https://video-site.com/video`!
</div>

## Opciones

### -f, --file <path>

Puedes pasar un archivo que contenga una lista de URL para descargar.

Esto se puede usar solo o especificando URLs desde los argumentos, lo que significa que los siguientes comandos son válidos:

- `video-manager get -f urls.txt`
- `video-manager get https://video-site.com/video.mp4 -f urls.txt`

Por ejemplo,

```text
# urls.txt
# No es necesario que sea un `.txt`
# Puede haber solo 1 URL por línea

https://video-site.com/video.mp4
https://my.other/video.mp4
```

```sh
video-manager get -f urls.txt
```

### -s, --strategy <synchronous|concurrent>


Puedes especificar la estrategia de descarga. Por defecto es `concurrente`.

- `synchronous`: Descargar todos los videos secuencialmente
- `concurrent`: Descargar videos de manera concurrente

```sh
video-manager get -s synchronous https://video-site.com/video.mp4 https://my.other/video.mp4
```

### -m, --max-concurrency [0 = unlimited] <integer>

Puedes especificar el número máximo de descargas concurrentes. Por defecto es `10` y es ilimitado si se establece en `0`.

<div class="admonition NOTE" markdown>
<p class="admonition-title">Nota</p>
Solo disponible cuando se utiliza la estrategia `concurrente`.
</div>

```sh
video-manager get -s concurrent -m 10 https://video-site.com/video.mp4 https://my.other/video.mp4
```

### Heredado

Mira las [opciones persistentes](./index.md#opciones-persistentes) para todos opciones heredadas soportadas.

<!-- See [persistent options](./index.md#opciones-persistentes) for all the supported inherited options. -->

## See also

Lista completa de [comandos](./index.md).
