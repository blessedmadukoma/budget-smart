# Shared Packages in the Backend Monorepo

The `common/` directory contains shared code that is used across multiple services in the monorepo. These packages are designed to promote code reuse, maintain consistency, and provide common utilities that every service might need. Here's what each package contains:

*Note: Think of `pkg/` as a collection of internal libraries, similar to how you might use third-party packages, but developed and maintained within your organization*

## 1. `common/auth/` - Authentication Utilities
- JWT token generation, validation, and parsing
- Google OAuth2 client implementation
- Authentication middleware for HTTP requests
- Authorization checks and RBAC (Role-Based Access Control)
- User identity context management
- Secure cookie handling
- Session management utilities

## 2. `common/models/` - Shared Data Models
- Common domain models used across services (User, Transaction, Budget, etc.)
- Data Transfer Objects (DTOs) for consistent API request/response formats
- Validation logic for data models
- Type conversion utilities
- JSON serialization/deserialization helpers
- Database entity definitions shared between services

## 3. `common/config/` - Configuration Utilities
- Environment variable loading and parsing
- Configuration file management
- Dynamic configuration with hot reloading
- Service discovery interfaces
- Feature flag management
- Environment-specific configuration (dev/staging/prod)
- Configuration validation and default values

## 4. `common/logger/` - Logging Utilities
- Structured logging implementation with consistent fields
- Log level management
- Context-aware logging
- Request ID propagation
- Correlation ID utilities for request tracing
- Log formatting and output configuration
- Integration with external logging platforms

## 5. `common/queue/` - Message Queue Clients
- Abstractions for RabbitMQ operations
- Message publishing utilities
- Consumer implementations with retry logic
- Dead letter queue handling
- Circuit breaker pattern implementations
- Message serialization/deserialization
- Topic/queue management utilities

## Benefits of Shared Packages
This approach provides several advantages:

1. **Consistency**: All services use the same authentication, logging, and messaging patterns.

2. **Maintainability**: Changes to core functionality can be made in one place.

3. **Development Speed**: Teams don't need to reinvent common utilities for each service.

4. **Reduced Bugs**: Well-tested shared code means fewer bugs across services.

5. **Standardization**: Enforces architectural standards across the codebase.

The shared packages act as the foundation for all services, providing proven, tested implementations of common functionality that each service can build upon while focusing on their specific business logic.