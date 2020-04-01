#define _GNU_SOURCE
#include <stdio.h>
#include <stdlib.h>
#include <sched.h>
#include <unistd.h>
#include <string.h>
#include <sys/wait.h>
#include <sys/utsname.h>
#include "util.h"

#define STACK_SIZE 1024 * 1024
#define errExit(msg) do { perror(msg); exit(EXIT_FAILURE); } while(0);

const int namespaces = CLONE_NEWUTS | CLONE_NEWPID | SIGCHLD;
static char child_stack[STACK_SIZE];

static int initNamespace(void *args) {
	struct utsname uts;


	char *hostname = (char *)malloc(sizeof(char) * UUID_LEN);
	getUuid(hostname);

	if (sethostname(hostname, strlen(hostname)) == -1)
		errExit("sethostname error");

	if (uname(&uts) == -1)
		errExit("uname");

	printf("uts nodename in child: %s\n", uts.nodename);
	printf("Child PID: %ld\n", (long)getpid());
	printf("Parent PID: %ld\n", (long)getppid());
	sleep(200);

	free(hostname);
	return 0;
}

void create_namespace() {
	struct utsname uts;

	pid_t child_pid = clone(initNamespace, child_stack + STACK_SIZE, namespaces, NULL);
	if (child_pid == -1)
		errExit("clone");
	printf("clone() returned %ld\n", (long) child_pid);

	sleep(1);

	if (uname(&uts) == -1)
    	errExit("uname");
    printf("uts.nodename in parent: %s\n", uts.nodename);

    printf("My PID: %ld\n", (long)getpid());
    printf("Container PID: %ld\n", (long)child_pid);

    if (waitpid(child_pid, NULL, 0) == -1)
    	errExit("waitpid");
    printf("child has terminated\n");

    printf("clone() returned %ld\n", (long) child_pid);
}