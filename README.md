# ASCII Art Web

A web application written in Go that converts user input into ASCII art using different banner styles. It improves the user interface of the ASCII Art Web project by adding responsive CSS styling while preserving the original business logic.

---

## Features

* Convert text to ASCII art
* Supports multiple banner styles:

  * Standard
  * Shadow
  * Thinkertoy
* Web-based user interface
* Input validation
* HTTP error handling
* Clean separation between handlers and business logic
* CSS styling


---

## Project Structure

```
ascii-art-web-stylize/
│
├── banner/
├── handler/
├── logic/
├── static/
│   └── css/
├── templates/
├── main.go
├── go.mod
└── README.md
```

## Requirements

- Go 1.22 or later
- Git


## Installation

Clone the repository

```bash
git clone <repository-url>
```

Enter the project directory

```bash
cd ascii-art-web
```

Run the application

```bash
go run .
```

## Usage

1. Open your browser.

2. Visit

http://localhost:8080

3. Enter text.

4. Select a banner.

5. Click Submit.

6. View the generated ASCII art.


## Example

Input

```
Hello
```

Banner

```
standard
```

Output

```
 _   _
| | | |
| |_| |
...
```

## Error Handling

The application returns appropriate HTTP status codes for:

- Empty input
- Invalid banner
- Unsupported HTTP methods
- Invalid routes
- Non-ASCII characters
- Internal server errors

## Technologies

- Go
- HTML
- CSS
- net/http


## Author

Emmanuel Okonkwo
