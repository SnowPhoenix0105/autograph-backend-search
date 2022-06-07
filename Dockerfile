FROM alpine:3.16.0
WORKDIR /app
COPY ./autograph-backend-search.exe ./autograph-backend-search.exe
CMD ./autograph-backend-search.exe