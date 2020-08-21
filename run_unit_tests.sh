#!/bin/bash

source scripts/colours.sh

set -e

PROJECT_ROOT=$(cd $(dirname $0) && pwd)

mkdir -p .gomodules/pkg

echo -e "${BLUE}Building Unit image${NC}"
docker build -t unit-image -f Test.Dockerfile .

{
    echo -e "${BLUE}Running go checks...${NC}"
    docker run --rm unit-image /bin/sh -c "go test -cover ./... | { grep -v 'no test files'; true; } >> result.txt && cat result.txt && if grep -q FAIL \"result.txt\"; then
    exit 1
    fi"
} || {
    echo -e "${RED}STAGE FAILED${NC}"
    echo -e "${BLUE}Deleting image${NC}"
    docker rmi unit-image
    exit 1
}

echo -e "${GREEN}STAGE PASSED${NC}"
echo "Deleting image"
docker rmi unit-image