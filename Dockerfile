# Stage 1: Build executable
FROM golang:1.11.4 as buildImage
 
ENV OS linux
ENV CGO_ENABLED 0
ENV PKGPATH github.com/inigohu/healthchecker


# copy current workspace
COPY . ${GOPATH}/src/${PKGPATH}
WORKDIR ${GOPATH}/src/${PKGPATH}

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o health-check 

# Stage 2: Create release image
FROM scratch as releaseImage

COPY --from=buildImage /go/src/github.com/inigohu/healthchecker/health-check ./healthcheck

ENV HOST localhost
ENV PORT 8080
ENV INTERVAL 10

ENTRYPOINT [ "/healthcheck" ]
