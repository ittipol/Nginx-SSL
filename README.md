# NGINX SSL Server

## Gen SSL private key & Cert 
``` bash
# Update hostname in gen.sh and ssl.cnf
# ex. hostname = localhost, myhost
./gen.sh
```

## NGINX SSL Server Configuration

``` nginx
events {

}

http {
  server {

    listen 80 default_server;


    server_name _;


    return 301 https://$host$request_uri;

  }

  server {
    listen 443 ssl;

    ssl_certificate     /etc/nginx/certs/nginx.crt;
    ssl_certificate_key /etc/nginx/certs/nginx.key;

    location / {
      proxy_set_header Host $http_host;
      proxy_pass       http://nextjs:3000/;
    }
  }
}

```

## Start server and application

``` bash
docker compose up -d --build
```

## Test

Open [http://localhost](http://localhost) or [https://localhost](https://localhost) with your browser to test service and application.