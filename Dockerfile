#FROM alpine:latest
#WORKDIR /go/src/github.com/kelseyhightower/app/
#COPY . .
#RUN CGO_ENABLED=0 GOOS=linux go build .

# FROM scratch

# EXPOSE 8080
# LABEL maintainer="marcos AT unitron DOT com DOT br"
# COPY ./go-helloword-http /app
# ENTRYPOINT ["/app"]
# RUN ["/app"]


FROM scratch
LABEL AUTHOR="marcos AT unitron DOT com DOT br"
ADD main /
CMD ["/main"]

# this is how you test it:
# docker run  -p 8080:8080 89bbcbfccc03
# curl