FROM mcr.microsoft.com/playwright:v1.49.0-jammy

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm ci

COPY ../src/ .

# folder for remotely managed auth states
RUN mkdir -p /auth

CMD ["node", "index.js"]
