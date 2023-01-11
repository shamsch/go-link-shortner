FROM node:alpine as build

# create a build for the svelte app and copy it to the nginx server

WORKDIR /app

COPY package.json .

RUN npm install

COPY . .

RUN npm run build



FROM nginx:alpine

# serve the build from previous stage 

COPY --from=build /app/public /usr/share/nginx/html

EXPOSE 80



