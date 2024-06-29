# serch and go

Search is definitely one of the most interesting domains in computer science.
Search itself could be a critical part of the business model of a company, like Google, or it could be a feature of a product, like the search feature in a code editor.

Implementing a search engine from scratch is a very challenging task.

## Running instructions

Before running the server, copy the `.env.example` file to `.env` and fill in the required values.

```bash
cp .env.example .env
```

Change the values in the `.env` file to match your environment.

Especially, you need to set the `DATABASE_URL` to the URL of your PostgreSQL database.
Change the URL to match your database configuration.

Also, change the `PORT` to the port you want the server to run on, and set the `SECRET_KEY` to a random string.

Then, compile and run the server:

```bash
# Install dependencies
go build

# Run the server
./search-and-go
```
