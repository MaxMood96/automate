+++
title = "AWS Deployment using EFS"

draft = false

gh_repo = "automate"

[menu]
  [menu.automate]
    title = "AWS Deployment using EFS"
    identifier = "automate/deploy_high_availability/backup_and_restore/ha_backup_restore_prerequisites.md Backup and Restore AWS Deployment - EFS"
    parent = "automate/deploy_high_availability/backup_and_restore"
    weight = 230
+++

{{< note >}}
{{% automate/ha-warn %}}
{{< /note >}}

{{< note >}}

- If the user chooses `backup_config` as `efs` in `config.toml` backup is already configured during deployment, **the below steps are not required and can be skipped**. i.e., **`backup_config = "efs"`** . If we have kept the `backup_config` blank, then the configuration needs to be configured manually.

{{< /note >}}

## Overview

A shared file system is always required to create **OpenSearch** snapshots. To register the snapshot repository using OpenSearch, it is necessary to mount the same shared filesystem to the exact location on all master and data nodes. Register the location in the `path.repo` setting on all master and data nodes.

### Setting up the backup configuration

- To create an EFS file system, please refer to the sample steps on the [Create your Amazon EFS file system](https://docs.aws.amazon.com/efs/latest/ug/gs-step-two-create-efs-resources.html) page.

- Let's create a folder structure `/mnt/automate_backups/` on all the Frontend and backend nodes, then we have to mount EFS to all the vm's manually. To do that please refer [this](https://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-helper-ec2-linux.html)

#### Configuration in OpenSearch Node

- Mount the EFS on all OpenSearch Node. For example you mount the EFS to folder structure `/mnt/automate_backups/`

- Create an `opensearch` sub-directory and set permissions as mention below  (all the OpenSearch nodes).

    ```sh
    sudo mkdir -p /mnt/automate_backups/opensearch
    sudo chown hab:hab /mnt/automate_backups/opensearch/
    ```

#### Configuration for OpenSearch Node from Provision host

Configure the OpenSearch `path.repo` attribute.

- Create a toml file (`os_config.toml`) and add below template

    ```toml
    [path]
    repo = "/mnt/automate_backups/opensearch"
    ```

- Patch the config `os_config.toml` from bastion to the OpenSearch cluster.

    ```sh
    chef-automate config patch --opensearch os_config.toml
    ```

- Above command will restart the OpenSearch cluster.

#### Healthcheck commands
- Get the OpenSearch Cluster status

    ```sh
    chef-automate status --os
    ```

- Following command can be run in the OpenSearch node

    ```sh
    hab svc status (check whether OpenSearch service is up or not)

    curl -k -X GET "<https://localhost:9200/_cat/indices/*?v=true&s=index&pretty>" -u admin:admin (Another way to check is to check whether all the indices are green or not)

    # Watch for a message about OpenSearch going from RED to GREEN
    `journalctl -u hab-sup -f | grep 'automate-ha-opensearch'
    ```

#### Configuration for Automate node from Bastion host

- Mount the EFS to all the Frontend node manually. For example you mount the EFS to folder structure `/mnt/automate_backups`
- Create an `automate.toml` file on the bastion host using the following command:

    ```bash
    touch automate.toml
    ```

- Add the following configuration to `automate.toml` on the bastion host:

    ```sh
    [global.v1.external.opensearch.backup]
    enable = true
    location = "fs"

    [global.v1.external.opensearch.backup.fs]
    # The `path.repo` setting you've configured on your OpenSearch nodes must be a parent directory of the setting you configure here:
    path = "/mnt/automate_backups/opensearch"

    [global.v1.backups.filesystem]
    path = "/mnt/automate_backups/backups"
    ```

- Patch the `config` using the below command.

    ```sh
    ./chef-automate config patch --frontend automate.toml
    ```

## Backup and Restore commands

### Backup

- Run the backup command from Bastion as shown below to create a backup:

    ```sh
    chef-automate backup create
    ```

### Restore

#### Pre-Restore Validation

Run the restore command with the --verify-restore-config flag to validate the configuration settings before initiating the restore process.
To perform the pre-check, run the following command from the bastion host:

`chef-automate backup restore <BACKUP-ID> --yes -b /mnt/automate_backups/backups --verify-restore-config`

The verification process ensures that the backup and restore configurations are correct and identifies potential issues so they can be addressed in advance.

#### Run Restore

To restore backed-up data of the Chef Automate High Availability (HA) using External File System (EFS), follow the steps given below:

- Check the status of all Chef Automate and Chef Infra Server front-end nodes by executing the `chef-automate status` command.

- Execute the restore command from bastion`chef-automate backup restore <BACKUP-ID> -b /mnt/automate_backups/backups --airgap-bundle </path/to/bundle>`.

{{< note >}}

- If you are restoring the backup from an older version, then you need to provide the `--airgap-bundle </path/to/current/bundle>`.
- Large Compliance Report is not supported in Automate HA
{{< /note >}}

## Troubleshooting

{{< readfile file = "content/automate/reusable/md/restore_troubleshooting.md" >}}
