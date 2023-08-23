![code coverage](https://github.com/doovel/cicd/actions/workflows/ci.yml/badge.svg)

# CICD-mini-course


## Local Development

Make sure you're on Go version 1.21+.

Create a `.env` file in the root of the project with the following contents:

```bash
PORT="8000"
```

Run the server:

```bash
go build -o notely && ./notely
```

*This starts the server in non-database mode.* It will serve a simple webpage at `http://localhost:8000`.

