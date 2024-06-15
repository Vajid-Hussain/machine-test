FROM golang:1.22-alpine AS first-build
COPY ./ /machine-test
WORKDIR /machine-test
RUN go mod download
RUN go build -o build cmd/main.go

FROM scratch
COPY --from=first-build /machine-test/build /
COPY --from=first-build /machine-test/env.json /
ENTRYPOINT [ "/build" ]
