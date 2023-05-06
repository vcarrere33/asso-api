# Rna OpenAPI

🛠 This is currentlty in Work In Progress

## Menu

- [Menu](#menu)
- [📘 Description](#️-📘-description)
- [👷‍♂️ How it's work](#️-how-its-work)
- [🛸 Next Steps](#-🛸-next-steps)

# 📘 Description

This OpenAPI aims to deliver a full text search on the french Rna (french association register).

# 👷‍♂️ How it's work

## 📺 API

This API provides a GET endpoint with a query parameter q to search on the association name using a Mongodb database.

The OpenAPI doc is in progress.

## 💽 Self host

If you want to self host the api, you can clone the repository and compile it using the following command :

```shell
go build
```

If you need more information you can look on golang [documentation](https://go.dev/doc/tutorial/compile-install).

You will also need a working mongodb running with a database named asso.
If you want to update the DB record you can use the following command :

```shell
go run ./cmd/importer/main.go
```

This command will download the last Rna csv, unzip it and upload all the datas to the database

# 🛸 Next Steps

The project is still in progress, this will be the next things that I will develop :

- Remove package with vulnerabilities
- Add environement variables for databse url
- Add OpenAPI documentation
- Find a way to automatically get the last up to date csv
- Add Ip check on request to avoid DDOS
- Add unit test
- Add search on other fields
