FROM node:alpine

WORKDIR /app

COPY package*.json .

EXPOSE 3000

RUN npm ci

COPY . .

CMD [ "npm", "run", "dev" ]
