FROM docker.repos.balad.ir/node:lts-alpine AS BUILDER

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

RUN npm run build

FROM docker.repos.balad.ir/node:lts-alpine

WORKDIR /app

COPY --from=BUILDER /app  .

EXPOSE 3000
CMD [ "npm", "start" ]