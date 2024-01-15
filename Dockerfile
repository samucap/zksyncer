# Support setting various labels on the final image
#ARG COMMIT=""
#ARG VERSION=""
#ARG BUILDNUM=""
# fix USER ?
# Build Geth in a stock Go builder container
FROM golang:alpine as builder

RUN apk add --no-cache gcc musl-dev linux-headers git make

RUN git clone https://github.com/ethereum/go-ethereum.git /go-ethereum
# Get dependencies - will also be cached if we won't change go.mod/go.sum
WORKDIR /go-ethereum
RUN go mod download
RUN make all

# Pull all binaries into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
#USER nobody
RUN mkdir -p /home/.ethereum/goerli
COPY --from=builder /go-ethereum/build/bin/clef /usr/local/bin/

# TODO setup variables $CHAIN_NAME $CHAIN_ID
CMD [ "/bin/sh" ]
#clef --keystore /home/.ethereum/goerli/keystore --configdir /home/.ethereum/goerli/clef --chainid 300
# Add some metadata labels to help programatic image consumption
#ARG COMMIT=""
#ARG VERSION=""
#ARG BUILDNUM=""
#
#LABEL commit="$COMMIT" version="$VERSION" buildnum="$BUILDNUM"
