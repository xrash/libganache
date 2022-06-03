package libganache

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type GanacheRuntime struct {
	cmd          *exec.Cmd
	accsFilename string
}

type RunGanacheOptions struct {
	Executable   string
	GasLimit     string
	GasPrice     string
	StdoutWriter io.Writer
}

var defaultOptions = &RunGanacheOptions{
	Executable:   "ganache",
	GasLimit:     "100000000000",
	GasPrice:     "2000000000",
	StdoutWriter: nil,
}

func RunGanache(options *RunGanacheOptions) (*GanacheRuntime, error) {
	if options == nil {
		options = defaultOptions
	}

	accsFilename := genAccountsFilename()

	var executable string
	if options.Executable != "" {
		executable = options.Executable
	} else {
		executable = defaultOptions.Executable
	}

	var gasLimit string
	if options.GasLimit != "" {
		gasLimit = options.GasLimit
	} else {
		gasLimit = defaultOptions.GasLimit
	}

	var gasPrice string
	if options.GasPrice != "" {
		gasPrice = options.GasPrice
	} else {
		gasPrice = defaultOptions.GasPrice
	}

	cmd := exec.Command(
		executable,
		"--wallet.accountKeysPath",
		accsFilename,
		"--miner.blockGasLimit",
		gasLimit,
		"--gasPrice",
		gasPrice,
	)

	if options.StdoutWriter != nil {
		cmd.Stdout = options.StdoutWriter
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("error running ganache: %w", err)
	}

	return &GanacheRuntime{
		cmd:          cmd,
		accsFilename: accsFilename,
	}, nil
}

func (gr *GanacheRuntime) AccountsFile() (*AccountsFile, error) {
	f, err := retryReadFile(gr.accsFilename)
	if err != nil {
		return nil, fmt.Errorf("couldn't open accounts file: %w", err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("couldn't read accounts file: %w", err)
	}

	af := &AccountsFile{}
	if err := json.Unmarshal(b, &af); err != nil {
		return nil, fmt.Errorf("couldn't unmarshal accounts file: %w", err)
	}

	return af, nil
}

func (gr *GanacheRuntime) Accounts() ([]*Account, error) {

	af, err := gr.AccountsFile()
	if err != nil {
		return nil, fmt.Errorf("error reading accounts file: %w", err)
	}

	accs := make([]*Account, 0)

	for k, v := range af.PrivateKeys {
		publicKeyString := k
		privateKeyString := v

		publicKey := common.HexToAddress(publicKeyString)

		// We take the initial 0x out because HexToECDSA expects this way.
		privateKeyStringWithout0x := privateKeyString[2:]
		privateKey, err := crypto.HexToECDSA(privateKeyStringWithout0x)
		if err != nil {
			return nil, fmt.Errorf("error making private key: %w", err)
		}

		acc := &Account{
			PublicKeyString:  publicKeyString,
			PrivateKeyString: privateKeyString,
			PublicKey:        publicKey,
			PrivateKey:       privateKey,
		}

		accs = append(accs, acc)
	}

	return accs, nil
}

func retryReadFile(filename string) (*os.File, error) {
	tries := 0
	max := 10

	for {
		tries++

		f, err := os.Open(filename)

		if err == nil {
			return f, nil
		}

		if tries >= max {
			return nil, fmt.Errorf("tried %v times and failed: %v", tries, err)
		}

		time.Sleep(time.Second * 1)
	}
}

func genAccountsFilename() string {
	dirname := os.TempDir()
	// @TODO append a random factor to the basename
	basename := strconv.FormatInt(time.Now().UnixMilli(), 10)
	return filepath.Join(dirname, basename)
}
