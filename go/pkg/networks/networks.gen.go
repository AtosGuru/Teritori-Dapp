// Code generated by "generateGoNetworks.ts"; DO NOT EDIT.

package networks

import (
	"encoding/json"
	"github.com/pkg/errors"
)

const (
	NetworkKindUnknown  = NetworkKind("Unknown")
	NetworkKindEthereum = NetworkKind("Ethereum")
	NetworkKindCosmos   = NetworkKind("Cosmos")
	NetworkKindSolana   = NetworkKind("Solana")
	NetworkKindGno      = NetworkKind("Gno")
)

func UnmarshalNetwork(b []byte) (Network, error) {
	var base NetworkBase
	if err := json.Unmarshal(b, &base); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal network base")
	}
	switch base.Kind {
	case NetworkKindEthereum:
		var n EthereumNetwork
		if err := json.Unmarshal(b, &n); err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal Ethereum network")
		}
		return &n, nil

	case NetworkKindCosmos:
		var n CosmosNetwork
		if err := json.Unmarshal(b, &n); err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal Cosmos network")
		}
		return &n, nil

	case NetworkKindSolana:
		var n SolanaNetwork
		if err := json.Unmarshal(b, &n); err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal Solana network")
		}
		return &n, nil

	case NetworkKindGno:
		var n GnoNetwork
		if err := json.Unmarshal(b, &n); err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal Gno network")
		}
		return &n, nil

	default:
		return &base, nil
	}
}

func (netstore NetworkStore) GetEthereumNetwork(id string) (*EthereumNetwork, error) {
	network, err := netstore.GetNetwork(id)
	if err != nil {
		return nil, err
	}
	cn, ok := network.(*EthereumNetwork)
	if !ok {
		return nil, ErrWrongType
	}
	return cn, nil
}

func (netstore NetworkStore) MustGetEthereumNetwork(id string) *EthereumNetwork {
	network, err := netstore.GetEthereumNetwork(id)
	if err != nil {
		panic(err)
	}
	return network
}

func (netstore NetworkStore) GetCosmosNetwork(id string) (*CosmosNetwork, error) {
	network, err := netstore.GetNetwork(id)
	if err != nil {
		return nil, err
	}
	cn, ok := network.(*CosmosNetwork)
	if !ok {
		return nil, ErrWrongType
	}
	return cn, nil
}

func (netstore NetworkStore) MustGetCosmosNetwork(id string) *CosmosNetwork {
	network, err := netstore.GetCosmosNetwork(id)
	if err != nil {
		panic(err)
	}
	return network
}

func (netstore NetworkStore) GetSolanaNetwork(id string) (*SolanaNetwork, error) {
	network, err := netstore.GetNetwork(id)
	if err != nil {
		return nil, err
	}
	cn, ok := network.(*SolanaNetwork)
	if !ok {
		return nil, ErrWrongType
	}
	return cn, nil
}

func (netstore NetworkStore) MustGetSolanaNetwork(id string) *SolanaNetwork {
	network, err := netstore.GetSolanaNetwork(id)
	if err != nil {
		panic(err)
	}
	return network
}

func (netstore NetworkStore) GetGnoNetwork(id string) (*GnoNetwork, error) {
	network, err := netstore.GetNetwork(id)
	if err != nil {
		return nil, err
	}
	cn, ok := network.(*GnoNetwork)
	if !ok {
		return nil, ErrWrongType
	}
	return cn, nil
}

func (netstore NetworkStore) MustGetGnoNetwork(id string) *GnoNetwork {
	network, err := netstore.GetGnoNetwork(id)
	if err != nil {
		panic(err)
	}
	return network
}