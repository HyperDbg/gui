package ntdll

//go:generate -command mkcode go run mkcode.go --
//go:generate mkcode $GOFILE

/*
func:
NTSTATUS LsaOpenPolicy(
  _In_    PLSA_UNICODE_STRING    SystemName,
  _In_    PLSA_OBJECT_ATTRIBUTES ObjectAttributes,
  _In_    ACCESS_MASK            DesiredAccess,
  _Inout_ PLSA_HANDLE            PolicyHandle
);
*/

/*
func:
NTSTATUS LsaAddAccountRights(
  _In_ LSA_HANDLE          PolicyHandle,
  _In_ PSID                AccountSid,
  _In_ PLSA_UNICODE_STRING UserRights,
  _In_ ULONG               CountOfRights
);
*/

/*
func:
NTSTATUS LsaEnumerateAccountRights(
  _In_  LSA_HANDLE          PolicyHandle,
  _In_  PSID                AccountSid,
  _Out_ PLSA_UNICODE_STRING *UserRights,
  _Out_ PULONG              CountOfRights
);
*/

/*
func:
NTSTATUS LsaRemoveAccountRights(
  _In_ LSA_HANDLE          PolicyHandle,
  _In_ PSID                AccountSid,
  _In_ BOOLEAN             AllRights,
  _In_ PLSA_UNICODE_STRING UserRights,
  _In_ ULONG               CountOfRights
);
*/

/*
func:
NTSTATUS LsaClose(
  _In_ LSA_HANDLE ObjectHandle
);
*/

type (
	LsaUnicodeString    UnicodeString
	LsaHandle           Handle
	LsaObjectAttributes ObjectAttributes
)
