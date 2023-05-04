runtest:
# run testdb in docker
	docker-compose  up testdb -d
	sleep 4
	go test $(path) -v || true
	docker-compose rm -sv testdb --force
