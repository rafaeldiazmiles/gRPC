#Go gRPC services course
Creating a gRPC service from scratch


Command line collection:

sudo docker run --name  rocket-inventory-db -e POSTGRES_PASSWORD=postgres -p 5433:5433 -d postgres
((running a postgres database locally using docker))

docker ps
((Process status, shows all docker processes actively running))

docker-compose up --build
((runs de docker compose file))

sudo docker exec -it ["first 3 letters"] bash
    - psql -U postgres
        - \dt


Go gRPC Services course
================================
sudo apt install protobuf-compiler
((installs protocol buffer library))