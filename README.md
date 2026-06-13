# This is a docker setup repo that will run a playwright app
## How to use this repo
1. Add this as a submodule to your repo
```bash
git submodule add https://github.com/ryanbays/playwright-docker.git
```
2. Add the js-helper package to your package.json file
```bash
npm install --save-dev playwright-docker/js-helper
```
4. Copy the compose file and env from this repo to your repo
```bash
cp playwright-docker/example-compose.yml docker-compose.yml
cp playwright-docker/example.env .env
```
> Note: the docker compose will use `.env` while the docker container will use `.projectenv` file.
5. Build the docker image
```bash
docker compose up -d --build
```
---
[API reference](/api-reference.md)
