FROM scratch
COPY go-discogs /go-discogs
ENTRYPOINT ["/go-discogs"]