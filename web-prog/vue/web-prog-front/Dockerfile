FROM 8.16.2-alpine3.9

RUN npm install -g http-server

WORKDIR /opt/
COPY src .
COPY public .
COPY babel.config.js .
COPY package.json .
COPY yarn.lock .
RUN npm install
RUN npm run build

EXPOSE :8080
CMD [ "http-server", "dist" ]