# Use official PostgreSQL image
FROM postgres:15-alpine

# Set environment variables to match DATABASE_URL
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=Holoras15
ENV POSTGRES_DB=postgres

# Ensure the schema exists in the database
RUN mkdir -p /docker-entrypoint-initdb.d
COPY initDB.sql /docker-entrypoint-initdb.d/

# Expose PostgreSQL default port
EXPOSE 5432
