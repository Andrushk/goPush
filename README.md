# Отправка PUSH-уведомлений через FCM
Посредник между вашим приложением и Firebase Cloud Messaging. 
Его задача - связь пользователя (из вашего приложения) с токенами девайсов FCM.

## Зарегистрировать
зарегистрировать пользователя/устройство

```
curl --location --request POST 'http://localhost:8009/register' \
--header 'key: <apikey from goPush settings>' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "UserId": "<user id from your system>",
    "FcmToken": "<Firebase device token>",
    "Device": "Web or Android or IOS"
}'
```

## Отправить
отправить сообщение на все устройства пользователя

```
curl --location --request POST 'http://localhost:8009/send/user' \
--header 'Content-Type: application/json' \
--header 'key: <apikey from goPush settings>' \
--data-raw '{
    "Notification": {
        "title": "Hello",
        "body": "Hello world"
    },
    "Userid" : "<user id from your system>"
}'
```

## Данные пользователя
проверить зарегистрирован ли пользователь и какие у него есть устройства

```
curl --location --request GET 'http://localhost:8009/user?id=<user id from your system>' \
--header 'key: <apikey from goPush settings>' \
--data-raw ''
```
