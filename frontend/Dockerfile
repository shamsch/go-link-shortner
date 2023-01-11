FROM node:alpine AS build

# create a build for the svelte app and copy it to the nginx server

WORKDIR /app

COPY package.json .

RUN npm install

COPY . .

RUN npm run build

ENTRYPOINT [ "node", "./build"] 



