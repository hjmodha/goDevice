###Device Detection Library in Golang

####Installing Package:
```
go get github.com/hiteshmodha/goDevice
```

####How To use Package to detect Device:
```
  if deviceType=="Mobile" {
    fmt.Fprintf(w,"<h1>Mobile</h1>")
  }else if deviceType=="Web"{
    fmt.Fprintf(w,"<h1>Web</h1>")
  }else if deviceType=="Tab"{
    fmt.Fprintf(w,"<h1>Tablet</h1>")
  }
  ```

For now goDevice is detecting Mobile,Desktop and Tablet.In future we will add support for SmartTV, Watch etc
