#!/bin/bash

# Source artifacts. Output path is derived from the contract name:
#   contracts go to pkg/contracts/<scope>/<name>/<name>.abi
#   structs   go to pkg/tee/structs/<name>/<name>.abi
# For contracts, <name> is the lowercased contract name with any "Facet" suffix dropped.
# For structs,   <name> is the lowercased contract name with any "Structs" suffix dropped.

CONTRACTS=(
  "tee/implementation/TeePaymentsFeeScheduleManager.sol/TeePaymentsFeeScheduleManager.json"
  "tee/implementation/TeePaymentsRegistry.sol/TeePaymentsRegistry.json"
  "tee/implementation/TeePaymentsLimitsManager.sol/TeePaymentsLimitsManager.json"
  "tee/facets/ExtensionManagerFacet.sol/ExtensionManagerFacet.json"
  "tee/facets/InstructionsFacet.sol/InstructionsFacet.json"
  "tee/facets/MachineManagerFacet.sol/MachineManagerFacet.json"
  "tee/facets/OwnerAllowlistFacet.sol/OwnerAllowlistFacet.json"
  "tee/implementation/TeePayments.sol/TeePayments.json"
  "tee/facets/VerificationFacet.sol/VerificationFacet.json"
  "tee/facets/VrfFacet.sol/VrfFacet.json"
  "tee/facets/WalletBackupManagerFacet.sol/WalletBackupManagerFacet.json"
  "tee/facets/WalletKeyManagerFacet.sol/WalletKeyManagerFacet.json"
  "tee/facets/WalletManagerFacet.sol/WalletManagerFacet.json"
  "tee/facets/WalletProjectManagerFacet.sol/WalletProjectManagerFacet.json"
  "tee/facets/OperationFeesFacet.sol/OperationFeesFacet.json"
  "tee/implementation/VrfVerifier.sol/VrfVerifier.json"
  "fdc2/implementation/Fdc2Hub.sol/Fdc2Hub.json"
  "fdc2/implementation/Fdc2RequestFeeConfigurations.sol/Fdc2RequestFeeConfigurations.json"
)

STRUCTS=(
  "tee/structs/TeeMachineStructs.sol/TeeMachineStructs.json"
  "tee/structs/TeePaymentsStructs.sol/TeePaymentsStructs.json"
  "tee/structs/TeeVrfStructs.sol/TeeVrfStructs.json"
  "tee/structs/TeeStructs.sol/TeeStructs.json"
  "tee/structs/TeeVerificationStructs.sol/TeeVerificationStructs.json"
  "tee/structs/TeeWalletStructs.sol/TeeWalletStructs.json"
  "fdc2/structs/Fdc2Structs.sol/Fdc2Structs.json"
)

lower() {
  echo "$1" | tr '[:upper:]' '[:lower:]'
}

extract() {
  local input_file="../../fsp/flare-smart-contracts-v2/artifacts/contracts/$1"
  local output_file="$2"

  if [ -f "$input_file" ]; then
    mkdir -p "$(dirname "$output_file")"
    jq '.abi' "$input_file" > "$output_file"
    echo "Extracted ABI from $input_file → $output_file"
  else
    echo "File not found: $input_file"
  fi
}

for entry in "${CONTRACTS[@]}"; do
  json="${entry##*/}"
  name="$(lower "${json%.json}")"
  name="${name%facet}"
  scope="${entry%%/*}"
  extract "$entry" "pkg/contracts/${scope}/${name}/${name}.abi"
done

for entry in "${STRUCTS[@]}"; do
  json="${entry##*/}"
  name="$(lower "${json%.json}")"
  name="${name%structs}"
  if [[ "$name" == tee?* ]]; then
    name="${name#tee}"
  fi
  extract "$entry" "pkg/tee/structs/${name}/${name}.abi"
done
