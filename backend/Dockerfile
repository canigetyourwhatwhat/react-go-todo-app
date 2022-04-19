# final version
# 10 MB

FROM golang:alpine As builder

RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

# Since scratch image doesn't have much stuff, this file 
# creates user to run the program from inside the container
ENV USER=appuser
ENV UID=10001

RUN adduser \    
    --disabled-password \    
    # don't use the password
    --gecos "" \    
    # additional information anout creating user
    --home "/goHome" \    
    # it becomes the home directory
    --shell "/sbin/nologin" \   
    # no one can modify the contianer by "nologin" 
    # the purpose of creating this user is to only run 
    # this container, so it's fine
    --no-create-home \    
    # tells not to create home directory
    --uid "${UID}" \    
    # uid is user ID, it sets up the user id for this user
    "${USER}"
    # the name of this user

WORKDIR /go/src/app/
COPY . .

RUN go get -u github.com/gin-gonic/gin


RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o webserver ./httpd
#  CGO_ENABLED=0
#  This means that the resulting binary will not be linked to any C libraries. 
#  If your application uses any system libraries (like the system’s cryptography 
#  library) then you will not be able to statically compile the binary like this.


FROM scratch

USER appuser:appuser
# adds the user and group as appuser

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# copy the certificate
COPY --from=builder /etc/passwd /etc/passwd
# copy the account information. That's why I did long useradd commmand on the top
COPY --from=builder /etc/group /etc/group
# copy the group as well

COPY --from=builder /go/src/app/webserver /bin/webserver
# It is executable, so I put under bin

ENTRYPOINT [ "/bin/webserver" ] 