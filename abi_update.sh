#!/bin/bash

# List of input/output pairs:
FILES=(
  "tee/implementation/TeeExtensionRegistry.sol/TeeExtensionRegistry.json:contracts/teeextensionregistry/teeextensionregistry.abi"
  "tee/implementation/TeeMachineRegistry.sol/TeeMachineRegistry.json:contracts/teemachineregistry/teemachineregistry.abi"
  "tee/implementation/TeeOwnerAllowlist.sol/TeeOwnerAllowlist.json:contracts/teeownerallowlist/teeownerallowlist.abi"
  "tee/implementation/TeePayments.sol/TeePayments.json:contracts/teepayment/teepayment.abi"
  "tee/implementation/TeeVerification.sol/TeeVerification.json:contracts/teeverification/teeverification.abi"
  "tee/implementation/TeeVRF.sol/TeeVRF.json:contracts/teevrf/teevrf.abi"
  "tee/implementation/TeeWalletBackupManager.sol/TeeWalletBackupManager.json:contracts/teewalletbackupmanager/teewalletbackupmanager.abi"
  "tee/implementation/TeeWalletKeyManager.sol/TeeWalletKeyManager.json:contracts/teewalletkeymanager/teewalletkeymanager.abi"
  "tee/implementation/TeeWalletManager.sol/TeeWalletManager.json:contracts/teewalletmanager/teewalletmanager.abi"
  "tee/implementation/TeeWalletProjectManager.sol/TeeWalletProjectManager.json:contracts/teewalletprojectmanager/teewalletprojectmanager.abi"
  "tee/implementation/TeeFeeCalculator.sol/TeeFeeCalculator.json:contracts/teefeecalculator/teefeecalculator.abi"
  "tee/lib/VRFVerifier.sol/VRFVerifier.json:contracts/teevrfverifier/teevrfverifier.abi"
  "fdc2/implementation/Fdc2Hub.sol/Fdc2Hub.json:contracts/fdc2hub/fdc2hub.abi"
  "fdc2/implementation/Fdc2RequestFeeConfigurations.sol/Fdc2RequestFeeConfigurations.json:contracts/fdc2requestfeeconfigurations/fdc2requestfeeconfigurations.abi"
  "tee/structs/TeeMachineRegistryStructs.sol/TeeMachineRegistryStructs.json:tee/structs/machineregistry/machineregistry.abi"
  "tee/structs/TeePaymentsStructs.sol/TeePaymentsStructs.json:tee/structs/payment/payment.abi"
  "tee/structs/TeeVrfStructs.sol/TeeVrfStructs.json:tee/structs/vrf/vrf.abi"
  "tee/structs/TeeStructs.sol/TeeStructs.json:tee/structs/tee/tee.abi"
  "tee/structs/TeeVerificationStructs.sol/TeeVerificationStructs.json:tee/structs/verification/verification.abi"
  "tee/structs/TeeVersionStructs.sol/TeeVersionStructs.json:tee/structs/version/version.abi"
  "tee/structs/TeeWalletStructs.sol/TeeWalletStructs.json:tee/structs/wallet/wallet.abi"
  "fdc2/structs/Fdc2Structs.sol/Fdc2Structs.json:tee/structs/connector/connector.abi"
)

for entry in "${FILES[@]}"; do
  INPUT_FILE="../flare-smart-contracts-v2/artifacts/contracts/${entry%%:*}"
  OUTPUT_FILE="pkg/${entry##*:}"

  if [ -f "$INPUT_FILE" ]; then
    jq '.abi' "$INPUT_FILE" > "$OUTPUT_FILE"
    echo "Extracted ABI from $INPUT_FILE → $OUTPUT_FILE"
  else
    echo "File not found: $INPUT_FILE"
  fi
done