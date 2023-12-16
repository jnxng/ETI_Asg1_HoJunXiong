# ETI_Asg1_HoJunXiong
Carpooling Microservices
1. Design Considerations of Microservices

1.1.  User Microservice
Manages user accounts for passengers and car owners.
Supports creating, updating, and deleting user accounts.
Allows users to switch between passenger and car owner profiles.
Handles user authentication and authorization.

1.2. Trip Microservice
Manages car-pooling trips published by car owners.
Enables users to search and enroll in available trips.
Supports starting and canceling trips by car owners.
Manages the number of available seats for each trip.

2. Architecture Diagram
The carpooling application was built using 2 microservices, with the 2 main microservices being the User microservice and the Trip microservice.
Each microservice has a MYSQL database as its persistent storage, keeping the data for each microservice respectively.
Each Microservice has functions built for them to work listed above.

3. Instructions to setting up and running microservices
move directory into each file, user and trip, then run the main.go for each file.