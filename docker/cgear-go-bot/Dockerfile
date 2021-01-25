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
FROM golang:1.15-alpine3.12 as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /tmp/cgear-go
ADD . .
RUN go build -o /tmp/cgear-go-bot ./go/bot

# Runtime
FROM scratch

ENV DISCORD_TOKEN=""

COPY --from=builder /tmp/cgear-go-bot /app/cgear-go-bot
ENTRYPOINT [ "/app/cgear-go-bot" ]
