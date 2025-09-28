<!-- markdownlint-disable MD013 -->

# Команди cache

Ці команди використовуються для управління кешем.

## Зміст

<!--toc:start-->
<!-- - [Команди cache](#команда-cache) -->
<!--   - [Зміст](#зміст) -->
<!--   - [Команди](#команди) -->
<!--     - [list|ls <pattern>](#listls-pattern) -->
<!--     - [remove|rm <pattern>](#removerm-pattern) -->
<!--     - [clear|clr|wipe](#clearclrwipe) -->
<!--   - [Опції](#опції) -->
<!--   - [Дивись також](#дивись-також) -->
<!--toc:end-->

## Команди

### list|ls <pattern>

Виводить всі записи кешу або ті, які відповідають вказаному шаблону.

Кожен запис кешу позбавляється свого протоколу (`https://`, `http://`, `www.`) перед порівнянням.

Наприклад, `https://video-site.com` буде позбавлено до `video-site.com`, а `http://www.other-site.com` буде позбавлено до `other-site.com`.

Шаблони порівнюються в такому порядку:

- Перевірка, чи рядок починається з шаблону
- Порівняння RegExp з рядком

```sh
video-manager cache list
video-manager cache ls myVideoTitle
```

### remove|rm <pattern>

Видаляє всі записи кешу, які відповідають вказаному шаблону.

<div class="admonition NOTE" markdown>
<p class="admonition-title">Note</p>
Протокол (`https://`, `http://`, `www.`) не видаляється перед порівнянням.
</div>

Шаблони порівнюються в такому порядку:

- Перевірка, чи рядок починається з шаблону
- Порівняння RegExp з рядком

```sh
video-manager cache remove https
```

### clear|clr|wipe

Очищає весь кеш.

```sh
video-manager cache clear
```

## Опції

Тільки [постійні опції](./index.md#_2) доступні для всіх команд.

## Дивись також

- [Команда get](./get.md)
