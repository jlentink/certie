# Certie

Certie is a simple and easy CLI tool that is able to pull 
information from a TLS (SSL) server and display the output in multiple formats (Text, Yaml, XML, JSON).

Certie will exit with status code 1 if the server is not reachable or the certificate is invalid.

examples:

output in text format:
```
$ certie www.google.com
Remote Addr: 142.251.36.36:443
Common:      www.google.com
SANs :       [www.google.com]
Issuer:      GTS CA 1C3
Not before:  2023-02-01 20:43:59 +0100 CET
Not After:   2023-04-26 21:43:58 +0200 CEST
Expired:     Valid for 67 days
```

output in JSON format:
- ```certie -f json https://www.google.com```

Define a diffrent port than 443:
- ```certie -f json https://www.google.com:4433```

Hope my hobby project to play arround with TLS and go is useful for you.