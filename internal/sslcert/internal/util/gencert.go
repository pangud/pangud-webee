package util

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"

	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/providers/dns/tencentcloud"
	"github.com/go-acme/lego/v4/registration"
)

// You'll need a user or account type that implements acme.User
type MyUser struct {
	Email        string
	Registration *registration.Resource
	key          crypto.PrivateKey
}

func (u *MyUser) GetEmail() string {
	return u.Email
}
func (u MyUser) GetRegistration() *registration.Resource {
	return u.Registration
}
func (u *MyUser) GetPrivateKey() crypto.PrivateKey {
	return u.key
}

func GenCertByDnspod(host, email, secretId, secretKey string) (privateKey, cert []byte, err error) {

	// Create a user. New accounts need an email and private key to start.
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	myUser := MyUser{
		Email: email,
		key:   key,
	}

	config := lego.NewConfig(&myUser)

	// This CA URL is configured for a local dev instance of Boulder running in Docker in a VM.
	// config.CADirURL = "http://192.168.99.100:4000/directory"
	config.Certificate.KeyType = certcrypto.EC256

	// A client facilitates communication with the CA server.
	client, err := lego.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	providerCfg := tencentcloud.NewDefaultConfig()

	providerCfg.SecretID = secretId
	providerCfg.SecretKey = secretKey

	dnsprovider, err := tencentcloud.NewDNSProviderConfig(providerCfg)

	if err != nil {
		return
	}

	err = client.Challenge.SetDNS01Provider(dnsprovider)

	if err != nil {
		return
	}

	// New users will need to register
	reg, err := client.Registration.Register(registration.RegisterOptions{TermsOfServiceAgreed: true})
	if err != nil {
		return
	}
	myUser.Registration = reg

	request := certificate.ObtainRequest{
		Domains: []string{host},
		Bundle:  true,
	}
	var certResource *certificate.Resource
	certResource, err = client.Certificate.Obtain(request)
	if err != nil {
		return
	}

	// Each certificate comes back with the cert bytes, the bytes of the client's
	// private key, and a certificate URL. SAVE THESE TO DISK.
	// fmt.Printf("private_key:\n%s\n", certResource.PrivateKey)
	// fmt.Printf("private_key:\n%s\n", certResource.Certificate)

	privateKey = certResource.PrivateKey
	cert = certResource.Certificate
	return
	// ... all done.
}
