#ifndef GSE_HTTP_SIGNATURES_H
#define GSE_HTTP_SIGNATURES_H

#ifdef __cplusplus 
extern "C" {
#endif

// Header type for defining what a header might look like.
typedef struct _Header {
   char* name;
   char* value;
} Header;

// Request type defining the basis of a request.
typedef struct _Request {
   char* method;
   char* url;
   char* body;
   Header headers[];
} Request;

#ifdef __cplusplus
}
#endif

#endif // GSE_HTTP_SIGNATURES_H
