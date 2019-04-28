FROM alpine
RUN apk add --update ca-certificates
COPY bin/et /usr/local/bin/
VOLUME [ "/data" ]
CMD [ "et"]
