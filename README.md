# Golang for Backend Developer

> For more detail, please visit [wiki](https://github.com/tpphu/golang-training/wiki)

## About the instructor

**[Phu, Tran Phong](https://www.linkedin.com/in/tpphu/)**
Developer and Instructor

I'm Phu, I'm a developer with a passion for teaching. I'm an academic advisor and instructor at Nordic Coder. My expertise and skills are about big data, distributed system, continuous deployment. I have been teaching many courses about backend and frontend such as Golang/Nodejs/PHP, ES/MySQL and Reactjs/Angular. I've helped students learn to code, improve their skills. I've been invited/collaborated by companies such as ABC to build the courses for training their employees.

I spend most of my time to researching latest methodologies, technologies, skills which should be have for developers to help them learn to make code better, code for fun and love, and make hard concepts easy to understand.

Welcome to my course, I am on way to help you.
## About this course

This course is designed to get you up and running as fast as possible with Go.  We'll quickly cover the basics, then dive into some of the more advanced features of the language. 

After finishing this course, students will be able to learn how to:

- Build high-performance concurrent programs with Golang
- Build modern programs with Golang and its ecosystem such as CI/CD, GRPC
- Build their own programs with hands-on mentoring from instructor


### Course syllabus

#### Topic 1

- An Introduction to Programming with Go and TDD
- Overview on the fastest growing language and why we should learn Golang
- Set up environments & code organization
- Go characteristics and principles
- How Go runs in single CPU and multiple CPUs
- Reflection and Go's reflect package
- Exercise: converting CSV to YAML using pointer
- Exercise: cloning Lodash/Underscore to deep dive into Go
- Exercise: practicing TDD methodology


#### Topic 2

- Pointer in Go
- Discover goroutines, channel, defer, panic, recover
- Error handling in Go
- Pipeline programming in Go
- Working with database/mysql
- Diving into Golang by designing a scripting/job 
- Strategy to watch scripting/job
- Guide to setting up Docker/Laradock to install MySQL
- Exercise: building a crawler
- Exercise: following TDD and HTTP mock test

#### Topic 3

- Diving into Golang by designing a RESTful API 
- CRUD service using Go with Gin/Echo framework
- Login and JWT Authentication
- HTTP Protocol, Cookie, Session
- Load Balancing with Nginx, high performance, and scalability
- Exercise: building a note service for Todo MVC application
- Exercise: building an auto increment number and issues with multiple-CPU problems.
- Exercise: writing Unit Testing Code and DB mock test

#### Topic 4

- Design and apply microservices/serverless architecture using Go
- RPC, GRPC and benefits from GRPC
- How to design a new system with GRPC 
- How to integrate GRPC to an existing system
- Exercise: using GRPC to refactor note service from topic 2
- Exercise: API caching, DB connection pool
- Exercise: writing a Voucher service

#### Topic 5

- Understand about string operations and UTF8 encoding in Go
- Best practices to develop a project with Go & final project
- CI/CD, logging and monitoring on Production environment with ELK stack
- Deploy a CI/CD for note services with Jenkins/K8S
- Overview and requirements of the Final Project 
- Review and feedback on the course

#### Topic 6

- Build your own open source software
- Build a helper lib (ex: porting underscore)
- Or build a http router
- Or build a (micro) service to generate fake data
- Or build an application with your own idea

### Pre-work

A mini project to demonstrate students’ knowledge at input level

- With a given URL of news website (Eg vnexpress.net, thesaigontimes.vn,...), create a crawler that:
- Task 1: Can fetch the whole data from a news website then parse the data into specific piece as: title, published date, summary, body and author
- Task 2: With any URLs found in the previously fetched data, repeat the task 1 until reaching maximum of 5 levels.	

### Final project

The final project is up to students with mentoring from instructor. The objective of the project is to: 

- Demonstrate understanding of Golang and its ecosystem.
- Apply knowledge gained during the course by building an application/program such as a helper lib (ex: porting underscore), an http router, a (micro) service to generate fake data or create API for blog, e-commerce, Todo app, etc.  
