# SPA base image

A very opinionated server with the only purpose of serving a single page application from the `public` folder where all assets in the sub-folder `public/assets` will be served with long term caching headers.

The server is written in Go and is distributed as a Docker image.

* serves files from `public`
* request that can't be mapped to a file it will serve `/public/index.html`
* any files in `public/assets` will be served with `Cache-Control: public, max-age=31556952`

## Usage

Make a Dockerfile builds a single page application with an index.html entry
point and all of it's assets inside of a public folder:

```
FROM node:22.6.0-alpine AS builder

WORKDIR /app

COPY . .

RUN npm install
RUN npm run build


FROM sunesimonsen/spa-base-image

COPY --from=builder /app/dist public
```

This Dockerfile will use a Node image to build a SPA into a `dist` folder, the `dist` folder is then copied into the server images `public` folder and served from there. Any request that doesn't match a file in the folder will return `index.html`.
