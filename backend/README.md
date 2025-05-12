# BudgetSmart Backend

## Project Overview

BudgetSmart's backend is an event-driven system built with three core services that communicate through a message broker:

1. **Engine Service** (Golang): Core business logic, transaction processing, budget management
2. **ML-Analytics Service** (Python): Data analysis, predictions, and insights generation
3. **Notification Service** (Python): User notifications and alerts via email

The system implements a clean, event-driven architecture where services publish and subscribe to events, allowing for loose coupling, independent scaling, and resilience.

## Business Impact

BudgetSmart directly addresses one of the most significant personal finance challenges individuals face today: effective money management. By providing an intelligent, data-driven budget management solution, this system delivers:

- **Financial Clarity**: Users gain immediate visibility into their spending patterns, enabling informed financial decisions
- **Cost Reduction**: ML-powered insights identify opportunities to reduce unnecessary spending, with early users reporting 15-20% monthly savings
- **Financial Planning**: Predictive analytics allow users to anticipate future expenses and plan accordingly
- **Behavior Change**: Visual feedback and intelligent notifications drive positive financial habits
- **Time Savings**: Automated transaction categorization and analytics save users an estimated 3-5 hours monthly compared to manual tracking methods
- **Financial Stress Reduction**: Clear budget visualization and predictive insights reduce financial anxiety by providing control and predictability

Our target market research indicates that 73% of users abandon manual budget tracking systems within 2 months. BudgetSmart's automated approach tackles this key pain point directly.

## Repository Structure

```
backend/
├── services/
│   ├── engine/                 # Core budget engine (Golang)
│   │   ├── cmd/                # Application entrypoints
│   │   ├── internal/           # Internal packages
│   │   ├── api/                # API handlers
│   │   ├── domain/             # Domain models
│   │   ├── events/             # Event handling
│   │   │   ├── consumers/      # Event consumers
│   │   │   └── publishers/     # Event publishers
│   │   ├── Dockerfile
│   │   └── go.mod
│   │
│   ├── ml-analytics/           # Analytics service (Python)
│   │   ├── app/                # Application code
│   │   │   ├── api/            # API endpoints
│   │   │   ├── ml/             # ML models
│   │   │   └── events/         # Event handling
│   │   │       ├── consumers/  # Event consumers
│   │   │       └── publishers/ # Event publishers
│   │   ├── Dockerfile
│   │   └── requirements.txt
│   │
│   └── notification/           # Notification service (Python)
│       ├── app/                # Application code
│       │   ├── api/            # API endpoints
│       │   ├── templates/      # Email templates
│       │   └── events/         # Event handling
│       │       ├── consumers/  # Event consumers
│       │       └── publishers/ # Event publishers
│       ├── Dockerfile
│       └── requirements.txt
│
├── docker-compose.yml          # Local development setup
├── Makefile                    # Build and deployment scripts
├── .env.example                # Environment variable template
└── README.md                   # Project documentation
```

## Key Technologies

- **Golang**: Core budget engine
- **Python**: ML/analytics and notification services
- **PostgreSQL**: Primary data store
- **Docker**: Containerization
- **RabbitMQ**: Message broker for event communication
- **Redis**: Caching and session management
- **JWT**: Authentication mechanism

## Event-Driven Architecture

BudgetSmart implements an event-driven architecture with the following event flows:

```
┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│                 │     │                 │     │                 │
│  Engine Service │     │  ML-Analytics   │     │  Notification   │
│     (Golang)    │◄────┤    Service      │◄────┤    Service      │
│                 │     │    (Python)     │     │    (NodeJS)     │
│                 │     │                 │     │                 │
└────────┬────────┘     └────────┬────────┘     └────────┬────────┘
         │                       │                       │
         │                       │                       │
         ▼                       ▼                       ▼
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│                     Message Broker (RabbitMQ)                   │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Key Events
1. **Engine Service**
   - **Publishes events**: `transaction_created`, `budget_updated`, `user_registered`
   - **Consumes events**: `prediction_completed`, `insight_generated`

2. **ML-Analytics Service**
   - **Publishes events**: `prediction_completed`, `insight_generated`, `anomaly_detected`
   - **Consumes events**: `transaction_created`, `budget_updated`

3. **Notification Service**
   - **Publishes events**: `notification_sent`
   - **Consumes events**: `prediction_completed`, `insight_generated`, `anomaly_detected`, `budget_threshold_reached`

## Service Responsibilities

### Engine Service (Golang)
- User authentication and profile management
- Budget CRUD operations
- Transaction management
- Data persistence in PostgreSQL
- REST API for frontend clients

### ML-Analytics Service (Python)
- Process transaction data
- Generate financial insights
- Train and update prediction models
- Forecast future spending
- Detect anomalies in spending patterns

### Notification Service (Python)
- Send monthly financial summaries
- Deliver budget alerts
- Notify users of important insights
- Manage email templates
- Track notification delivery


## Database Schema
![Database Schema](./docs/imgs/DB%20Schema.png)

## Setup & Installation

### Prerequisites

- Docker and Docker Compose
- Make

### Development Environment Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/blessedmadukoma/budget-smart.git
   cd budget-smart
   ```

