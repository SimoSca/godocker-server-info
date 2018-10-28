################ Versione leggera, con peso ~ 1.7 MB con scratch e ~ 6.1 MB con alpine
# Compile all in this full (and heavy) image
FROM golang as build

COPY ./go/src/github.com/SimoSca/godocker-server-info /go/src/github.com/SimoSca/godocker-server-info
WORKDIR /go/src/github.com/SimoSca/godocker-server-info

# Install MailHog as statically compiled binary:
# ldflags explanation (see `go tool link`):
#   -s  disable symbol table
#   -w  disable DWARF generation
RUN CGO_ENABLED=0 go install -ldflags='-s -w'
# eventualli use the get...
# RUN go get && go install
# # use some compile-time parameters in the build stage to instruct the go compiler to statically link the runtime libraries into the binary itself:
# # RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

# Now use a thin image
# # with scratch the final image is ~ 1.7 MB, ma non ho shell a disposizione
# FROM scratch 
# # with scratch the final image is ~ 6.1 MB, ed ho delle funzionalita' di base (come la shell)
FROM alpine
# ca-certificates are required for the "release message" feature:
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# eventually copy all /go/folder if you wanna preserver the original files
COPY --from=build /go/bin/godocker-server-info /bin/
# User ID 1000 is a workaround for boot2docker issue #581, see
# https://github.com/boot2docker/boot2docker/issues/581
USER 1000
# ENTRYPOINT ["godocker-server-info"]
CMD ["godocker-server-info"]
