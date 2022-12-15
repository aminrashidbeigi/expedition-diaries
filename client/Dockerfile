FROM node:16.3.0-alpine AS BUILDER

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

RUN npm run build

FROM node:16.3.0-alpine

WORKDIR /app

COPY --from=BUILDER /app  .

EXPOSE 3000
CMD [ "npm", "start" ]