package xmpp

import (
    "fmt"
)

func (c* Client) GetRoster(jid string) {
    fmt.Fprintf(c.conn, "<iq from='%s' type='get' id='roster_1'>\n" +
            "<query xmlns='jabber:iq:roster'>\n" +
            "</iq>",
            xmlEscape(jid))
}
