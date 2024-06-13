# GO-KMS-SSHKEYS
**Description:** The purpose of this project is to handle encrypted ssh keys using AWS KMS with their content stored on AWS SecretsManager.

## HOW-TO
To it works properly, you've to configure the ```config.json```, this file must be at the root level, use the example below and change the content to fit what you need.

## Config File example: ```config.json```

```
{
    
    "aws": {
        "kmsid": "mrk-xxxxxxxxxxxxxxxxxxxxxxx",
        "secret_name": "prod/kms_key",
        "region": "us-east-1"
    }
}
```