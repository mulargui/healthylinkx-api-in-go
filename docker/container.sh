sudo docker stop golangapi
sudo docker rm golangapi
sudo docker run -ti -p :8081:8081 -v /vagrant/apps/healthylinkx-api-in-go:/myapp/api --name golangapi --link MySQLDB:MySQLDB golang /bin/bash