FROM mcr.microsoft.com/playwright:v1.60.0-jammy

WORKDIR /app

COPY ./src/package.json ./src/package-lock.json ./
COPY ./src/ ./src/
COPY .env .env
COPY /playwright-docker/ /playwright-docker/

RUN npm ci
RUN npx playwright install --with-deps



# folder for remotely managed auth states
RUN mkdir -p /auth

CMD ["node", "src/index.js"]
