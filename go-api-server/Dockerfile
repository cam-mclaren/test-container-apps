FROM ubuntu:noble 

WORKDIR /usr/local/app

COPY main ./

EXPOSE 8080

RUN useradd app
USER app

CMD ["/usr/local/app/main"]