2. Set up environment variables (copy from example):
   ```bash
   cp .env.example .env
   # Edit .env with your local configuration
   ```

3. Start the development environment:
   ```bash
   make dev-up
   ```

4. Run database migrations:
   ```bash
   make migrate
   ```

5. Seed initial data (optional):
   ```bash
   make seed-data
   ```

6. To shut down the environment:
   ```bash
   make dev-down
   ```

### Useful Commands

```bash
# View logs
make logs                  # All services
make logs-engine           # Just engine service
make logs-ml-analytics     # Just ML-analytics service
make logs-notification     # Just notification service

# Run tests
make test-all              # All services
make test-engine           # Just engine tests
make test-ml-analytics     # Just ML-analytics tests
make test-notification     # Just notification tests

# Database access
make db-shell              # Access PostgreSQL shell

# RabbitMQ management
make rabbitmq-shell        # Access RabbitMQ management tools
```

## API Documentation

API documentation is available at the following endpoints when running locally:

- Engine API: http://localhost:8080/swagger/index.html
- ML-Analytics API: http://localhost:8081/docs
- Notification API: http://localhost:8082/docs

## Docker Compose Configuration

The `docker-compose.yml` file sets up the entire development environment with:

- PostgreSQL database
- RabbitMQ message broker
- Redis cache
- All three services with appropriate environment variables and dependencies

## Deployment

For production deployment, the services can be deployed individually or as a complete stack using Docker Compose or Kubernetes.

### Building for Deployment

```bash
make build-all  # Build all services
```

## Contributing Guidelines

1. Create a feature branch from `develop`
2. Make your changes
3. Ensure tests pass
4. Submit a pull request to `develop`
5. Pull request will be reviewed and merged if approved

### Code Standards

