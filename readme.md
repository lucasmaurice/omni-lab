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

## Setup Backup cronjob

```
sudo crontab -e
0  *  *  *  *  docker compose --project-directory /home/lmaurice/omni up bk
```

## Restore from Backup

> You will maybe have to adjust the IP in the omni service parameters.

1. Clone this repository
2. Complete the configs/*.env files with required secreds (B2 keys and GPG passphrase)
3. Create the required folders

   ```
   mkdir data
   mkdir data/etcd
   mkdir data/etcd-backup
   mkdir gpg
   ```
4. Get the GPG keys for Backup and etcd from 1Pasword to ./gpg/
5. Run the restore script `docker compose up prep`
6. Remove the private key and passphrase from the .env files
