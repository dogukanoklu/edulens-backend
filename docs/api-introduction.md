# **Introduction to EduLens API**

### Overview

Welcome to the EduLens API reference documentation! This API provides a comprehensive interface for interacting with EduLens's core backend services, enabling developers to manage organizations, users, and related data efficiently.

The API is written in **Go**, ensuring high performance and scalability. For additional insights about our backend, visit the [source code repository](https://github.com/dogukanoklu/edulens-backend/tree/main/apps/api).

Explore this documentation to understand the endpoints, their usage, and how to integrate them into your application.

---

### Base URL

All API requests should be made to the following base URL unless otherwise specified:

```bash
http://localhost:8000
```

---

### **Purpose of the API**

This API enables seamless integration with your applications to manage users and organizations efficiently. It supports the following key functionalities:

- **Attendance Management:**  
  Create, retrieve, and update attendance within the system.

- **Access Control:**  
  Role-based access control (RBAC) ensures that only authorized users can perform sensitive operations.

---

### **Authentication and Authorization**

- **Token-Based Authentication:**
  All endpoints require a valid JWT token passed in the Authorization header as a Bearer token. Tokens include claims that the API uses to verify the identity and permissions of the user.


---

### **API Endpoint Overview**

The API is organized into logical groups to facilitate user and organization management:

Here’s the table with a bulleted list in the right column:

| **Endpoints**  | **Operations**                |
| -------------- | ------------------------------|
| **attendance** | • Create<br>• Get<br>• Update |
| **class**      | • Create                      |

---

### **How to Use This Documentation**

Each endpoint is described with the following sections:

- **Description:** Explains the endpoint's purpose and typical use cases.
- **Endpoint Details:** Specifies the HTTP method, path, and required headers.
- **Path Detils:** Defines the parameters included in the URL path, such as level, branch and date
- **Request Payload:** Describes the format and structure of the request body.
- **Response Format:** Provides examples of successful and error responses.
- **cURL Example:** Demonstrates how to call the endpoint using `cURL`.

This documentation is intended to be a complete guide for developers who need to integrate this API into their applications. By following the examples and reference materials provided, developers can ensure seamless implementation and robust integration.