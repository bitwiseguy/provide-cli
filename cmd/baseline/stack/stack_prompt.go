package stack

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/provideservices/provide-cli/cmd/common"
	"github.com/spf13/cobra"
)

const promptStepRun = "Run"
const promptStepStop = "Stop"
const promptStepLogs = "Logs"

var emptyPromptArgs = []string{promptStepRun, promptStepStop, promptStepLogs}
var emptyPromptLabel = "What would you like to do"

var boolPromptArgs = []string{"No", "Yes"}
var tunnelPromptLabel = "Would you like to set up a tunnel"
var tunnelAPIPromptLabel = "Would you like to set up a API tunnel"
var tunnelMessagingPromptLabel = "Would you like to set up a messaging tunnel"
var autoRemovePromptLabel = "Would you like to automatically remove"
var localVaultPromptLabel = "Would you like to set up vault locally"
var localIdentPromptLabel = "Would you like to set up ident locally"
var localNchainPromptLabel = "Would you like to set up nachain locally"
var localPrivacyPromptLabel = "Would you like to set up privacy locally"

var SoRPromptArgs = []string{"SAP", "Service Now", "Sales Force"}
var SoRPromptLabel = "Select a Sor"

// General Endpoints
func generalPrompt(cmd *cobra.Command, args []string, currentStep string) {
	switch step := currentStep; step {
	case promptStepRun:
		common.RequireOrganization()
		if Optional {
			if name == "" {
				name = common.FreeInput("Name", "", common.NoValidation)
			}
			if common.APIEndpoint == "" {
				common.APIEndpoint = common.FreeInput("API endpoint", "", common.NoValidation)
			}
			if common.MessagingEndpoint == "" {
				common.MessagingEndpoint = common.FreeInput("Messaging endpoint", "", common.NoValidation)
			}
			if !common.Tunnel {
				common.Tunnel = common.SelectInput(boolPromptArgs, tunnelPromptLabel) == "Yes"
			}
			if !common.ExposeAPITunnel {
				common.ExposeAPITunnel = common.SelectInput(boolPromptArgs, tunnelAPIPromptLabel) == "Yes"
			}
			if !common.ExposeMessagingTunnel {
				common.ExposeMessagingTunnel = common.SelectInput(boolPromptArgs, tunnelMessagingPromptLabel) == "Yes"
			}
			if sorID == "" {
				sorID = common.SelectInput(SoRPromptArgs, SoRPromptLabel)
			}
			if sorURL == "" {
				sorURL = common.FreeInput("SoR URL", "", common.NoValidation)
			}
			if apiHostname == "" {
				apiHostname = common.FreeInput("API Hostname", "", common.NoValidation)
			}
			if port == 8080 {
				port, _ = strconv.Atoi(common.FreeInput("Port", "8080", common.NumberValidation))
			}
			if consumerHostname == name+"-consumer" {
				consumerHostname = common.FreeInput("Consumer Hostname", name+"-consumer", common.NoValidation)
			}
			if natsHostname == name+"-nats" {
				natsHostname = common.FreeInput("Nats Hostname", name+"-nats", common.NoValidation)
			}
			if natsPort == 4222 {
				natsPort, _ = strconv.Atoi(common.FreeInput("Nats Port", "4222", common.NumberValidation))
			}
			if natsWebsocketPort == 4221 {
				natsWebsocketPort, _ = strconv.Atoi(common.FreeInput("Nats Websocket Port", "4221", common.NumberValidation))
			}
			if natsAuthToken == "testtoken" {
				natsAuthToken = common.FreeInput("Nats Auth Token", "testtoken", common.NoValidation)
			}
			if natsStreamingHostname == name+"-nats-streaming" {
				natsStreamingHostname = common.FreeInput("Nats Streaming Token", name+"-nats-streaming", common.NoValidation)
			}
			if natsStreamingPort == 4220 {
				natsStreamingPort, _ = strconv.Atoi(common.FreeInput("Nats Streaming Port", "4221", common.NumberValidation))
			}
			if redisHostname == name+"-reddis" {
				redisHostname = common.FreeInput("Reddis Host Name", name+"-reddis", common.NoValidation)
			}
			if redisPort == 6379 {
				redisPort, _ = strconv.Atoi(common.FreeInput("Reddis Port", "6379", common.NumberValidation))
			}
			if redisHosts == redisHostname+":"+strconv.Itoa(redisContainerPort) {
				redisPort, _ = strconv.Atoi(common.FreeInput("Reddis Port", redisHostname+":"+strconv.Itoa(redisContainerPort), common.NoValidation))
			}
			if !autoRemove {
				autoRemove = common.SelectInput(boolPromptArgs, autoRemovePromptLabel) == "Yes"
			}
			if logLevel == "DEBUG" {
				logLevel = common.FreeInput("Reddis Host Name", "DEBUG", common.NoValidation)
			}
			if jwtSignerPublicKey == "" {
				jwtSignerPublicKey = common.FreeInput("JWT Signer Public Key", "", common.NoValidation)
			}
			if identAPIHost == "ident.provide.services" {
				nchainAPIHost = common.FreeInput("Ident API Host", "ident.provide.services", common.NoValidation)
			}
			if identAPIScheme == "https" {
				nchainAPIScheme = common.FreeInput("Ident API Scheme", "https", common.NoValidation)
			}
			if nchainAPIHost == "nchain.provide.services" {
				nchainAPIHost = common.FreeInput("Nchain API Host", "nchain.provide.services", common.NoValidation)
			}
			if nchainAPIScheme == "https" {
				nchainAPIScheme = common.FreeInput("Nchain API Scheme", "https", common.NoValidation)
			}
			if privacyAPIHost == "privacy.provide.services" {
				privacyAPIHost = common.FreeInput("Privacy API Host", "privacy.provide.services", common.NoValidation)
			}
			if privacyAPIScheme == "https" {
				privacyAPIScheme = common.FreeInput("Privacy API Scheme", "https", common.NoValidation)
			}
			if vaultAPIHost == "vault.provide.services" {
				vaultAPIHost = common.FreeInput("Vault API Host", "vault.provide.services", common.NoValidation)
			}
			if vaultAPIScheme == "https" {
				vaultAPIScheme = common.FreeInput("Vault API Scheme", "https", common.NoValidation)
			}
			if vaultRefreshToken == os.Getenv("VAULT_REFRESH_TOKEN") {
				vaultRefreshToken = common.FreeInput("Vault API Refresh Token", os.Getenv("VAULT_REFRESH_TOKEN"), common.NoValidation)
			}
			if vaultSealUnsealKey == os.Getenv("VAULT_SEAL_UNSEAL_KEY") {
				vaultSealUnsealKey = common.FreeInput("Vault Un/Seal Token", os.Getenv("VAULT_SEAL_UNSEAL_KEY"), common.NoValidation)
			}
			if !withLocalVault {
				withLocalVault = strings.ToLower(common.SelectInput(boolPromptArgs, localVaultPromptLabel)) == "yes"
			}
			if !withLocalIdent {
				withLocalIdent = strings.ToLower(common.SelectInput(boolPromptArgs, localIdentPromptLabel)) == "yes"
			}
			if !withLocalNChain {
				withLocalNChain = strings.ToLower(common.SelectInput(boolPromptArgs, localNchainPromptLabel)) == "yes"
			}
			if !withLocalPrivacy {
				withLocalPrivacy = strings.ToLower(common.SelectInput(boolPromptArgs, localPrivacyPromptLabel)) == "yes"
			}
			if organizationRefreshToken == os.Getenv("PROVIDE_ORGANIZATION_REFRESH_TOKEN") {
				organizationRefreshToken = common.FreeInput("Organization Refresh Token", os.Getenv("PROVIDE_ORGANIZATION_REFRESH_TOKEN"), common.NoValidation)
			}
			if baselineOrganizationAddress == "0x" {
				baselineOrganizationAddress = common.FreeInput("Baseline Organization Address", "0x", common.NoValidation)
			}
			if baselineRegistryContractAddress == "0x" {
				baselineOrganizationAddress = common.FreeInput("Baseline Registry Contract Address", "0x", common.HexValidation)
			}
			if baselineWorkgroupID == "" {
				baselineOrganizationAddress = common.FreeInput("Baseline Workgroup ID", "", common.HexValidation)
			}
			if nchainBaselineNetworkID == "0x" {
				baselineOrganizationAddress = common.FreeInput("Nchain Baseline Network ID", "0x", common.HexValidation)
			}
		}
		common.ManifestSave(content)
		runProxyRun(cmd, args)
	case promptStepStop:
		if Optional {
			fmt.Println("Optional Flags:")
			if name == "" {
				name = common.FreeInput("Name", "", common.NoValidation)
			}
		}
		stopProxyRun(cmd, args)
	case promptStepLogs:
		if Optional {
			fmt.Println("Optional Flags:")
			if name == "" {
				name = common.FreeInput("Name", "", common.NoValidation)
			}
		}
		logsProxyRun(cmd, args)
	case "":
		result := common.SelectInput(emptyPromptArgs, emptyPromptLabel)
		generalPrompt(cmd, args, result)
	}
}

