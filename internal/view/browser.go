package view

import (
	"context"
	"fmt"
	"sync"

	"github.com/gdamore/tcell/v2"
	"github.com/one2nc/cloud-lens/internal"
	"github.com/one2nc/cloud-lens/internal/ui"
	"github.com/rs/zerolog/log"
)

// Browser represents a generic resource browser.
type Browser struct {
	*Table
	context  context.Context
	cancelFn context.CancelFunc
	mx       sync.RWMutex
}

// NewBrowser returns a new browser.
func NewBrowser(resource string, ctx context.Context) ResourceViewer {
	return &Browser{
		Table: NewTable(resource),
		context: ctx,
	}
}

// Init watches all running pods in given namespace.
func (b *Browser) Init(ctx context.Context) error {
	b.context = ctx
	if err := b.Table.Init(ctx); err != nil {
		return err
	}

	b.bindKeys(b.Actions())
	for _, f := range b.bindKeysFn {
		f(b.Actions())
	}

	row, _ := b.GetSelection()
	if row == 0 && b.GetRowCount() > 0 {
		b.Select(1, 0)
	}
	b.GetModel().SetRefreshRate(DefaultRefreshRate)
	return nil
}

func (b *Browser) bindKeys(aa ui.KeyActions) {
	aa.Add(ui.KeyActions{
		tcell.KeyEscape: ui.NewSharedKeyAction("Filter Reset", b.resetCmd, false),
		tcell.KeyHelp:   ui.NewSharedKeyAction("Help", b.helpCmd, false),
	})
}

// Start initializes browser updates.
func (b *Browser) Start() {
	log.Info().Msg(fmt.Sprintf("59b:watch ctx type: %T", b.context.Value(internal.KeySession)))

	b.Stop()
	//b.GetModel().AddListener(b)
	b.Table.Start()
	if err := b.GetModel().Watch(b.context); err != nil {
		b.App().Flash().Err(fmt.Errorf("Watcher failed for %s -- %w", b.Resource(), err))
	}
}

// Stop terminates browser updates.
func (b *Browser) Stop() {
	b.mx.Lock()
	{
		if b.cancelFn != nil {
			b.cancelFn()
			b.cancelFn = nil
		}
	}
	b.mx.Unlock()
	//b.GetModel().RemoveListener(b)
	b.Table.Stop()
}

// Name returns the component name.
func (b *Browser) Name() string { return b.Table.Resource() }

// SetContextFn populates a custom context.
func (b *Browser) SetContext(ctx context.Context) {
	log.Info().Msg(fmt.Sprintf("set ctx type: %T", ctx.Value(internal.KeySession)))
	b.context = ctx
}

// GetTable returns the underlying table.
func (b *Browser) GetTable() *Table { return b.Table }

func (b *Browser) helpCmd(evt *tcell.EventKey) *tcell.EventKey {

	return evt
}

func (b *Browser) resetCmd(evt *tcell.EventKey) *tcell.EventKey {
	b.Table.GetModel().Refresh(b.context)
	b.Refresh()
	return nil
}