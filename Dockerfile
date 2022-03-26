FROM node:alpine as build
WORKDIR /koodi
COPY package.json /koodi
RUN npm i
COPY . /koodi
RUN npm run build

FROM node:alpine
COPY ./package.json ./package.json
COPY --from=build /koodi/lib/*js /app/
RUN npm i --release
ENV ES_NODE "http://localhost:9200"
ENV PORT "3000"
EXPOSE 3000

CMD [ "/usr/local/bin/node", "/app" ]