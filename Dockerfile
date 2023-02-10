# lightweight container for go
FROM golang:1.19-alpine

# update container's packages and install git
RUN apk update && apk add --no-cache git

# set /app to be the active directory
WORKDIR /src

# copy all files from outside container, into the container
COPY . .

# download dependencies
RUN go mod tidy

# build binary
RUN go build ./cmd/app/

# set the entry point of the binary
ENTRYPOINT ["/src/app","-prod=true"]