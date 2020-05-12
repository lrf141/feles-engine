FROM ubuntu:18.04

RUN apt-get update && apt-get -y upgrade && \
    	apt-get install -y software-properties-common && \
	add-apt-repository -y ppa:longsleep/golang-backports

RUN apt-get update && \
	apt-get -y install git curl golang-1.12-go uuid-dev
	
ENV PATH $PATH:/usr/lib/go-1.12/bin
ENV GO111MODULE on

WORKDIR /src 
