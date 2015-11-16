all:
	go build goreader.go
	gcc -o creader creader.c
	gcc -o cwriter cwriter.c

fclean:
	rm -f goreader creader cwriter

clean: fclean
