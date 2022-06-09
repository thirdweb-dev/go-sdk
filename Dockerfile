FROM node:14-buster

WORKDIR /code/hardhat

COPY hardhat.config.js /code/hardhat/hardhat.config.js
COPY package.json /code/hardhat/package.json

RUN yarn

COPY /scripts/test/start-hardhat.sh /code/hardhat/start-hardhat.sh
RUN chmod +x /code/hardhat/start-hardhat.sh

ENTRYPOINT ["/code/hardhat/start-hardhat.sh"]