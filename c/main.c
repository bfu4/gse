#include "include/go-types.h"
#include <stdio.h>

#define HASH "771,4865-4867-4866-49195-49199-52393-52392-49196-49200-49162-49161-49171-49172-156-157-47-53-10,0-23-65281-10-11-35-16-5-51-43-13-45-28-21,29-23-24-25-256-257,0"

int main(void) {
   long transport = CreateJA3Transport(HASH);
   printf("%lu\n", transport);
   return 0;
}
