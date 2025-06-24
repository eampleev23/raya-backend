.PHONY: all
all: ;

.PHONY: pg
pg:
	docker run --rm \
		--name=raya_local \
		-v $(abspath ./db/init/):/docker-entrypoint-initdb.d \
		-v $(abspath ./db/data/):/var/lib/postgresql/data \
		-e POSTGRES_PASSWORD="P@ssw0rd" \
		-d \
		-p 5432:5432 \
		postgres:16.3

.PHONY: stop-pg
stop-pg:
	docker stop raya_local

.PHONY: clean-data
clean-data:
	sudo rm -rf ./db/data/