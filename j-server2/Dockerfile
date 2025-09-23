FROM node:22.14-alpine3.20

WORKDIR /app

COPY package.json ./
RUN npm install && npm cache clean --force

COPY . ./

EXPOSE 4002

CMD ["npm", "run", "start"]