FROM node:18-alpine

ENV MODE_ENV=development

WORKDIR /app

COPY ["package.json", "package-lock.json*", "./"]

RUN npm install -y

COPY . .

CMD ["node", "index.js"]