#include <unistd.h>
#include <fcntl.h>

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
		data[i].id = i;
		data[i].comment[0] = 't';
		data[i].comment[1] = 'e';
		data[i].comment[2] = 's';
		data[i].comment[3] = 't';
		data[i].comment[4] = '\0';
	}
	if (argc != 2)
	{
		write(1, "need file path\n", 15);
		return (1);
	}
	fd = open(argv[1], O_CREAT | O_TRUNC | O_RDWR);
	if (fd == -1)
	{
		write(1, "couldn't open file\n", 19);
		return (1);
	}
	write(fd, &data, sizeof(t_foo) * 10);
	close(fd);
	return (0);
}
