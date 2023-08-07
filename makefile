mongodb:
	docker run --name freetime-together-mongo -e MONGODB_DATABASE=admin -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=qwerty -p 27017:27017 -d mongo