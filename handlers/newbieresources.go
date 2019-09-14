package handlers

import (
	"context"
	"strings"

	"github.com/gobridge/gopher/bot"
)

// NewbieResources responds with some beginner resources when Message.TrimmedText
// has prefix.
//
// If Message.TrimmedText also contains "pvt" the response will be sent as a direct
// message.
func NewbieResources(prefix string) bot.Handler {
	return bot.HandlerFunc(func(ctx context.Context, m bot.Message, r bot.Responder) {
		if !strings.HasPrefix(m.TrimmedText, prefix) {
			return
		}

		msg := "Here are some resources you should check out if you are learning / new to Go:"
		if strings.Contains(m.TrimmedText, "pvt") { // TODO: Is this useful? Direct messaging the bot works as well.
			r.RespondPrivateWithAttachment(ctx, msg, newbieResources)
			return
		}

		r.RespondWithAttachment(ctx, msg, newbieResources)
	})
}

var newbieResources = `First you should take the language tour: <https://tour.golang.org/>

Then, you should visit:
- <https://golang.org/doc/code.html> to learn how to organize your Go workspace
- <https://golang.org/doc/effective_go.html> be more effective at writing Go
- <https://golang.org/ref/spec> learn more about the language itself
- <https://golang.org/doc/#articles> a lot more reading material

There are some awesome websites as well:
- <https://blog.gopheracademy.com> great resources for Gophers in general
- <http://gotime.fm> awesome weekly podcast of Go awesomeness
- <https://gobyexample.com> examples of how to do things in Go
- <http://go-database-sql.org> how to use SQL databases in Go
- <https://dmitri.shuralyov.com/idiomatic-go> tips on how to write more idiomatic Go code
- <https://divan.github.io/posts/avoid_gotchas> will help you avoid gotchas in Go
- <https://golangbot.com> tutorials to help you get started in Go

There's also an exhaustive list of videos <http://gophervids.appspot.com> related to Go from various authors.

If you prefer books, you can try these:
- <http://www.golangbootcamp.com/book>
- <http://gopl.io/>
- <https://www.manning.com/books/go-in-action> (if you e-mail @wkennedy at bill@ardanlabs.com you can get a free copy for being part of this Slack)

If you want to learn how to organize your Go project, make sure to read: <https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1#.ds38va3pp>.
Once you are accustomed to the language and syntax, you can read this series of articles for a walkthrough the various standard library packages: <https://medium.com/go-walkthrough>.

Finally, <https://github.com/golang/go/wiki#learning-more-about-go> will give a list of even more resources to learn Go`
