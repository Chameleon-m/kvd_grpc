FROM mysql:8.0.32

# Copy starting scripts file
COPY ./deployments/docker/mysql/docker-entrypoint-initdb.d/start.sh /root/start.sh
COPY ./deployments/docker/mysql/conf.d/my_precedence.cnf /etc/mysql/conf.d/my_precedence.cnf

# Run necessary services
CMD ["/bin/bash", "/root/start.sh"]