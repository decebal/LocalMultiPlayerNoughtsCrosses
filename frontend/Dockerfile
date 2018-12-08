FROM node:alpine as builder

RUN npm i -g yarn

ENV APP_PATH /var/www/frontend

COPY . $APP_PATH
WORKDIR $APP_PATH

RUN yarn

RUN yarn build

FROM node:alpine

RUN npm i -g serve

ENV APP_PATH /var/www/frontend

WORKDIR $APP_PATH

COPY --from=builder /var/www/frontend/dist /var/www/frontend

CMD ["sh", "-c", "serve -p 3000"]