// var portMappings = common.PortMapping{
// 	[]portMapping{
// 		hostPort:      port,
// 		containerPort: apiContainerPort,
// 	},
// }
var content, err = json.Marshal(common.Keys{
	common.APIEndpoint,
	apiContainerPort,
	natsContainerPort,
	natsWebsocketContainerPort,
	natsStreamingContainerPort,
	postgresContainerPort,
	redisContainerPort,
	//portMappings,
	dockerNetworkID,
	Optional,
	name,
	port,
	identPort,
	nchainPort,
	privacyPort,
	vaultPort,
	natsPort,
	natsWebsocketPort,
	natsStreamingPort,
	postgresPort,
	redisPort,

	apiHostname,
	consumerHostname,
	identHostname,
	identConsumerHostname,
	nchainHostname,
	nchainConsumerHostname,
	nchainStatsdaemonHostname,
	nchainReachabilitydaemonHostname,
	privacyHostname,
	privacyConsumerHostname,
	vaultHostname,
	natsHostname,
	natsServerName,
	natsStreamingHostname,
	postgresHostname,
	redisHostname,
	redisHosts,

	autoRemove,
	logLevel,

	baselineOrganizationAddress,

	//  baselineOrganizationAPIEndpoint
	baselineRegistryContractAddress,
	baselineWorkgroupID,

	nchainBaselineNetworkID,

	jwtSignerPublicKey,
	natsAuthToken,

	identAPIHost,
	identAPIScheme,

	nchainAPIHost,
	nchainAPIScheme,

	workgroupAccessToken,
	organizationRefreshToken,

	privacyAPIHost,
	privacyAPIScheme,

	sorID,
	sorURL,

	vaultAPIHost,
	vaultAPIScheme,
	vaultRefreshToken,
	vaultSealUnsealKey,

	sapAPIHost,
	sapAPIScheme,
	sapAPIUsername,
	sapAPIPassword,
	sapAPIPath,

	serviceNowAPIHost,
	serviceNowAPIScheme,
	serviceNowAPIUsername,
	serviceNowAPIPassword,
	serviceNowAPIPath,

	salesforceAPIHost,
	salesforceAPIScheme,
	salesforceAPIPath,

	withLocalVault,
	withLocalIdent,
	withLocalNChain,
	withLocalPrivacy,
})
