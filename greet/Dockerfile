############################
# STEP 1 build executable binary
############################
# golang alpine 1.13.5
FROM golang@sha256:0991060a1447cf648bab7f6bb60335d1243930e38420bee8fec3db1267b84cfa as builder


WORKDIR /opt/greet



COPY . .


ENV GO111MODULE=on
RUN go mod download
RUN go mod verify


RUN go build -i -o server ./greet_server

EXPOSE 50051/tcp



ENTRYPOINT ["/opt/greet/server"]
