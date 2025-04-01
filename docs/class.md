# Class Management Endpoints

## Overview

This document details the API endpoints for retrieving class information. All endpoints require authentication.

---

## 1. Get Classes

### **Description**

This endpoint retrieves a list of classes.

### **Endpoint Details**

- **Method:** `GET`
- **Path:** `/classes`
- **Headers:**
  - `Authorization: Bearer <token>` (Required)

### **Response**

- **Success (200 OK):**

```json
[
  {
    "id": "fado-8400-4c29-a1d3-0604f822994a",
    "level": 11,
    "branch": "AT"
  },
  {
    "id": "fado-b1a2-4c3d-9e4f-567890abcdef",
    "level": 12,
    "branch": "AT"
  }
]
```

- **Error Responses:**
  - `401 Unauthorized`: Token claims missing or invalid.
  - `500 Internal Server Error`: Failed to retrieve class information.

### **cURL Example**

```bash
curl -X GET http://localhost:8000/v1/classes \
-H "Authorization: Bearer <token>" \
```