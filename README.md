# omni-cmd

Lightweight shortcut/bookmark tool for Chrome omnibox.

## Getting Started

### Prerequisites

Go compiler.

### Running locally

Download codes from Github.

Next, add .env file to the project root directory.

```.env
ENV={ runtime environment e.g. "development", "test" or "production" }

# Default port 8008, free to change.
PORT=8008

# Optional, default Google.
SEARCH_ENGINE_URL={ search query string for your preferred search engine }
```

## Deployment

Deploy in container (Build from Dockerfile or pull the pre-build image) or use go compiler to build binary.

### Docker command example

```bash
docker run -d \
--name omni-cmd \
--restart unless-stopped \
-p 8008:8008 \
-v { directory you store your config.yml }:/config
--env-file { your environment file path } \
qiyihuang/omni-cmd:latest
```

## Feature

TBD

## Author

**Qiyi Huang** - [@qiyihuang](https://github.com/qiyihuang)

## Licence

MIT
