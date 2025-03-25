#!/bin/bash

NETWORK_NAME="app_network"

NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY="<API_KEY>"
CLERK_SECRET_KEY="<API_KEY>"

NEXT_PUBLIC_CLOUDINARY_CLOUD_NAME="<CLOUD_NAME>"
NEXT_PUBLIC_COUDINARY_PRESET="<PRESET>"

STRIPE_API_KEY="<API_KEY>"
STRIPE_WEBHOOK_SECRET="<WEBHOOK_SECRET>"

CLERK_PUBLIC_KEY_JWKS="<API_JWKS_KEY>"
ENV_FILE="./front-end/.env"

if [ -f "$ENV_FILE" ]; then
  echo "⚠️  .env file already exists — it will be overwritten..."
else
  echo "🆕 .env file does not exist — creating a new one..."
fi

cat > "$ENV_FILE" <<EOF
NEXT_PUBLIC_CLERK_SIGN_IN_URL=/sign-in
NEXT_PUBLIC_CLERK_SIGN_UP_URL=/sign-up
NEXT_PUBLIC_CLERK_AFTER_SIGN_IN_URL=/
NEXT_PUBLIC_CLERK_AFTER_SIGN_UP_URL=/

NEXT_PUBLIC_GO_URL=http://back-end:8080
FRONTEND_STORE_URL=http://localhost:3002

NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY=$NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY
CLERK_SECRET_KEY=$CLERK_SECRET_KEY

NEXT_PUBLIC_CLOUDINARY_CLOUD_NAME=$NEXT_PUBLIC_CLOUDINARY_CLOUD_NAME
NEXT_PUBLIC_COUDINARY_PRESET=$NEXT_PUBLIC_COUDINARY_PRESET

DATABASE_URL=postgresql://postgres:Testing12345@db:5432/postgres?schema=public

STRIPE_API_KEY=$STRIPE_API_KEY
STRIPE_WEBHOOK_SECRET=$STRIPE_WEBHOOK_SECRET
EOF


BACKEND_ENV_FILE="./back-end/.bin/.env"

if [ -f "$BACKEND_ENV_FILE" ]; then
  echo "⚠️  Backend .env file already exists — it will be overwritten..."
else
  echo "🆕 Backend .env file does not exist — creating a new one..."
fi

cat > "$BACKEND_ENV_FILE" <<EOF
DB_URL="postgresql://postgres:Testing12345@db:5432/postgres?sslmode=disable"
CLERK_PUBLIC_KEY_JWKS=$CLERK_PUBLIC_KEY_JWKS
EOF

echo "Backend .env file written at $BACKEND_ENV_FILE"

echo " Starting Docker Compose in E-Manager..."
docker-compose up -d