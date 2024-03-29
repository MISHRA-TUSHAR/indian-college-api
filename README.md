# Indian College API

Indian College API is a web service providing information about colleges in India. It allows users to search for colleges based on various criteria such as state, district, etc.

This API uses data sourced from [data.gov.in/catalog/institutions-aishe-survey] data.gov.in/catalog/institutions-aishe-survey , which has been made publicly available by the government of India.

## Features

- Search for colleges based on different criteria.
- Retrieve a list of all states and districts in India.
- Fetch colleges by state or district.

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/MISHRA-TUSHAR/indian-college-api.git
   ```

2. Install dependencies:
```
go mod tidy
```


3. Run the application:

```
go run main.go
```


The API will start running on `http://localhost:8000`.

## API Endpoints

- `GET /`: Welcome message.
- `GET /colleges/`: Search for colleges.
- `GET /colleges/states`: Get all states in India.
- `GET /colleges/:state/districts`: Get districts by state.
- `GET /colleges/:state`: Get all colleges in a specific state.
- `GET /colleges/:state/:district`: Get all colleges in a specific district of a state.

## Usage

You can send HTTP requests to the above endpoints to interact with the API. Here's an example using cURL:

```
curl http://localhost:8000/colleges/

```

