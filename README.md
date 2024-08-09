# Houses

[Тестовое задание](https://github.com/avito-tech/backend-bootcamp-assignment-2024/tree/main) для отбора на Avito Backend Bootcamp 2024

## Модель данных

![img](https://github.com/KhuzinT/Houses/blob/main/assets/model.png)

## Endpoint

![img](https://github.com/KhuzinT/Houses/blob/main/assets/endpoint.png)

## Решение

1. Использован данный [API](https://github.com/KhuzinT/Houses/blob/main/api/api.yaml), который немного отличается от исходного
2. Реализована авторизация по почте и паролю через endpoint-ы `/register` и `/login`
   - При регистрации можно передать желаемый тип пользователя: client или moderator
   - При успешной авторизации возвращается токен пользователя
   - Возникла проблема с передачей токена в endpoint-ы, требующие авторизации. Поэтому пока что в некоторых методах `handler` логика запрета доступа закомментирована

## Стек

- **Язык сервиса:** Go, в качестве HTTP-маршрутизатора использовался gorilla/mux
- **База данных:** SQLite + GORM 

## ToDo

- [ ] Решить проблему с передачей токена авторизации в endpoint-ы
- [ ] Реализовать метод `subscribe`
- [ ] Разработать интеграционные и модульные тесты
- [ ] Настроить CI с использованием badges
- [ ] Настроить Docker & DockerCompose
- [ ] Избавиться от GORM и поменять базу на PostgreSQL

Вы можете вернуться сюда завтра и возможно какие-то из этих пунктов уже будут закрыты ¯\_(ツ)_/¯