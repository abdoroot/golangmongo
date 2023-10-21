# Go MongoDB Practice

This Go program is a basic practice project for working with MongoDB using the Go programming language. It demonstrates various MongoDB operations like inserting, updating, deleting, and querying documents in a MongoDB collection.

## Prerequisites

Before running this program, make sure you have the following installed:

- Go: [Download and install Go](https://golang.org/dl/)

- MongoDB: [Download and install MongoDB](https://www.mongodb.com/try/download/community)

- MongoDB Go Driver: Install the official MongoDB Go driver with the following command:

  ```bash
  go get go.mongodb.org/mongo-driver/mongo
  go get go.mongodb.org/mongo-driver/mongo/options
  ```

## Usage

1. Clone this repository or copy the `main.go` code into your Go project.

2. Ensure your MongoDB server is running at `localhost:27017`. You can adjust the MongoDB connection URI in the code if needed.

3. In your project directory, run the following command to execute the program:

   ```bash
   go run main.go
   ```

4. The program will perform MongoDB operations and display the results in the console.

## Functions

The program showcases MongoDB operations including:

- Inserting a single document.
- Inserting multiple documents.
- Updating a document.
- Deleting a document.
- Querying documents in the collection.

## Customization

Feel free to customize or expand this project to practice more MongoDB operations or tailor it to your specific needs.

## License

This project is open-source and is available under the [MIT License](LICENSE).

Enjoy practicing MongoDB with Go! If you encounter issues or have questions, please don't hesitate to contact us.