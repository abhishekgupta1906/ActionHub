# Live Demo
[Visit this website](https://actionhub-production.up.railway.app/)


# ActionHub
ActionHub is an industrial-grade task management app, designed to streamline operations and enhance productivity across teams. With a robust backend built in Go and a responsive frontend in React, ActionHub provides a seamless experience for managing tasks, monitoring workflows, and improving collaboration.

# Features
- Task Management: Create, edit, and delete tasks with specific deadlines and priority levels.
- User Authentication: Secure login and registration powered by JWT for authenticated access.
- Real-Time Updates: Real-time task updates to keep all team members in sync.
- Task Assignment: Assign tasks to team members and track task completion.
- Analytics Dashboard: Visual insights into task completion rates, user activity, and more.


# Tech Stack
- Backend: Go (Golang) - REST API
- Frontend: React with Hooks, Context API
- Database:MONGODB (or other supported databases)
- Authentication: JWT (JSON Web Tokens)
- Hosting: Backend on Azure VM, Frontend on Vercel or Railway
# Getting Started
Prerequisites
Go 1.20+
Node.js 18+
PostgreSQL (or any supported database)
Docker (for optional containerized setup)

# Set Up your .env file
```shell
MONGODB_URL=<your mongo db url>
PORT=5000
ENV=development
```

# Command for Running
```shell
# Install dependencies for backend
go mod tidy

# Run backend server
go run main.go

# Install dependencies for frontend
cd client 

npm install

# Run frontend server
npm start
```

# Using Docker (Recommended)
```shell
docker compose up
```

