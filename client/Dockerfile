FROM node:alpine AS build

WORKDIR usr/src/app

COPY package.json package-lock.json ./

RUN npm ci

COPY . ./

RUN npm run build --prod

FROM nginx:alpine

COPY --from=build /usr/src/app/dist/app/browser /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]