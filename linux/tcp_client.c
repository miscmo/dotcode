#include <unistd.h>

#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/tcp.h>
#include <arpa/inet.h>

#include <stdio.h>
#include <string.h>
#include <errno.h>

#define PORT			8888

const char *ip = "127.0.0.1";

int main() {

	int sockfd = socket(AF_INET, SOCK_STREAM, 0);
	if (sockfd < 0) {
		perror("sockfd create failed");
		return 0;
	}

	struct sockaddr_in addr;
	memset(&addr, 0, sizeof(struct sockaddr_in));
	addr.sin_family = AF_INET;
	addr.sin_port = htons(PORT);
	inet_pton(AF_INET, ip, &addr.sin_addr);

	if (connect(sockfd, (struct sockaddr *)&addr, sizeof(struct sockaddr_in)) < 0) {
		perror("connect error");
		return 0;
	}

	char *buffer = "helloworld";

	int ret = send(sockfd, buffer, strlen(buffer), 0);
	if (ret <= 0) {
		perror("send data error");
		return 0;
	}

	close(sockfd);

	return 0;
}