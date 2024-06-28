# Home Omni Deployment

## Overview

> TODO

## Generate new GPG key

See the Sidero documentation: <https://omni.siderolabs.com/docs/how-to-guides/self-hosted/how-to-deploy-omni-on-prem/#create-etcd-encryption-key>

### TLDR

```bash
# Generate with NO passphrase
gpg --quick-generate-key "Omni (Used for etcd data encryption) how-to-guide@siderolabs.com" rsa4096 cert never
gpg --list-secret-keys
gpg --quick-add-key <fingerprint> rsa4096 encr never
gpg --export-secret-key --armor how-to-guide@siderolabs.com > omni.asc
```

