
# E-Manager

## Project Overview

E-Manager allows users to create and manage multiple stores, add products, and handle store-related operations.
Each store can contain multiple products. However, before adding a product, you must first create a billboard, category, and size.

Once a product is added, it becomes available for purchase in the corresponding store. The dashboard will then display real-time analytics, including total revenue and sales metrics.

This guide walks you through how to launch this project on a Windows machine. Most of the steps are automated using Docker and .sh files.

## Get Environment Variables

Before running the application, you need to create accounts and obtain API keys from the following services:

* [Cloudinary](https://cloudinary.com/)
* [Stripe](https://stripe.com)
* [Clerk](https://clerk.com/)

Then, populate the environment variables in your .env or directly in the start-project.sh file:

```shell

NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY=<api_key>
CLERK_SECRET_KEY=<api_key>

NEXT_PUBLIC_CLOUDINARY_CLOUD_NAME=<name>
NEXT_PUBLIC_COUDINARY_PRESET=<preset>

STRIPE_API_KEY=<api_key>
STRIPE_WEBHOOK_SECRET=<web_hook>
CLERK_API_KEY_JWKS=<api_key>

```

## Run application

Make sure Docker is installed on your local machine. Once installed:
  
* Start Docker Desktop.

* Open a terminal such as Git Bash (recommended on Windows).

* Run the following command to start the project:

```shell

./start-project.sh

```

To connect the backend with the storefront, clone and run the companion project:

* [E-Commerce](https://github.com/BFaras/E-Commerce)

Follow the setup instructions in that repository to complete the full e-commerce experience.

## Showcase of the whole project

Want to see what the final result looks like? Watch this video showcasing all the features of the E-Manager and E-Commerce apps:





