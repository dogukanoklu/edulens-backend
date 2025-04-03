# Class Attendance Management Endpoints

## Overview

This document details the endpoints for managing students within an class, including adding, updating, and rostering attendance. All endpoints require authentication.

---

## 1. Add Attendance

### **Description**

This endpoint add attendance.
Only students who are present are sent.

### **Endpoint Details**

- **Method:** `POST`
- **Path:** `/v1/attendance/{classID}`
- **Headers:**
  - `Authorization: Bearer <token>` (Required)
 
### **Parameters**:
- **classID**: (Required) The unique identifier of the class in UUID format. (a1b2c3d4-e5f6-7890-1234-567890abcdef).
 
### **Request Payload**

```json
[
  {
    "userID": "fado-cdef-4012-8345-6789abcd0123",
    "isPresent": true,
  },
  {
    "userID": "yk9a7b8c-d0e1-4f23-8a9b-c5d6e7f80123",
    "isPresent": false,
  }  
]
```

### **Response**

- **Success (204 No Content)**

  HTTP/1.1 204 No Content

- **Error Responses:**
  - `401 Unauthorized`: Token claims missing or invalid.
  - `400 Bad Request`: Invalid request payload or parameters.
  - `500 Internal Server Error`: Failed to add attendance.
 
### **cURL Example**

```bash
curl -X POST http://localhost:8000/v1/attendance/11/AT \
-H "Authorization: Bearer <token>" \
-H "Content-Type: application/json" \
-d '[
  {
    "userID": "fado-cdef-4012-8345-6789abcd0123",
    "isPresent": true,
  },
  {
    "userID": "yk9a7b8c-d0e1-4f23-8a9b-c5d6e7f80123",
    "isPresent": false,
  }  
]
```

## 2. Get Attendance

### **Description**

This endpoint get attendance.

### **Endpoint Details**

- **Method:** `GET`
- **Path:** `/attendance/v1/{classID}?date=1202335200`
- **Headers:**
  - `Authorization: Bearer <token>` (Required)

### **Parameters**:
- **classID**: (Required) The unique identifier of the class in UUID format. (a1b2c3d4-e5f6-7890-1234-567890abcdef).
- **date**: (Optional) A timestamp representing the date. If provided, fetches attendance data for the given date. If the date is not provided, the data for the current date will be returned.

### **Response**

- **Success (200 OK):**

```json
{
  "id": "fado-1234-4567-89ab-cdef01234567",
  "level": 11,
  "branch": "AT",
  "students": [
  {
    "id": "fado-cdef-4012-8345-6789abcd0123",
    "studentImage": "http://localhost:8000/image/1",
    "schollNumber": 100,
    "firstName": "Doğukan",
    "lastName": "OKLU",
    "isPresent": true
  },
  {
    "id": "yk9a7b8c-d0e1-4f23-8a9b-c5d6e7f80123",
    "studentImage": "http://localhost:8000/image/2",
    "schollNumber": 101,
    "firstName": "Yiğit",
    "lastName": "Gülmez",
    "isPresent": false
  }
],
  "createdAt": 1202335200
}
```

- **Error Responses:**
  - `401 Unauthorized`: Token claims missing or invalid.
  - `400 Bad Request`: Invalid parameters.
  - `404 Not Found`: The requested attendance data could not be found for the specified parameters.
  - `500 Internal Server Error`: Failed to retrieve attendance information.

### **cURL Example**

```bash
curl -X GET http://localhost:8000/v1/attendance/11/AT?date=1202335200 \
-H "Authorization: Bearer <token>" \
```

## 3. Update Attendance

### **Description**

This endpoint update attendance. 
Sends student IDs and their present statuses for updating.

### **Endpoint Details**

- **Method:** `PUT`
- **Path:** `/v1/attendance/{attendanceID}`
- **Headers:**
  - `Authorization: Bearer <token>` (Required)

### **Parameters**:
- **attendanceID**: (Required) The attendance id (e.g., 1, 2).

### **Request Payload**

```json
[
  {
    "userID": "fado-cdef-4012-8345-6789abcd0123",
    "isPresent": true,
  },
  {
    "userID": "yk9a7b8c-d0e1-4f23-8a9b-c5d6e7f80123",
    "isPresent": false,
  }  
]
```

### **Response**

- **Success (204 No Content)**

  HTTP/1.1 204 No Content

- **Error Responses:**
  - `401 Unauthorized`: Token claims missing or invalid.
  - `400 Bad Request`: Invalid request payload.
  - `404 Not Found`: Attendance record not found.
  - `500 Internal Server Error`: Failed to update attendance.
 
### **cURL Example**

```bash
curl -X PUT http://localhost:8000/v1/attendance/{attendanceID} \
-H "Authorization: Bearer <token>" \
-H "Content-Type: application/json" \
-d '{
  "presentStudents": [123, 124]
}'
```