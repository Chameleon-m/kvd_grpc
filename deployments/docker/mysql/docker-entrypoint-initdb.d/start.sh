#!/bin/sh
chmod 0444 /etc/mysql/conf.d/my_precedence.cnf

/entrypoint.sh mysqld