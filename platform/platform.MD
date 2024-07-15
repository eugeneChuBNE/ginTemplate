# Platform Directory

The `platform` directory contains code and configurations for external dependencies or services your application interacts with, such as databases, third-party services, and infrastructure as code.

## Common Uses

1. **Database Configuration:**
   - Initialisation scripts and connection management.

2. **Third-Party Services:**
   - Integration code and configuration files.

3. **Infrastructure as Code (IaC):**
   - Deployment scripts using tools like Terraform and Kubernetes.

## Example Files

### Database Initialisation

**platform/database/db.go**

```go
package database

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

func InitDB() *sql.DB {
    db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:0000)/dbname")
    if err != nil {
        log.Fatal(err)
    }

    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }

    return db
}
```

### Third-Party Service Integration

**platform/third_party/email_service.go**

```go
package third_party

import (
    "fmt"
    "net/smtp"
)

func SendEmail(to, subject, body string) error {
    from := "email@example.com"
    password := "email-password"
    smtpHost := "smtp.example.com"
    smtpPort := "587"

    msg := "From: " + from + "\n" +
        "To: " + to + "\n" +
        "Subject: " + subject + "\n\n" +
        body

    auth := smtp.PlainAuth("", from, password, smtpHost)

    err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
    if err != nil {
        return err
    }

    fmt.Println("Email sent successfully")
    return nil
}
```

### Infrastructure as Code (IaC)

**Infrastructure as Code (IaC):**
   - Deployment scripts using tools like Terraform and Kubernetes.

**platform/infrastructure/terraform/main.tf**

```hcl
provider "aws" {
  region = "us-west-2"
}

resource "aws_instance" "example" {
  ami           = "ami-0c55b159cbfafe1f0"
  instance_type = "t2.micro"

  tags = {
    Name = "example-instance"
  }
}
```
