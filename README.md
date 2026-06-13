# This is a docker setup repo that will run a playwright app
## How to use this repo
1. Add this as a submodule to your repo
2. Add the js-helper package to your package.json file
3. Run this to build the docker image:
```bash
cd playwright-docker
docker build -t playwright-docker .
```
---
[API reference](/api-reference.md)
