package compliance

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/asaskevich/govalidator"
	"github.com/fonero-project/fonero-golang/clients/fonerotoml"
	"github.com/fonero-project/fonero-golang/keypair"
	"github.com/fonero-project/fonero-golang/support/errors"
)

func (r *AuthRequest) Populate(request *http.Request) *AuthRequest {
	r.DataJSON = request.PostFormValue("data")
	r.Signature = request.PostFormValue("sig")
	return r
}

// Validate is using govalidator to check if fields are valid and also
// runs Validate method on authData.
// This method only performs data validation. You should also call
// VerifySignature to confirm that signature is valid.
func (r *AuthRequest) Validate() error {
	valid, err := govalidator.ValidateStruct(r)

	if !valid {
		return err
	}

	authData := AuthData{}
	err = json.Unmarshal([]byte(r.DataJSON), &authData)
	if err != nil {
		return errors.Wrap(err, "Data is not valid JSON")
	}

	// Validate DataJSON
	err = authData.Validate()
	if err != nil {
		return errors.New("Invalid Data: " + err.Error())
	}

	return nil
}

// VerifySignature verifies if signature is valid. It makes a network connection
// to sender server in order to obtain fonero.toml file and signing key.
func (r *AuthRequest) VerifySignature(sender string) error {
	signatureBytes, err := base64.StdEncoding.DecodeString(r.Signature)
	if err != nil {
		return errors.New("Signature is not base64 encoded")
	}

	senderFoneroToml, err := fonerotoml.GetFoneroTomlByAddress(sender)
	if err != nil {
		return errors.Wrap(err, "Cannot get fonero.toml of sender domain")
	}

	if senderFoneroToml.SigningKey == "" {
		return errors.New("No SIGNING_KEY in fonero.toml of sender")
	}

	kp, err := keypair.Parse(senderFoneroToml.SigningKey)
	if err != nil {
		return errors.New("SigningKey is invalid")
	}

	err = kp.Verify([]byte(r.DataJSON), signatureBytes)
	if err != nil {
		return errors.New("Signature is invalid")
	}

	return nil
}

// Data returns AuthData from the request.
func (r *AuthRequest) Data() (data AuthData, err error) {
	err = json.Unmarshal([]byte(r.DataJSON), &data)
	if err != nil {
		err = errors.Wrap(err, "Error unmarshalling JSON data")
	}
	return
}

// ToURLValues returns AuthData encoded as url.Values.
func (r *AuthRequest) ToURLValues() url.Values {
	return url.Values{
		"data": []string{r.DataJSON},
		"sig":  []string{r.Signature},
	}
}
