# Banking API

- Pre-requisites
  
  Go Version: 1.19 <br>
  Docker <br>
  Any testable tool with TCP protocol like Postman or Insomnia.

- Project Resume

  Banking API project for TEST proposes.
  The API will be bound at 3001 of your localhost.  

- Setting up the project

  Step 1: $ docker-compose up -d

  Step 2: $ docker network inspect dock-goapi_banking-network
    
  Step 3: Copy the Getway IP, that IP will be use to log in into our mysql database. Ex: "Gateway": "192.168.80.1"

  Step 4: Open the project folder on terminal and execute the follow command: $ cat infra/database/repositories/mysqlRepositories/.sql | mysql -h 192.168.80.1 -u root -P 3307 -p

  Step 5: $ docker-compose down

  <b>The Project is Ready ! Open your Postman</b>

- Run the project

  $ docker-compose up -d

MIT LICENSE 172.27.0.1