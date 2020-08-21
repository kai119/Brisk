#!/bin/bash

source scripts/colours.sh

set -e

PROJECT_ROOT=$(cd $(dirname $0) && pwd)

mkdir -p .gomodules/pkg

echo -e "${BLUE}Building Lint image${NC}"
docker build -t lint-image -f Test.Dockerfile .

{
    echo -e "${BLUE}Running go checks...${NC}"
    docker run --rm lint-image /bin/sh -c "go get golang.org/x/lint/golint && golangci-lint run ./..."
} || {
    echo -e "${RED}STAGE FAILED${NC}"
    echo -e "${BLUE}Deleting image${NC}"
    docker rmi lint-image
    exit 1
}

echo -e "${GREEN}STAGE PASSED${NC}"
echo -e "${BLUE}Deleting image${NC}"
docker rmi lint-image

