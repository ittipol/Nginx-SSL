###################
# BUILD FOR LOCAL DEVELOPMENT
###################

FROM node:18-alpine As development

WORKDIR /var/www/nextjs-app

COPY package*.json ./

RUN npm ci

###################
# BUILD FOR PRODUCTION
###################

FROM node:18-alpine As build

WORKDIR /var/www/nextjs-app

COPY package*.json ./

COPY --from=development /var/www/nextjs-app/node_modules ./node_modules 

COPY . .

RUN npm run build

ENV NODE_ENV production

RUN npm ci --omit=dev && npm cache clean --force

###################
# PRODUCTION
###################

FROM node:18-alpine As production

WORKDIR /var/www/nextjs-app

COPY --from=build /var/www/nextjs-app/node_modules ./node_modules
COPY --from=build /var/www/nextjs-app/package*.json ./
COPY --from=build /var/www/nextjs-app/.next ./.next

EXPOSE 3000

CMD [ "npm", "start" ]