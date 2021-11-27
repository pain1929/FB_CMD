package command

import (
	"fmt"
	"phoenixbuilder/minecraft/mctype"
	"phoenixbuilder/minecraft"
	//"github.com/google/uuid"
	"encoding/json"
	"strings"
)

func TitleRequest(target mctype.Target, lines ...string) string {
	var items []TellrawItem
	for _, text := range lines {
		items=append(items,TellrawItem{Text:strings.Replace(text, "schematic", "sc***atic", -1)})
	}
	final := &TellrawStruct {
		RawText: items,
	}
	content, _ := json.Marshal(final)
	cmd := fmt.Sprintf("titleraw %v actionbar %s", target, content)
	return cmd
}

func Title(conn *minecraft.Conn, lines ...string) error {
	return SendSizukanaCommand(TitleRequest(mctype.AllPlayers, lines...), conn)
}