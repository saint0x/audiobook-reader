# PDF Player Backend

This is the backend service for the PDF Player application. It provides APIs for uploading PDFs, extracting text, managing books, and tracking reading progress.

## Prerequisites

- Go 1.16 or later
- SQLite3

## Setup

1. Install dependencies:
```bash
go mod download
```

2. Run the server:
```bash
go run *.go
```

The server will start on port 8080 by default. You can change the port by setting the `PORT` environment variable.

## API Endpoints

### Books
- **POST** `/api/upload` - Upload a PDF file
  - Content-Type: `multipart/form-data`
  - Form field: `file` (PDF file)
  - Returns: Book object with extracted text

- **GET** `/api/books` - Get all books
  - Returns: Array of book objects

- **GET** `/api/book/{id}` - Get a specific book
  - Returns: Single book object

### Reading Progress
- **PUT** `/api/book/{id}/progress` - Update reading progress
  - Body: `{ "currentPage": number, "completion": number }`

- **GET** `/api/book/{id}/progress` - Get reading progress
  - Returns: Reading progress object

### Bookmarks
- **POST** `/api/book/{id}/bookmarks` - Create a bookmark
  - Body: `{ "pageNumber": number, "note": string }`
  - Returns: Created bookmark object

- **GET** `/api/book/{id}/bookmarks` - Get all bookmarks for a book
  - Returns: Array of bookmark objects

### Audio Segments
- **POST** `/api/book/{id}/audio` - Create an audio segment
  - Body: `{ "startPage": number, "endPage": number }`
  - Returns: Created audio segment object

- **PUT** `/api/book/{id}/audio/{segmentId}` - Update audio segment status
  - Body: `{ "status": string, "audioUrl": string, "duration": number }`

### Categories
- **POST** `/api/categories` - Create a category
  - Body: `{ "name": string, "description": string }`
  - Returns: Created category object

- **PUT** `/api/book/{id}/categories/{categoryId}` - Add book to category

### Tags
- **POST** `/api/tags` - Create a tag
  - Body: `{ "name": string }`
  - Returns: Created tag object

- **PUT** `/api/book/{id}/tags/{tagId}` - Add tag to book

## Data Models

### Book
```json
{
  "id": "string",
  "title": "string",
  "author": "string",
  "coverUrl": "string",
  "content": "string",
  "filePath": "string",
  "pageCount": number,
  "currentPage": number,
  "language": "string",
  "createdAt": "datetime",
  "updatedAt": "datetime",
  "categories": ["string"],
  "tags": ["string"]
}
```

### Reading Progress
```json
{
  "id": "string",
  "bookId": "string",
  "currentPage": number,
  "completionPercent": number,
  "lastReadAt": "datetime"
}
```

### Bookmark
```json
{
  "id": "string",
  "bookId": "string",
  "pageNumber": number,
  "note": "string",
  "createdAt": "datetime"
}
```

### Audio Segment
```json
{
  "id": "string",
  "bookId": "string",
  "segmentNumber": number,
  "content": "string",
  "audioUrl": "string",
  "duration": number,
  "status": "string",
  "createdAt": "datetime"
}
```

## Future Enhancements

1. PDF metadata extraction for better book information
2. Custom book cover image upload
3. Text-to-speech synthesis integration using Kokoro TTS
4. Book categories and tags for better organization
5. Full-text search functionality 