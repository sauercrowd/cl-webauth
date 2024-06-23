(defpackage :cl-webauthn
  (:use :cl)
  (:import-from #:SB-SYS
		#:INT-SAP)
  (:export :validate-login))

(in-package :cl-webauthn)

(pushnew (asdf:system-source-directory 'cl-webauthn)
	cffi:*foreign-library-directories*
	:test #'equal)

(cffi:define-foreign-library webauthn
  (t (:default "libwebauthn")))
      

(cffi:use-foreign-library webauthn)

(cffi:defcstruct validate-login-result
  "result for ValidateLogin"
  (isValid :int)
  (error :string))

(cffi:defcfun "ValidateLogin" (:pointer
			       (:struct validate-login-result))
    "Take credetials and return 1 if they're invalid, and 0 if they're valid"
    (cId :string)
    (cName :string)
    (cCredentialJSONs :string)
    (cSessionDataJson :string))


(defun validate-login (id name credential-jsons session-data-json)
  (let ((result  (validatelogin id name credential-jsons session-data-json)))
    (unwind-protect
	 (cffi:with-foreign-slots ((error isValid) result (:struct validate-login-result))
	   (values (= isValid 1) error))
      (cffi:foreign-free result))))


;  (if (= 0 (validatelogin id name credential-jsons session-data-json))
;      t
;      nil))
