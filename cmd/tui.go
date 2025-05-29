package cmd

import (
	"aliasvault/vault"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

var visibleAliases []vault.Alias

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Launch interactive terminal UI to manage saved commands",
	Run: func(cmd *cobra.Command, args []string) {
		app := tview.NewApplication()
		table := tview.NewTable().SetBorders(false).SetSelectable(true, false)
		searchInput := tview.NewInputField().SetLabel("Search: ").SetFieldWidth(50)
		footerText := tview.NewTextView().SetText("[Enter] Run   [a] Add   [e] Edit   [d] Delete   [/] Search   [q] Quit").SetTextColor(tcell.ColorGray)

		layout := tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(table, 0, 1, true).
			AddItem(searchInput, 1, 0, false).
			AddItem(footerText, 1, 0, false)

		headerStyle := tcell.StyleDefault.Foreground(tcell.ColorYellow).Bold(true)

		var fullAliases []vault.Alias
		var searchMode bool

		loadTableData := func(aliases []vault.Alias) {
			table.Clear()
			table.SetCell(0, 0, tview.NewTableCell("Alias").SetSelectable(false).SetStyle(headerStyle))
			table.SetCell(0, 1, tview.NewTableCell("Command").SetSelectable(false).SetStyle(headerStyle))
			table.SetCell(0, 2, tview.NewTableCell("Tags").SetSelectable(false).SetStyle(headerStyle))

			for i, a := range aliases {
				aliasCell := tview.NewTableCell(fmt.Sprintf(" %-20s", a.Alias)).SetTextColor(tcell.ColorLightCyan)
				cmdCell := tview.NewTableCell(fmt.Sprintf(" %-40s", a.Command)).SetTextColor(tcell.ColorWhite)
				tagsCell := tview.NewTableCell(fmt.Sprintf(" %-30s", strings.Join(a.Tags, ", "))).SetTextColor(tcell.ColorLightGreen)
				table.SetCell(i+1, 0, aliasCell)
				table.SetCell(i+1, 1, cmdCell)
				table.SetCell(i+1, 2, tagsCell)
			}
		}

		refreshData := func() {
			fullAliases, _ = vault.GetAllAliases()
			visibleAliases = fullAliases
			loadTableData(visibleAliases)
		}

		filterData := func(query string) {
			var filtered []vault.Alias
			q := strings.ToLower(query)
			for _, a := range fullAliases {
				if strings.Contains(strings.ToLower(a.Alias), q) ||
					strings.Contains(strings.ToLower(a.Command), q) ||
					strings.Contains(strings.ToLower(strings.Join(a.Tags, ",")), q) {
					filtered = append(filtered, a)
				}
			}
			visibleAliases = filtered
			loadTableData(visibleAliases)
		}

		refreshData()

		table.SetSelectedFunc(func(row, column int) {
			if row == 0 || row > len(fullAliases) {
				return
			}
			cmdStr := table.GetCell(row, 1).Text
			app.Suspend(func() {
				fmt.Println("Running:", cmdStr)
				c := exec.Command("cmd", "/C", cmdStr)
				c.Stdout = os.Stdout
				c.Stderr = os.Stderr
				c.Stdin = os.Stdin
				err := c.Run()
				if err != nil {
					fmt.Println("Error executing command:", err)
				}
				fmt.Print("Press Enter to return...")
				fmt.Scanln()
			})
		})

		table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if searchMode {
				return event
			}
			switch event.Rune() {
			case 'q':
				app.Stop()
			case 'd':
				row, _ := table.GetSelection()
				if row > 0 && row <= len(fullAliases) {
					_ = vault.DeleteAlias(fullAliases[row-1].Alias)
					refreshData()
				}
			case 'a':
				form := tview.NewForm()
				var alias, command, tags string
				form.AddInputField("Alias", "", 20, nil, func(val string) { alias = val })
				form.AddInputField("Command", "", 50, nil, func(val string) { command = val })
				form.AddInputField("Tags (comma separated)", "", 50, nil, func(val string) { tags = val })
				form.AddButton("Save", func() {
					tagList := strings.Split(tags, ",")
					_ = vault.SaveAlias(alias, command, tagList)
					refreshData()
					app.SetRoot(layout, true).SetFocus(table)
				})
				form.AddButton("Cancel", func() {
					app.SetRoot(layout, true).SetFocus(table)
				})
				form.SetBorder(true).SetTitle("Add New Alias").SetTitleAlign(tview.AlignLeft)
				form.SetFocus(0)
				app.SetRoot(form, true).SetFocus(form)
			case 'e':
				row, _ := table.GetSelection()
				if row > 0 && row <= len(visibleAliases) {
					a := visibleAliases[row-1]

					var command, tags string

					form := tview.NewForm()
					form.AddInputField("Command", a.Command, 50, nil, func(val string) { command = val })
					form.AddInputField("Tags (comma separated)", strings.Join(a.Tags, ", "), 50, nil, func(val string) { tags = val })

					form.AddButton("Save", func() {
						tagList := strings.Split(tags, ",")
						for i := range tagList {
							tagList[i] = strings.TrimSpace(tagList[i])
						}
						err := vault.SaveAlias(a.Alias, command, tagList)
						if err != nil {
							fmt.Println("Error saving alias:", err)
						}
						refreshData()
						app.SetRoot(layout, true).SetFocus(table)
					})
					form.AddButton("Cancel", func() {
						app.SetRoot(layout, true).SetFocus(table)
					})

					form.SetBorder(true).SetTitle(fmt.Sprintf("Edit Alias: %s", a.Alias)).SetTitleAlign(tview.AlignLeft)
					form.SetFocus(0)
					app.SetRoot(form, true).SetFocus(form)
				}

			case '/':
				searchMode = true
				app.SetFocus(searchInput)
			}
			return event
		})

		searchInput.SetDoneFunc(func(key tcell.Key) {
			searchMode = false
			app.SetFocus(table)
		})

		searchInput.SetChangedFunc(func(text string) {
			filterData(text)
		})

		app.SetRoot(layout, true).SetFocus(table)
		if err := app.Run(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
