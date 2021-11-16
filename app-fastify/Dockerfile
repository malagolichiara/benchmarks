FROM node:latest as builder
WORKDIR /app
RUN npm install -g npm@8.1.3
COPY server.js package.json package-lock.json /app/
RUN npm install

FROM node:latest  
WORKDIR /root/
COPY --from=builder /app .
EXPOSE 8000
CMD ["node", "server.js"]

