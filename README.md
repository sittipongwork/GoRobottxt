# The robots.txt for golang make your web have robots.txt


### Installing
Use Golang CLI
```go
go get github.com/sittipongwork/GoRobottxt
```
or

Github Repository

[https://github.com/sittipongwork/GoRobottxt](https://github.com/sittipongwork/GoRobottxt)

### Import Package
example : Import to go
```go
  import "github.com/sittipongwork/GoRobottxt"
```

### Create Robot Data Container
example
```go
  robot := gorobottxt.CreateRobot()
```



### Add UserAgent to robot container
Agent Allowed for keep data your website in robots.txt | All name of agent: [UserAgent List](http://www.user-agents.org/)

example
```go
  //Add UserAgent to robot (Robot Container)
  //"*" that mean all of robot
  robot.AddUserAgent("*")
```

example
```go
//Create Robot Container.
robot := gorobottxt.CreateRobot()

//Add UserAgent to robot (Robot Container).
//"*" That mean all of robot.
robot.AddUserAgent("*")

//Add Allowed URL for robot can access.
//That will be ...
//Allow: /blog/*
robot.AddAllowUrl("/blog/*")
//You can add more than one. Example
robot.AddAllowUrl("/url2")
robot.AddAllowUrl("/url3/*")

//Add Disallow URL for robot can't access.
//That will be ...
//Disallow: /admin/*
robot.AddDisallowUrl("/admin/*")
//You can add more than one. Example
//robot.AddDisallowUrl("/disallowurl")

//Add Sitemap Url Website.
//That will be ...
//Sitemap: https://website.com/blog-sitemap.xml
robot.AddSitemapUrl("https://website.com/blog-sitemap.xml")

//Add Sitemap Url Website.
//That will be ...
//Sitemap: https://website.com/profiles-sitemap.xml
robot.AddSitemapUrl("https://website.com/profiles-sitemap.xml")

//Generate Robots Container to String/text.
robottxt := robot.RobotGenerateString()
```

output
```text
User-agent: *
Allow: /blog/*
Allow: /url2
Allow: /url3

Disallow: /admin/*

Sitemap: https://website.com/blog-sitemap.xml
Sitemap: https://website.com/profiles-sitemap.xml
```
