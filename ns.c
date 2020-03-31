#define _GNU_SOURCE
#include<stdio.h>
#include<stdlib.h>
#include<sched.h>
#include<unistd.h>
#include<sys/wait.h>

#define STACK_SIZE 1024 * 1024

static char child_stack[STACK_SIZE];

static int child() {
	printf("Child PID: %ld\n", (long)getpid());
	printf("Parent PID: %ld\n", (long)getppid());
	return 0;
}

void clone_ns() {
	pid_t child_pid = clone(child, child_stack + STACK_SIZE, CLONE_NEWPID | SIGCHLD, NULL);
	printf("clone() = %lds\n", (long)child_pid);
	printf("My PID: %ld\n", (long)getpid());
	waitpid(child_pid, NULL, 0);
}