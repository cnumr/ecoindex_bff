FROM golang:buster AS build

WORKDIR /app
ADD ./ ./
RUN go mod download
RUN go build -o /ecoindex-bff


FROM gcr.io/distroless/base-debian10

WORKDIR /
COPY --from=build /ecoindex-bff /ecoindex-bff
EXPOSE 3000
USER nonroot:nonroot
ENTRYPOINT ["/ecoindex-bff"]
