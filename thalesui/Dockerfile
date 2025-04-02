FROM node:latest AS builder

WORKDIR /app

COPY . .
RUN npm install --legacy-peer-deps
RUN npm run build --prod

# Production stage using Nginx to serve the app
FROM nginx:latest
# Adjust the path to match your Angular build output folder
COPY --from=builder /app/dist/browser /usr/share/nginx/html
EXPOSE 80
