# Отправка PUSH-уведомлений через FCM

## Зарегистрировать
зарегистрировать пользователя/устройство

```
curl --location --request POST 'http://localhost:8009/register' \
--header 'key: <your goPush server application key (look goPush settings)>' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "UserId": "<user id from your system>",
    "FcmToken": "<Firebase device token>",
    "Device": "Web or Android or IOS"
}'
```

## Отправить

POST:send - отправить сообщение пользователю
