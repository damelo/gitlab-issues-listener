FROM ubuntu:16.04
MAINTAINER Daniel Melo "daniel.melo@serpro.gov.br"
RUN apt-get -y update && apt-get -y upgrade && apt-get -y install python-pip python-requests python-yaml && apt-get -y clean
RUN pip install flask 

COPY *.py start.sh /app/
#COPY ./ca/ca.estaleiro.serpro.pem /usr/local/share/ca-certificates/ca.estaleiro.serpro.crt
#RUN chmod 644 /usr/local/share/ca-certificates/ca.estaleiro.serpro.crt
#COPY ./ca/ca.crt /usr/share/ca-certificates/ca.crt

#RUN echo "ca.crt" >> /etc/ca-certificates.conf
#RUN update-ca-certificates

WORKDIR /app
ENV FLASK_APP="webapp.py"
RUN chmod +x start.sh
#ENTRYPOINT "./start.sh"
