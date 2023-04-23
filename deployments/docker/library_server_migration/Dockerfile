FROM migrate/migrate:v4.15.2

COPY ./migrations ./migrations
COPY ./Makefile ./Makefile

RUN apk update
RUN apk add make # install make