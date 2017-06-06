# Lucca Leaves, Slack reminder

[![CircleCI](https://circleci.com/gh/bureaux-a-partager/lucca-leaves-slack-reminder/tree/master.svg?style=svg&circle-token=639a55cc96c45f36e04335f57531b3e81cd576f7)](https://circleci.com/gh/bureaux-a-partager/lucca-leaves-slack-reminder/tree/master)

This project is used to notify a team in Slack of the leaves in Lucca.


## Get started

```bash
# Required
export SLACK_TOKEN=xxxx-xxxxxxxxxx-xxxxxxxxxxx-xxxxxxxxxxxx-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
export LUCCA_TOKEN=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

# Optional
export SLACK_CHANNEL=random      # Default: lifeteam
export SLACK_ICON_EMOJI=:rocket: # Default: :sunglasses:
export SLACK_USERNAME=Bob        # Default: Life team

./build/linux/lucca-leaves-slack-reminder
```


## License

See [LICENSE](https://github.com/bureaux-a-partager/gh-slack-reminder/blob/master/LICENSE).


## GPG Signatures

You can download Julien Breux's public key to verify the signature.

    gpg --keyserver hkp://pgp.mit.edu --recv-keys 951C3F93B6A8C22C
