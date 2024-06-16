
#ifndef WEBAUTHN_H
#define WEBAUTHN_H

#ifdef __cplusplus
extern "C" {
#endif

//cId, cName, cCredentialJSONs, cSessionDataJson
int ValidateLogin(char* cId, char* cName, char* cCredentialsJSONs, char* cSessionDataJson);

#ifdef __cplusplus
}
#endif

#endif // WEBAUTHN_H
