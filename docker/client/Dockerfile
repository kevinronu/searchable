# build stage
FROM node:18.17.1-alpine3.18 AS build-stage
WORKDIR /app
COPY .env ./.env.production
COPY ./client/package*.json ./
RUN npm install
COPY ./client .
RUN npm run build

# production stage
FROM nginx:1.25.2-alpine3.18 as production-stage
RUN mkdir /app
COPY --from=build-stage /app/dist /app
COPY ./docker/client/nginx.conf /etc/nginx/nginx.conf
LABEL Name="email-indexer-client"