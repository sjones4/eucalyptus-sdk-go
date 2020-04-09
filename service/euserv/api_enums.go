// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package euserv

type CertificateFormatEnum string

// Enum values for CertificateFormatEnum
const (
	CertificateFormatEnumPem CertificateFormatEnum = "pem"
	CertificateFormatEnumDer CertificateFormatEnum = "der"
)

func (enum CertificateFormatEnum) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CertificateFormatEnum) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type StateEnum string

// Enum values for StateEnum
const (
	StateEnumDisable StateEnum = "DISABLE"
	StateEnumEnable  StateEnum = "ENABLE"
	StateEnumStart   StateEnum = "START"
	StateEnumStop    StateEnum = "STOP"
)

func (enum StateEnum) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum StateEnum) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}