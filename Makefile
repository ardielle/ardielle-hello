all:: 
	@echo 'targets:'
	@echo '  go-server - run the go server'
	@echo '  go-client - run the go client'
	@echo '  java-client - run the java client'
	@echo '  java-server - run the java server'

go-server::
	@go run go/cmd/hellod/main.go

go-client::
	@go run go/cmd/hello/main.go

java-client: java-build
	@(cd java; mvn -q compile exec:java -Dexec.mainClass="Hello")

java-server: java-build
	@(cd java; mvn -q compile exec:java -Dexec.mainClass="HelloImpl")

java-build: java/src/main/java/HelloClient.java

java/src/main/java/HelloClient.java:
	@(cd java; rdl generate -o target/generated-sources/rdl java-client ../hello.rdl)

regenerate-go:
	rdl generate -o go go-server-project hello.rdl

clean::
	@(cd java; mvn -q clean)
