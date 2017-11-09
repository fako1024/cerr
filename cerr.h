#ifndef CERR_H
#define CERR_H

#include <string.h>

#define CERR_BUFFER_SIZE 8192

inline void setCErr(char* errPtr, const char* errString) {
	if (strlen(errString) >= CERR_BUFFER_SIZE)
		strncpy(errPtr, errString, CERR_BUFFER_SIZE);
	else
		strcpy(errPtr, errString);
}

#endif /* CERR_H */
