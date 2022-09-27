FROM ubuntu:bionic

# Устанавливаем утилиту wget для скачивания файлов с командной строки
RUN apt update && apt install wget -y

# Скачиваем из оф. сайта go бинарник и перемещаем его в директорию /usr/local/bin
RUN wget https://go.dev/dl/go1.19.1.linux-amd64.tar.gz && \
    tar -C /usr/local/ -xzf go1.19.1.linux-amd64.tar.gz && \
    rm go1.19.1.linux-amd64.tar.gz

# Добавляем путь по бинарника go, чтобы запускать go в терминале
ENV PATH="$PATH:/usr/local/go/bin"

# Копируем go.mod и go.sum для установки зависимостей
COPY go.mod /code/
COPY go.sum /code/

# Устанавливаем в качестве рабочей директорий папку /code
WORKDIR /code

# Скачиваем все зависимости проекта
RUN go mod download

# Копируем все файлы с расширение .go внутр образа
COPY *.go /code/

# Собираем go бинарник из нашего проекта. Называем бинарник app
RUN go build -o /code/app

# Говорим докеру чтобы при запуске контейнера запустилась эта команда
CMD ["/code/app"]