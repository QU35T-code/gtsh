#!/bin/bash

SRV_PEM="server.pem"
SRV_KEY="server.key"

openssl req -subj '/CN=localhost/O=Localhost/C=US' -new -newkey rsa:4096 -days 3650 -nodes -x509 -keyout ${SRV_KEY} -out ${SRV_PEM}
mv ${SRV_PEM} ${SRV_KEY} certs/