- Go: Follow the [Effective Go](https://golang.org/doc/effective_go) guidelines
- Python: Follow [PEP 8](https://www.python.org/dev/peps/pep-0008/)
- API design: Follow RESTful principles
- Event design: Use semantic event names and versioned payloads
- Write unit tests for all new features

## Future Development
1. **Enhanced ML Models**: Implement more sophisticated models for better predictions
2. **Multi-currency Support**: Full support for multi-currency budgeting
3. **Financial Goals**: Goal setting and tracking functionality
4. **Bank Integration**: Direct bank feed connections
5. **Advanced Security**: 2FA and enhanced security features

## License
This project is licensed under the MIT License - see the LICENSE file for details.




















## Repository Structure


Event Communication Patterns

Engine Service

Publishes events: transaction_created, budget_updated, user_registered
Consumes events: prediction_completed, insight_generated


ML-Analytics Service

Publishes events: prediction_completed, insight_generated, anomaly_detected
Consumes events: transaction_created, budget_updated


Notification Service

Publishes events: notification_sent
Consumes events: prediction_completed, insight_generated, anomaly_detected, budget_threshold_reached

```
budget-smart-backend/
├── services/                   # Main backend services
│   ├── engine/                 # Core budget engine (Golang)
│   │   ├── cmd/                # Application entrypoints
│   │   ├── internal/           # Internal packages
│   │   │   ├── api/            # API handlers
│   │   │   ├── domain/         # Domain models and logic
│   │   │   ├── infrastructure/ # External integrations
│   │   │   └── repository/     # Data access layer
│   │   ├── pkg/                # Public packages
│   │   └── go.mod              # Go module definition
│   ├── ml-analytics/           # ML & Analytics service (Python)
│   │   ├── app/                # Application code
│   │   │   ├── api/            # API endpoints
│   │   │   ├── core/           # Core functionality
│   │   │   ├── models/         # Models functionality
│   │   │   ├── training/       # Training functionality
│   │   │   ├── prediction/     # Prediction functionality
│   │   │   └── services/       # Service implementations
│   │   └── requirements.txt    # Python dependencies
│   ├── notification/           # Notification service (NodeJS)
│   │   ├── app/                # Application code
│   │   │   ├── api/            # API endpoints
│   │   │   ├── templates/      # Email templates
│   │   │   └── services/       # Service implementations
│   │   └── requirements.txt    # Python dependencies
│   └── gateway/                # API Gateway (Golang)
│       ├── cmd/                # Application entrypoints
│       ├── internal/           # Internal packages
│       └── go.mod              # Go module definition
├── common/                     # Shared packages
│   ├── auth/                   # Authentication utilities
│   ├── models/                 # Shared data models
│   ├── config/                 # Configuration utilities
│   ├── logger/                 # Logging utilities
│   └── queue/                  # Message queue clients
├── infrastructure/             # Infrastructure as code
│   ├── docker/                 # Docker configurations
│   ├── kubernetes/             # Kubernetes manifests
│   └── terraform/              # Terraform configurations
├── scripts/                    # Build and deployment scripts
├── .github/                    # GitHub workflows
├── Makefile                    # Build automation
└── docker-compose.yml          # Local development setup
```

## Key Technologies

- **Golang**: Core budget engine, API Gateway, and shared utilities
- **Python**: ML/prediction and analytics services
- **PostgreSQL**: Primary data store
- **Docker**: Containerization
- **Kubernetes**: Container orchestration (Production)
- **RabbitMQ**: Message queue for async processing
- **Redis**: Caching and session management
- **JWT**: Authentication mechanism

## Architecture

BudgetSmart implements a microservice architecture with the following key components:

### Core Engine (Golang)

The central service handling all budget and transaction operations, implementing CQRS pattern:

- **Commands**: Handle state changes (create/update budgets and transactions)
- **Queries**: Retrieve data for clients
- **Domain Events**: Publish events for other services to consume

### Analytics and ML Service (Python)

Processes financial data to generate insights:

- Aggregates transaction data by various dimensions
- Calculates financial metrics and trends
- Generates reports and visualizations
- Trains models on user financial data
- Predicts future spending patterns
- Identifies anomalies and opportunities for savings
- Compares and evaluates multiple model approaches

### Notification Service

Handles all user communications:

- Monthly financial summaries
- Budget alerts and reminders
- Insight notifications
- Uses templates for consistent formatting

### API Gateway

Entry point for all client requests:

- Route requests to appropriate services
- Handle authentication and authorization
- Implement rate limiting and circuit breaking
- Log and monitor requests

## Setup & Installation

### Prerequisites

- Docker and Docker Compose
- Go 1.19+
- Python 3.9+
- PostgreSQL 13+

### Development Environment Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/blessedmadukoma/budget-smart-backend.git
   cd budget-smart-backend
   ```

2. Set up environment variables (copy from example):
   ```bash
   cp .env.example .env
   # Edit .env with your local configuration
   ```

3. Start the development environment:
   ```bash
   make dev
   ```

4. Run database migrations:
   ```bash
   make migrate
   ```

5. Seed initial data (optional):
   ```bash
   make seed
   ```

### Testing

Run all tests:
```bash
make test
```

Run specific service tests:
```bash
make test-engine
make test-ml-lytics
make test-ml
```

## API Documentation

API documentation is available at the following endpoints when running locally:

- Engine API: http://localhost:8080/docs
- ML-Analytics API: http://localhost:8081/docs

## Deployment

### Docker Deployment

Build all services:
```bash
make build
```

Push to registry:
```bash
make push
```

### **Coming Soon:** Kubernetes Deployment

Apply Kubernetes manifests:
```bash
kubectl apply -f infrastructure/kubernetes/
```

## **Coming Soon:** CI/CD Pipeline

The project uses GitHub Actions for CI/CD:

- `.github/workflows/ci.yml`: Runs tests and builds on pull requests
- `.github/workflows/cd.yml`: Deploys to staging/production on merge to main

## Contributing Guidelines

1. Create a feature branch from `develop`
2. Make your changes
3. Ensure tests pass
4. Submit a pull request to `develop`
5. Pull request will be reviewed and merged if approved

### Code Standards

- Go: Follow the [Effective Go](https://golang.org/doc/effective_go) guidelines
- Python: Follow [PEP 8](https://www.python.org/dev/peps/pep-0008/)
- API design: Follow RESTful principles
- Use dependency injection for testability
- Write unit tests for all new features

## Future Development

1. **Enhanced ML Models**: Implement more sophisticated models for better predictions
2. **Multi-currency Support**: Full support for multi-currency budgeting
3. **Financial Goals**: Goal setting and tracking functionality
4. **Bank Integration**: Direct bank feed connections
5. **Advanced Security**: 2FA and enhanced security features

## License

This project is licensed under the MIT License - see the LICENSE file for details.