package cappuccino

import (
	"context"
	"fmt"
	"image/color"
	"time"

	"barista.run/bar"
	"barista.run/outputs"
	"barista.run/pango"
	"github.com/abnt713/cappuccino/pkg/lib/github"
)

// NewGithubViewer creates a new Github Viewer instance.
func NewGithubViewer(
	token string,
	ignoreRead bool,
	icon GithubIcon,
	color GithubColor,
) *GithubViewer {
	return &GithubViewer{
		token:      token,
		ignoreRead: ignoreRead,
		icon:       icon,
		color:      color,
	}
}

// GithubViewer creates the github notification viewer.
type GithubViewer struct {
	token      string
	ignoreRead bool
	icon       GithubIcon
	color      GithubColor
}

// GenerateBaristaModule generates the github barista module.
func (gv *GithubViewer) GenerateBaristaModule() (bar.Module, error) {
	return gv, nil
}

// Stream receives barista informations.
func (gv *GithubViewer) Stream(s bar.Sink) {
	client := github.NewClient(context.Background(), gv.token, 60)
	lastModified := ""
	for {
		var pollInterval int
		lastModified, pollInterval = gv.countNotifications(client, lastModified, s)
		time.Sleep(time.Duration(pollInterval) * time.Second)
	}
}

func (gv *GithubViewer) countNotifications(client github.Client, lastModified string, s bar.Sink) (newLastModified string, pollInterval int) {
	result, err := client.GetNotifications(lastModified, !gv.ignoreRead)
	if err != nil {
		s.Error(err)
		return "", result.PollInterval
	}

	if result.NotModified {
		return result.LastModified, result.PollInterval
	}

	counter := map[string]int{"all": 0, "len": len(result.Notifications)}
	for _, notification := range result.Notifications {
		if !notification.IsUnread() && gv.ignoreRead {
			continue
		}

		counter[notification.GetReason()]++
		counter["all"]++
	}

	gv.printNotifications(s, counter["review_requested"], counter["all"])
	return result.LastModified, result.PollInterval
}

func (gv *GithubViewer) printNotifications(s bar.Sink, reviews int, all int) {
	content := fmt.Sprintf("%d | %d", reviews, all)
	s.Output(outputs.Pango(
		gv.icon.Github(),
		space,
		pango.Text(content),
	).Color(gv.color.Github()))
}

// GithubIcon retrieves the github icon.
type GithubIcon interface {
	Github() *pango.Node
}

// GithubColor retrieves the github viewer color.
type GithubColor interface {
	Github() color.Color
}
