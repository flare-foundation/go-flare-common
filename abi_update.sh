#!/bin/bash
set -euo pipefail

# Resolve paths relative to this script so it works from any CWD.
cd "$(dirname "$0")"

# Source artifacts. Output path is derived from the contract name:
#   contracts go to pkg/contracts/<scope>/<name>/<name>.abi
#   structs   go to pkg/tee/structs/<name>/<name>.abi
# For contracts, <name> is the lowercased contract name with any "Facet" suffix dropped.
# For structs,   <name> is the lowercased contract name with any "Structs" suffix dropped.
#
# Contracts listed in BIN_CONTRACTS additionally get their deployment bytecode
# written to <name>.bin, for abigen --bin (deployable bindings).

CONTRACTS=(
  "tee/implementation/TeePayments.sol/TeePayments.json"
  "tee/implementation/TeePaymentsBase.sol/TeePaymentsBase.json"
  "tee/implementation/TeePaymentsConfigVerifier.sol/TeePaymentsConfigVerifier.json"
  "tee/implementation/TeePaymentsFeeScheduleManager.sol/TeePaymentsFeeScheduleManager.json"
  "tee/implementation/TeePaymentsRegistry.sol/TeePaymentsRegistry.json"
  "tee/implementation/TeePaymentsUtxo.sol/TeePaymentsUtxo.json"
  "tee/implementation/TeeRewardOffersManager.sol/TeeRewardOffersManager.json"
  "tee/implementation/VrfVerifier.sol/VrfVerifier.json"
  "tee/facets/DiamondGovernanceFacet.sol/DiamondGovernanceFacet.json"
  "tee/facets/ExtensionGovernanceFacet.sol/ExtensionGovernanceFacet.json"
  "tee/facets/ExtensionManagerFacet.sol/ExtensionManagerFacet.json"
  "tee/facets/ExternalAddressesFacet.sol/ExternalAddressesFacet.json"
  "tee/facets/FlareTeeManagerInit.sol/FlareTeeManagerInit.json"
  "tee/facets/InstructionsFacet.sol/InstructionsFacet.json"
  "tee/facets/MachineEmergencyPauseFacet.sol/MachineEmergencyPauseFacet.json"
  "tee/facets/MachineManagerFacet.sol/MachineManagerFacet.json"
  "tee/facets/MachinePathManagerFacet.sol/MachinePathManagerFacet.json"
  "tee/facets/OperationFeesFacet.sol/OperationFeesFacet.json"
  "tee/facets/OwnerAllowlistFacet.sol/OwnerAllowlistFacet.json"
  "tee/facets/VerificationFacet.sol/VerificationFacet.json"
  "tee/facets/VrfFacet.sol/VrfFacet.json"
  "tee/facets/WalletBackupManagerFacet.sol/WalletBackupManagerFacet.json"
  "tee/facets/WalletKeyManagerFacet.sol/WalletKeyManagerFacet.json"
  "tee/facets/WalletManagerFacet.sol/WalletManagerFacet.json"
  "tee/facets/WalletProjectManagerFacet.sol/WalletProjectManagerFacet.json"
  "tee/facets/WalletProjectPauseFacet.sol/WalletProjectPauseFacet.json"
  "fdc2/implementation/Fdc2Hub.sol/Fdc2Hub.json"
  "fdc2/implementation/Fdc2RequestFeeConfigurations.sol/Fdc2RequestFeeConfigurations.json"
)

STRUCTS=(
  "tee/structs/TeeInstructionsStructs.sol/TeeInstructionsStructs.json"
  "tee/structs/TeeMachinePathStructs.sol/TeeMachinePathStructs.json"
  "tee/structs/TeeMachineStructs.sol/TeeMachineStructs.json"
  "tee/structs/TeePaymentsStructs.sol/TeePaymentsStructs.json"
  "tee/structs/TeeStructs.sol/TeeStructs.json"
  "tee/structs/TeeVerificationStructs.sol/TeeVerificationStructs.json"
  "tee/structs/TeeVrfStructs.sol/TeeVrfStructs.json"
  "tee/structs/TeeWalletStructs.sol/TeeWalletStructs.json"
  "fdc2/structs/Fdc2Structs.sol/Fdc2Structs.json"
)

# Contracts that additionally need deployment bytecode extracted to <name>.bin.
# Entries must match a value in CONTRACTS.
BIN_CONTRACTS=(
  "tee/implementation/VrfVerifier.sol/VrfVerifier.json"
)

lower() {
  echo "$1" | tr '[:upper:]' '[:lower:]'
}

extract() {
  local input_file="../../fsp/flare-smart-contracts-v2/artifacts/contracts/$1"
  local output_file="$2"

  if [ -f "$input_file" ]; then
    mkdir -p "$(dirname "$output_file")"
    local abi
    abi="$(jq '.abi' "$input_file")"
    if [ "$abi" = "null" ]; then
      echo "Missing .abi in $input_file" >&2
      return 1
    fi
    printf '%s\n' "$abi" > "$output_file"
    echo "Extracted ABI from $input_file → $output_file"
  else
    echo "File not found: $input_file"
  fi
}

# needsBin reports whether the given CONTRACTS entry is listed in BIN_CONTRACTS.
needsBin() {
  local entry="$1" b
  for b in ${BIN_CONTRACTS[@]+"${BIN_CONTRACTS[@]}"}; do
    [ "$b" = "$entry" ] && return 0
  done
  return 1
}

# extractBin writes the artifact's deployment bytecode (0x-prefixed, no trailing
# newline) to the output file, matching what abigen --bin expects.
extractBin() {
  local input_file="../../fsp/flare-smart-contracts-v2/artifacts/contracts/$1"
  local output_file="$2"

  if [ -f "$input_file" ]; then
    mkdir -p "$(dirname "$output_file")"
    local bin
    bin="$(jq -r '.bytecode' "$input_file")"
    if [ "$bin" = "null" ] || [ -z "$bin" ]; then
      echo "Missing .bytecode in $input_file" >&2
      return 1
    fi
    printf '%s' "$bin" > "$output_file"
    echo "Extracted bytecode from $input_file → $output_file"
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
  if needsBin "$entry"; then
    extractBin "$entry" "pkg/contracts/${scope}/${name}/${name}.bin"
  fi
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
