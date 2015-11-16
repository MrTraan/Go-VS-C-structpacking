all:
	go build goreader.go
	go build gowriter.go
	gcc -o creader creader.c
	gcc -o cwriter cwriter.c

fclean:
	rm -f goreader gowriter creader cwriter

clean: fclean
