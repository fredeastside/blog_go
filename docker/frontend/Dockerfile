FROM node:alpine AS frontend
WORKDIR /app
COPY ./frontend ./
RUN apk add --no-cache --virtual .gyp \
        python \
        make \
        g++
RUN npm install
RUN npm run build

FROM nginx:latest
COPY --from=frontend /app/dist /usr/share/nginx/html