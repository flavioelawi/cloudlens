package view

import (
	"github.com/gdamore/tcell/v2"
	"github.com/one2nc/cloudlens/internal/ui"
)

type EcsClusters struct {
	ResourceViewer
}

func NewEcs(resource string) ResourceViewer {
	var ecs EcsClusters
	ecs.ResourceViewer = NewBrowser(resource)
	ecs.AddBindKeysFn(ecs.bindKeys)
	return &ecs
}
func (ecs *EcsClusters) bindKeys(aa ui.KeyActions) {
	aa.Add(ui.KeyActions{
		ui.KeyShiftB: ui.NewKeyAction("Sort Cluster-Arn", ecs.GetTable().SortColCmd("Cluster-Arn", true), true),
		// ui.KeyD:      ui.NewKeyAction("Describe", ecs.describeCluster, true),
		tcell.KeyEscape: ui.NewKeyAction("Back", ecs.App().PrevCmd, false),
	})
}

// func (ecs *Ecs) describeCluster(evt *tcell.EventKey) *tcell.EventKey {
// 	bName := ecs.GetTable().GetSelectedItem()
// 	f := describeResource
// 	if ecs.GetTable().enterFn != nil {
// 		f = ecs.GetTable().enterFn
// 	}
// 	if bName != "" {
// 		ecs.App().Flash().Info("Cluster-Name: " + bName)
// 		f(ecs.App(), ecs.GetTable().GetModel(), ecs.Resource(), bName)
// 	}

// 	return nil
// }

// func (ecs *Ecs) enterCmd(evt *tcell.EventKey) *tcell.EventKey {
// 	bName := ecs.GetTable().GetSelectedItem()
// 	if bName != "" {
// 		o := NewS3FileViewer("s3://", bName)
// 		ctx := context.WithValue(ecs.App().GetContext(), internal.BucketName, bName)
// 		ecs.App().SetContext(ctx)
// 		ctx = context.WithValue(ecs.App().GetContext(), internal.FolderName, "")
// 		ecs.App().SetContext(ctx)
// 		ecs.App().Flash().Info("Bucket Name: " + bName)
// 		ecs.App().inject(o)
// 		o.GetTable().SetTitle(o.path)
// 	}

// 	return nil
// }
