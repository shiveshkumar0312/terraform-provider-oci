// Copyright (c) 2020, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	oci_bds "github.com/oracle/oci-go-sdk/bds"

	oci_common "github.com/oracle/oci-go-sdk/common"
)

func init() {
	RegisterOracleClient("oci_bds.BdsClient", &OracleClient{initClientFn: initBdsBdsClient})
}

func initBdsBdsClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_bds.NewBdsClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) bdsClient() *oci_bds.BdsClient {
	return m.GetClient("oci_bds.BdsClient").(*oci_bds.BdsClient)
}
