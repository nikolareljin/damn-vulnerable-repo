#!/usr/bin/env bash
GITHUB_TOKEN="*****"
curl -H "Authorization: token $GITHUB_TOKEN" https://api.github.com/user
