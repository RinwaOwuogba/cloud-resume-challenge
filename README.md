# Cloud Resume Project -- without the certificate

This is my take on [Forrest Brazeal's cloud resume challenge](https://cloudresumechallenge.dev). A critical component remains missing; the cloud certification, but this was fun to build nonetheless.

TLDR; It is a three-tiered web application that showcases your resume.

Live url: [https://bolarinwa.dev/index.html](https://bolarinwa.dev/index.html)

### Frontend

Built using HTML, CSS with tailwindCSS and vanilla JavaScript.

### Backend

- Implemented using Golang.
- Battered with tests using the go standard testing package.
- Application data is persisted on Firestore.

### Infrastructure

Infrastructure is completely automated with Terraform. It consists of:

- An HTTP Load balancer
- IAM resources
- Storage buckets
- VPC network
- DNS records
- Cloud functions, etcetera

### CI/CD

Continuous integration and delivery is configured with Github actions. Cheers!
