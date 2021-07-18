#ifndef GSE_GOTYPES_H
#define GSE_GOTYPES_H

#include "http-signatures.h"

#ifdef __cplusplus
extern "C" {
#endif

// Create a JA3 Transport utilizing a JA3 hash.
extern void* CreateJA3Transport(char*);


// Send a request utilising a transport and a request.
extern void* SendRequest(void*, Request*);

#ifdef __cplusplus
}
#endif

#endif // GSE_GOTYPES_H
