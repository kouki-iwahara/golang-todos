FROM mysql:5.7

RUN touch /var/log/mysql/mysqld.log

COPY ./my.cnf /etc/mysql/conf.d/charset.cnf

# COPY ./sql /docker-entrypoint-initdb.d

EXPOSE 3306