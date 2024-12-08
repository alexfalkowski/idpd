include bin/build/make/http.mak
include bin/build/make/git.mak

# Run the application.
run:
	(cd test && CONFIG_FILE=.config/server.yml ../idpd server)
