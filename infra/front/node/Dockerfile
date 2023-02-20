FROM node:18.14-alpine

WORKDIR /var/www/html

RUN apk update && \
  npm install -g npm

EXPOSE 8000
ENV HOST 0.0.0.0

CMD ["npm", "run", "dev"]