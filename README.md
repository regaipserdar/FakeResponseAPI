# VulnApi - A Simple Fake API for Web Security Testing

VulnApi is a simple fake API created for web security testing and demonstrating how attackers might identify technologies used by a web application. This project is inspired by the blog post [Power-BCheck: How Can It Enhance Your Web Security Testing](https://www.linkedin.com/pulse/power-bcheck-how-can-enhance-your-web-security-testing-serdar-k%2525C3%2525B6k%2525C3%2525A7%2525C3%2525BC-mc3ef%3FtrackingId=sR8woToZRH2G8T4nSRHBhw%253D%253D) by Serdar Kökçü.

## Features

- Emulates a simple web server with fake ASP.NET response headers.
- Demonstrates how attackers can identify server technologies and versions.
- Designed for educational and testing purposes.
- Easily customizable to add more fake headers or responses.

## Usage

To use VulnApi, follow these steps:

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/regaipserdar/VulnApi.git
   cd VulnApi


**Prerequisites**

Before getting started, make sure you have Go (Golang) installed on your system.

**Running VulnApi**

To run the VulnApi server, open your terminal and execute the following command:

```bash
    go run main.go  
```

Then, open your web browser and navigate to http://localhost:9090 to see the welcome message.

To see the fake ASP.NET response headers, visit http://localhost:9090/fake-asp-response.

Contributions

Contributions are welcome! If you have any ideas for improvements or new features, feel free to create an issue or submit a pull request.