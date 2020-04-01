#include<stdio.h>
#include<stdlib.h>
#include<sched.h>
#include<unistd.h>
#include<sys/wait.h>

extern void clone_ns();
extern void create_namespace();