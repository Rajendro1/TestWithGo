# AccuKnox


1.Clone the github registry image
```
sudo docker pull ghcr.io/rajendro1/accuknox-api-main:latest
```

2. Pull postgres image
```
docker pull postgres
```

3. Run postgres image
```
sudo docker run --name="postgres" --rm -d -p 5433:5432 -e POSTGRES_PASSWORD=perfectPassword -e POSTGRES_USER=postgres -e POSTGRES_DB=accuknox -d postgres
```
4. copy .env file file form my repo and modify DB_HOST with your current docker host
```

ip addr show docker0 | grep -Po 'inet \K[\d.]+'
```

5.Run the image in your loclhost in 8081 port number
```
sudo docker run --name="api" --env-file=.env -p 8081:8081 --net=bridge -d ghcr.io/rajendro1/accuknox-api-main
```
