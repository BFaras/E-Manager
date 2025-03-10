
# E-Manager

## Project Overview

This project allows users to create multiple stores, add products, and manage store-related operations. Each store can have multiple products, but before adding a product, a billboard, category, and size must be created. Once a product is added, it becomes available for purchase in the store, and the dashboard is updated to provide analytics for the user, including total revenue and sales metrics.

## Environment variables

First step should be to get the environment variables for testing this application. You can do it by creating an account in each of one of the next webiste getting your own api key:

* [Cloudinary](https://cloudinary.com/)
* [Stripe](https://stripe.com)
* [Clerk](https://clerk.com/)

Then we need to choose a password that we will set in the .env file within the front-end, db and backend directorty.

/back-end/.bin/.env

```shell

DB_URL="postgresql://postgres:<password25>@localhost:5432/postgres?sslmode=disable"

```

/front-end/.env

```shell

NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY=<api_key>
CLERK_SECRET_KEY=<api_key>

NEXT_PUBLIC_CLOUDINARY_CLOUD_NAME=<name>
NEXT_PUBLIC_COUDINARY_PRESET=<preset>

DATABASE_URL="postgresql://postgres:<password>@localhost:5432/postgres?schema=public"

STRIPE_API_KEY=<api_key>
STRIPE_WEBHOOK_SECRET=<web_hook>

```

/db/Dockerfile

```shell

ENV POSTGRES_PASSWORD=<password>

```

Once this has been done, we can finally start testing the application. We use a Docker compose and do the following command :

```shell

docker-compose up --build -d

```
Once this has been done, use your browser and enter the url leading you to the front-end of the applicaiton. Then, connect yourself using your own google account and you can finally start testing the CMS application called E-Manager.

The next step should be to set up the e-commerce side using the next link to my next project:

* [E-Commerce](https://github.com/BFaras/E-Commerce)

## Showcase of the whole project

In case you want to see what the final result looks like, you can watch the following video that showcases what it looks like:







