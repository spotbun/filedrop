# filedrop
Simple file drop server.

<p align="center">
<a href="https://hub.docker.com/r/spotbun/filedrop"><img alt="Docker Cloud Build Status" src="https://img.shields.io/github/actions/workflow/status/spotbun/filedrop/docker-build.yml?style=for-the-badge"></a>
<a href="https://hub.docker.com/r/spotbun/filedrop"><img alt="Docker Image Size (tag)" src="https://img.shields.io/docker/image-size/spotbun/filedrop/latest?style=for-the-badge"></a>
<a href="https://hub.docker.com/r/spotbun/filedrop"><img alt="Docker Pulls" src="https://img.shields.io/docker/pulls/spotbun/filedrop?style=for-the-badge"></a>
<a href="https://github.com/spotbun/filedrop/blob/main/LICENCE"><img alt="GitHub" src="https://img.shields.io/github/license/spotbun/filedrop?style=for-the-badge"></a>
</p>


## Installation

### Docker

```bash
#                            ┌-- Path to store uploaded files
docker run -d -p 80:80 -v ${pwd}/uploads:/uploads spotbun/filedrop
```

### Docker-Compose

```yaml
version: "3.7"

services:
  filedrop:
    image: spotbun/filedrop
    ports:
      - "80:80"
    volumes:
      - ./uploads:/uploads
```

## Description

Filedrop is a simple file collecting server.
It can be used to quickly spin up a website that allows everyone to anonymously upload files to your server.


## Use Cases
Filedrop provides a range of use cases, including:

- **Anonymous Uploading**: You can easily upload files to the server without revealing your identity. This feature makes it an excellent tool for sensitive information sharing.
- **Uploading Files to Your Own Server**: Filedrop can be used to upload files to your own server.
- **Saving Files on the Go**: Filedrop allows you to quickly save files on the go without the need to log in - no matter on what device you are.

## Configuration

Filedrop does not need any configuration.
If you want to add authentication, you can use a reverse proxy like [traefik](https://traefik.io/traefik/) or [nginx](https://www.nginx.com/).
