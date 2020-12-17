FROM golang:1.15-alpine

COPY . /src 
WORKDIR /src 
RUN go install /src
ENTRYPOINT ["alphabet-echo"]
