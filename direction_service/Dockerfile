FROM golang:latest

RUN wget https://github.com/Masterminds/glide/releases/download/v0.13.2/glide-v0.13.2-linux-amd64.tar.gz && \
  tar xzvf glide-v0.13.2-linux-amd64.tar.gz && \
  cp linux-amd64/glide /go/bin/glide && \
  rm -r linux-amd64 && rm glide-v0.13.2-linux-amd64.tar.gz

ENV INSTALL_PATH /go/src/direction_service
RUN mkdir $INSTALL_PATH
WORKDIR $INSTALL_PATH
COPY . .
RUN glide install
RUN go build -o /app/main .
CMD ["/app/main"]
