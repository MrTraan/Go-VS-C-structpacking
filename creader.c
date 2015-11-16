#include <unistd.h>
#include <fcntl.h>
#include <sys/types.h>
#include <sys/uio.h>
#include <stdio.h>

typedef struct	s_foo
{
	int		id;
	char	comment[5];
}				t_foo;

int		main(int argc, char **argv)
{
	int		fd;
	t_foo	data[10];

	for (int i = 0; i < 10; i++)
	{
		data[i].id = 0;
		data[i].comment[0] = '\0';
	}
	if (argc != 2)
	{
		write(1, "need file path\n", 15);
		return (1);
	}
	fd = open(argv[1], O_RDONLY);
	if (fd == -1)
	{
		write(1, "couldn't open file\n", 19);
		return (1);
	}
	read(fd, &data, sizeof(t_foo) * 10);
	close(fd);
	for (int i = 0; i < 10; i++)
		printf("id: %d\tcomment: %s\t\n", data[i].id, data[i].comment);
	return (0);
}
