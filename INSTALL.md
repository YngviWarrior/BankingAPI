# Banking API

- Pre-requisites
  
  Go Version: 1.19 <br>
  Docker <br>
  Any testable tool with TCP protocol like Postman or Insomnia.

- Project Resume

  Banking API project for TEST proposes.
  The Holder MicroService will be bind at 3001 port.  
  The Account MicroService will be bind at 3002 port.

- Setting up the project

  Step 1: $ docker-compose up -d

  Step 2: $ docker network inspect bankingapi_banking-network
    
  Step 3: Copy the Getway IP, that IP will be use to log in into our mysql database. Ex: "Gateway": "192.168.32.1"

  Step 4: Open the project folder on terminal and execute the follow command: $ cat .sql | mysql -h 192.168.32.1 -u root -P 3307 -p

  Step 5: $ docker-compose down

- Run the project

  $ docker-compose up -d

<b>The Project is Ready ! Open your Postman</b>


