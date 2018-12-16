FROM wrfly/glide
ENV PKG /go/src/github.com/wrfly/et
COPY . ${PKG}
RUN cd ${PKG} && \
    glide i && \
    make test && \
    make build && \
    mv ${PKG}/bin/et /

FROM alpine
RUN apk add --update ca-certificates
COPY --from=0 /et /usr/local/bin/
VOLUME [ "/data" ]
CMD [ "et"]
