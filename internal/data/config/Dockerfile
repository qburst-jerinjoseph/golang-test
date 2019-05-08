FROM postgres:9.6.6-alpine

ARG POSTGRES_USER
ARG POSTGRES_PASSWORD
ARG POSTGRES_DB

ENV POSTGRES_USER ${POSTGRES_USER}
ENV POSTGRES_PASSWORD ${POSTGRES_PASSWORD}
ENV POSTGRES_DB ${POSTGRES_DB}
ENV PGDATA /var/lib/postgresql/data/pgdata

COPY set-config.sh /docker-entrypoint-initdb.d/_set-config.sh

# install all the schema and migration table
ADD ./migrations/*up.sql /docker-entrypoint-initdb.d/
RUN echo $(ls /docker-entrypoint-initdb.d | egrep -o '^[0-9]+' | tail -1) > /.latest_version
ADD ./seeds/*up.sql /docker-entrypoint-initdb.d/
RUN sed -i -e "s/__VER__/$(cat /.latest_version)/g" /docker-entrypoint-initdb.d/9999999999_initial_migration.up.sql