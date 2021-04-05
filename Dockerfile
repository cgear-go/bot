#   Copyright 2020 Pokémon GO Nancy
#
#   Licensed under the Apache License, Version 2.0 (the "License");
#   you may not use this file except in compliance with the License.
#   You may obtain a copy of the License at
#
#       http://www.apache.org/licenses/LICENSE-2.0
#
#   Unless required by applicable law or agreed to in writing, software
#   distributed under the License is distributed on an "AS IS" BASIS,
#   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#   See the License for the specific language governing permissions and
#   limitations under the License.

# Build
FROM golang:1.16-alpine3.12 as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

ENV USER=cgear-go
ENV UID=10001
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

WORKDIR /tmp/cgear-go

ADD discord     discord
ADD raid        raid
ADD bot.go      bot.go
ADD go.mod      go.mod
ADD go.sum      go.sum

RUN go mod download
RUN go mod verify
RUN go build -ldflags="-w -s" -o /tmp/cgear-go-bot .

# Runtime
FROM scratch

ENV DISCORD_TOKEN=""

# Runtime dependencies
COPY --from=builder /etc/ssl/certs/ca-certificates.crt  /etc/ssl/certs/
COPY --from=builder /etc/passwd                         /etc/passwd
COPY --from=builder /etc/group                          /etc/group

# Binary
COPY --from=builder /tmp/cgear-go-bot /app/cgear-go-bot

USER cgear-go:cgear-go

ENTRYPOINT [ "/app/cgear-go-bot" ]
