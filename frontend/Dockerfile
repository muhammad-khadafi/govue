# Based from: https://cli.vuejs.org/guide/deployment.html#docker-nginx
# Bundle the built application into an nginx image
FROM docker.io/nginx:1.20.2-alpine
WORKDIR /application
COPY nginx.conf /etc/nginx/nginx.conf
COPY dist ./
