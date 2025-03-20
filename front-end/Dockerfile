# Use official Node.js image
FROM node:18-alpine

# Set working directory
WORKDIR /app

# Copy package.json and package-lock.json
COPY package.json package-lock.json ./

# Install dependencies
RUN npm install && npm install jwt-decode

# Copy the rest of the application
COPY . .

# Expose port 3000
EXPOSE 3000

# Run Next.js in development mode
CMD ["npm", "run", "dev"]
