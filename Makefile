OBJ_NAME = 
LDFLAGS = 
install:
	$(eval OBJ_NAME += chess)
	$(eval LDFLAGS += "-w -s")
	cd ./src/; go build -v -ldflags $(LDFLAGS) -o $(OBJ_NAME); mv $(OBJ_NAME) ../bin 
run:
	./bin/chess

doc:
	cd ./cmd/; godoc -http=:6060
