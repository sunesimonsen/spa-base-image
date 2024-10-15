# SPA base image

A very opinionated server with the only purpose of serving a single page application from the public folder.

The server is written in Go and is distributed as a Docker image.

## Usage

Make a Dockerfile builds a single page application with an index.html entry
point and all of it's assets inside of a public folder:

```
FROM node:22.6.0-alpine AS builder

WORKDIR /app

COPY . .

RUN npm install
RUN npm run build


FROM sunesimonsen/spa-base

COPY --from=builder /app/dist public
```

This Dockerfile will use a Node image to build a SPA into a `dist` folder, the `dist` folder is then copied into the server images `public` folder and served from there. Any request that doesn't match a file in the folder will return `index.html`.
