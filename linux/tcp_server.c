#include <unistd.h>

#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/tcp.h>
#include <arpa/inet.h>
#include <errno.h>

#include <stdio.h>
#include <string.h>

#define PORT 			8888
#define BUFFER_LENGTH 	20480

int main() {
	int sockfd = socket(AF_INET, SOCK_STREAM, 0);
	if (sockfd < 0) {
		perror("socket create failed");
		return 0;
	}

	struct sockaddr_in addr;
	memset(&addr, 0, sizeof(struct sockaddr_in));

	addr.sin_family = AF_INET;
	addr.sin_port = htons(PORT);
	addr.sin_addr.s_addr = INADDR_ANY;

	if (bind(sockfd, (struct sockaddr *)&addr, sizeof(struct sockaddr_in)) < 0) {
		perror("socket bind failed");
		return 0;
	}

	if (listen(sockfd, 5) < 0) {
		perror("listen failed");
		return 0;
	}

	struct sockaddr_in client_addr;
	memset(&client_addr, 0, sizeof(struct sockaddr_in));

	socklen_t client_len = sizeof(client_addr);
	int clientfd = accept(sockfd, (struct sockaddr *)&client_addr, &client_len);
	if (clientfd <= 0) {
		perror("accept failed");
		return 0;
	}

	while (1) {
		char buffer[BUFFER_LENGTH] = {0};		
		int ret = recv(clientfd, buffer, BUFFER_LENGTH, 0);
		if (ret <= 0) {
			if (errno == EAGAIN || errno == EWOULDBLOCK) {
				printf("read all data");
			} else {
				perror("recv error");
			}
		} else {
			printf("read: %s\n", buffer);
		}
	}

	close(sockfd);

	return 0;
}