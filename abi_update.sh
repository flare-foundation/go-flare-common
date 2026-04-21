#!/bin/bash

# List of input/output pairs:
FILES=(
  "tee/implementation/TeePaymentsFeeScheduleManager.sol/TeePaymentsFeeScheduleManager.json:contracts/teepaymentsfeeschedulemanager/teepaymentsfeeschedulemanager.abi"
  "tee/implementation/TeePaymentsRegistry.sol/TeePaymentsRegistry.json:contracts/teepaymentsregistry/teepaymentsregistry.abi"
  "tee/implementation/TeePaymentsLimitsManager.sol/TeePaymentsLimitsManager.json:contracts/teepaymentslimitsmanager/teepaymentslimitsmanager.abi"
  "tee/facets/ExtensionManagerFacet.sol/ExtensionManagerFacet.json:contracts/teeextensionregistry/teeextensionregistry.abi"
  "tee/facets/InstructionsFacet.sol/InstructionsFacet.json:contracts/teeinstructions/teeinstructions.abi"
  "tee/facets/MachineManagerFacet.sol/MachineManagerFacet.json:contracts/teemachineregistry/teemachineregistry.abi"
  "tee/facets/OwnerAllowlistFacet.sol/OwnerAllowlistFacet.json:contracts/teeownerallowlist/teeownerallowlist.abi"
  "tee/implementation/TeePayments.sol/TeePayments.json:contracts/teepayment/teepayment.abi"
  "tee/facets/VerificationFacet.sol/VerificationFacet.json:contracts/teeverification/teeverification.abi"
  "tee/facets/VrfFacet.sol/VrfFacet.json:contracts/teevrf/teevrf.abi"
  "tee/facets/WalletBackupManagerFacet.sol/WalletBackupManagerFacet.json:contracts/teewalletbackupmanager/teewalletbackupmanager.abi"
  "tee/facets/WalletKeyManagerFacet.sol/WalletKeyManagerFacet.json:contracts/teewalletkeymanager/teewalletkeymanager.abi"
  "tee/facets/WalletManagerFacet.sol/WalletManagerFacet.json:contracts/teewalletmanager/teewalletmanager.abi"
  "tee/facets/WalletProjectManagerFacet.sol/WalletProjectManagerFacet.json:contracts/teewalletprojectmanager/teewalletprojectmanager.abi"
  "tee/facets/OperationFeesFacet.sol/OperationFeesFacet.json:contracts/teefeecalculator/teefeecalculator.abi"
  "tee/implementation/VrfVerifier.sol/VrfVerifier.json:contracts/teevrfverifier/teevrfverifier.abi"
  "fdc2/implementation/Fdc2Hub.sol/Fdc2Hub.json:contracts/fdc2hub/fdc2hub.abi"
  "fdc2/implementation/Fdc2RequestFeeConfigurations.sol/Fdc2RequestFeeConfigurations.json:contracts/fdc2requestfeeconfigurations/fdc2requestfeeconfigurations.abi"
  "tee/structs/TeeMachineStructs.sol/TeeMachineStructs.json:tee/structs/machineregistry/machineregistry.abi"
  "tee/structs/TeePaymentsStructs.sol/TeePaymentsStructs.json:tee/structs/payment/payment.abi"
  "tee/structs/TeeVrfStructs.sol/TeeVrfStructs.json:tee/structs/vrf/vrf.abi"
  "tee/structs/TeeStructs.sol/TeeStructs.json:tee/structs/tee/tee.abi"
  "tee/structs/TeeVerificationStructs.sol/TeeVerificationStructs.json:tee/structs/verification/verification.abi"
  "tee/structs/TeeWalletStructs.sol/TeeWalletStructs.json:tee/structs/wallet/wallet.abi"
  "fdc2/structs/Fdc2Structs.sol/Fdc2Structs.json:tee/structs/connector/connector.abi"
)

for entry in "${FILES[@]}"; do
  INPUT_FILE="../../fsp/flare-smart-contracts-v2/artifacts/contracts/${entry%%:*}"
  OUTPUT_FILE="pkg/${entry##*:}"

  if [ -f "$INPUT_FILE" ]; then
    mkdir -p "$(dirname "$OUTPUT_FILE")"
    jq '.abi' "$INPUT_FILE" > "$OUTPUT_FILE"
    echo "Extracted ABI from $INPUT_FILE → $OUTPUT_FILE"
  else
    echo "File not found: $INPUT_FILE"
  fi
done
