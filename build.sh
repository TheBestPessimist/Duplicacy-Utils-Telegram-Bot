#!/bin/sh

docker stop telegram_bot
docker rm telegram_bot

docker build -t telegram_bot --no-cache .

# this is used for production
docker run -dt --restart always -p 13337:13337 --name telegram_bot telegram_bot

# this is used when testing as it attaches to the console
#docker run -p 13337:13337 --name telegram_bot telegram_bot
