.PHONY: clean
clean:
	-rm -fr thingsboard

.PHONY: server
run/server:
	docker run --rm -it -p 8080:4010 -v $$PWD:/tmp stoplight/prism:4 mock -h 0.0.0.0 /tmp/thingsboard.json

.PHONY: run/client
run/client:
	SERVER='http://localhost:8080' go run ./main.go

.PHONY: edit
edit: clean
	sed -i -E -e 's/\{\?[A-Za-z,]+\}//g' ./thingsboard.json
	docker run --rm -it -v $$PWD:$$PWD -w $$PWD node:lts-alpine ./edit.js ./thingsboard.json

.PHONY: generate
generate: edit
	mkdir -p thingsboard
	docker run --rm -it -v $$PWD:$$PWD -w $$PWD hidori/oapi-codegen \
		-o ./thingsboard/types.gen.go -package thingsboard -generate types ./thingsboard.json

.PHONY: update
update:
	curl -o ./thingsboard.json https://demo.thingsboard.io/v3/api-docs/thingsboard
