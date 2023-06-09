# Pillar

<!-- BADGES -->
<p align="center">
  <img src="https://img.shields.io/github/actions/workflow/status/mohammadne/pillar/ci.yaml?label=ci&logo=github&style=for-the-badge&branch=main" alt="GitHub Workflow Status" />
  <img src="https://img.shields.io/github/release/mohammadne/pillar.svg?style=for-the-badge" alt="Release">
  <a href="https://codecov.io/gh/mohammadne/pillar">
    <img src="https://img.shields.io/codecov/c/gh/mohammadne/pillar?logo=codecov&style=for-the-badge" alt="Codecov" />
  </a>
  <img src="https://img.shields.io/github/license/mohammadne/pillar?style=for-the-badge" alt="License">
  <img src="https://img.shields.io/github/stars/mohammadne/pillar?style=for-the-badge" alt="Starts">
  <img src="https://img.shields.io/github/downloads/mohammadne/pillar/total.svg?style=for-the-badge" alt="Downloads">
</p>

## Introduction

Simple Go application for using as starting point for writing Go applications.

## Usage

1. Clone the repository

    ```bash
    git clone git@github.com:mohammadne/pillar.git
    ```

2. Change the remote git-origin (replace your own)

    ```bash
    REMOTE_ORIGIN="git@github.com:mohammadne/pillar.git"
    git remote set-url origin $REMOTE_ORIGIN
    ```

3. Update package-name (replace your own)

    ```bash
    REPLACE="username/repository"
    LC_ALL=C find . -type f -not -path ".git/*" -exec sed -i "s#mohammadne/pillar#$REPLACE#g" {} +
    ```

4. Update command descriptions

    ```bash
    vim main.go
    vim cmd/server.go
    ```

5. Update deployment name and create helm charts

    ```bash
    cd ./deployments

    # update 
    vim ./tearup.sh
    vim ./teardown.sh
    
    # replace your own application name.
    helm create pillar

    # update folder names based on the commands.
    mv -rf ./pillar ./server
    ```

6. Update or delete this README.md file based on your needs
