FROM scratch

WORKDIR /app

COPY ./eco /app/run
COPY ./static /app/static
COPY ./keys /app/keys

CMD ["/app/run", "http"]
