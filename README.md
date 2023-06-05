# Pillar

<!-- BADGES -->
<p align="center">
  <img src="https://img.shields.io/github/actions/workflow/status/mohammadne/pillar/ci.yaml?label=ci&logo=github&style=for-the-badge&branch=main" alt="GitHub Workflow Status" />
  <img src="https://img.shields.io/github/release/mohammadne/pillar.svg?style=for-the-badge" alt="Release">
  <img src="https://img.shields.io/codecov/c/gh/mohammadne/pillar?logo=codecov&style=for-the-badge" alt="Codecov">
  <img src="https://img.shields.io/github/license/mohammadne/pillar?style=for-the-badge" alt="Licanse">
  <img src="https://img.shields.io/github/stars/mohammadne/pillar?style=for-the-badge" alt="Starts">
  <img src="https://img.shields.io/github/downloads/mohammadne/pillar/total.svg?style=for-the-badge" alt="Releases">
</p>

## Introduction

Simple Go application for using as starting point for writing Go applications.

## Usage

```bash
# Define which package-name to use
REPLACE="username/repository"

# Update package-name
# MacOS note: use gsed instead of sed (brew install gnu-sed)
LC_ALL=C find . -type f -exec sed -i "s#mohammadne/pillar#$REPLACE#g" {} +

# Update descriptions
vim main.go
vim cmd/server.go

# Troubleshooting
docker pull "ghcr.io/$REPLACE:latest"
docker run -d -p 8090:8090 "ghcr.io/$REPLACE:latest"

# Delete This file and write your own README.md file
rm README.md
```
