package singlesig_test

import (
	"strconv"
	"testing"

	"github.com/DharitriOne/drt-chain-core-go/hashing/sha256"
	"github.com/DharitriOne/drt-chain-crypto-go/signing"
	"github.com/DharitriOne/drt-chain-crypto-go/signing/mcl"
	"github.com/DharitriOne/drt-chain-crypto-go/signing/mcl/singlesig"
	"github.com/stretchr/testify/require"
)

func BenchmarkBlsSingleSigner_Sign(b *testing.B) {
	signer := singlesig.NewBlsSigner()
	suite := mcl.NewSuiteBLS12()
	kg := signing.NewKeyGenerator(suite)
	privKey, _ := kg.GeneratePair()

	var err error
	nbMessages := 10000
	messages := make([][]byte, 0, 10000)
	hasher := sha256.NewSha256()

	for i := 0; i < nbMessages; i++ {
		strIdx := strconv.Itoa(i)
		messages = append(messages, hasher.Compute(strIdx))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err = signer.Sign(privKey, messages[i%nbMessages])
		require.Nil(b, err)
	}
}

func BenchmarkBlsSingleSigner_Verify(b *testing.B) {
	signer := singlesig.NewBlsSigner()
	suite := mcl.NewSuiteBLS12()
	kg := signing.NewKeyGenerator(suite)
	privKey, pubKey := kg.GeneratePair()

	var err error
	nbMessages := 10000
	messages := make([][]byte, 0, 10000)
	signatures := make([][]byte, 0, 10000)
	hasher := sha256.NewSha256()

	for i := 0; i < nbMessages; i++ {
		strIdx := strconv.Itoa(i)
		messages = append(messages, hasher.Compute(strIdx))
		signature, err := signer.Sign(privKey, messages[i%nbMessages])
		require.Nil(b, err)
		signatures = append(signatures, signature)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err = signer.Verify(pubKey, messages[i%nbMessages], signatures[i%nbMessages])
		require.Nil(b, err)
	}
}
