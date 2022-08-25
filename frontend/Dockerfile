FROM docker.repos.balad.ir/node:lts AS BUILDER

WORKDIR /app

COPY . .

RUN yarn install \
  --prefer-offline \
  --frozen-lockfile \
  --non-interactive \
  --production=false

RUN yarn build

RUN rm -rf node_modules && \
  NODE_ENV=production yarn install \
  --prefer-offline \
  --pure-lockfile \
  --non-interactive \
  --production=true

FROM docker.repos.balad.ir/node:lts

WORKDIR /app

COPY --from=BUILDER /app  .

EXPOSE 3000
CMD [ "yarn", "start" ]