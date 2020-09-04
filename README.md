# send-slack-message

![Go](https://github.com/mrturkmen06/notifyme/workflows/Go/badge.svg)

Super simple Github Action to send message to Slack. 

Usage: 

```yaml
- name: send-slack-message
  uses: mrturkmencom/notifyme@v1
  env:
    web_hook: ${{ secrets.HOOK }}
    message: "Hi, this is first message from my `send-slack-message` action ⚔️ "
```

- **web_hook** : It is required to be able to send given message to Slack. You can retrieve it from Slack, more information could be found here: [https://slack.com/intl/en-dk/help/articles/115005265063-Incoming-webhooks-for-Slack](https://slack.com/intl/en-dk/help/articles/115005265063-Incoming-webhooks-for-Slack)

- **message** : It is basically a message to send. 
