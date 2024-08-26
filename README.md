# Web Server in Golang
## Running locally with docker compose
1. In the root directory, run the following command:
```
docker compose up
```
or to run in detached mode (run containers in background):
```
docker compose up -d
```
2. The web server will be running on http://localhost:8080.
![image](https://github.com/user-attachments/assets/da25db3d-2000-42c4-9dea-cb1e0a072506)

3. To stop the web server, run:
```
docker compose down
```