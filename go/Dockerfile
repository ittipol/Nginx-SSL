FROM golang:1.19  as build

WORKDIR /go-app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o ./app

# =================================================
FROM golang:1.19  as production

WORKDIR /go-app

COPY --from=build /go-app/app ./
COPY --from=build /go-app/config.yaml ./

ENV APP_PORT=5000
ENV APP_JWT_ACCESS_TOKEN_SECRET=3b5b95bf3fe81c8e1021835d71cc81d3f94c224f915bb2d583143b5d7db7ded8
ENV APP_JWT_REFRESH_TOKEN_SECRET=a47ec1e9063d2dca33c08fb2d21d65f39daf0a05c88f1959f1985c83c1015929

ENV DATABASE_HOST=db

EXPOSE 5000

ENTRYPOINT ["./app"]