FROM golang:alpine AS base

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app
RUN mkdir /out

COPY go.mod .
COPY go.sum .
RUN go mod download


FROM base as build
ENV APP_NAME="my-app"

WORKDIR /app
ADD . /app
# Компилируем Go-приложение
RUN go build \
    -o /out/${APP_NAME} cmd/${APP_NAME}/main.go

# Финальный образ, содержащий только исполняемый файл
FROM alpine as relise
WORKDIR /app
COPY --from=build /out/${APP_NAME} /app/

CMD ["/app/my-app"]