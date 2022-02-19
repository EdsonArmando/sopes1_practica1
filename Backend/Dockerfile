FROM golang
WORKDIR /app/src/hardware-store
COPY . /app/src/hardware-store
RUN go build -o main .
EXPOSE 6000
CMD [ "./main" ]