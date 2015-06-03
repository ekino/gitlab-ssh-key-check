# Check SSH Keys on a GitLab server

This is a quick code to check if ssh keys available on a Gitlab server is exposed to similar thread as described in the blog ["Auditing GitHub users’ SSH key quality"](https://blog.benjojo.co.uk/post/auditing-github-users-keys).

## Installation

    go get github.com/ekino/gitlab-ssh-key-check
    
## Configuration

create a config.json file in the ``$GOPATH/src/github.com/ekino/gitlab-ssh-key-check folder``

The configuration is:

    {
      "host": "https://your.gitlab-server.com",
      "api_path": "/api/v3",
      "token": "your admin api token",
      "weak_key": 1023
    }

## Usage

    go run main.go
    
