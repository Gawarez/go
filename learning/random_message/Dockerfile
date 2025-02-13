FROM golang:alpine AS build
WORKDIR /app
COPY . .
RUN go mod download && go build -o main

FROM scratch AS run
COPY --from=build /app/main ./
CMD [ "/main" ]
