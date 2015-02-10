FROM ubuntu:latest

RUN DEBIAN_FRONTEND=noninteractive 

# Install git and supplementary programs
RUN apt-get -qq -y install git
RUN git config --global user.name "mulargui"
RUN git config --global user.email mauricio@ulargui.com

# app dir
RUN sudo mkdir /myapp
RUN sudo mkdir /myapp/src
WORKDIR /myapp/api
ENV GOPATH /myapp/src

# Install golang and supplementary programs
RUN apt-get -qq update && apt-get -y -qq install golang
RUN go get github.com/go-sql-driver/mysql
 
EXPOSE 8081

# By default, simply start the app
CMD /bin/bash /myapp/api/run.sh
