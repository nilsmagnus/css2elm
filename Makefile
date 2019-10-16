.PHONY: clean test

.SILENT: watch-test

WATCH_EXEC=$(if $(shell which inotifywait),"inotifywait","fswatch")
WATCH_ARGS=$(if $(shell which inotifywait),"-qre" "close_write" ".","*/*.elm" "-1")

all: test css2elm

test: *.go *.css
	go test *.go

css2elm: *.go
	go build -o css2elm

clean:
	rm -f css2elm
	rm -f *.elm

watch-test: test
	$(if $(shell which $(WATCH_EXEC)),,echo "No file-watcher found, cannot enter watch-loop"; exit 1)
	date
	while true ; do \
		$(WATCH_EXEC) $(WATCH_ARGS) ; \
		date ; \
		make test; \
		sleep 1; \
	done

integration-test: css2elm sample.css
	./css2elm -input sample.css > example/ElmFromCss.elm
	cd example && elm make ExampleMain.elm
