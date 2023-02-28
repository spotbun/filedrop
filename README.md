<h1 align="center">Filedrop</h1>
<p align="center">Simple file drop server.</p>

<p align="center">
<a href="https://hub.docker.com/r/spotbun/filedrop"><img alt="Docker Cloud Build Status" src="https://img.shields.io/github/actions/workflow/status/spotbun/filedrop/docker-build.yml?style=for-the-badge"></a>
<a href="https://hub.docker.com/r/spotbun/filedrop"><img alt="Docker Image Size (tag)" src="https://img.shields.io/docker/image-size/spotbun/filedrop/latest?style=for-the-badge"></a>
<a href="https://hub.docker.com/r/spotbun/filedrop"><img alt="Docker Pulls" src="https://img.shields.io/docker/pulls/spotbun/filedrop?style=for-the-badge"></a>
<a href="https://github.com/spotbun/filedrop/blob/main/LICENCE"><img alt="GitHub" src="https://img.shields.io/github/license/spotbun/filedrop?style=for-the-badge"></a>
<br>
<br>
<img alt="Screenshot" width="750" src="https://user-images.githubusercontent.com/31022056/221932050-88935d5c-1f53-400d-9dc0-61b29b0ab6cf.png">
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
