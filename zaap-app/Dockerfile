FROM node AS builder

WORKDIR /app

COPY package.json .
COPY yarn.lock .
RUN yarn install

ENV API_ENDPOINT http://api.zaap.sh
COPY . .
RUN yarn build

FROM nginx

COPY config/nginx.conf /etc/nginx/nginx.conf
COPY --from=builder /app/build /usr/share/nginx/html

RUN chown -R nginx.nginx /usr/share/nginx/html/
