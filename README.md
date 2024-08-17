## VaultKeeper: A Secure Secrets Management Server

**Introduction**

VaultKeeper is a simple yet secure secrets management server written in Go. It provides a way to store, retrieve, and manage sensitive information in a local SQLite database.

**Features**

* **Secure Storage:** Stores secrets using AES encryption with a randomly generated key.
* **Role-Based Access Control:** Implements basic RBAC to control access to secrets.
* **Secret Retrieval:** Allows authorized users to retrieve secrets.
* **Audit Logging:** Tracks all actions performed on secrets.
* **API Key Management:** Generates and revokes API keys for authentication.
* **Project Management:** Groups secrets under projects for better organization.
* **Output Formats:** Supports various output formats (dotenv, TOML, YAML, JSON) for easy integration.

**Getting Started**

1. **Clone the Repository:** `git clone https://github.com/thelamedev/vault-keeper`
2. **Install Dependencies:** `go mod tidy`
3. **Run the Server:** `go run main.go`
4. **Access the API:** Use a tool like Postman or curl to interact with the API endpoints.

**API Endpoints**

* **POST /projects:** Create a new project.
* **GET /projects/{id}:** Retrieve a specific project.
* **DELETE /projects/{id}:** Delete a project.
* **POST /projects/{id}/secrets:** Add a secret to a project.
* **GET /projects/{id}/secrets:** Retrieve all secrets within a project.
* **POST /secrets:** Create a new secret.
* **GET /secrets/{id}:** Retrieve a specific secret.
* **DELETE /secrets/{id}:** Delete a secret.
* **POST /api-keys:** Generate a new API key.
* **DELETE /api-keys/{key}:** Revoke an API key.

**Security**

* Secrets are stored using AES encryption with a randomly generated key.
* Basic RBAC is implemented to control access.
* Consider using a more robust database and encryption mechanism for production environments.

**Additional Features**

* **Rate Limiting:** Prevents abuse and protects against denial-of-service attacks.
* **Error Handling:** Provides informative error messages.
* **Documentation:** Comprehensive documentation guides users on how to use the server.
* **Testing:** Thorough unit and integration tests ensure the server's reliability.
* **Export/Import:** Exports/imports projects and their associated secrets in various formats.
* **Search and Filtering:** Allows users to search for secrets based on keywords and filter them by project or other criteria.
* **Accessibility:** Ensures the server is accessible from different platforms and devices.
* **Performance Optimization:** Optimizes the server's performance to handle large numbers of secrets and users.

**License**

MIT

