![tiktok-tech-immersion-march-2023-edm](https://github.com/cskang0121/tiktok-tech-immersion-2023-instant-messaging-system/assets/79074359/a8d4232a-3a50-41f1-ae05-72123a9a0ebc)
# TikTok Tech Immersion Programme 2023 

## Individual Project Completed By 
* SMU Computer Science Year 3 – Kang Chin Shen (cskang.2020@scis.smu.edu.sg)

## Backend Project Problem Statement
* Instant Messaging System (Pull Mode)

### Overview
![1280X1280](https://github.com/cskang0121/tiktok-tech-immersion-2023-instant-messaging-system/assets/79074359/628c97cc-a328-4456-b91b-c436136f9dac)
1. Develop an IM (Instant Messaging) System by implementing a set of APIs using Golang.
2. Develop a backend system using RPC (Remote Procedure Call) and microservices.

### Requirements

#### Architecture
1. HTTP Server – IDL (Interface Definition Language) in Protobuf.
2. RPC Server – IDL in Thrift. 

#### Storage
1. Database (MySQL) is used to store message data. Receivers can access them at any time.
2. Database schema design is not restricted to any form.

#### Message Delivery
1. Deliver messages to the intended recipients by "Pull" mode in a timely and consistent manner. <br>
  a. No need to maintain the connection. <br>
  b. No need to push new messages to receivers in real-time.
2. Pull API must be implemented for receivers to fetch messages.

#### Performance & Scalability
1. Support >20 concurrency in load testing.

#### Bonus
1. Elastic deployment of backend services.
2. Clear presentation in README.md.
3. Pass stress testing. [In Progress ...]

## Repository High Level Architecture

Assignment repository: https://github.com/TikTokTechImmersion/assignment_demo_2023/

```
| tiktok-tech-immersion-2023-instant-messaging-system      # Root folder

    | .github
        |workflows
            test.yml    # Build the go project on push, pull request

    | http-server       # HTTP server code implemented using Go

    | rpc-server        # RPC server code implemented using Go
    
    | test              # Contains load-testing.jmx for JMeter load test

    README.md           # Code documentation
    
    idl_http.proto      # IDL for HTTP server
    
    idl_rpc.thrift      # IDL for communication between HTTP and RPC server
    
    docker-compose.yml  # Docker compose files for specifying dependencies between services

    Other folders & files 
```

## Database Design

1. The code automatically creates a database, ```tiktok``` upon successful docker compose.
2. The code automatically creates a table, ```messages``` under the ```tiktok``` database, and drops table of the same name, if any.
3. The table ```messages``` has the following definition:
```
CREATE TABLE messages (
  id INT PRIMARY KEY AUTO_INCREMENT, 
  chat VARCHAR(255), 
  sender VARCHAR(255), 
  send_time INT, 
  message TEXT
);
```
4. Check ```tiktok-tech-immersion-2023-instant-messaging-system/test/db-testing.sql``` for more information.

## Running The Code

### Step 1 : Clone The Application
1. Before running the project, make sure the following are installed on your local machine: <br>
  a. Install Golang (```go version go1.20.4 darwin/arm64``` is used for this project) <br>
  b. Install Docker (```Docker version 23.0.5, build bc4487a``` is used for this project) <br>
  c. Install JMeter (```Binaries – apache-jmeter-5.5.zip``` is used for this project, note that JMeter requires Java 8+) and place this application in your desktop <br>
  d. Install MySQL Workbench for visualization of database (optional) <br>
  e. Install Insomnia for API testing [click here to download](https://insomnia.rest/download) or, Postman (optional) 
2. Run the command ```git clone https://github.com/cskang0121/tiktok-tech-immersion-2023-instant-messaging-system.git``` in the terminal in the prefered destination on your local machine.

### Step 2 : Run The Application
1. Open a new terminal, ```cd``` to ```tiktok-tech-immersion-2023-instant-messaging-system``` folder.
2. Run the command ```docker-compose up``` to start the instant messaging application on your local machine. Ensure all services are running on the corresponding ports:

<img width="350" alt="Screenshot 2023-06-14 at 12 07 18 PM" src="https://github.com/cskang0121/tiktok-tech-immersion-2023-instant-messaging-system/assets/79074359/919830c4-a7f7-4b76-990d-feb5d1da0cee">

3. Open a new terminal, run ```curl localhost:8080/ping``` to test if the application is running successfully, you should receive the response: ```{"message":"pong"}```.

### Step 3 : Satisfy The Project Requirements
1. Test the ```SEND``` function using ```Insomnia```:
```
# Request: Fire request to localhost:8080/api/send using HTTP POST with the following body:
{
	"chat": "a1:a2",
	"text": "hello",
	"sender": "a1"
}
```
```
# Response: 200 OK
```
2. Test the ```PULL``` function using ```Insomnia```:
```
# Request: Fire request to localhost:8080/api/pull using HTTP GET with the following body:
{
	"chat": "a1:a2",
	"cursor": 0,
	"limit": 1,
	"reverse": false
}
```
```
# Response: 
{
	"messages": [
		{
			"chat": "a1:a2",
			"text": "hello",
			"sender": "a1",
			"send_time": 1686672457
		}
	],
	"next_cursor": 1
}
```
3. Satisfy the load test (20 concurrency). <br>
  a. Open a new terminal, ```cd``` to ```Desktop/apache-jmeter-5.5/bin``` folder. <br>
  b. Run the command ```sh jmeter.sh``` to start JMeter application. <br>
  b. Open the file ```tiktok-tech-immersion-2023-instant-messaging-system/test/load-testing.jmx``` in JMeter application. <br>
  d. Run the load test and view the result:

![Screenshot 2023-06-14 at 1 01 55 AM](https://github.com/cskang0121/tiktok-tech-immersion-2023-instant-messaging-system/assets/79074359/b0d1fdce-6935-4484-9d7e-7f8168472c46)

- Based on the result, with 20 concurrent (threads) users, the application achieved 0% error and 88ms max response time.



  
