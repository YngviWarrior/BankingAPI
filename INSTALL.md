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

  Step 2: Open the project folder on terminal and execute the follow command: $ cat .sql | mysql -h 127.0.0.1 -u root -P 3307

  Obs: the loopback IP (127.0.0.1) it's necessary !

<b>The Project is Running ! Open your Postman.</b>

- If is needed Reboot the project

  $ docker-compose down && docker-compose up -d



