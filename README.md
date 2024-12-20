<p align="center">
<img src="docs/logo.svg" width="400">
</p>


# Документация по фреймворку

main.go для запуска приложения

    go run main.go

Файл main трогать не надо. Вся рабочая часть проекта сосредоточена в app domains



## Добавление сервиса

Сервисы лежат в папке domains/ в виде user_service employ_service...
У каждого сервиса свой роутер контроллеры и модеи

    cd  project\cmd\dev\main

    go run struct.go --name=user


## Запуск


    go run main.go


Демо роут будет находится по юрл
 http://localhost:8080/v1/template

## Документация по фреймворку

[-Про деплой](./docs/deploy.md)


 

 

 
