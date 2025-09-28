<!-- markdownlint-disable MD033 -->

# Встановлення

Ця інструкція допоможе вам встановити та налаштувати програмне забезпечення для роботи з проектом.

## Зміст

<!--toc:start-->
<!-- - [Встановлення](#встановлення) -->
<!--   - [Зміст](#зміст) -->
<!--   - [Нативно з Go](#нативно-з-go) -->
<!--   - [За допомогою Homebrew](#за-допомогою-homebrew) -->
<!--   - [Для всіх операційних систем](#для-всіх-операційних-систем) -->
<!--     - [GNU/Linux або MacOS](#gnulinux-або-macos) -->
<!--     - [WindowsOS](#windowsos) -->
<!--toc:end-->

## Нативно з Go

```sh
# Встановлює бінарний файл в $GOPATH/bin
go install github.com/caffeine-addictt/video-manager@latest

# Додати до PATH
echo "alias video-manager=$(go env | grep GOPATH)/bin/video-manager" >> ~/.bashrc

# Перезавантажити ресурси
source ~/.bashrc
```

## За допомогою Homebrew

```sh
brew tap caffeine-addictt/tap
bew install caffeine-addictt/tap/video-manager
```

## Для всіх операційних систем

Завантажте відповідні tarballs для вашої ОС на нашій [сторінці релізів](https://github.com/caffeine-addictt/video-manager/releases).

### GNU/Linux або MacOS

```sh
tar video-manager_* video-manager && mv video-manager /usr/local/bin/video-manager
```

### Windows

```sh
# Розпакувати tarball
unzip video-manager_*.zip -d video-manager && setx path "%path%;%cd%\video-manager\"
```

Далі перейдіть до [початку роботи](./getting-started.md).
