import wandb.apis.reports as wr

report = wr.Report(
    project='github-actions-report',
    title="Hello from Github Actions!",
    description="This report was generated from a github action on the marketplace",
    blocks=[
        wr.H1([
            'To learn more, check out the action ',
            wr.Link("on the github marketplace", "https://github.com/marketplace/actions/generate-weights-biases-report")
        ]),
        wr.PanelGrid(
            runsets=[wr.Runset(project="report-api-quickstart")],
            panels=[
                wr.LinePlot(y='val_acc', title='Validation Accuracy over Time'),
                wr.BarPlot(metrics='acc'),
                wr.MediaBrowser(media_keys='img', num_columns=1)
            ],
        ),
    ]
).save()

print(report.url)
