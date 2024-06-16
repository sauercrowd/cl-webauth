(in-package :cl-webauthn)

(cffi:defcfun "ValidateLogin" :int
    "Take credetials and return 1 if they're invalid, and 0 if they're valid"
    (cId :string)
	    (cName :string)
	    (cCredentialJSONs :string)
	    (cSessionDataJson :string))


(defun validate-login (id name credential-jsons session-data-json)
  (if (= 0 (validatelogin id name credential-jsons session-data-json))
      t
      nil))
