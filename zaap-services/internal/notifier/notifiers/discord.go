package notifier

type DiscordNotifier struct {
	webhookUrl string
}

func New(webhookUrl string) *DiscordNotifier {
	return &DiscordNotifier{webhookUrl}
}

