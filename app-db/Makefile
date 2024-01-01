PHONY: down up clog
USER_NAME=marvin

clog: #changelog
	@echo "Changelog name:";\
	read changelog;\
	/bin/sh ./scripts/create-changelog.sh -c $$changelog -u ${USER_NAME}

up:
	cd docs/liquibase && liquibase update

down:
	@echo "Count:";\
	read count;\
	cd docs/liquibase && liquibase rollback-count $$count