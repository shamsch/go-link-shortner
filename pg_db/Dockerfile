# postgres database
FROM postgres:alpine

# set environment variables

ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=postgres
ENV POSTGRES_DB=postgres

# initialize database

COPY ./init.sql /docker-entrypoint-initdb.d/

# expose port 5432 to the host

EXPOSE 5432

# start postgres

CMD ["postgres"]