echo MAJOR_MINOR_REVISION: $MAJOR_MINOR_REVISION

current=$(git describe --tags)

IFS='-' read -ra tags <<< "$current"
current2=${tags[0]}
echo current2: $current2
current_without_v=${current2:1}
echo current_without_v: ${current_without_v}

IFS='.' read -ra nums <<< "$current_without_v"

export MAJOR=${nums[0]}
export MINOR=${nums[1]}
export REVISION=${nums[2]}

if [[ $MAJOR_MINOR_REVISION == "major" ]]; then
    export TRAVIS_TAG=v$((MAJOR+1)).${MINOR}.${REVISION}
fi
if [[ $MAJOR_MINOR_REVISION == "minor" ]]; then
    export TRAVIS_TAG=v${MAJOR}.$((MINOR+1)).${REVISION}
fi
if [[ $MAJOR_MINOR_REVISION == "revision" ]]; then
    export TRAVIS_TAG=v${MAJOR}.${MINOR}.$((REVISION+1))
fi

echo tag: $TRAVIS_TAG
