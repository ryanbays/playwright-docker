# API Reference

## Overview

This API provides:
- status checking
- upload/remove auth state files
- read/delete API logs
- read/delete worker logs

All responses are JSON unless otherwise noted.

---

## `GET /status`

Returns service health status.

### Response
```json
{
  "status": "ok"
}
```

### Status Codes
- `200 OK`

---

## `POST /upload`

Uploads a file and saves it as a new JSON file under `/auth/` with a generated UUID filename.

### Request
- `Content-Type: multipart/form-data`
- Form field:
  - `file`: file to upload

### Response
```json
{
  "message": "File uploaded successfully",
  "file": "original-filename.ext",
  "id": "generated-uuid"
}
```

### Status Codes
- `200 OK` — file saved successfully
- `400 Bad Request` — no file provided
```json
{
  "error": "No file is received"
}
```
- `500 Internal Server Error` — upload/save failed
```json
{
  "error": "Unable to save the file"
}
```

### Notes
- File is saved to `/auth/<uuid>.json`
- The returned `id` is the generated UUID

---

## `POST /remove/:id`

Removes a saved auth state file by ID.

### Path Parameters
- `id`: UUID of the file to remove

### Behavior
Attempts to delete:
- `/auth/<id>.json`

### Response
```json
{
  "message": "File removed successfully",
  "id": "generated-uuid"
}
```

### Status Codes
- `200 OK` — file removed successfully
- `500 Internal Server Error` — file could not be removed
```json
{
  "error": "Unable to remove the file"
}
```

---
## `GET /states`
Returns a list of all saved auth state files.
---
## `GET /logs/api`

Returns the contents of the API log file as plain text.

### Response
- `Content-Type: text/plain; charset=utf-8`

### Status Codes
- `200 OK` — log contents returned
- `500 Internal Server Error` — log file could not be read
```json
{
  "error": "Unable to read the log file"
}
```

### Notes
- Reads from `/logs/api.log`

---

## `DELETE /logs/api`

Deletes the API log file.

### Response
```json
{
  "message": "Log file deleted successfully"
}
```

### Status Codes
- `200 OK` — log deleted successfully
- `500 Internal Server Error` — log file could not be deleted
```json
{
  "error": "Unable to delete the log file"
}
```

### Notes
- Deletes `/logs/api.log`

---

## `GET /logs/worker`

Returns the contents of the worker log file as plain text.

### Response
- `Content-Type: text/plain; charset=utf-8`

### Status Codes
- `200 OK` — log contents returned
- `500 Internal Server Error` — log file could not be read
```json
{
  "error": "Unable to read the log file"
}
```

### Notes
- Reads from `/logs/worker.log`

---

## `DELETE /logs/worker`

Deletes the worker log file.

### Response
```json
{
  "message": "Log file deleted successfully"
}
```

### Status Codes
- `200 OK` — log deleted successfully
- `500 Internal Server Error` — log file could not be deleted
```json
{
  "error": "Unable to delete the log file"
}
```

### Notes
- Deletes `/logs/worker.log`

---

## Logging Behavior

The API writes request logs asynchronously to:
- `/logs/api.log` (This is accessible via the `/logs/api` endpoint)


Log entries use the format:
```text
YYYY-MM-DD HH:MM:SS | METHOD | PATH | STATUS | LATENCY
```

Example:
```text
2026-06-13 10:15:30 | GET | /status | 200 | 1.2ms
```

Except for errors, which are logged as descriptions currrently.

---
