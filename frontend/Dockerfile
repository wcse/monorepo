FROM node:alpine as build

WORKDIR /usr/src/app

COPY . .

RUN yarn install --pure-lockfile --non-interactive

WORKDIR /usr/src/app/wetalk
RUN yarn build

#FROM node:alpine
#
#WORKDIR /usr/src/app
#
#COPY package.json .
#COPY yarn.lock .
#
#RUN mkdir -p /usr/src/app/wetalk
#COPY --from=build /usr/src/app/package.json /usr/src/app/
#COPY --from=build /usr/src/app/node_modules /usr/src/app/
#COPY --from=build /usr/src/app/wetalk/.next /usr/src/app/wetalk/
#COPY --from=build /usr/src/app/wetalk/package.json /usr/src/app/wetalk/
#COPY --from=build /usr/src/app/wetalk/node_modules /usr/src/app/wetalk/
#
#ENV NODE_ENV production
#
#WORKDIR /usr/src/app/wetalk

CMD ["yarn", "start"]
