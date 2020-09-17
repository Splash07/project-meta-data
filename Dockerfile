FROM golang:alpine AS builder

RUN apk add --no-cache gcc libc-dev
RUN apk add tzdata ca-certificates

ENV SOURCES /src
WORKDIR ${SOURCES}
COPY . ${SOURCES}

RUN CGO_ENABLED=0 GOOS=linux go build -a

FROM alpine AS final
ARG PROJECT_NAME
ARG APP_ENV

ENV APPLICATION_ENV ${APP_ENV}

COPY --from=builder ${SOURCES}/${PROJECT_NAME} .
COPY --from=builder ${SOURCES}/config/{$APP_ENV} ./config/${APP_ENV}/

ENTRYPOINT ["${SOURCES}/${PROJECT_NAME}"]
