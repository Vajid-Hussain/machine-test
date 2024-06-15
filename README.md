# machine-test

# Technologies Used

- Go
- Fiber Framework
- Docker
- Swagger
- GORM
- PostgreSQL (with a trigger to increase the apply quantity of a job during a single job seeker's application)

## Installation and Running Instructions

1. **Install Docker**Follow the instructions provided in the official Docker documentation to install Docker on your machine: [Docker Engine Install](https://docs.docker.com/engine/install/)
2. **Clone the Repository**Clone the GitHub repository using the following command: `git clone https://github.com/Vajid-Hussain/machine-test`
3. **Run Docker Compose**Navigate to the cloned repository's directory and start the services using Docker Compose.
 Navigate to the project directory: `cd machine-test`
 Run the Docker containers in detached mode: `docker compose up -d`
4. **Access Swagger**
   Once the services are up and running, you can access the Swagger documentation at: `http://localhost:3000/swagger/index.html`

## Default Admin Credentials

To log in as the admin, use the following credentials:

- Email: synlabs@gmail.com
- Password: synlabs

## Repository

The code for this project can be found in the following GitHub repository: [machine-test](https://github.com/Vajid-Hussain/machine-test)
