# ASCII Art Web

A web application written in Go that converts user input into ASCII art using different banner styles. The application provides a simple browser interface where users can enter text, choose a banner, and instantly generate ASCII art.

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

---

## Project Structure

```text
ascii-art-web/
│
├── banner/
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
│
├── handler/
│   └── handler.go
│
├── logic/
│   └── ascii.go
│
├── templates/
│   └── index.html
│
├── main.go
├── go.mod
└── README.md
```

---

## Architecture

The application follows a simple layered architecture.

```text
Browser
    │
    ▼
HTTP Server (main.go)
    │
    ▼
Handlers
    │
    ▼
Business Logic
    │
    ▼
Banner Files
```

### Responsibilities

**main.go**

* Starts the web server.
* Registers application routes.

**handler**

Responsible for:

* Receiving HTTP requests.
* Validating user input.
* Calling the business logic.
* Returning HTTP responses.

**logic**

Responsible for:

* Loading banner files.
* Validating characters.
* Converting text into ASCII art.
* Returning the generated output.

---

## Routes

### GET /

Displays the home page.

### POST /ascii-art

Receives:

* text
* banner

Returns the generated ASCII art.

---

## Supported Banner Styles

* standard
* shadow
* thinkertoy

---

## Error Handling

| Status Code | Description                                                      |
| ----------- | ---------------------------------------------------------------- |
| 200         | Request completed successfully                                   |
| 400         | Invalid input (empty text, invalid banner, non-ASCII characters) |
| 404         | Route not found                                                  |
| 405         | Invalid HTTP method                                              |
| 500         | Internal server error                                            |

---

## Running the Project

Clone the repository.

```bash
git clone <repository-url>
```

Navigate to the project.

```bash
cd ascii-art-web
```

Run the application.

```bash
go run .
```

Open your browser.

```text
http://localhost:8080
```

---

## Example

Input

```text
Hello
```

Banner

```text
shadow
```

Output

```text
<ASCII Art Output>
```

---

## Engineering Decisions

* Reused the ASCII generation logic from the CLI project instead of duplicating code.
* Separated HTTP handling from business logic.
* Kept banner loading inside the business logic layer.
* Used HTML templates for rendering.
* Implemented HTTP status codes according to REST principles.

---

## Lessons Learned

During this project I learned:

* Building HTTP servers in Go.
* Routing requests.
* Creating HTTP handlers.
* Rendering HTML templates.
* Reusing business logic.
* Input validation.
* Debugging file path issues.
* Returning appropriate HTTP status codes.
* Organizing Go projects using packages.

---

## Future Improvements

* Improve UI styling.
* Add downloadable ASCII art.
* Add support for additional fonts.
* Cache banner files for improved performance.
* Add automated tests.

---

## Author

**Emmanuel Okonkwo**

Backend Software Engineering Student

Learn2Earn Fellowship
