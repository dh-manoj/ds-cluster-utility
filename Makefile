
install:
	go build -o ds-cluster-list
	cp ./ds-cluster-list /usr/local/bin
	cp ./scripts/use /usr/local/bin/
	cp ./scripts/con /usr/local/bin/
	cp ./scripts/country /usr/local/bin/
	cp ./scripts/backup-db /usr/local/bin/
	cp ./scripts/restore-db /usr/local/bin/

