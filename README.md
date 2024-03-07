Creating a README file for your project is a great way to document its setup and usage. Here's a basic template you can use:

---

# Go Blog API

This is a simple blog API built with Go and Fiber framework, using MongoDB as the database.

## Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/go-blog-api.git
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Set up environment variables:

   Rename `.env.example` into `.env` and put down the important credentials.

   ```plaintext
   # Database configuration
   MONGO_URI=
   MONGO_DATABASE_NAME=

   # Server configuration
   SERVER_PORT= 

   # JWT configuration
   JWT_SECRET=
   ```

4. Start the MongoDB server:

   Make sure MongoDB is running on your machine or update the `MONGO_URI` in the `.env` file with the MongoDB connection string.

5. Run the application:

   ```bash
   go run main.go
   ```

## Endpoints (To be written soon)

