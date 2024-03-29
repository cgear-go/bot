#!/usr/bin/env bash
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

if [ ! -f .makerc ]; then
    echo "Please run this command from repository root directory."
    exit 1
fi

echo "Building Docker image..."
docker build \
	--rm \
    --quiet \
	--label "project=cgear-go-bot" \
	--tag   cgear-go-bot:latest . > /dev/null


echo "Running Docker image..."
echo ""
source .makerc
docker run \
    --rm \
    -e DISCORD_TOKEN="${DISCORD_TOKEN}" \
    -e RAID_CATEGORY_ID="${RAID_CATEGORY_ID}" \
    -e RAID_CHANNEL_ID="${RAID_CHANNEL_ID}" \
    cgear-go-bot

docker image prune \
    --force \
    --filter "label=project=cgear-go-bot" > /dev/null