#!/bin/bash

version=$(grep -E -o "[0-9]+\.[0-9]+\.[0-9]+\+git" ./version/version.go 
)

if [[ "$version" =~ ^(.*\.)([0-9]+)\+git$ ]];
then
	. ./build.sh
	inc=$((BASH_REMATCH[2]+1))
	currentversion="${BASH_REMATCH[1]}${BASH_REMATCH[2]}";
	nextversion="${BASH_REMATCH[1]}$inc";

	git tag -a "v$currentversion" -m "Version $currentversion"
	git push origin "v$currentversion"
	sed -i -e "s/${currentversion}/${nextversion}/" ./version/version.go
	git add ./version/version.go
	git commit -m "Setting version to ${nextversion}"
	git push

	github-release release \
	--user pharmpress \
	--repo helloworld \
	--tag "v$currentversion" \
	--name "v$currentversion" \
	--description "first release!"

    github-release upload \
	--user pharmpress \
	--repo helloworld \
	--tag "v$currentversion" \
	--name "helloworld-linux-amd64" \
	--file bin/helloworld-linux64-static
else
	echo "Version wrong format"
	exit 1
fi
