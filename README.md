# TravelAgencyGOLANG
Agência de viagens utilizando Golang and React

![screenshot](assets/banner.png)

## Stack:
- GIN web framework
- Prisma ORM for GO
- Postgres

## Steps to run the project
- docker compose up
    > backend:   http://localhost:8080/api/itineraries

    > frontend(pending) :  http://localhost:3000

## Steps to run without docker
- Set DATABASE_URL in the .env file
- Run the script to generate prisma client and create the database:
    > go run github.com/steebchen/prisma-client-go db push
- Run the app
    > go run main.go
- Follow the link:
    > http://localhost:8080/api/itineraries

## Deploy the project to a Kubernetes cluster
- Run
    > kubectl apply -f deploy/travelagency.yaml
