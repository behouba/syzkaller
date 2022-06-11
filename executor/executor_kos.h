// Copyright 2017 syzkaller project authors. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// Kaspersky OS
#include <errno.h>
#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/ioctl.h>
#include <sys/mman.h>
#include <sys/prctl.h>
#include <sys/syscall.h>
#include <unistd.h>
// #include "nocover.h"
#include "core/VMM.cdl.h"

static void os_init(int argc, char** argv, void* data, size_t data_size)
{
	if (mmap(data, data_size, PROT_READ | PROT_WRITE | PROT_EXEC, MAP_ANON | MAP_PRIVATE | MAP_FIXED, -1, 0) != data)
		fprintf(stderr, "mmap of data segment failed");
}

static intptr_t execute_syscall(const call_t* c, intptr_t a[kMaxArgs])
{
	// To do.
	return;
}